package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料货架库位
type MaterialShelfBin struct {
	ModelID
	Code            string         `gorm:"size:50;comment:编号"`
	Description     string         `gorm:"size:500;comment:描述"`
	Identifier      string         `gorm:"size:50;comment:识别码"`
	CurrentState    string         `gorm:"size:50;comment:当前状态"`
	LastUpdateTime  time.Time      `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark          string         `gorm:"size:500;comment:备注"`
	MaterialShelfID string         `gorm:"size:36;comment:隶属料架ID"`
	MaterialShelf   *MaterialShelf `gorm:"constraint:OnDelete:CASCADE"` //隶属料架
	MaterialInfoID  *string        `gorm:"size:36;comment:当前物料ID"`
	MaterialInfo    *MaterialInfo  `gorm:"constraint:OnDelete:SET NULL"` //当前物料
}

func PBToMaterialShelfBins(in []*proto.MaterialShelfBinInfo) []*MaterialShelfBin {
	var result []*MaterialShelfBin
	for _, c := range in {
		result = append(result, PBToMaterialShelfBin(c))
	}
	return result
}

func PBToMaterialShelfBin(in *proto.MaterialShelfBinInfo) *MaterialShelfBin {
	if in == nil {
		return nil
	}

	var materialInfoID *string
	if in.MaterialInfoID != "" {
		materialInfoID = &in.MaterialInfoID
	}

	return &MaterialShelfBin{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		CurrentState:    in.CurrentState,
		Remark:          in.Remark,
		MaterialShelfID: in.MaterialShelfID,
		MaterialInfoID:  materialInfoID,
	}
}

func MaterialShelfBinsToPB(in []*MaterialShelfBin) []*proto.MaterialShelfBinInfo {
	var list []*proto.MaterialShelfBinInfo
	for _, f := range in {
		list = append(list, MaterialShelfBinToPB(f))
	}
	return list
}

func MaterialShelfBinToPB(in *MaterialShelfBin) *proto.MaterialShelfBinInfo {
	if in == nil {
		return nil
	}

	var materialInfoID string
	if in.MaterialInfoID != nil {
		materialInfoID = *in.MaterialInfoID
	}

	m := &proto.MaterialShelfBinInfo{
		Id:              in.ID,
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		CurrentState:    in.CurrentState,
		LastUpdateTime:  utils.FormatTime(in.LastUpdateTime),
		Remark:          in.Remark,
		MaterialShelfID: in.MaterialShelfID,
		MaterialShelf:   MaterialShelfToPB(in.MaterialShelf),
		MaterialInfoID:  materialInfoID,
		MaterialInfo:    MaterialInfoToPB(in.MaterialInfo),
	}
	return m
}
