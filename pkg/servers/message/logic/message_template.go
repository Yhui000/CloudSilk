package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMessageTemplate(m *model.MessageTemplate) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMessageTemplate(m *model.MessageTemplate) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MessageParameter{}, "`message_template_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMessageTemplate(req *proto.QueryMessageTemplateRequest, resp *proto.QueryMessageTemplateResponse, preload bool) {
	db := model.DB.DB().Model(&model.MessageTemplate{}).Preload("MessageType")
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.MessageTypeID != "" {
		db = db.Where("`message_type_id` = ?", req.MessageTypeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MessageTemplate
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageTemplatesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMessageTemplates() (list []*model.MessageTemplate, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMessageTemplateByID(id string) (*model.MessageTemplate, error) {
	m := &model.MessageTemplate{}
	err := model.DB.DB().Preload("MessageType").Preload("MessageParameters").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMessageTemplateByIDs(ids []string) ([]*model.MessageTemplate, error) {
	var m []*model.MessageTemplate
	err := model.DB.DB().Preload("MessageType").Preload("MessageParameters").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMessageTemplate(id string) (err error) {
	return model.DB.DB().Delete(&model.MessageTemplate{}, "`id` = ?", id).Error
}
