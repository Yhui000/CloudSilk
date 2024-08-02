package logic

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"regexp"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductModel(m *model.ProductModel) (string, error) {
	var count int64
	err := model.DB.DB().Model(m).Where(" material_no  = ? ", m.MaterialNo).Count(&count).Error
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("存在相同产品型号")
	}

	//新增型号配置关联同产品类别下的产品类别特性
	var productCategoryAttributes []*model.ProductCategoryAttribute
	if err := model.DB.DB().Where("`product_category_id` = ?", m.ProductCategoryID).Find(&productCategoryAttributes).Error; err != nil {
		return "", err
	}
	var productModelAttributeValues []*model.ProductModelAttributeValue
	for _, _productCategoryAttribute := range productCategoryAttributes {
		productModelAttributeValues = append(productModelAttributeValues, &model.ProductModelAttributeValue{
			ProductAttributeID: _productCategoryAttribute.ProductAttributeID,
			AssignedValue:      _productCategoryAttribute.DefaultValue,
		})
	}
	m.ProductModelAttributeValues = productModelAttributeValues

	err = model.DB.DB().Create(m).Error

	return m.ID, err
}

func OnDeleteProductModelBoms(tx *gorm.DB, old, m *model.ProductModel) error {
	var deleteIDs []string
	for _, oldObj := range old.ProductModelBoms {
		flag := false
		for _, newObj := range m.ProductModelBoms {
			if newObj.ID == oldObj.ID {
				flag = true
			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		if err := tx.Delete(&model.ProductModelBom{}, "id in ?", deleteIDs).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProductModel(m *model.ProductModel) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductModel := &model.ProductModel{}
		if err := tx.Preload("ProductModelAttributeValues").Preload("ProductModelBoms").Preload(clause.Associations).Where("id = ?", m.ID).First(oldProductModel).Error; err != nil {
			return err
		}
		if err := tx.Delete(model.ProductModelAttributeValue{}, "product_model_id = ?", m.ID).Error; err != nil {
			return err
		}

		if err := OnDeleteProductModelBoms(tx, oldProductModel, m); err != nil {
			return err
		}

		// omits := []string{"created_at"}
		// if m.ProductCategoryID == "" {
		// 	omits = append(omits, "ProductCategoryID")
		// }
		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  AND  material_no  = ? ", m.ID, m.MaterialNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品型号")
		}
		return nil
	})
}

