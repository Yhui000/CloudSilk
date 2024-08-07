package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料退料记录
type MaterialReturnRequestForm struct {
	ModelID
	FormNo                   string                  `gorm:"size:50;comment:申请单号"`
	CreateTime               time.Time               `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID             string                  `gorm:"size:36;comment:创建人员ID"`
	MaterialSupplierID       *string                 `gorm:"size:36;comment:物料供应商ID"`
	MaterialSupplier         *MaterialSupplier       `gorm:"constraint:OnDelete:SET NULL"` //物料供应商
	MaterialInfoID           string                  `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo             *MaterialInfo           `gorm:"constraint:OnDelete:CASCADE"` //物料信息
	ProductionStationID      string                  `gorm:"size:36;comment:发现工站ID"`
	ProductionStation        *ProductionStation      `gorm:"constraint:OnDelete:CASCADE"` //发现工站
	MaterialTraceNo          string                  `gorm:"size:100;comment:物料追溯号"`
	Quantity                 float32                 `gorm:"comment:退料数量"`
	ReturnID                 string                  `gorm:"size:36;comment:退料来源ID"`
	ReturnSource             string                  `gorm:"size:50;comment:退料来源"`
	ReturnReason             string                  `gorm:"size:500;comment:退料原因"`
	MaterialReturnTypeID     *string                 `gorm:"size:36;comment:退料类型ID"`
	MaterialReturnType       *MaterialReturnType     `gorm:"constraint:OnDelete:SET NULL"` //退料类型
	MaterialReturnCauseID    *string                 `gorm:"size:36;comment:退料原因ID"`
	MaterialReturnCause      *MaterialReturnCause    `gorm:"constraint:OnDelete:SET NULL"` //退料原因
	MaterialReturnSolutionID *string                 `gorm:"size:36;comment:处理方案ID"`
	MaterialReturnSolution   *MaterialReturnSolution `gorm:"constraint:OnDelete:SET NULL"` //处理方案
	ReturnBrief              string                  `gorm:"size:500;comment:退料简述"`
	CheckTime                sql.NullTime            `gorm:"comment:复核时间"`
	CheckUserID              *string                 `gorm:"size:36;comment:复核人员ID"`
	HandleMethod             string                  `gorm:"size:50;comment:旧件处理"`
	CurrentState             int32                   `gorm:"comment:当前状态"`
	Remark                   string                  `gorm:"size:500;comment:备注"`
}

func PBToMaterialReturnRequestForms(in []*proto.MaterialReturnRequestFormInfo) []*MaterialReturnRequestForm {
	var result []*MaterialReturnRequestForm
	for _, c := range in {
		result = append(result, PBToMaterialReturnRequestForm(c))
	}
	return result
}

func PBToMaterialReturnRequestForm(in *proto.MaterialReturnRequestFormInfo) *MaterialReturnRequestForm {
	if in == nil {
		return nil
	}

	var materialSupplierID, materialReturnTypeID, materialReturnCauseID, materialReturnSolutionID, checkUserID *string
	if in.MaterialSupplierID != "" {
		materialSupplierID = &in.MaterialSupplierID
	}
	if in.MaterialReturnTypeID != "" {
		materialReturnTypeID = &in.MaterialReturnTypeID
	}
	if in.MaterialReturnCauseID != "" {
		materialReturnCauseID = &in.MaterialReturnCauseID
	}
	if in.MaterialReturnSolutionID != "" {
		materialReturnSolutionID = &in.MaterialReturnSolutionID
	}
	if in.CheckUserID != "" {
		checkUserID = &in.CheckUserID
	}

	return &MaterialReturnRequestForm{
		ModelID:                  ModelID{ID: in.Id},
		FormNo:                   in.FormNo,
		CreateUserID:             in.CreateUserID,
		MaterialSupplierID:       materialSupplierID,
		MaterialInfoID:           in.MaterialInfoID,
		ProductionStationID:      in.ProductionStationID,
		MaterialTraceNo:          in.MaterialTraceNo,
		Quantity:                 in.Quantity,
		ReturnID:                 in.ReturnID,
		ReturnSource:             in.ReturnSource,
		ReturnReason:             in.ReturnReason,
		MaterialReturnTypeID:     materialReturnTypeID,
		MaterialReturnCauseID:    materialReturnCauseID,
		MaterialReturnSolutionID: materialReturnSolutionID,
		ReturnBrief:              in.ReturnBrief,
		CheckUserID:              checkUserID,
		HandleMethod:             in.HandleMethod,
		CurrentState:             in.CurrentState,
		Remark:                   in.Remark,
	}
}

func MaterialReturnRequestFormsToPB(in []*MaterialReturnRequestForm) []*proto.MaterialReturnRequestFormInfo {
	var list []*proto.MaterialReturnRequestFormInfo
	for _, f := range in {
		list = append(list, MaterialReturnRequestFormToPB(f))
	}
	return list
}

func MaterialReturnRequestFormToPB(in *MaterialReturnRequestForm) *proto.MaterialReturnRequestFormInfo {
	if in == nil {
		return nil
	}

	var materialSupplierID, materialReturnTypeID, materialReturnCauseID, materialReturnSolutionID, checkUserID string
	if in.MaterialSupplierID != nil {
		materialSupplierID = *in.MaterialSupplierID
	}
	if in.MaterialReturnTypeID != nil {
		materialReturnTypeID = *in.MaterialReturnTypeID
	}
	if in.MaterialReturnCauseID != nil {
		materialReturnCauseID = *in.MaterialReturnCauseID
	}
	if in.MaterialReturnSolutionID != nil {
		materialReturnSolutionID = *in.MaterialReturnSolutionID
	}
	if in.CheckUserID != nil {
		checkUserID = *in.CheckUserID
	}

	m := &proto.MaterialReturnRequestFormInfo{
		Id:                       in.ID,
		FormNo:                   in.FormNo,
		CreateTime:               utils.FormatTime(in.CreateTime),
		CreateUserID:             in.CreateUserID,
		MaterialSupplierID:       materialSupplierID,
		MaterialSupplier:         MaterialSupplierToPB(in.MaterialSupplier),
		MaterialInfoID:           in.MaterialInfoID,
		MaterialInfo:             MaterialInfoToPB(in.MaterialInfo),
		ProductionStationID:      in.ProductionStationID,
		ProductionStation:        ProductionStationToPB(in.ProductionStation),
		MaterialTraceNo:          in.MaterialTraceNo,
		Quantity:                 in.Quantity,
		ReturnID:                 in.ReturnID,
		ReturnSource:             in.ReturnSource,
		ReturnReason:             in.ReturnReason,
		MaterialReturnTypeID:     materialReturnTypeID,
		MaterialReturnType:       MaterialReturnTypeToPB(in.MaterialReturnType),
		MaterialReturnCauseID:    materialReturnCauseID,
		MaterialReturnCause:      MaterialReturnCauseToPB(in.MaterialReturnCause),
		MaterialReturnSolutionID: materialReturnSolutionID,
		MaterialReturnSolution:   MaterialReturnSolutionToPB(in.MaterialReturnSolution),
		ReturnBrief:              in.ReturnBrief,
		CheckTime:                utils.FormatSqlNullTime(in.CheckTime),
		CheckUserID:              checkUserID,
		HandleMethod:             in.HandleMethod,
		CurrentState:             in.CurrentState,
		Remark:                   in.Remark,
	}
	return m
}
