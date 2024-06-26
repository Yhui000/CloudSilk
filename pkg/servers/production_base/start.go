package production_base

import (
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/provider"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/http"
	"github.com/gin-gonic/gin"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.ProductionCrosswayProvider{})
	config.SetProviderService(&provider.ProductionLineProvider{})
	config.SetProviderService(&provider.ProductionStationProvider{})
}
