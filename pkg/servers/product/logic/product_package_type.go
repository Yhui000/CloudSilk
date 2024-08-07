package logic

import (
	"errors"
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductPackageType(m *model.ProductPackageType) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, "`code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同包装类型")
	}
	return m.ID, nil
}

func UpdateProductPackageType(m *model.ProductPackageType) error {
	// omits := []string{"created_at"}
	// if m.LabelTypeID == "" {
	// 	omits = append(omits, "LabelTypeID")
	// }
	// if m.SystemEventID == "" {
	// 	omits = append(omits, "SystemEventID")
	// }
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{"created_at"}, "`id` != ? and  `code`  = ? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同包装类型")
	}

	return nil
}

func QueryProductPackageType(req *proto.QueryProductPackageTypeRequest, resp *proto.QueryProductPackageTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductPackageType{}).Preload("SystemEvent").Preload("LabelType")
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductPackageType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductPackageTypes() (list []*model.ProductPackageType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductPackageTypeByID(id string) (*model.ProductPackageType, error) {
	m := &model.ProductPackageType{}
	err := model.DB.DB().Preload("LabelType").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductPackageTypeByIDs(ids []string) ([]*model.ProductPackageType, error) {
	var m []*model.ProductPackageType
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductPackageType(id string) (err error) {
	var count int64
	if err = model.DB.DB().Where(&model.ProductPackage{ProductPackageTypeID: &id}).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		return fmt.Errorf("数据冲突，请情况此类型下属的产品包装")
	}

	return model.DB.DB().Delete(&model.ProductPackageType{}, "`id` = ?", id).Error
}

func UpdateProductPackageTypeAll(m *model.ProductPackageType) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, true, []string{}, "`id` <> ? and  `code`  = ? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同包装类型")
	}

	return nil
}
