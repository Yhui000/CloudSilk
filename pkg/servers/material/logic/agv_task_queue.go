package logic

import (
	"fmt"
	"runtime/debug"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateAGVTaskQueue(m *model.AGVTaskQueue) (string, error) {
	departure := &model.AGVParkingSpace{}
	if err := model.DB.DB().First(departure, "`id` = ?", m.DepartureID).Error; err != nil {
		return "", err
	}
	if departure.MaterialShelfID == nil {
		return "", fmt.Errorf("创建失败！当前选择的起点泊位的货架已被移走，请重新选择其他泊位。")
	}

	destination := &model.AGVParkingSpace{}
	if err := model.DB.DB().First(destination, "`id` = ?", m.DestinationID).Error; err != nil {
		return "", err
	}
	if destination.MaterialShelfID != nil {
		return "", fmt.Errorf("创建失败！当前选择的终点泊位已被其他货架占用，请重新选择其他泊位。")
	}
	if destination.ShelfType != departure.ShelfType {
		return "", fmt.Errorf("创建失败！当前选择的起点和终点的货架类型不一致，请重新选择其他泊位。")
	}

	agvTaskType := &model.AGVTaskType{}
	if err := model.DB.DB().First(agvTaskType, "`shelf_type` = ? AND `shelf_type` = ?", destination.ShelfType, destination.SpaceType).Error; err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("创建失败！当前选择的终点泊位缺少对应的AGV任务类型，请先创建AGV任务类型或者选择其他泊位。")
	} else if err != nil {
		return "", err
	}
	m.MaterialShelfID = departure.MaterialShelfID
	m.AGVTaskTypeID = &agvTaskType.ID
	m.CurrentState = types.AGVTaskQueueStateWaitDispatch

	err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(departure).Update("current_state", types.AGVParkingSpaceStateEmpting).Error; err != nil {
			return err
		}
		if err := tx.Model(destination).Update("current_state", types.AGVParkingSpaceStateParking).Error; err != nil {
			return err
		}

		if err := tx.Create(m).Error; err != nil {
			return err
		}

		return nil
	})

	return m.ID, err
}

func UpdateAGVTaskQueue(m *model.AGVTaskQueue) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryAGVTaskQueue(req *proto.QueryAGVTaskQueueRequest, resp *proto.QueryAGVTaskQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.AGVTaskQueue{}).Preload("AGVTaskType").Preload("Departure").Preload("Destination")
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.TaskNo != "" {
		db = db.Where("`task_no` = ?", req.TaskNo)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.AGVTaskQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVTaskQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllAGVTaskQueues() (list []*model.AGVTaskQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetAGVTaskQueueByID(id string) (*model.AGVTaskQueue, error) {
	m := &model.AGVTaskQueue{}
	err := model.DB.DB().Preload("Departure").Preload("Destination").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetAGVTaskQueueByIDs(ids []string) ([]*model.AGVTaskQueue, error) {
	var m []*model.AGVTaskQueue
	err := model.DB.DB().Preload("Departure").Preload("Destination").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteAGVTaskQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.AGVTaskQueue{}, "`id` = ?", id).Error
}

func SubmitAGVTaskQueue(userID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err = model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/material/agvtaskqueue/submit",
				Level:        types.EventTypeError,
				Source:       "AGV任务回调",
				Message:      fmt.Sprintf("%v", r),
				StackTrace:   string(debug.Stack()),
				ReportUserID: userID,
			}).Error; err != nil {
				return
			}
		}
	}()

	agvTaskQueues := []*model.AGVTaskQueue{}
	if err = model.DB.DB().Where("`transaction_state` = ? AND (`current_state` = ? OR `current_state` = ?)", types.TransactionStateWaitSubmit, types.AGVTaskQueueStateCompleted, types.AGVTaskQueueStateCancelled).Find(&agvTaskQueues).Error; err != nil {
		return
	}
	for _, agvTaskQueue := range agvTaskQueues {
		switch agvTaskQueue.CurrentState {
		case types.AGVTaskQueueStateCompleted:
			if err = OnAGVTaskQueueComplete(agvTaskQueue); err != nil {
				return
			}
		case types.AGVTaskQueueStateCancelled:
			if err = OnAGVTaskQueueCancel(agvTaskQueue); err != nil {
				return
			}
		}
		//TODO TaskQueueExecution
	}

	return
}

func OnAGVTaskQueueComplete(agvTaskQueue *model.AGVTaskQueue) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(agvTaskQueue).Update("`transaction_state` = ?", types.TransactionStateSubmitted).Error; err != nil {
			return
		}

		if err = tx.Model(agvTaskQueue.Departure).Updates(map[string]interface{}{
			"current_state":     types.AGVParkingSpaceStateEmptied,
			"material_shelf_id": nil,
		}).Error; err != nil {
			return
		}

		if err = tx.Model(agvTaskQueue.Destination).Updates(map[string]interface{}{
			"current_state":     types.AGVParkingSpaceStateParked,
			"material_shelf_id": agvTaskQueue.MaterialShelfID,
		}).Error; err != nil {
			return
		}

		if err = tx.Model(agvTaskQueue.MaterialShelf).Update("agv_parking_space_id = ?", agvTaskQueue.Destination.ID).Error; err != nil {
			return
		}

		return
	})
}

func OnAGVTaskQueueCancel(agvTaskQueue *model.AGVTaskQueue) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(agvTaskQueue).Update("`transaction_state` = ?", types.TransactionStateSubmitted).Error; err != nil {
			return
		}
		if err = tx.Model(agvTaskQueue.Departure).Update("`current_state` = ?", types.AGVParkingSpaceStateParked).Error; err != nil {
			return
		}
		if err = tx.Model(agvTaskQueue.Destination).Update("`current_state` = ?", types.AGVParkingSpaceStateEmptied).Error; err != nil {
			return
		}
		return
	})
}
