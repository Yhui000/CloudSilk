package logic

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMessageSendQueue(m *model.MessageSendQueue) (string, error) {
	m.CurrentState = types.MessageSendQueueStateWaitSend
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMessageSendQueue(m *model.MessageSendQueue) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MessageSendQueueParameter{}, "`message_send_queue_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Delete(&model.MessageSendQueueExecution{}, "`message_send_queue_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at", "create_time").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMessageSendQueue(req *proto.QueryMessageSendQueueRequest, resp *proto.QueryMessageSendQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.MessageSendQueue{}).Preload("MessageType").Preload("RemoteServiceTask")
	if req.TaskNo != "" {
		db = db.Where("`task_no` LIKE ?", "%"+req.TaskNo+"%")
	}
	if req.MessageTypeID != "" {
		db = db.Where("`message_type_id` = ?", req.MessageTypeID)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MessageSendQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MessageSendQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMessageSendQueues() (list []*model.MessageSendQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMessageSendQueueByID(id string) (*model.MessageSendQueue, error) {
	m := &model.MessageSendQueue{}
	err := model.DB.DB().Preload("MessageType").Preload("RemoteServiceTask").Preload("MessageParameters").Preload("SendExecutions").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMessageSendQueueByIDs(ids []string) ([]*model.MessageSendQueue, error) {
	var m []*model.MessageSendQueue
	err := model.DB.DB().Preload("MessageType").Preload("RemoteServiceTask").Preload("MessageParameters").Preload("SendExecutions").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMessageSendQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.MessageSendQueue{}, "`id` = ?", id).Error
}

func SendMessageSendQueue(userID string, ids []string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err = model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/message/messagesendqueue/send",
				Level:        types.EventTypeError,
				Source:       "消息发送队列执行",
				Message:      fmt.Sprintf("%v", r),
				StackTrace:   string(debug.Stack()),
				ReportUserID: userID,
			}).Error; err != nil {
				return
			}
		}
	}()

	return model.DB.DB().Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&model.MessageSendQueue{}).Where("`id` IN ?", ids).Update("current_state", types.MessageSendQueueStateWaitSend).Error; err != nil {
			return
		}
		if err = OnMessageSendQueueExecute(tx, ids); err != nil {
			return
		}
		//TODO TaskQueueExecution

		return
	})
}

func OnMessageSendQueueExecute(tx *gorm.DB, ids []string) (err error) {
	taskNoArray := []string{}
	if err = tx.Model(model.MessageSendQueue{}).Where("`id` IN ?", ids).Pluck("task_no", taskNoArray).Error; err != nil {
		return
	}
	// if err = tx.Model(model.MessageSendQueue{}).Where("`current_state` = ?", types.MessageSendQueueStateWaitSend).Pluck("task_no", taskNoArray).Error; err != nil {
	// 	return
	// }

	for _, taskNo := range taskNoArray {
		messageSendQueue := &model.MessageSendQueue{}
		if err = tx.Preload("MessageType").Preload("MessageParameters").First(messageSendQueue, "`task_no` = ?", taskNo).Error; err != nil {
			return
		}
		content := strings.ReplaceAll(messageSendQueue.Content, "&amp;", "&")
		content = strings.ReplaceAll(content, "&lt;", "<")
		content = strings.ReplaceAll(content, "&gt;", ">")
		content = strings.ReplaceAll(content, "&#39;", "'")
		content = strings.ReplaceAll(content, "&quot;", "\"")

		for _, v := range messageSendQueue.MessageParameters {
			content = strings.ReplaceAll(content, "{"+v.Name+"}", v.FixedValue)
		}

		// TODO 发送消息
		configs := []*model.SystemParamsConfig{}
		if err = tx.Where("`code` = ?", messageSendQueue.MessageType.Code).Find(&configs).Error; err != nil {
			return
		}
		configMap := make(map[string]string)
		for _, v := range configs {
			configMap[v.Key] = v.Value
		}
		var currentState = types.MessageSendQueueStateSent
		var failureReason string
		var success = true
		if err = tool.SendEmail(configMap, messageSendQueue.Cc, messageSendQueue.To, messageSendQueue.Title, content); err != nil {
			currentState = types.MessageSendQueueStateCancelled
			success = false
			failureReason = err.Error()
			// *message = err.Error()
		}

		//更新队列状态
		messageSendQueue.CurrentState = currentState
		messageSendQueue.SendTimes += 1

		if err = tx.Create(&model.MessageSendQueueExecution{
			MessageSendQueueID: messageSendQueue.ID,
			Success:            success,
			FailureReason:      failureReason,
		}).Error; err != nil {
			return
		}

		if err = tx.Save(messageSendQueue).Error; err != nil {
			return
		}

		// if message != "" {
		// 	if err = tx.Create(&model.TaskQueueExecution{
		// 		TaskQueueID: "",
		// 		Success:       false,
		// 		FailureReason: *message,
		// 		DataTrace:     fmt.Sprintf("数据表: MessageSendQueues, 索引: %s", taskNo),
		// 	}).Error; err != nil {
		// 		return
		// 	}
		// }
	}

	return
}
