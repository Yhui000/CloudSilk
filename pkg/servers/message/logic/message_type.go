package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMessageType(m *model.MessageType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMessageType(m *model.MessageType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMessageType(req *proto.QueryMessageTypeRequest, resp *proto.QueryMessageTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.MessageType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MessageType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMessageTypes() (list []*model.MessageType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMessageTypeByID(id string) (*model.MessageType, error) {
	m := &model.MessageType{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMessageTypeByIDs(ids []string) ([]*model.MessageType, error) {
	var m []*model.MessageType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMessageType(id string) (err error) {
	return model.DB.DB().Delete(&model.MessageType{}, "`id` = ?", id).Error
}
