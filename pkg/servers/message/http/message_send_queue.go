package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/message/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMessageSendQueue godoc
// @Summary 新增
// @Description 新增
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageSendQueueInfo true "Add MessageSendQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendqueue/add [post]
func AddMessageSendQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageSendQueueInfo{CreateUserID: middleware.GetUserID(c)}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建消息发送队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMessageSendQueue(model.PBToMessageSendQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMessageSendQueue godoc
// @Summary 更新
// @Description 更新
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MessageSendQueueInfo true "Update MessageSendQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendqueue/update [put]
func UpdateMessageSendQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MessageSendQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新消息发送队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMessageSendQueue(model.PBToMessageSendQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMessageSendQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param taskNo query string false "任务编号"
// @Param messageTypeID query string false "消息类型ID"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryMessageSendQueueResponse
// @Router /api/mom/message/messagesendqueue/query [get]
func QueryMessageSendQueue(c *gin.Context) {
	req := &proto.QueryMessageSendQueueRequest{}
	resp := &proto.QueryMessageSendQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMessageSendQueue(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllMessageSendQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMessageSendQueueResponse
// @Router /api/mom/message/messagesendqueue/all [get]
func GetAllMessageSendQueue(c *gin.Context) {
	resp := &proto.GetAllMessageSendQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMessageSendQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MessageSendQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMessageSendQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMessageSendQueueDetailResponse
// @Router /api/mom/message/messagesendqueue/detail [get]
func GetMessageSendQueueDetail(c *gin.Context) {
	resp := &proto.GetMessageSendQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMessageSendQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageSendQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMessageSendQueue godoc
// @Summary 删除
// @Description 删除
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MessageSendQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendqueue/delete [delete]
func DeleteMessageSendQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除消息发送队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMessageSendQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// SendMessageSendQueue godoc
// @Summary 发送
// @Description 发送
// @Tags 消息发送队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Send MessageSendQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/message/messagesendqueue/send [get]
func SendMessageSendQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,发送消息发送队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.SendMessageSendQueue(middleware.GetUserID(c), req.Ids)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMessageSendQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/message/messagesendqueue")

	g.POST("add", AddMessageSendQueue)
	g.PUT("update", UpdateMessageSendQueue)
	g.GET("query", QueryMessageSendQueue)
	g.DELETE("delete", DeleteMessageSendQueue)
	g.GET("all", GetAllMessageSendQueue)
	g.GET("detail", GetMessageSendQueueDetail)
	g.GET("send", SendMessageSendQueue)
}
