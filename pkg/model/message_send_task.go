package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type MessageSendTask struct {
	ModelID
	Code                string             `gorm:"size:50;comment:代号"`
	Description         string             `gorm:"size:500;comment:描述"`
	Enable              bool               `gorm:"comment:是否启用"`
	ProductionLineID    *string            `gorm:"size:36;comment:工厂产线ID"`
	ProductionLine      *ProductionLine    `gorm:"constraint:OnDelete:SET NULL"`
	MessageTemplateID   string             `gorm:"size:36;comment:消息模板ID"`
	MessageTemplate     *MessageTemplate   `gorm:"constraint:OnDelete:CASCADE"`
	To                  string             `gorm:"size:-1;comment:收件地址"`
	Cc                  string             `gorm:"size:-1;comment:抄送地址"`
	RemoteServiceTaskID *string            `gorm:"size:36;comment:远程服务ID"`
	RemoteServiceTask   *RemoteServiceTask `gorm:"constraint:OnDelete:SET NULL"`
	Remark              string             `gorm:"size:500;comment:备注"`
}

func PBToMessageSendTasks(in []*proto.MessageSendTaskInfo) []*MessageSendTask {
	var result []*MessageSendTask
	for _, c := range in {
		result = append(result, PBToMessageSendTask(c))
	}
	return result
}

func PBToMessageSendTask(in *proto.MessageSendTaskInfo) *MessageSendTask {
	if in == nil {
		return nil
	}

	var productionLineID, remoteServiceTaskID *string
	if in.ProductionLineID != "" {
		productionLineID = &in.ProductionLineID
	}
	if in.RemoteServiceTaskID != "" {
		remoteServiceTaskID = &in.RemoteServiceTaskID
	}

	return &MessageSendTask{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		ProductionLineID:    productionLineID,
		MessageTemplateID:   in.MessageTemplateID,
		To:                  in.To,
		Cc:                  in.Cc,
		RemoteServiceTaskID: remoteServiceTaskID,
		Remark:              in.Remark,
	}
}

func MessageSendTasksToPB(in []*MessageSendTask) []*proto.MessageSendTaskInfo {
	var list []*proto.MessageSendTaskInfo
	for _, f := range in {
		list = append(list, MessageSendTaskToPB(f))
	}
	return list
}

func MessageSendTaskToPB(in *MessageSendTask) *proto.MessageSendTaskInfo {
	if in == nil {
		return nil
	}

	var productionLineID, remoteServiceTaskID string
	if in.ProductionLineID != nil {
		productionLineID = *in.ProductionLineID
	}
	if in.RemoteServiceTaskID != nil {
		remoteServiceTaskID = *in.RemoteServiceTaskID
	}

	m := &proto.MessageSendTaskInfo{
		Id:                  in.ID,
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		ProductionLineID:    productionLineID,
		ProductionLine:      ProductionLineToPB(in.ProductionLine),
		MessageTemplateID:   in.MessageTemplateID,
		MessageTemplate:     MessageTemplateToPB(in.MessageTemplate),
		To:                  in.To,
		Cc:                  in.Cc,
		RemoteServiceTaskID: remoteServiceTaskID,
		RemoteServiceTask:   RemoteServiceTaskToPB(in.RemoteServiceTask),
		Remark:              in.Remark,
	}
	return m
}
