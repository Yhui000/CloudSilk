package logic

import (
	"fmt"
	"runtime/debug"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateSystemEventTrigger(m *model.SystemEventTrigger) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSystemEventTrigger(m *model.SystemEventTrigger) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QuerySystemEventTrigger(req *proto.QuerySystemEventTriggerRequest, resp *proto.QuerySystemEventTriggerResponse, preload bool) {
	db := model.DB.DB().Model(&model.SystemEventTrigger{}).Preload("SystemEvent")
	if req.EventNo != "" {
		db = db.Where("`event_no` = ?", req.EventNo)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.SystemEvent != "" {
		db = db.Joins("JOIN system_events ON system_event_triggers.system_event_id = system_events.id").
			Where("system_events.description = ?", req.SystemEvent)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SystemEventTrigger
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventTriggersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemEventTriggers() (list []*model.SystemEventTrigger, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSystemEventTriggerByID(id string) (*model.SystemEventTrigger, error) {
	m := &model.SystemEventTrigger{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetSystemEventTriggerByIDs(ids []string) ([]*model.SystemEventTrigger, error) {
	var m []*model.SystemEventTrigger
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSystemEventTrigger(id string) (err error) {
	return model.DB.DB().Delete(&model.SystemEventTrigger{}, "`id` = ?", id).Error
}

func ExecuteSystemEventTrigger(userID string, ids []string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err = model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/system/systemeventtrigger/execute",
				Level:        types.EventTypeError,
				Source:       "系统事件触发执行",
				Message:      fmt.Sprintf("%v", r),
				StackTrace:   string(debug.Stack()),
				ReportUserID: userID,
			}).Error; err != nil {
				return
			}
		}
	}()

	return model.DB.DB().Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(model.SystemEventTrigger{}).Where("`id` IN ?", ids).Update("current_state", types.SystemEventTriggerStateWaitExecute).Error; err != nil {
			return
		}
		if err = OnSystemEventTriggerExecute(tx, ids); err != nil {
			return
		}
		//TODO TaskQueueExecution

		return
	})
}

func OnSystemEventTriggerExecute(tx *gorm.DB, ids []string) (err error) {
	systemEventTriggers := []*model.SystemEventTrigger{}
	if err = tx.Preload("SystemEvent.SystemEventSubscriptions").Where("`id` IN ?", ids).Find(&systemEventTriggers).Error; err != nil {
		return
	}

	for _, systemEventTrigger := range systemEventTriggers {
		for _, systemEventSubscription := range systemEventTrigger.SystemEvent.SystemEventSubscriptions {
			remoteServiceTaskQueueParameters := make([]*model.RemoteServiceTaskQueueParameter, len(systemEventTrigger.SystemEventTriggerParameters))
			for i, systemEventTriggerParameter := range systemEventTrigger.SystemEventTriggerParameters {
				remoteServiceTaskQueueParameters[i] = &model.RemoteServiceTaskQueueParameter{
					DataType:    systemEventTriggerParameter.DataType,
					Name:        systemEventTriggerParameter.Name,
					Description: systemEventTriggerParameter.Description,
					Value:       systemEventTriggerParameter.Value,
				}
			}
			remoteServiceTaskQueue := &model.RemoteServiceTaskQueue{
				TaskNo:                           uuid.NewString(),
				CurrentState:                     types.RemoteServiceTaskQueueStateWaitExecute,
				RemoteServiceTaskID:              systemEventSubscription.RemoteServiceTaskID,
				RemoteServiceTaskQueueParameters: remoteServiceTaskQueueParameters,
			}
			if err = tx.Create(remoteServiceTaskQueue).Error; err != nil {
				return
			}

			if err = tx.Create(&model.SystemEventTriggerExecution{
				RemoteServiceTaskQueueID: remoteServiceTaskQueue.ID,
				SystemEventTriggerID:     systemEventTrigger.ID,
			}).Error; err != nil {
				return
			}
		}

		if err = tx.Model(systemEventTrigger).Update("current_state", types.SystemEventTriggerStateExecuted).Error; err != nil {
			return
		}
	}

	return nil
}
