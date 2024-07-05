package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// AGV停靠泊位
type AGVParkingSpace struct {
	ModelID
	Code             string          `gorm:"size:50;comment:代号"`
	Description      string          `gorm:"size:500;comment:描述"`
	Identifier       string          `gorm:"size:50;comment:识别码"`
	SpaceType        int32           `gorm:"comment:泊位类型"`
	ShelfType        int32           `gorm:"comment:货架类型"`
	EnablePark       bool            `gorm:"comment:允许停靠"`
	EnableDispatch   bool            `gorm:"comment:允许签派"`
	CurrentState     string          `gorm:"size:50;comment:当前状态"`
	LastUpdateTime   time.Time       `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark           string          `gorm:"size:500;comment:备注"`
	ProductionLineID string          `gorm:"size:36;comment:隶属产线ID"`
	ProductionLine   *ProductionLine `gorm:"constraint:OnDelete:CASCADE"`
	MaterialShelfID  *string         `gorm:"size:36;comment:当前货架ID"`
	MaterialShelf    *MaterialShelf  `gorm:"constraint:OnDelete:SET NULL"`
}

func PBToAGVParkingSpaces(in []*proto.AGVParkingSpaceInfo) []*AGVParkingSpace {
	var result []*AGVParkingSpace
	for _, c := range in {
		result = append(result, PBToAGVParkingSpace(c))
	}
	return result
}

func PBToAGVParkingSpace(in *proto.AGVParkingSpaceInfo) *AGVParkingSpace {
	if in == nil {
		return nil
	}

	var materialShelfID *string
	if in.MaterialShelfID != "" {
		materialShelfID = &in.MaterialShelfID
	}

	return &AGVParkingSpace{
		ModelID:          ModelID{ID: in.Id},
		Code:             in.Code,
		Description:      in.Description,
		Identifier:       in.Identifier,
		SpaceType:        in.SpaceType,
		ShelfType:        in.ShelfType,
		EnablePark:       in.EnablePark,
		EnableDispatch:   in.EnableDispatch,
		CurrentState:     in.CurrentState,
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		MaterialShelfID:  materialShelfID,
	}
}

func AGVParkingSpacesToPB(in []*AGVParkingSpace) []*proto.AGVParkingSpaceInfo {
	var list []*proto.AGVParkingSpaceInfo
	for _, f := range in {
		list = append(list, AGVParkingSpaceToPB(f))
	}
	return list
}

func AGVParkingSpaceToPB(in *AGVParkingSpace) *proto.AGVParkingSpaceInfo {
	if in == nil {
		return nil
	}

	var materialShelfID string
	if in.MaterialShelfID != nil {
		materialShelfID = *in.MaterialShelfID
	}

	m := &proto.AGVParkingSpaceInfo{
		Id:               in.ID,
		Code:             in.Code,
		Description:      in.Description,
		Identifier:       in.Identifier,
		SpaceType:        in.SpaceType,
		ShelfType:        in.ShelfType,
		EnablePark:       in.EnablePark,
		EnableDispatch:   in.EnableDispatch,
		CurrentState:     in.CurrentState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		ProductionLine:   ProductionLineToPB(in.ProductionLine),
		MaterialShelfID:  materialShelfID,
		MaterialShelf:    MaterialShelfToPB(in.MaterialShelf),
	}
	return m
}
