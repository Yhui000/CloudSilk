package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateDataMapping(m *model.DataMapping) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, "`group` = ? AND `code` = ?", m.Group, m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同数据词条")
	}
	return m.ID, nil
}

func UpdateDataMapping(m *model.DataMapping) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryDataMapping(req *proto.QueryDataMappingRequest, resp *proto.QueryDataMappingResponse, preload bool) {
	db := model.DB.DB().Model(&model.DataMapping{})
	if req.Group != "" {
		db = db.Where("`group` LIKE ?", "%"+req.Group+"%")
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.DataMapping
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.DataMappingsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllDataMappings() (list []*model.DataMapping, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetDataMappingByID(id string) (*model.DataMapping, error) {
	m := &model.DataMapping{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetDataMappingByIDs(ids []string) ([]*model.DataMapping, error) {
	var m []*model.DataMapping
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteDataMapping(id string) (err error) {
	return model.DB.DB().Delete(&model.DataMapping{}, "`id` = ?", id).Error
}
