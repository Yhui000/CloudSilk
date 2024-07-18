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

// AddMessageSendTask godoc
// @Summary 新增
// @Description 新增
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageSendTaskInfo true "Add MessageSendTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendtask/add [post]
func AddMessageSendTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageSendTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建消息发送任务请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMessageSendTask(model.PBToMessageSendTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMessageSendTask godoc
// @Summary 更新
// @Description 更新
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageSendTaskInfo true "Update MessageSendTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendtask/update [put]
func UpdateMessageSendTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageSendTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新消息发送任务请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMessageSendTask(model.PBToMessageSendTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMessageSendTask godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param messageTypeID query string false "消息类型ID"
// @Param productionLineID query string false "隶属产线ID"
// @Success 200 {object} proto.QueryMessageSendTaskResponse
// @Router /api/mom/message/messagesendtask/query [get]
func QueryMessageSendTask(c *gin.Context) {
	req := &proto.QueryMessageSendTaskRequest{}
	resp := &proto.QueryMessageSendTaskResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMessageSendTask(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMessageSendTask godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMessageSendTaskResponse
// @Router /api/mom/message/messagesendtask/all [get]
func GetAllMessageSendTask(c *gin.Context) {
	resp := &proto.GetAllMessageSendTaskResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMessageSendTasks()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MessageSendTasksToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMessageSendTaskDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMessageSendTaskDetailResponse
// @Router /api/mom/message/messagesendtask/detail [get]
func GetMessageSendTaskDetail(c *gin.Context) {
	resp := &proto.GetMessageSendTaskDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMessageSendTaskByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageSendTaskToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMessageSendTask godoc
// @Summary 删除
// @Description 删除
// @Tags 消息发送任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MessageSendTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendtask/delete [delete]
func DeleteMessageSendTask(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除消息发送任务请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMessageSendTask(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMessageSendTaskRouter(r *gin.Engine) {
	g := r.Group("/api/mom/message/messagesendtask")

	g.POST("add", AddMessageSendTask)
	g.PUT("update", UpdateMessageSendTask)
	g.GET("query", QueryMessageSendTask)
	g.DELETE("delete", DeleteMessageSendTask)
	g.GET("all", GetAllMessageSendTask)
	g.GET("detail", GetMessageSendTaskDetail)
}
