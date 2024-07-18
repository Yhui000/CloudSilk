package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 消息发送队列
type MessageSendQueue struct {
	ModelID
	TaskNo              string                       `gorm:"size:50;comment:任务编号"`
	MessageTypeID       string                       `gorm:"size:36;comment:消息类型ID"`
	MessageType         *MessageType                 `gorm:"constraint:OnDelete:CASCADE"` //消息类型
	To                  string                       `gorm:"size:-1;comment:收件地址"`
	Cc                  string                       `gorm:"size:-1;comment:抄送地址"`
	Title               string                       `gorm:"size:500;comment:标题"`
	Content             string                       `gorm:"size:-1;comment:内容"`
	SendTimes           int32                        `gorm:"comment:发送次数"`
	CreateTime          time.Time                    `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        string                       `gorm:"size:36;comment:创建人员"`
	CurrentState        string                       `gorm:"size:50;comment:当前状态"`
	LastUpdateTime      time.Time                    `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark              string                       `gorm:"size:500;comment:备注"`
	RemoteServiceTaskID *string                      `gorm:"size:36;comment:远程任务ID"`
	RemoteServiceTask   *RemoteServiceTask           `gorm:"constraint:OnDelete:SET NULL"` //远程任务
	MessageParameters   []*MessageSendQueueParameter `gorm:"constraint:OnDelete:CASCADE"`  //消息发送队列参数
	SendExecutions      []*MessageSendQueueExecution `gorm:"constraint:OnDelete:CASCADE"`  //消息发送队列执行
}

// 消息发送队列参数
type MessageSendQueueParameter struct {
	ModelID
	MessageSendQueueID string `gorm:"index;size:36;comment:消息发送队列ID"`
	Name               string `gorm:"size:50;comment:名称"`
	FixedValue         string `gorm:"size:5000;comment:设定值"`
	Remark             string `gorm:"size:500;comment:备注"`
}

// 消息发送队列执行
type MessageSendQueueExecution struct {
	ModelID
	MessageSendQueueID string    `gorm:"index;size:36;comment:消息发送队列ID"`
	Success            bool      `gorm:"comment:是否成功"`
	FailureReason      string    `gorm:"size:5000;comment:失败原因"`
	CreateTime         time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
}

func PBToMessageSendQueues(in []*proto.MessageSendQueueInfo) []*MessageSendQueue {
	var result []*MessageSendQueue
	for _, c := range in {
		result = append(result, PBToMessageSendQueue(c))
	}
	return result
}

func PBToMessageSendQueue(in *proto.MessageSendQueueInfo) *MessageSendQueue {
	if in == nil {
		return nil
	}

	var remoteServiceTaskID *string
	if in.RemoteServiceTaskID != "" {
		remoteServiceTaskID = &in.RemoteServiceTaskID
	}

	return &MessageSendQueue{
		ModelID:             ModelID{ID: in.Id},
		TaskNo:              in.TaskNo,
		MessageTypeID:       in.MessageTypeID,
		To:                  in.To,
		Cc:                  in.Cc,
		Title:               in.Title,
		Content:             in.Content,
		SendTimes:           in.SendTimes,
		CreateUserID:        in.CreateUserID,
		CurrentState:        in.CurrentState,
		Remark:              in.Remark,
		RemoteServiceTaskID: remoteServiceTaskID,
		MessageParameters:   PBToMessageSendQueueParameters(in.MessageParameters),
		SendExecutions:      PBToMessageSendQueueExecutions(in.SendExecutions),
	}
}

func MessageSendQueuesToPB(in []*MessageSendQueue) []*proto.MessageSendQueueInfo {
	var list []*proto.MessageSendQueueInfo
	for _, f := range in {
		list = append(list, MessageSendQueueToPB(f))
	}
	return list
}

