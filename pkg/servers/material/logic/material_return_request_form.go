package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnRequestForm(m *model.MaterialReturnRequestForm) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialReturnRequestForm(m *model.MaterialReturnRequestForm) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryMaterialReturnRequestForm(req *proto.QueryMaterialReturnRequestFormRequest, resp *proto.QueryMaterialReturnRequestFormResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnRequestForm{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnRequestForm
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnRequestFormsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnRequestForms() (list []*model.MaterialReturnRequestForm, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnRequestFormByID(id string) (*model.MaterialReturnRequestForm, error) {
	m := &model.MaterialReturnRequestForm{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnRequestFormByIDs(ids []string) ([]*model.MaterialReturnRequestForm, error) {
	var m []*model.MaterialReturnRequestForm
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnRequestForm(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnRequestForm{}, "`id` = ?", id).Error
}
