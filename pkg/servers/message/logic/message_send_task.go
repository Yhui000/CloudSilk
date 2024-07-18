package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMessageSendTask(m *model.MessageSendTask) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMessageSendTask(m *model.MessageSendTask) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMessageSendTask(req *proto.QueryMessageSendTaskRequest, resp *proto.QueryMessageSendTaskResponse, preload bool) {
	db := model.DB.DB().Model(&model.MessageSendTask{}).Preload("ProductionLine").Preload("MessageTemplate").Preload("MessageTemplate.MessageType").Preload("RemoteServiceTask")
	if req.Code != "" {
		db = db.Where("message_send_tasks.code LIKE ? OR message_send_tasks.description LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.MessageTypeID != "" {
		db = db.Joins("JOIN message_templates ON message_send_tasks.message_template_id=message_templates.id").
			Where("message_templates.message_type_id = ?", req.MessageTypeID)
	}
	if req.ProductionLineID != "" {
		db = db.Where("message_send_tasks.production_line_id = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MessageSendTask
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageSendTasksToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMessageSendTasks() (list []*model.MessageSendTask, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMessageSendTaskByID(id string) (*model.MessageSendTask, error) {
	m := &model.MessageSendTask{}
	err := model.DB.DB().Preload("ProductionLine").Preload("MessageTemplate").Preload("RemoteServiceTask").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMessageSendTaskByIDs(ids []string) ([]*model.MessageSendTask, error) {
	var m []*model.MessageSendTask
	err := model.DB.DB().Preload("ProductionLine").Preload("MessageTemplate").Preload("RemoteServiceTask").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMessageSendTask(id string) (err error) {
	return model.DB.DB().Delete(&model.MessageSendTask{}, "`id` = ?", id).Error
}
