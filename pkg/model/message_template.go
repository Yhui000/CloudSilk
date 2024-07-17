package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 消息模版
type MessageTemplate struct {
	ModelID
	Code              string              `gorm:"size:50;comment:代号"`
	Description       string              `gorm:"size:500;comment:描述"`
	Title             string              `gorm:"size:500;comment:标题"`
	Content           string              `gorm:"size:-1;comment:内容"`
	Remark            string              `gorm:"size:500;comment:备注"`
	MessageTypeID     string              `gorm:"size:36;comment:消息类型ID"`
	MessageType       *MessageType        `gorm:"constraint:OnDelete:CASCADE"`
	MessageParameters []*MessageParameter `gorm:"constraint:OnDelete:CASCADE"`
}

// 消息参数
type MessageParameter struct {
	ModelID
	MessageTemplateID string `gorm:"index;size:36;comment:消息模版ID"`
	Name              string `gorm:"size:50;comment:名称"`
	DefaultValue      string `gorm:"size:1000;comment:预设值"`
}

func PBToMessageTemplates(in []*proto.MessageTemplateInfo) []*MessageTemplate {
	var result []*MessageTemplate
	for _, c := range in {
		result = append(result, PBToMessageTemplate(c))
	}
	return result
}

func PBToMessageTemplate(in *proto.MessageTemplateInfo) *MessageTemplate {
	if in == nil {
		return nil
	}
	return &MessageTemplate{
		ModelID:           ModelID{ID: in.Id},
		Code:              in.Code,
		Description:       in.Description,
		Title:             in.Title,
		Content:           in.Content,
		Remark:            in.Remark,
		MessageTypeID:     in.MessageTypeID,
		MessageParameters: PBToMessageParameters(in.MessageParameters),
	}
}

func MessageTemplatesToPB(in []*MessageTemplate) []*proto.MessageTemplateInfo {
	var list []*proto.MessageTemplateInfo
	for _, f := range in {
		list = append(list, MessageTemplateToPB(f))
	}
	return list
}

func MessageTemplateToPB(in *MessageTemplate) *proto.MessageTemplateInfo {
	if in == nil {
		return nil
	}
	m := &proto.MessageTemplateInfo{
		Id:                in.ID,
		Code:              in.Code,
		Description:       in.Description,
		Title:             in.Title,
		Content:           in.Content,
		Remark:            in.Remark,
		MessageTypeID:     in.MessageTypeID,
		MessageType:       MessageTypeToPB(in.MessageType),
		MessageParameters: MessageParametersToPB(in.MessageParameters),
	}
	return m
}

func PBToMessageParameters(in []*proto.MessageParameterInfo) []*MessageParameter {
	var result []*MessageParameter
	for _, c := range in {
		result = append(result, PBToMessageParameter(c))
	}
	return result
}

func PBToMessageParameter(in *proto.MessageParameterInfo) *MessageParameter {
	if in == nil {
		return nil
	}
	return &MessageParameter{
		ModelID:           ModelID{ID: in.Id},
		MessageTemplateID: in.MessageTemplateID,
		Name:              in.Name,
		DefaultValue:      in.DefaultValue,
	}
}

func MessageParametersToPB(in []*MessageParameter) []*proto.MessageParameterInfo {
	var list []*proto.MessageParameterInfo
	for _, f := range in {
		list = append(list, MessageParameterToPB(f))
	}
	return list
}

func MessageParameterToPB(in *MessageParameter) *proto.MessageParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.MessageParameterInfo{
		Id:                in.ID,
		MessageTemplateID: in.MessageTemplateID,
		Name:              in.Name,
		DefaultValue:      in.DefaultValue,
	}
	return m
}
