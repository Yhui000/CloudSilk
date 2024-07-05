package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddAGVParkingSpace godoc
// @Summary 新增
// @Description 新增
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVParkingSpaceInfo true "Add AGVParkingSpace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvparkingspace/add [post]
func AddAGVParkingSpace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVParkingSpaceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建AGV停靠泊位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateAGVParkingSpace(model.PBToAGVParkingSpace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAGVParkingSpace godoc
// @Summary 更新
// @Description 更新
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVParkingSpaceInfo true "Update AGVParkingSpace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvparkingspace/update [put]
func UpdateAGVParkingSpace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVParkingSpaceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新AGV停靠泊位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateAGVParkingSpace(model.PBToAGVParkingSpace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryAGVParkingSpace godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productionLineID query string false "隶属产线ID"
// @Success 200 {object} proto.QueryAGVParkingSpaceResponse
// @Router /api/mom/material/agvparkingspace/query [get]
func QueryAGVParkingSpace(c *gin.Context) {
	req := &proto.QueryAGVParkingSpaceRequest{}
	resp := &proto.QueryAGVParkingSpaceResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryAGVParkingSpace(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllAGVParkingSpace godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllAGVParkingSpaceResponse
// @Router /api/mom/material/agvparkingspace/all [get]
func GetAllAGVParkingSpace(c *gin.Context) {
	resp := &proto.GetAllAGVParkingSpaceResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllAGVParkingSpaces()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.AGVParkingSpacesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetAGVParkingSpaceDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAGVParkingSpaceDetailResponse
// @Router /api/mom/material/agvparkingspace/detail [get]
func GetAGVParkingSpaceDetail(c *gin.Context) {
	resp := &proto.GetAGVParkingSpaceDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetAGVParkingSpaceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVParkingSpaceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAGVParkingSpace godoc
// @Summary 删除
// @Description 删除
// @Tags AGV停靠泊位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete AGVParkingSpace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvparkingspace/delete [delete]
func DeleteAGVParkingSpace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.DelRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除AGV停靠泊位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteAGVParkingSpace(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterAGVParkingSpaceRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/agvparkingspace")

	g.POST("add", AddAGVParkingSpace)
	g.PUT("update", UpdateAGVParkingSpace)
	g.GET("query", QueryAGVParkingSpace)
	g.DELETE("delete", DeleteAGVParkingSpace)
	g.GET("all", GetAllAGVParkingSpace)
	g.GET("detail", GetAGVParkingSpaceDetail)
}
