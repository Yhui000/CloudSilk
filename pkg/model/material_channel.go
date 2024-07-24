package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料通道
type MaterialChannel struct {
	ModelID
	SortIndex              int32                 `gorm:"comment:顺序"`
	Code                   string                `gorm:"size:50;comment:代号"`
	Description            string                `gorm:"size:100;comment:描述"`
	Size                   string                `gorm:"size:50;comment:尺寸"`
	Spec                   float32               `gorm:"comment:规格"`
	EnableMonitor          bool                  `gorm:"comment:启用监控"`
	CurrentState           string                `gorm:"size:50;comment:当前状态"`
	LastUpdateTime         time.Time             `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                 string                `gorm:"size:500;comment:备注"`
	MaterialChannelLayerID string                `gorm:"size:36;comment:物料通道层ID"`
	MaterialChannelLayer   *MaterialChannelLayer `gorm:"constraint:OnDelete:CASCADE"` //物料通道层
	MaterialInfoID         string                `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo           *MaterialInfo         `gorm:"constraint:OnDelete:CASCADE"` //物料通道层
}

func PBToMaterialChannels(in []*proto.MaterialChannelInfo) []*MaterialChannel {
	var result []*MaterialChannel
	for _, c := range in {
		result = append(result, PBToMaterialChannel(c))
	}
	return result
}

func PBToMaterialChannel(in *proto.MaterialChannelInfo) *MaterialChannel {
	if in == nil {
		return nil
	}
	return &MaterialChannel{
		ModelID:                ModelID{ID: in.Id},
		SortIndex:              in.SortIndex,
		Code:                   in.Code,
		Description:            in.Description,
		Size:                   in.Size,
		Spec:                   in.Spec,
		EnableMonitor:          in.EnableMonitor,
		CurrentState:           in.CurrentState,
		Remark:                 in.Remark,
		MaterialChannelLayerID: in.MaterialChannelLayerID,
		MaterialInfoID:         in.MaterialInfoID,
	}
}

func MaterialChannelsToPB(in []*MaterialChannel) []*proto.MaterialChannelInfo {
	var list []*proto.MaterialChannelInfo
	for _, f := range in {
		list = append(list, MaterialChannelToPB(f))
	}
	return list
}

func MaterialChannelToPB(in *MaterialChannel) *proto.MaterialChannelInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialChannelInfo{
		Id:                     in.ID,
		SortIndex:              in.SortIndex,
		Code:                   in.Code,
		Description:            in.Description,
		Size:                   in.Size,
		Spec:                   in.Spec,
		EnableMonitor:          in.EnableMonitor,
		CurrentState:           in.CurrentState,
		LastUpdateTime:         utils.FormatTime(in.LastUpdateTime),
		Remark:                 in.Remark,
		MaterialChannelLayerID: in.MaterialChannelLayerID,
		MaterialChannelLayer:   MaterialChannelLayerToPB(in.MaterialChannelLayer),
		MaterialInfoID:         in.MaterialInfoID,
		MaterialInfo:           MaterialInfoToPB(in.MaterialInfo),
	}
	return m
}
