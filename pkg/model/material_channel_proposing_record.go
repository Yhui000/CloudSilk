package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 通道叫料记录
type MaterialChannelProposingRecord struct {
	ModelID
	MaterialChannelID string           `gorm:"size:36;comment:物料通道ID"`
	MaterialChannel   *MaterialChannel `gorm:"constraint:OnDelete:CASCADE"`
	RequestTime       time.Time        `gorm:"autoCreateTime:nano;comment:叫料时间"`
	ResponseTime      sql.NullTime     `gorm:"comment:补料时间"`
	Duration          int64            `gorm:"comment:候料时长"`
	CurrentState      string           `gorm:"size:50;comment:当前状态"`
	LastUpdateTime    time.Time        `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
}

func PBToMaterialChannelProposingRecords(in []*proto.MaterialChannelProposingRecordInfo) []*MaterialChannelProposingRecord {
	var result []*MaterialChannelProposingRecord
	for _, c := range in {
		result = append(result, PBToMaterialChannelProposingRecord(c))
	}
	return result
}

func PBToMaterialChannelProposingRecord(in *proto.MaterialChannelProposingRecordInfo) *MaterialChannelProposingRecord {
	if in == nil {
		return nil
	}
	return &MaterialChannelProposingRecord{
		ModelID:           ModelID{ID: in.Id},
		MaterialChannelID: in.MaterialChannelID,
		ResponseTime:      utils.ParseSqlNullTime(in.ResponseTime),
		Duration:          in.Duration,
		CurrentState:      in.CurrentState,
	}
}

func MaterialChannelProposingRecordsToPB(in []*MaterialChannelProposingRecord) []*proto.MaterialChannelProposingRecordInfo {
	var list []*proto.MaterialChannelProposingRecordInfo
	for _, f := range in {
		list = append(list, MaterialChannelProposingRecordToPB(f))
	}
	return list
}

func MaterialChannelProposingRecordToPB(in *MaterialChannelProposingRecord) *proto.MaterialChannelProposingRecordInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialChannelProposingRecordInfo{
		Id:                in.ID,
		MaterialChannelID: in.MaterialChannelID,
		MaterialChannel:   MaterialChannelToPB(in.MaterialChannel),
		RequestTime:       utils.FormatTime(in.RequestTime),
		ResponseTime:      utils.FormatSqlNullTime(in.ResponseTime),
		Duration:          in.Duration,
		CurrentState:      in.CurrentState,
		LastUpdateTime:    utils.FormatTime(in.LastUpdateTime),
	}
	return m
}
