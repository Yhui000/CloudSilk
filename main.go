package main

import (
	"flag"
	"fmt"
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/CloudSilk/CloudSilk/docs"
	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/servers/label"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system"
	"github.com/CloudSilk/CloudSilk/pkg/servers/trace"
	"github.com/CloudSilk/CloudSilk/pkg/servers/user"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/http"
	"github.com/CloudSilk/curd/gen"
	curdhttp "github.com/CloudSilk/curd/http"
	curdmodel "github.com/CloudSilk/curd/model"
	curdservice "github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/sqlite"
	"github.com/CloudSilk/pkg/utils"
	ucconfig "github.com/CloudSilk/usercenter/config"
	uchttp "github.com/CloudSilk/usercenter/http"
	ucmodel "github.com/CloudSilk/usercenter/model"
	"github.com/CloudSilk/usercenter/model/token"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type IServer interface {
	Start(*gin.Engine)
}

func StartAll(webPath, dbType string, port int) {
	err := ucconfig.InitFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println("数据库类型：", dbType)
	var dbClient db.DBClientInterface
	if dbType == "sqlite" {
		dbClient = sqlite.NewSqlite2("", "", "./mom.s3db", "mom", ucconfig.DefaultConfig.Debug)
	} else {
		dbClient = mysql.NewMysql(ucconfig.DefaultConfig.Mysql, ucconfig.DefaultConfig.Debug)
	}

	fmt.Println(webPath)
	model.InitDB(dbClient, ucconfig.DefaultConfig.Debug)
	curdmodel.InitDB(dbClient, ucconfig.DefaultConfig.Debug)
	curdservice.Init()
	ucmodel.InitDB(dbClient, ucconfig.DefaultConfig.Debug)
	token.InitTokenCache(ucconfig.DefaultConfig.Token.Key, ucconfig.DefaultConfig.Token.RedisAddr, ucconfig.DefaultConfig.Token.RedisName, ucconfig.DefaultConfig.Token.RedisPwd, ucconfig.DefaultConfig.Token.Expired)
	constants.SetPlatformTenantID(ucconfig.DefaultConfig.PlatformTenantID)
	constants.SetSuperAdminRoleID(ucconfig.DefaultConfig.SuperAdminRoleID)
	constants.SetDefaultRoleID(ucconfig.DefaultConfig.DefaultRoleID)
	constants.SetEnabelTenant(ucconfig.DefaultConfig.EnableTenant)
	ucmodel.SetDefaultPwd(ucconfig.DefaultConfig.DefaultPwd)

	gen.LoadCache()

	r := gin.Default()
	r.Use(ucmiddleware.AuthRequired)
	r.Use(utils.Cors())

	uchttp.RegisterAuthRouter(r)
	curdhttp.RegisterRouter(r)

	startMom(r)

	r.Static("/web", webPath)
	r.Run(fmt.Sprintf(":%d", port))
}

func startMom(r *gin.Engine) {
	docs.SwaggerInfo.Title = "mom API"
	docs.SwaggerInfo.Description = "This is a mom server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/mom/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	http.RegisterAdminRouter(r)

	servers := []IServer{
		&production_base.Server{},
		&production.Server{},
		&product_base.Server{},
		&product.Server{},
		&label.Server{},
		&system.Server{},
		&material.Server{},
		&user.Server{},
		&trace.Server{},
	}
	for _, server := range servers {
		server.Start(r)
	}
	clients.Init(os.Getenv("SERVICE_MODE"))
}

func StartOne(dbType string, port int) {
	r := gin.Default()
	r.Use(utils.Cors())

	if os.Getenv("MOM_DISABLE_AUTH") != "true" {
		ucmiddleware.InitIdentity()
		r.Use(ucmiddleware.AuthRequiredWithRPC)
	}

	startMom(r)
	if err := config.Load(); err != nil {
		panic(err)
	}
	params := config.GetRootConfig().ConfigCenter.Params
	var dbClient db.DBClientInterface
	debug := params["debug"] == "true"
	if dbType == "sqlite" {
		dbClient = sqlite.NewSqlite2("", "", "./mom.s3db", "mom", debug)
	} else {
		dbClient = mysql.NewMysql(params["mysql"], debug)
	}

	model.InitDB(dbClient, debug)
	fmt.Println("started server")

	r.Run(fmt.Sprintf(":%d", port))
}

func main() {
	webPath := flag.String("ui", "./web", "web路径")
	dbType := flag.String("db_type", "mysql", "数据库类型：sqlite和mysql两种")
	serviceMode := flag.String("service_mode", "One", "运行模式：ALL、One")
	port := flag.Int("port", 48900, "端口")
	flag.Parse()
	if *serviceMode == "ALL" {
		StartAll(*webPath, *dbType, *port)
	} else {
		StartOne(*dbType, *port)
	}
}