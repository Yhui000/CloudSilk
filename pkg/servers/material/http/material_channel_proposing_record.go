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

// AddMaterialChannelProposingRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelProposingRecordInfo true "Add MaterialChannelProposingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannelproposingrecord/add [post]
func AddMaterialChannelProposingRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelProposingRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建通道叫料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialChannelProposingRecord(model.PBToMaterialChannelProposingRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialChannelProposingRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelProposingRecordInfo true "Update MaterialChannelProposingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannelproposingrecord/update [put]
func UpdateMaterialChannelProposingRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelProposingRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新通道叫料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialChannelProposingRecord(model.PBToMaterialChannelProposingRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialChannelProposingRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialChannelProposingRecordResponse
// @Router /api/mom/material/materialchannelproposingrecord/query [get]
func QueryMaterialChannelProposingRecord(c *gin.Context) {
	req := &proto.QueryMaterialChannelProposingRecordRequest{}
	resp := &proto.QueryMaterialChannelProposingRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialChannelProposingRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialChannelProposingRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialChannelProposingRecordResponse
// @Router /api/mom/material/materialchannelproposingrecord/all [get]
func GetAllMaterialChannelProposingRecord(c *gin.Context) {
	resp := &proto.GetAllMaterialChannelProposingRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialChannelProposingRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialChannelProposingRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialChannelProposingRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialChannelProposingRecordDetailResponse
// @Router /api/mom/material/materialchannelproposingrecord/detail [get]
func GetMaterialChannelProposingRecordDetail(c *gin.Context) {
	resp := &proto.GetMaterialChannelProposingRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialChannelProposingRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelProposingRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialChannelProposingRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 通道叫料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialChannelProposingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannelproposingrecord/delete [delete]
func DeleteMaterialChannelProposingRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除通道叫料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialChannelProposingRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialChannelProposingRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialchannelproposingrecord")

	g.POST("add", AddMaterialChannelProposingRecord)
	g.PUT("update", UpdateMaterialChannelProposingRecord)
	g.GET("query", QueryMaterialChannelProposingRecord)
	g.DELETE("delete", DeleteMaterialChannelProposingRecord)
	g.GET("all", GetAllMaterialChannelProposingRecord)
	g.GET("detail", GetMaterialChannelProposingRecordDetail)
}
