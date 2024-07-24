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

// AddMaterialChannel godoc
// @Summary 新增
// @Description 新增
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelInfo true "Add MaterialChannel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannel/add [post]
func AddMaterialChannel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialChannel(model.PBToMaterialChannel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialChannel godoc
// @Summary 更新
// @Description 更新
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelInfo true "Update MaterialChannel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannel/update [put]
func UpdateMaterialChannel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialChannel(model.PBToMaterialChannel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialChannel godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 料架通道管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialChannelResponse
// @Router /api/mom/material/materialchannel/query [get]
func QueryMaterialChannel(c *gin.Context) {
	req := &proto.QueryMaterialChannelRequest{}
	resp := &proto.QueryMaterialChannelResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialChannel(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialChannel godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialChannelResponse
// @Router /api/mom/material/materialchannel/all [get]
func GetAllMaterialChannel(c *gin.Context) {
	resp := &proto.GetAllMaterialChannelResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialChannels()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialChannelsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialChannelDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialChannelDetailResponse
// @Router /api/mom/material/materialchannel/detail [get]
func GetMaterialChannelDetail(c *gin.Context) {
	resp := &proto.GetMaterialChannelDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialChannelByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialChannel godoc
// @Summary 删除
// @Description 删除
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialChannel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannel/delete [delete]
func DeleteMaterialChannel(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialChannel(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialChannelRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialchannel")

	g.POST("add", AddMaterialChannel)
	g.PUT("update", UpdateMaterialChannel)
	g.GET("query", QueryMaterialChannel)
	g.DELETE("delete", DeleteMaterialChannel)
	g.GET("all", GetAllMaterialChannel)
	g.GET("detail", GetMaterialChannelDetail)
}
