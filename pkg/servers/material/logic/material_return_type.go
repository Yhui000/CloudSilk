package logic

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnType(m *model.MaterialReturnType) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同物料退料类型")
	}
	return m.ID, nil
}

func UpdateMaterialReturnType(m *model.MaterialReturnType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialReturnType(req *proto.QueryMaterialReturnTypeRequest, resp *proto.QueryMaterialReturnTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnTypes() (list []*model.MaterialReturnType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnTypeByID(id string) (*model.MaterialReturnType, error) {
	m := &model.MaterialReturnType{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnTypeByIDs(ids []string) ([]*model.MaterialReturnType, error) {
	var m []*model.MaterialReturnType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnType(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnType{}, "`id` = ?", id).Error
}

func UploadMaterialReturnType(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("物料退料类型")
	if err != nil {
		return err
	}

	type materialReturnType struct {
		Code        string //退料类型
		Description string //类型描述
		Remark      string //备注
	}

	var count int
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if len(rows) > 0 {
			for i, row := range rows[1:] {
				m := materialReturnType{
					Code:        row[0],
					Description: row[1],
					Remark:      row[2],
				}
				if m.Code == "" {
					return fmt.Errorf("第%d行不良类型代号为空", i+1)
				}

				materialReturnType := &model.MaterialReturnType{}
				if err := tx.First(materialReturnType, "`code` = ?", m.Code).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				materialReturnType.Description = m.Description
				materialReturnType.Remark = m.Remark
				if materialReturnType.ID == "" {
					materialReturnType.Code = m.Code
					if err := tx.Create(materialReturnType).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Omit("created_at").Save(materialReturnType).Error; err != nil {
						return err
					}
				}

				count++
			}
		}
		fmt.Printf("总计%d条记录，成功%d条记录，因无法识别物料退料类型而忽略%d条记录", len(rows), count, len(rows)-count)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
