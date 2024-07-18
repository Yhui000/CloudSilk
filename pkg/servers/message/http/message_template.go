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

// AddMessageTemplate godoc
// @Summary 新增
// @Description 新增
// @Tags 消息模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageTemplateInfo true "Add MessageTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetemplate/add [post]
func AddMessageTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageTemplateInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建消息模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMessageTemplate(model.PBToMessageTemplate(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMessageTemplate godoc
// @Summary 更新
// @Description 更新
// @Tags 消息模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageTemplateInfo true "Update MessageTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetemplate/update [put]
func UpdateMessageTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageTemplateInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新消息模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMessageTemplate(model.PBToMessageTemplate(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMessageTemplate godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 消息模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param messageTypeID query string false "消息类型ID"
// @Success 200 {object} proto.QueryMessageTemplateResponse
// @Router /api/mom/message/messagetemplate/query [get]
func QueryMessageTemplate(c *gin.Context) {
	req := &proto.QueryMessageTemplateRequest{}
	resp := &proto.QueryMessageTemplateResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMessageTemplate(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMessageTemplate godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 消息模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMessageTemplateResponse
// @Router /api/mom/message/messagetemplate/all [get]
func GetAllMessageTemplate(c *gin.Context) {
	resp := &proto.GetAllMessageTemplateResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMessageTemplates()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MessageTemplatesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMessageTemplateDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 消息模版管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMessageTemplateDetailResponse
// @Router /api/mom/message/messagetemplate/detail [get]
func GetMessageTemplateDetail(c *gin.Context) {
	resp := &proto.GetMessageTemplateDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMessageTemplateByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageTemplateToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMessageTemplate godoc
// @Summary 删除
// @Description 删除
// @Tags 消息模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MessageTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagetemplate/delete [delete]
func DeleteMessageTemplate(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除消息模版请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMessageTemplate(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMessageTemplateRouter(r *gin.Engine) {
	g := r.Group("/api/mom/message/messagetemplate")

	g.POST("add", AddMessageTemplate)
	g.PUT("update", UpdateMessageTemplate)
	g.GET("query", QueryMessageTemplate)
	g.DELETE("delete", DeleteMessageTemplate)
	g.GET("all", GetAllMessageTemplate)
	g.GET("detail", GetMessageTemplateDetail)
}
