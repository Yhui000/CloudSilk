package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料货架
type MaterialShelf struct {
	ModelID
	Code              string                         `gorm:"size:50;comment:代号"`
	Description       string                         `gorm:"size:500;comment:描述"`
	ShelfType         int32                          `gorm:"comment:货架类型"`
	Enable            bool                           `gorm:"comment:是否启用"`
	CurrentState      string                         `gorm:"size:50;comment:当前状态"`
	LastUpdateTime    time.Time                      `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark            string                         `gorm:"size:500;comment:备注"`
	ProductionLineID  string                         `gorm:"size:36;comment:隶属产线ID"`
	ProductionLine    *ProductionLine                `gorm:"constraint:OnDelete:CASCADE"` //隶属产线
	AGVParkingSpaceID *string                        `gorm:"size:36;comment:当前泊位ID"`
	AGVParkingSpace   *AGVParkingSpace               `gorm:"constraint:OnDelete:SET NULL"`
	ParkableSpaces    []*MaterialShelfAvailableSpace `gorm:"constraint:OnDelete:CASCADE"` //可停靠泊位
}

type MaterialShelfAvailableSpace struct {
	ModelID
	MaterialShelfID   string           `gorm:"index;size:36;comment:物料货架ID"`
	AGVParkingSpaceID string           `gorm:"size:36;comment:AGV停靠泊位ID"`
	AGVParkingSpace   *AGVParkingSpace `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialShelfs(in []*proto.MaterialShelfInfo) []*MaterialShelf {
	var result []*MaterialShelf
	for _, c := range in {
		result = append(result, PBToMaterialShelf(c))
	}
	return result
}

func PBToMaterialShelf(in *proto.MaterialShelfInfo) *MaterialShelf {
	if in == nil {
		return nil
	}

	var agvParkingSpaceID *string
	if in.AGVParkingSpaceID != "" {
		agvParkingSpaceID = &in.AGVParkingSpaceID
	}

	return &MaterialShelf{
		ModelID:           ModelID{ID: in.Id},
		Code:              in.Code,
		Description:       in.Description,
		ShelfType:         in.ShelfType,
		Enable:            in.Enable,
		CurrentState:      in.CurrentState,
		Remark:            in.Remark,
		ProductionLineID:  in.ProductionLineID,
		AGVParkingSpaceID: agvParkingSpaceID,
		ParkableSpaces:    PBToMaterialShelfAvailableSpaces(in.ParkableSpaces),
	}
}

func MaterialShelfsToPB(in []*MaterialShelf) []*proto.MaterialShelfInfo {
	var list []*proto.MaterialShelfInfo
	for _, f := range in {
		list = append(list, MaterialShelfToPB(f))
	}
	return list
}

func MaterialShelfToPB(in *MaterialShelf) *proto.MaterialShelfInfo {
	if in == nil {
		return nil
	}

	var agvParkingSpaceID string
	if in.AGVParkingSpaceID != nil {
		agvParkingSpaceID = *in.AGVParkingSpaceID
	}

	var agvParkingSpaceIDs []string
	for _, v := range in.ParkableSpaces {
		agvParkingSpaceIDs = append(agvParkingSpaceIDs, v.AGVParkingSpaceID)
	}

	m := &proto.MaterialShelfInfo{
		Id:                in.ID,
		Code:              in.Code,
		Description:       in.Description,
		ShelfType:         in.ShelfType,
		Enable:            in.Enable,
		CurrentState:      in.CurrentState,
		LastUpdateTime:    utils.FormatTime(in.LastUpdateTime),
		Remark:            in.Remark,
		ProductionLineID:  in.ProductionLineID,
		ProductionLine:    ProductionLineToPB(in.ProductionLine),
		AGVParkingSpaceID: agvParkingSpaceID,
		ParkableSpaceIDs:  agvParkingSpaceIDs,
		ParkableSpaces:    MaterialShelfAvailableSpacesToPB(in.ParkableSpaces),
	}

	return m
}

func PBToMaterialShelfAvailableSpaces(in []*proto.MaterialShelfAvailableSpaceInfo) []*MaterialShelfAvailableSpace {
	var result []*MaterialShelfAvailableSpace
	for _, c := range in {
		result = append(result, PBToMaterialShelfAvailableSpace(c))
	}
	return result
}

func PBToMaterialShelfAvailableSpace(in *proto.MaterialShelfAvailableSpaceInfo) *MaterialShelfAvailableSpace {
	if in == nil {
		return nil
	}
	return &MaterialShelfAvailableSpace{
		ModelID:           ModelID{ID: in.Id},
		MaterialShelfID:   in.MaterialShelfID,
		AGVParkingSpaceID: in.AGVParkingSpaceID,
	}
}

func MaterialShelfAvailableSpacesToPB(in []*MaterialShelfAvailableSpace) []*proto.MaterialShelfAvailableSpaceInfo {
	var list []*proto.MaterialShelfAvailableSpaceInfo
	for _, f := range in {
		list = append(list, MaterialShelfAvailableSpaceToPB(f))
	}
	return list
}

func MaterialShelfAvailableSpaceToPB(in *MaterialShelfAvailableSpace) *proto.MaterialShelfAvailableSpaceInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialShelfAvailableSpaceInfo{
		Id:                in.ID,
		MaterialShelfID:   in.MaterialShelfID,
		AGVParkingSpaceID: in.AGVParkingSpaceID,
		AGVParkingSpace:   AGVParkingSpaceToPB(in.AGVParkingSpace),
	}
	return m
}
