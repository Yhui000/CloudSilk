package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 消息类型
type MessageType struct {
	ModelID
	Code        string `gorm:"size:50;comment:代号"`
	Description string `gorm:"size:500;comment:描述"`
	Remark      string `gorm:"size:500;comment:备注"`
}

func PBToMessageTypes(in []*proto.MessageTypeInfo) []*MessageType {
	var result []*MessageType
	for _, c := range in {
		result = append(result, PBToMessageType(c))
	}
	return result
}

func PBToMessageType(in *proto.MessageTypeInfo) *MessageType {
	if in == nil {
		return nil
	}
	return &MessageType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func MessageTypesToPB(in []*MessageType) []*proto.MessageTypeInfo {
	var list []*proto.MessageTypeInfo
	for _, f := range in {
		list = append(list, MessageTypeToPB(f))
	}
	return list
}

func MessageTypeToPB(in *MessageType) *proto.MessageTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.MessageTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
