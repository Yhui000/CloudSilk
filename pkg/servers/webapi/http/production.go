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

// OnlineProductInfo godoc
// @Summary 设定产品信息状态为上线装配
// @Description 设定产品信息状态为上线装配
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.OnlineProductInfoRequest true "OnlineProductInfoRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/onlineproductinfo [post]
func OnlineProductInfo(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.OnlineProductInfoRequest{}
	resp := &proto.CommonResponse{Code: 20000}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求上线装配接口参数无效:%v", transID, err)
		return
	}

	if err = middleware.Validate.Struct(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, logic.OnlineProductInfo(req))
}

// EnterProductionStation godoc
// @Summary 请求进站接口
// @Description 请求进站接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EnterProductionStationRequest true "EnterProductionStationRequest"
// @Success 200 {object} proto.EnterProductionStationResponse
// @Router /api/mom/webapi/production/enterproductionstation [post]
func EnterProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EnterProductionStationRequest{}
	resp := &proto.EnterProductionStationResponse{Code: 0}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求进站接口参数无效:%v", transID, err)
		return
	}

	if err = middleware.Validate.Struct(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, logic.EnterProductionStation(req))
}

// GetProductionStationExhibition godoc
// @Summary 获取工站的当前生产工单工序信息
// @Description 获取工站的当前生产工单工序信息
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param productionStation query string true "生产工站代号"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductAttributeValuateRuleDetailResponse
// @Router /api/mom/webapi/production/getproductionstationexhibition [get]
// func GetProductionStationExhibition(c *gin.Context) {
// 	resp := &proto.GetProductionStationExhibitionResponse{
// 		Code: types.ServiceResponseCodeSuccess,
// 	}
// 	productionStationCode := c.Query("productionStation")
// 	if productionStationCode == "" {
// 		resp.Code = types.ServiceResponseCodeFailure
// 		c.JSON(http.StatusOK, resp)
// 		return
// 	}
// 	var err error

// 	data, err := logic.GetProductionStationExhibition(productionStationCode)
// 	if err != nil {
// 		resp.Code = types.ServiceResponseCodeFailure
// 		resp.Message = err.Error()
// 	} else {
// 		resp.Data = data
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// ExitProductionStation godoc
// @Summary 请求出站接口
// @Description 请求出站接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ExitProductionStationRequest true "ExitProductionStationRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/exitproductionstation [post]
func ExitProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ExitProductionStationRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求进站接口参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	_, err := logic.ExitProductionStation(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
	}

	c.JSON(http.StatusOK, resp)
}

// CreateProductTestRecord godoc
// @Summary 创建产品测试记录
// @Description 创建产品测试记录
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CreateProductTestRecordRequest true "CreateProductTestRecordRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/createproducttestrecord [post]
func CreateProductTestRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CreateProductTestRecordRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,创建产品测试记录参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.CreateProductTestRecord(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CheckProductProcessRouteFailure godoc
// @Summary 设置失败后续处理接口
// @Description 设置失败后续处理接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CheckProductProcessRouteFailureRequest true "CheckProductProcessRouteFailureRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/checkproductprocessroutefailure [post]
func CheckProductProcessRouteFailure(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CheckProductProcessRouteFailureRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,设置失败后续处理接口参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.CheckProductProcessRouteFailure(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetProductionProcessStepWithParameter godoc
// @Summary 获取作业步骤和作业参数
// @Description 获取作业步骤和作业参数
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.GetProductionProcessStepWithParameterRequest true "GetProductionProcessStepWithParameterRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/getproductionprocessstepwithparameter [post]
func GetProductionProcessStepWithParameter(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetProductionProcessStepWithParameterRequest{}
	resp := map[string]interface{}{"code": 0}

	if err := c.BindJSON(req); err != nil {
		resp["code"] = 1
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,获取作业步骤和作业参数接口参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp["code"] = 1
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := logic.GetProductionProcessStepWithParameter(req)
	if err != nil {
		resp["code"] = 1
		resp["message"] = err.Error()
	} else {
		resp["data"] = data
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterProductionRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/production")

	g.POST("onlineproductinfo", OnlineProductInfo)
	g.POST("enterproductionstation", EnterProductionStation)
	// g.GET("getproductionstationexhibition", GetProductionStationExhibition)
	g.POST("exitproductionstation", ExitProductionStation)
	g.POST("createproducttestrecord", CreateProductTestRecord)
	g.POST("checkproductprocessroutefailure", CheckProductProcessRouteFailure)
	g.POST("getproductionprocessstepwithparameter", GetProductionProcessStepWithParameter)
}
