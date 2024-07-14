package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialChannelProposingRecord(m *model.MaterialChannelProposingRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialChannelProposingRecord(m *model.MaterialChannelProposingRecord) error {
	return model.DB.DB().Omit("created_at", "request_time").Save(m).Error
}

func QueryMaterialChannelProposingRecord(req *proto.QueryMaterialChannelProposingRecordRequest, resp *proto.QueryMaterialChannelProposingRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialChannelProposingRecord{}).Preload("MaterialChannel").Preload("MaterialChannel.MaterialChannelLayer").Preload("MaterialChannel.MaterialChannelLayer.ProductionStation")
	if req.Code != "" {
		db = db.Joins("JOIN material_channels ON material_channel_proposing_record.material_channel_id=material_channels.id").
			Where("material_channels.code` LIKE ? OR `material_channels.description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.RequestTime0 != "" && req.RequestTime1 != "" {
		db = db.Where("material_channel_proposing_record.request_time BETWEEN ? AND ?", req.RequestTime0, req.RequestTime1)
	}
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN material_channels ON material_channel_proposing_record.material_channel_id=material_channels.id").
			Joins("JOIN material_channel_layers ON material_channels.material_channel_layer_id=material_channel_layers.id").
			Joins("JOIN production_stations ON material_channel_layers.production_station_id=production_stations.id").
			Where("production_stations.production_line_id` = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialChannelProposingRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelProposingRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialChannelProposingRecords() (list []*model.MaterialChannelProposingRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialChannelProposingRecordByID(id string) (*model.MaterialChannelProposingRecord, error) {
	m := &model.MaterialChannelProposingRecord{}
	err := model.DB.DB().Preload("ProductionStation").Preload("MaterialChannels").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMaterialChannelProposingRecordByIDs(ids []string) ([]*model.MaterialChannelProposingRecord, error) {
	var m []*model.MaterialChannelProposingRecord
	err := model.DB.DB().Preload("ProductionStation").Preload("MaterialChannels").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialChannelProposingRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialChannelProposingRecord{}, "`id` = ?", id).Error
}