func MessageSendQueueToPB(in *MessageSendQueue) *proto.MessageSendQueueInfo {
	if in == nil {
		return nil
	}

	var remoteServiceTaskID string
	if in.RemoteServiceTaskID != nil {
		remoteServiceTaskID = *in.RemoteServiceTaskID
	}

	m := &proto.MessageSendQueueInfo{
		Id:                  in.ID,
		TaskNo:              in.TaskNo,
		MessageTypeID:       in.MessageTypeID,
		MessageType:         MessageTypeToPB(in.MessageType),
		To:                  in.To,
		Cc:                  in.Cc,
		Title:               in.Title,
		Content:             in.Content,
		SendTimes:           in.SendTimes,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		CurrentState:        in.CurrentState,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		Remark:              in.Remark,
		RemoteServiceTaskID: remoteServiceTaskID,
		RemoteServiceTask:   RemoteServiceTaskToPB(in.RemoteServiceTask),
		MessageParameters:   MessageSendQueueParametersToPB(in.MessageParameters),
		SendExecutions:      MessageSendQueueExecutionsToPB(in.SendExecutions),
	}
	return m
}

func PBToMessageSendQueueParameters(in []*proto.MessageSendQueueParameterInfo) []*MessageSendQueueParameter {
	var result []*MessageSendQueueParameter
	for _, c := range in {
		result = append(result, PBToMessageSendQueueParameter(c))
	}
	return result
}

func PBToMessageSendQueueParameter(in *proto.MessageSendQueueParameterInfo) *MessageSendQueueParameter {
	if in == nil {
		return nil
	}
	return &MessageSendQueueParameter{
		ModelID:            ModelID{ID: in.Id},
		MessageSendQueueID: in.MessageSendQueueID,
		Name:               in.Name,
		FixedValue:         in.FixedValue,
		Remark:             in.Remark,
	}
}

func MessageSendQueueParametersToPB(in []*MessageSendQueueParameter) []*proto.MessageSendQueueParameterInfo {
	var list []*proto.MessageSendQueueParameterInfo
	for _, f := range in {
		list = append(list, MessageSendQueueParameterToPB(f))
	}
	return list
}

func MessageSendQueueParameterToPB(in *MessageSendQueueParameter) *proto.MessageSendQueueParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.MessageSendQueueParameterInfo{
		Id:                 in.ID,
		MessageSendQueueID: in.MessageSendQueueID,
		Name:               in.Name,
		FixedValue:         in.FixedValue,
		Remark:             in.Remark,
	}
	return m
}

func PBToMessageSendQueueExecutions(in []*proto.MessageSendQueueExecutionInfo) []*MessageSendQueueExecution {
	var result []*MessageSendQueueExecution
	for _, c := range in {
		result = append(result, PBToMessageSendQueueExecution(c))
	}
	return result
}

func PBToMessageSendQueueExecution(in *proto.MessageSendQueueExecutionInfo) *MessageSendQueueExecution {
	if in == nil {
		return nil
	}
	return &MessageSendQueueExecution{
		ModelID:            ModelID{ID: in.Id},
		MessageSendQueueID: in.MessageSendQueueID,
		Success:            in.Success,
		FailureReason:      in.FailureReason,
	}
}

func MessageSendQueueExecutionsToPB(in []*MessageSendQueueExecution) []*proto.MessageSendQueueExecutionInfo {
	var list []*proto.MessageSendQueueExecutionInfo
	for _, f := range in {
		list = append(list, MessageSendQueueExecutionToPB(f))
	}
	return list
}

func MessageSendQueueExecutionToPB(in *MessageSendQueueExecution) *proto.MessageSendQueueExecutionInfo {
	if in == nil {
		return nil
	}
	m := &proto.MessageSendQueueExecutionInfo{
		Id:                 in.ID,
		MessageSendQueueID: in.MessageSendQueueID,
		Success:            in.Success,
		FailureReason:      in.FailureReason,
		CreateTime:         utils.FormatTime(in.CreateTime),
	}
	return m
}
