package logic

import (
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateAGVParkingSpace(m *model.AGVParkingSpace) (string, error) {
	m.CurrentState = types.AGVParkingSpaceStateDisabled
	if m.EnablePark {
		m.CurrentState = types.AGVParkingSpaceStateEmptied
	}
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateAGVParkingSpace(m *model.AGVParkingSpace) error {
	if !m.EnablePark && m.MaterialShelfID != nil {
		return fmt.Errorf("禁止停靠失败，当前泊位已停靠编号为%s的货架", m.MaterialShelf.Code)
	}
	m.CurrentState = types.AGVParkingSpaceStateDisabled
	if m.EnablePark {
		m.CurrentState = types.AGVParkingSpaceStateEmptied
	}
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryAGVParkingSpace(req *proto.QueryAGVParkingSpaceRequest, resp *proto.QueryAGVParkingSpaceResponse, preload bool) {
	db := model.DB.DB().Model(&model.AGVParkingSpace{}).Preload("ProductionLine").Preload("MaterialShelf")
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.AGVParkingSpace
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVParkingSpacesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllAGVParkingSpaces() (list []*model.AGVParkingSpace, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetAGVParkingSpaceByID(id string) (*model.AGVParkingSpace, error) {
	m := &model.AGVParkingSpace{}
	err := model.DB.DB().Preload("ProductionLine").Preload("MaterialShelf").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetAGVParkingSpaceByIDs(ids []string) ([]*model.AGVParkingSpace, error) {
	var m []*model.AGVParkingSpace
	err := model.DB.DB().Preload("ProductionLine").Preload("MaterialShelf").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteAGVParkingSpace(id string) (err error) {
	return model.DB.DB().Delete(&model.AGVParkingSpace{}, "`id` = ?", id).Error
}
