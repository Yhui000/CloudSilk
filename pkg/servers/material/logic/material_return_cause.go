package logic

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnCause(m *model.MaterialReturnCause) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同物料退料原因")
	}
	return m.ID, nil
}

func UpdateMaterialReturnCause(m *model.MaterialReturnCause) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialReturnCauseAvailableCategory{}, "`material_return_cause_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Delete(&model.MaterialReturnCauseAvailableType{}, "`material_return_cause_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialReturnCause(req *proto.QueryMaterialReturnCauseRequest, resp *proto.QueryMaterialReturnCauseResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnCause{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnCause
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnCausesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnCauses() (list []*model.MaterialReturnCause, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnCauseByID(id string) (*model.MaterialReturnCause, error) {
	m := &model.MaterialReturnCause{}
	err := model.DB.DB().
		Preload("MaterialCategories").Preload("MaterialCategories.MaterialCategory").
		Preload("MaterialReturnTypes").Preload("MaterialReturnTypes.MaterialReturnType").
		Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnCauseByIDs(ids []string) ([]*model.MaterialReturnCause, error) {
	var m []*model.MaterialReturnCause
	err := model.DB.DB().
		Preload("MaterialCategories").Preload("MaterialCategories.MaterialCategory").
		Preload("MaterialReturnTypes").Preload("MaterialReturnTypes.MaterialReturnType").
		Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnCause(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnCause{}, "`id` = ?", id).Error
}

func UploadMaterialReturnCause(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("物料退料类型")
	if err != nil {
		return err
	}

	type materialReturnCause struct {
		Code               string //原因代号
		Description        string //原因描述
		Remark             string //备注
		MaterialReturnType string //归属类型
		MaterialCategory   string //归属物料类别
	}

	var count int
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if len(rows) > 0 {
			for i, row := range rows[1:] {
				m := materialReturnCause{
					Code:               row[0],
					Description:        row[1],
					Remark:             row[2],
					MaterialReturnType: row[3],
					MaterialCategory:   row[4],
				}
				if m.MaterialCategory == "" {
					return fmt.Errorf("第%d条物料类别代号为空", i+1)
				}
				if m.Code == "" {
					return fmt.Errorf("第%d条数据原因代号为空", i+1)
				}

				materialCategories := []*model.MaterialReturnCauseAvailableCategory{}
				materialReturnTypes := []*model.MaterialReturnCauseAvailableType{}

				for _, code := range strings.Split(m.MaterialReturnType, ",") {
					materialReturnType := &model.MaterialReturnType{}
					if err := tx.First(materialReturnType, "`code` = ?", code).Error; err == gorm.ErrRecordNotFound {
						return fmt.Errorf("第%d条数据归属类型代号存在错误", i+1)
					} else if err != nil {
						return err
					}
					materialCategories = append(materialCategories, &model.MaterialReturnCauseAvailableCategory{
						MaterialCategoryID: materialReturnType.ID,
					})
				}

				for _, code := range strings.Split(m.MaterialCategory, ",") {
					materialCategory := &model.MaterialCategory{}
					if err := tx.First(materialCategory, "`code` = ?", code).Error; err == gorm.ErrRecordNotFound {
						return fmt.Errorf("第%d条数据归属物料类别代号存在错误", i+1)
					} else if err != nil {
						return err
					}
					materialReturnTypes = append(materialReturnTypes, &model.MaterialReturnCauseAvailableType{
						MaterialReturnTypeID: materialCategory.ID,
					})
				}

				materialReturnCause := &model.MaterialReturnCause{}
				if err := tx.First(materialReturnCause, "`code` = ?", m.Code).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				materialReturnCause.Description = m.Description
				materialReturnCause.Remark = m.Remark
				materialReturnCause.MaterialCategories = materialCategories
				materialReturnCause.MaterialReturnTypes = materialReturnTypes
				if materialReturnCause.ID == "" {
					materialReturnCause.Code = m.Code
					if err := tx.Create(materialReturnCause).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Delete(model.MaterialReturnCauseAvailableCategory{}, "`material_return_cause_id` = ?", materialReturnCause.ID).Error; err != nil {
						return err
					}
					if err := tx.Delete(model.MaterialReturnCauseAvailableType{}, "`material_return_cause_id` = ?", materialReturnCause.ID).Error; err != nil {
						return err
					}
					if err := tx.Omit("created_at").Save(materialReturnCause).Error; err != nil {
						return err
					}
				}

				count++
			}
		}
		fmt.Printf("总计%d条记录，成功%d条记录，因无法识别物料退料原因而忽略%d条记录", len(rows), count, len(rows)-count)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
