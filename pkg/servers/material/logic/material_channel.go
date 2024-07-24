package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialChannel(m *model.MaterialChannel) (string, error) {
	m.CurrentState = types.MaterialChannelStateNormal
	if m.EnableMonitor {
		m.CurrentState = types.MaterialChannelStateUnknown
	}

	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialChannel(m *model.MaterialChannel) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialChannel{}, "`material_channel_layer_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialChannel(req *proto.QueryMaterialChannelRequest, resp *proto.QueryMaterialChannelResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialChannel{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialChannel
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialChannels() (list []*model.MaterialChannel, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialChannelByID(id string) (*model.MaterialChannel, error) {
	m := &model.MaterialChannel{}
	err := model.DB.DB().Preload("ProductionStation").Preload("MaterialChannels").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMaterialChannels(req *proto.GetMaterialChannelRequest) ([]*model.MaterialChannel, error) {
	var m []*model.MaterialChannel
	err := model.DB.DB().
		Preload("MaterialChannel").
		Preload("MaterialInfo").
		Joins("JOIN material_channel_layers ON material_channels.material_channel_layer_id=material_channel_layers.id").
		Where("material_channel_layers.production_station_id=?", req.ProductionStationID).Find(&m).Error

	return m, err
}

func GetMaterialChannelByIDs(ids []string) ([]*model.MaterialChannel, error) {
	var m []*model.MaterialChannel
	err := model.DB.DB().Preload("ProductionStation").Preload("MaterialChannels").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialChannel(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialChannel{}, "`id` = ?", id).Error
}