func QueryProductModel(req *proto.QueryProductModelRequest, resp *proto.QueryProductModelResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductModel{})
	if req.ProductCategoryID != "" {
		db = db.Where("`product_category_id` = ?", req.ProductCategoryID)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `material_no` LIKE ? OR `material_description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.IsPrefabricated {
		db = db.Where("`is_prefabricated` = ?", req.IsPrefabricated)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductModel
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductModels() (list []*model.ProductModel, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductModelByID(id string) (*model.ProductModel, error) {
	m := &model.ProductModel{}
	err := model.DB.DB().Preload("ProductModelAttributeValues.ProductAttribute.ProductCategoryAttributes").
		Preload("ProductModelBoms").Preload(clause.Associations).Where("id = ?", id).First(m).Error

	for _, productModelAttributeValue := range m.ProductModelAttributeValues {
		for _, productCategoryAttribute := range productModelAttributeValue.ProductAttribute.ProductCategoryAttributes {
			if productCategoryAttribute.ProductCategoryID == m.ProductCategoryID {
				productModelAttributeValue.AllowNullOrBlank = productCategoryAttribute.AllowNullOrBlank
			}
		}
	}
	return m, err
}

func GetProductModelByIDs(ids []string) ([]*model.ProductModel, error) {
	var m []*model.ProductModel
	err := model.DB.DB().Preload("ProductModelAttributeValues.ProductAttribute.ProductCategoryAttributes").
		Preload("ProductModelBoms").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductModel(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductModel{}, "`id` = ?", id).Error
}

func ParamProductModelByID(id string) (err error) {
	var productModel model.ProductModel
	if err := model.DB.DB().Preload("ProductModelAttributeValues").Preload("ProductModelBoms").Preload(clause.Associations).First(&productModel, id).Error; err != nil {
		return err
	}
	//物料描述
	materialDescription := productModel.MaterialDescription

	var productCategory model.ProductCategory
	if err := model.DB.DB().First(&productCategory, productModel.ProductCategoryID).Error; err != nil {
		return err
	}
	//正则特性表达式
	attributeExpression := productCategory.AttributeExpression
	if attributeExpression == "" {
		return fmt.Errorf("缺少此产品类别的特性表达式配置，无法解析")
	}

	//TODO 原正则在go语言报错，需替换字符
	attributeExpression = strings.ReplaceAll(attributeExpression, "(?<", "(?P<")
	pattern := regexp.MustCompile(attributeExpression)

	//正则匹配
	if !pattern.MatchString(materialDescription) {
		return fmt.Errorf("此产品的型号物料描述不匹配此产品类别的特性表达式的格式，解析失败")
	}

	matches := pattern.FindStringSubmatch(materialDescription)

	productModelAttributeValues := []*model.ProductModelAttributeValue{}
	for i, code := range pattern.SubexpNames() {
		if strings.TrimSpace(code) != "" {
			var productCategoryAttribute model.ProductCategoryAttribute
			if err := model.DB.DB().Joins("JOIN product_attributes ON product_category_attributes.product_attribute_id=product_attributes.id").
				Where("`product_category_id` = ? and `code` = ?", productModel.ProductCategoryID, code).First(&productCategoryAttribute).Error; err != nil {
				continue
			}

			var productModelAttributeValue model.ProductModelAttributeValue
			if err := model.DB.DB().Where("`product_model_id` = ? and `product_attribute_id` = ?", productModel.ID, productCategoryAttribute.ProductAttributeID).First(&productModelAttributeValue).Error; err == gorm.ErrRecordNotFound {
				productModelAttributeValue = model.ProductModelAttributeValue{
					ProductAttributeID: productCategoryAttribute.ProductAttributeID,
					ProductModelID:     productModel.ID,
					AssignedValue:      productCategoryAttribute.DefaultValue,
				}
			}

			if len(strings.TrimSpace(matches[i])) > 0 {
				productModelAttributeValue.AssignedValue = strings.TrimSpace(matches[i])
			}

			productModelAttributeValues = append(productModelAttributeValues, &productModelAttributeValue)
		}
	}

	productModel.ProductModelAttributeValues = productModelAttributeValues
	if err := UpdateProductModel(&productModel); err != nil {
		return err
	}

	return nil
}

func UploadProductModel(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("产品型号")
	if err != nil {
		return err
	}

	type productModel struct {
		ProductCategory     string //类别
		Code                string //型号
		MaterialNo          string //物料号
		MaterialDescription string //物料描述
		Identifier          string //识别码
		IsAdvanced          bool   //需前置生产
	}

	type productModelAttribute struct {
		MaterialNo    string //产品物料号
		Code          string //特性代号
		AssignedValue string //特性值
	}

	var productModelCount int
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if len(rows) > 0 {
			for _, row := range rows[1:] {
				m := productModel{
					ProductCategory:     row[0],
					Code:                row[1],
					MaterialNo:          row[2],
					MaterialDescription: row[3],
					Identifier:          row[4],
				}

				productModel := &model.ProductModel{}
				if err := tx.First(productModel, "`material_no` = ?", m.MaterialNo).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if productModel.ID == "" {
					//查找产品类别
					productCategory := &model.ProductCategory{}
					if err := tx.Preload("ProductCategoryAttributes").Where("`code` = ?", m.ProductCategory).First(productCategory).Error; err == gorm.ErrRecordNotFound {
						productCategorys := []*model.ProductCategory{}
						if err := tx.Preload("ProductCategoryAttributes").Where("`attribute_expression` IS NOT NULL", m.ProductCategory).Find(&productCategorys).Error; err != nil {
							return err
						}
						for _, v := range productCategorys {
							if regexp.MustCompile(strings.ReplaceAll(v.AttributeExpression, "(?<", "(?P<")).MatchString(m.MaterialDescription) {
								productCategory = v
								break
							}
						}
					} else if err != nil {
						return err
					}

					if productCategory.ID == "" {
						continue
					}

					//根据当前产品类别特性，创建产品型号特性值
					productModelAttributeValues := make([]*model.ProductModelAttributeValue, len(productCategory.ProductCategoryAttributes))
					for i, v := range productCategory.ProductCategoryAttributes {
						productModelAttributeValues[i] = &model.ProductModelAttributeValue{
							ProductAttributeID: v.ProductAttributeID,
							AssignedValue:      v.DefaultValue,
						}
					}
					//创建产品型号
					if err := tx.Create(&model.ProductModel{
						ProductCategoryID:           productCategory.ID,
						MaterialNo:                  m.MaterialNo,
						Code:                        m.Code,
						ProductModelAttributeValues: productModelAttributeValues,
					}).Error; err != nil {
						return err
					}

					//根据当前产品类别特性表达式，自动解析
					attributeExpression := strings.ReplaceAll(productCategory.AttributeExpression, "(?<", "(?P<")
					if attributeExpression != "" {
						matches := regexp.MustCompile(productCategory.AttributeExpression).FindStringSubmatch(m.MaterialDescription)
						groups := regexp.MustCompile(productCategory.AttributeExpression).SubexpNames()
						for i, code := range groups {
							productAttribute := &model.ProductAttribute{}
							if err := model.DB.DB().First(productAttribute, "`code` = ?", code).Error; err == gorm.ErrRecordNotFound {
								continue
							} else if err != nil {
								return err
							}

							productModelAttributeValue := &model.ProductModelAttributeValue{}
							if err := model.DB.DB().Where(model.ProductModelAttributeValue{ProductModelID: productModel.ID, ProductAttributeID: productAttribute.ID}).First(productModelAttributeValue).Error; err == gorm.ErrRecordNotFound {
								if err := tx.Create(&model.ProductModelAttributeValue{
									ProductAttributeID: productAttribute.ID,
									ProductModelID:     productModel.ID,
								}).Error; err != nil {
									return err
								}
							} else if err != nil {
								return err
							}

							assigndValue := ""
							if matches[i] != "" {
								assigndValue = strings.TrimSpace(matches[i])
							} else if productModelAttributeValue.AssignedValue != "" {
								assigndValue = productModelAttributeValue.AssignedValue
							} else {
								assigndValue = productAttribute.DefaultValue
							}
							productModelAttributeValue.AssignedValue = assigndValue

							rv := reflect.ValueOf(productModel).Elem()
							for i := 0; i < rv.NumField(); i++ {
								field := rv.Type().Field(i)
								if field.Name == code || strings.HasPrefix(field.Tag.Get("binding"), code+":") {
									if rv.Field(i).CanSet() {
										value := reflect.ValueOf(assigndValue)
										if value.Type().AssignableTo(field.Type) {
											rv.Field(i).Set(value)
										}

									}
								}
							}
						}
					}
				}

				//更新识别码和物料描述
				productModel.Identifier = m.Identifier
				productModel.MaterialDescription = m.MaterialDescription

				productModelCount++
			}

			fmt.Printf("总计%d条记录，成功%d条记录，因无法识别产品类别而忽略%d条记录", len(rows), productModelCount, len(rows)-productModelCount)
			return nil
		}

		var productModelAttributeCount int
		rows, err = f.GetRows("产品型号特性")
		if err != nil {
			return err
		}
		if len(rows) > 0 {
			for _, td := range rows {
				m := productModelAttribute{
					MaterialNo:    td[0],
					Code:          td[1],
					AssignedValue: td[2],
				}

				productModel := &model.ProductModel{}
				if err := tx.First(productModel, "`material_no` = ?", m.MaterialNo).Error; err == gorm.ErrRecordNotFound {
					continue
				} else if err != nil {
					return err
				}

				productAttribute := &model.ProductAttribute{}
				if err := tx.Joins("JOIN product_attribute_identifiers ON product_attributes.id=product_attribute_identifiers.product_attribute_id").
					Where("product_attributes.code = ?", m.Code).
					Or("product_attribute_identifiers.identifier LIKE ?", "%"+m.Code+"%").First(productAttribute).Error; err == gorm.ErrRecordNotFound {
					continue
				} else if err != nil {
					return err
				}

				productModelAttributeValue := &model.ProductModelAttributeValue{}
				if err := tx.Where(model.ProductModelAttributeValue{ProductAttributeID: productAttribute.ID, ProductModelID: productModel.ID}).First(productModelAttributeValue).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}
				if productModelAttributeValue.ID == "" {
					productModelAttributeValue = &model.ProductModelAttributeValue{
						ProductAttributeID: productAttribute.ID,
						ProductModelID:     productModel.ID,
					}
					if err := tx.Create(productModelAttributeValue).Error; err != nil {
						return err
					}
				}

				if err := tx.Model(productModelAttributeValue).Update("assigned_value", m.AssignedValue).Error; err != nil {
					return err
				}
				productModelAttributeCount++
			}

			fmt.Printf("总计%d条记录，成功%d条记录，，因无法识别产品物料号或特性代号%d条记录", len(rows), productModelAttributeCount, len(rows)-productModelAttributeCount)
			return nil
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func CancelProductModel(ids []string) error {
	return model.DB.DB().Model(model.ProductModel{}).Where("`id` IN ?", ids).Update("is_authorized", false).Error
}

func AuthorizeProductModel(ids []string) error {
	return model.DB.DB().Model(model.ProductModel{}).Where("`id` IN ?", ids).Update("is_authorized", true).Error
}

func SynchronizeProductModel(productCategoryID string) error {
	productCategory := &model.ProductCategory{}
	if err := model.DB.DB().First(productCategory, "`id` = ?", productCategoryID).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的产品类别ID")
	} else if err != nil {
		return err
	}

	taskCode := "RS205"
	if err := model.DB.DB().Model(model.SystemParamsConfig{}).Where("`code` = ?", types.SystemConfigKeyProductModelSynchronizeTask).Pluck("value", &taskCode).Error; err != nil {
		return err
	}

	remoteServiceTask := &model.RemoteServiceTask{}
	if err := model.DB.DB().First(remoteServiceTask, "`code` = ?", taskCode).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无法找到标签打印的远程任务(%s)", taskCode)
	} else if err != nil {
		return err
	}

	if err := model.DB.DB().Create(&model.RemoteServiceTaskQueue{
		RemoteServiceTaskID: remoteServiceTask.ID,
		CurrentState:        types.RemoteServiceTaskQueueStateWaitExecute,
		TaskNo:              uuid.NewString(),
		RemoteServiceTaskQueueParameters: []*model.RemoteServiceTaskQueueParameter{
			{
				DataType:    types.DataTypeString,
				Name:        "ProductCategory",
				Value:       productCategory.Code,
				Description: "标签代号",
			},
		},
	}).Error; err != nil {
		return err
	}

	return nil
}
