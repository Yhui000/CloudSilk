package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料通道层
type MaterialChannelLayer struct {
	ModelID
	SortIndex             int32              `gorm:"comment:顺序"`
	Code                  string             `gorm:"size:50;comment:代号"`
	StatusRegisterAddress int32              `gorm:"comment:状态寄存器地址"`
	LightRegisterAddress  int32              `gorm:"comment:亮灯寄存器地址"`
	Remark                string             `gorm:"size:500;comment:备注"`
	ProductionStationID   string             `gorm:"size:36;comment:隶属工站ID"`
	ProductionStation     *ProductionStation `gorm:"constraint:OnDelete:CASCADE"` //隶属工站
	MaterialChannels      []*MaterialChannel `gorm:"constraint:OnDelete:CASCADE"` //物料通道
}

func PBToMaterialChannelLayers(in []*proto.MaterialChannelLayerInfo) []*MaterialChannelLayer {
	var result []*MaterialChannelLayer
	for _, c := range in {
		result = append(result, PBToMaterialChannelLayer(c))
	}
	return result
}

func PBToMaterialChannelLayer(in *proto.MaterialChannelLayerInfo) *MaterialChannelLayer {
	if in == nil {
		return nil
	}
	return &MaterialChannelLayer{
		ModelID:               ModelID{ID: in.Id},
		SortIndex:             in.SortIndex,
		Code:                  in.Code,
		StatusRegisterAddress: in.StatusRegisterAddress,
		LightRegisterAddress:  in.LightRegisterAddress,
		Remark:                in.Remark,
		ProductionStationID:   in.ProductionStationID,
		MaterialChannels:      PBToMaterialChannels(in.MaterialChannels),
	}
}

func MaterialChannelLayersToPB(in []*MaterialChannelLayer) []*proto.MaterialChannelLayerInfo {
	var list []*proto.MaterialChannelLayerInfo
	for _, f := range in {
		list = append(list, MaterialChannelLayerToPB(f))
	}
	return list
}

func MaterialChannelLayerToPB(in *MaterialChannelLayer) *proto.MaterialChannelLayerInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialChannelLayerInfo{
		Id:                    in.ID,
		SortIndex:             in.SortIndex,
		Code:                  in.Code,
		StatusRegisterAddress: in.StatusRegisterAddress,
		LightRegisterAddress:  in.LightRegisterAddress,
		Remark:                in.Remark,
		ProductionStationID:   in.ProductionStationID,
		ProductionStation:     ProductionStationToPB(in.ProductionStation),
		MaterialChannels:      MaterialChannelsToPB(in.MaterialChannels),
	}
	return m
}
