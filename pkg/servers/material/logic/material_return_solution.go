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

func CreateMaterialReturnSolution(m *model.MaterialReturnSolution) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同物料退料方案")
	}
	return m.ID, nil
}

func UpdateMaterialReturnSolution(m *model.MaterialReturnSolution) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialReturnSolutionAvailableCause{}, "`material_return_solution_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialReturnSolution(req *proto.QueryMaterialReturnSolutionRequest, resp *proto.QueryMaterialReturnSolutionResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnSolution{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnSolution
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnSolutionsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnSolutions() (list []*model.MaterialReturnSolution, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnSolutionByID(id string) (*model.MaterialReturnSolution, error) {
	m := &model.MaterialReturnSolution{}
	err := model.DB.DB().Preload("MaterialReturnCauses").Preload("MaterialReturnCauses.MaterialReturnCause").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnSolutionByIDs(ids []string) ([]*model.MaterialReturnSolution, error) {
	var m []*model.MaterialReturnSolution
	err := model.DB.DB().Preload("MaterialReturnCauses").Preload("MaterialReturnCauses.MaterialReturnCause").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnSolution(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnSolution{}, "`id` = ?", id).Error
}

func UploadMaterialReturnSolution(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("物料退料类型")
	if err != nil {
		return err
	}

	type materialReturnSolution struct {
		Code                string //方案代号
		Description         string //方案描述
		Remark              string //备注
		MaterialReturnCause string //归属原因
	}

	var count int
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if len(rows) > 0 {
			for i, row := range rows[1:] {
				m := materialReturnSolution{
					Code:                row[0],
					Description:         row[1],
					Remark:              row[2],
					MaterialReturnCause: row[3],
				}
				if m.Code == "" {
					return fmt.Errorf("第%d行数据方案代号为空", i+1)
				}

				materialReturnCauses := []*model.MaterialReturnSolutionAvailableCause{}

				for _, code := range strings.Split(m.MaterialReturnCause, ",") {
					materialReturnCause := &model.MaterialReturnCause{}
					if err := tx.First(materialReturnCause, "`code` = ?", code).Error; err == gorm.ErrRecordNotFound {
						return fmt.Errorf("第%d条数据归属原因代号存在错误", i+1)
					} else if err != nil {
						return err
					}
					materialReturnCauses = append(materialReturnCauses, &model.MaterialReturnSolutionAvailableCause{
						MaterialReturnCauseID: materialReturnCause.ID,
					})
				}

				materialReturnSolution := &model.MaterialReturnSolution{}
				if err := tx.First(materialReturnSolution, "`code` = ?", m.Code).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				materialReturnSolution.Description = m.Description
				materialReturnSolution.Remark = m.Remark
				materialReturnSolution.MaterialReturnCauses = materialReturnCauses
				if materialReturnSolution.ID == "" {
					materialReturnSolution.Code = m.Code
					if err := tx.Create(materialReturnSolution).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Delete(model.MaterialReturnSolutionAvailableCause{}, "`material_return_solution_id` = ?", materialReturnSolution.ID).Error; err != nil {
						return err
					}
					if err := tx.Omit("created_at").Save(materialReturnSolution).Error; err != nil {
						return err
					}
				}

				count++
			}
		}
		fmt.Printf("总计%d条记录，成功%d条记录，因无法识别物料退料方案而忽略%d条记录", len(rows), count, len(rows)-count)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
