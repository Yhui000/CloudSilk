package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// BindMaterialTray godoc
// @Summary 绑定物料托盘
// @Description 绑定物料托盘
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.BindMaterialTrayRequest true "BindMaterialTrayRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/material/bindmaterialtray [post]
func BindMaterialTray(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.BindMaterialTrayRequest{}
	resp := &proto.CommonResponse{Code: 20000}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求绑定物料托盘接口参数无效:%v", transID, err)
		return
	}

	if err = middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := logic.BindMaterialTray(req); err != nil {
		resp.Code = 500
		resp.Message = err.Error()
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/material")

	g.POST("bindmaterialtray", BindMaterialTray)
}
