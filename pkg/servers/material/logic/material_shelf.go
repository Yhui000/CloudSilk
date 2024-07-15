package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialShelf(m *model.MaterialShelf) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialShelf(m *model.MaterialShelf) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialShelfAvailableSpace{}, "`material_shelf_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialShelf(req *proto.QueryMaterialShelfRequest, resp *proto.QueryMaterialShelfResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialShelf{}).Preload("ProductionLine").Preload("AGVParkingSpace")
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialShelf
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialShelfsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialShelfs() (list []*model.MaterialShelf, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialShelfByID(id string) (*model.MaterialShelf, error) {
	m := &model.MaterialShelf{}
	err := model.DB.DB().Preload("AGVParkingSpace").Preload("ParkableSpaces").Preload("ParkableSpaces.AGVParkingSpace").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialShelfByIDs(ids []string) ([]*model.MaterialShelf, error) {
	var m []*model.MaterialShelf
	err := model.DB.DB().Preload("AGVParkingSpace").Preload("ParkableSpaces").Preload("ParkableSpaces.AGVParkingSpace").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialShelf(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialShelf{}, "`id` = ?", id).Error
}
