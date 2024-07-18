package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/message/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMessageType godoc
// @Summary 新增
// @Description 新增
// @Tags 消息类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageTypeInfo true "Add MessageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetype/add [post]
func AddMessageType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建消息类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMessageType(model.PBToMessageType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMessageType godoc
// @Summary 更新
// @Description 更新
// @Tags 消息类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageTypeInfo true "Update MessageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetype/update [put]
func UpdateMessageType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新消息类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMessageType(model.PBToMessageType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMessageType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 消息类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMessageTypeResponse
// @Router /api/mom/message/messagetype/query [get]
func QueryMessageType(c *gin.Context) {
	req := &proto.QueryMessageTypeRequest{}
	resp := &proto.QueryMessageTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMessageType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMessageType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 消息类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMessageTypeResponse
// @Router /api/mom/message/messagetype/all [get]
func GetAllMessageType(c *gin.Context) {
	resp := &proto.GetAllMessageTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMessageTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MessageTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMessageTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 消息类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMessageTypeDetailResponse
// @Router /api/mom/message/messagetype/detail [get]
func GetMessageTypeDetail(c *gin.Context) {
	resp := &proto.GetMessageTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMessageTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMessageType godoc
// @Summary 删除
// @Description 删除
// @Tags 消息类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MessageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetype/delete [delete]
func DeleteMessageType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除消息类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMessageType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMessageTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/message/messagetype")

	g.POST("add", AddMessageType)
	g.PUT("update", UpdateMessageType)
	g.GET("query", QueryMessageType)
	g.DELETE("delete", DeleteMessageType)
	g.GET("all", GetAllMessageType)
	g.GET("detail", GetMessageTypeDetail)
}
