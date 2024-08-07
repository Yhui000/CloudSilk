package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductTestRecord(m *model.ProductTestRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductTestRecord(m *model.ProductTestRecord) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductTestRecord(req *proto.QueryProductTestRecordRequest, resp *proto.QueryProductTestRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductTestRecord{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload("ProductInfo.ProductOrder").Preload("TestProject")
	if req.ProductOrderNo != "" || req.ProductSerialNo != "" {
		db.Joins("JOIN product_infos ON product_test_records.product_info_id=product_infos.ID")

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id=product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}
	}
	if req.TestStartTime0 != "" && req.TestStartTime1 != "" {
		db = db.Where("product_test_records.test_start_time BETWEEN ? AND ?", req.TestStartTime0, req.TestStartTime1)
	}
	if req.ProductionLineID != "" {
		db.Joins("JOIN production_stations ON product_test_records.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductTestRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductTestRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductTestRecords() (list []*model.ProductTestRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductTestRecordByID(id string) (*model.ProductTestRecord, error) {
	m := &model.ProductTestRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductTestRecordByIDs(ids []string) ([]*model.ProductTestRecord, error) {
	var m []*model.ProductTestRecord
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductTestRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductTestRecord{}, "`id` = ?", id).Error
}
