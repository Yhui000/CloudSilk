package logic

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductModelBom(m *model.ProductModelBom) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `material_no`  = ? ", m.MaterialNo)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品型号Bom")
	}
	return m.ID, nil
}

func UpdateProductModelBom(m *model.ProductModelBom) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductModelBom := &model.ProductModelBom{}
		if err := tx.Preload(clause.Associations).Where("`id` = ?", m.ID).First(oldProductModelBom).Error; err != nil {
			return err
		}
		// omits := []string{}
		// if m.ProductModelID == "" {
		// 	omits = append(omits, "ProductModelID")
		// }
		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "`id` <> ?  and  `material_no`  = ? ", m.ID, m.MaterialNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品型号Bom")
		}

		return nil
	})
}

func QueryProductModelBom(req *proto.QueryProductModelBomRequest, resp *proto.QueryProductModelBomResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductModelBom{})
	if req.ProductCategoryID != "" {
		db = db.Joins("JOIN product_models ON product_model_boms.product_model_id=product_models.id").
			Where("product_models.product_category_id = ?", req.ProductCategoryID)
	}
	if req.ProductModelID != "" {
		db = db.Where("`product_model_id`=?", req.ProductModelID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductModelBom
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelBomsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductModelBoms() (list []*model.ProductModelBom, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductModelBomByID(id string) (*model.ProductModelBom, error) {
	m := &model.ProductModelBom{}
	err := model.DB.DB().Preload("ProductModel").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductModelBomByIDs(ids []string) ([]*model.ProductModelBom, error) {
	var m []*model.ProductModelBom
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductModelBom(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductModelBom{}, "`id` = ?", id).Error
}

func DeleteProductModelBoms(ids []string) (err error) {
	return model.DB.DB().Delete(&model.ProductModelBom{}, "`id` in ?", ids).Error
}

func UploadProductModelBom(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("产品型号BOM")
	if err != nil {
		return err
	}

	type productModelBom struct {
		ProductModel        string  //产品型号
		ItemNo              string  //项目号
		MaterialNo          string  //物料号
		MaterialDescription string  //物料描述
		RequireQTY          float32 //需求数量
		Unit                string  //单位
		ProductionProcess   string  //需求工序
	}

	var count int
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if len(rows) > 0 {
			for _, row := range rows[1:] {
				m := productModelBom{
					ProductModel:        row[0],
					ItemNo:              row[1],
					MaterialNo:          row[2],
					MaterialDescription: row[3],
					// RequireQTY:          row[4],
					Unit:              row[5],
					ProductionProcess: row[6],
				}
				requireQTY, err := strconv.ParseFloat(row[4], 32)
				if err != nil {
					return err
				}
				m.RequireQTY = float32(requireQTY)

				productModel := &model.ProductModel{}
				if err := tx.First(productModel, "`material_no` = ?", m.ProductModel).Error; err == gorm.ErrRecordNotFound {
					continue
				} else if err != nil {
					return err
				}

				productModelBom := &model.ProductModelBom{}
				if err := tx.Where("`product_model_id` = ? AND `material_no` = ?", productModel.ID, m.MaterialNo).First(productModelBom).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				productModelBom.ItemNo = "00"
				if m.ItemNo != "" {
					productModelBom.ItemNo = m.ItemNo
				}
				productModelBom.MaterialDescription = m.MaterialDescription
				productModelBom.RequireQTY = m.RequireQTY
				productModelBom.Unit = "EA"
				if m.Unit != "" {
					productModelBom.Unit = m.Unit
				}
				productModelBom.ProductionProcess = m.ProductionProcess
				if productModelBom.ID == "" {
					productModelBom.ProductModelID = productModel.ID
					productModelBom.MaterialNo = m.MaterialNo
					if err := tx.Create(productModelBom).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Omit("created_at").Save(productModelBom).Error; err != nil {
						return err
					}
				}

				count++
			}
		}
		fmt.Printf("总计%d条记录，成功%d条记录，因无法识别产品型号而忽略%d条记录", len(rows), count, len(rows)-count)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
