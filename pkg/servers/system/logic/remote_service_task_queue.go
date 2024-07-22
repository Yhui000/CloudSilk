package logic

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateRemoteServiceTaskQueue(m *model.RemoteServiceTaskQueue) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateRemoteServiceTaskQueue(m *model.RemoteServiceTaskQueue) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryRemoteServiceTaskQueue(req *proto.QueryRemoteServiceTaskQueueRequest, resp *proto.QueryRemoteServiceTaskQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.RemoteServiceTaskQueue{}).Preload("RemoteServiceTask").Preload("RemoteServiceTask.RemoteService")
	if req.TaskNo != "" {
		db.Where("`task_no` LIKE ? OR `request_text` LIKE ? OR `response_text` LIKE ?", "%"+req.TaskNo+"%", "%"+req.TaskNo+"%", "%"+req.TaskNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.RemoteServiceTaskQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceTaskQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllRemoteServiceTaskQueues() (list []*model.RemoteServiceTaskQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetRemoteServiceTaskQueueByID(id string) (*model.RemoteServiceTaskQueue, error) {
	m := &model.RemoteServiceTaskQueue{}
	err := model.DB.DB().Preload("RemoteServiceTaskQueueParameters").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetRemoteServiceTaskQueueByIDs(ids []string) ([]*model.RemoteServiceTaskQueue, error) {
	var m []*model.RemoteServiceTaskQueue
	err := model.DB.DB().Preload("RemoteServiceTaskQueueParameters").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteRemoteServiceTaskQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.RemoteServiceTaskQueue{}, "`id` = ?", id).Error
}

func RemoteServiceTaskQueueCallback(userID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err = model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/system/remoteservicetaskqueue/callback",
				Level:        types.EventTypeError,
				Source:       "远程服务任务回调",
				Message:      fmt.Sprintf("%v", r),
				StackTrace:   string(debug.Stack()),
				ReportUserID: userID,
			}).Error; err != nil {
				return
			}
		}
	}()

	remoteServiceTaskQueues := []*model.RemoteServiceTaskQueue{}
	if err = model.DB.DB().Preload("RemoteServiceTask").Where("`current_state` = ? AND `transaction_state` = ?", types.RemoteServiceTaskQueueStateCompleted, types.TransactionStateWaitSubmit).Find(&remoteServiceTaskQueues).Error; err != nil {
		return
	}
	for _, v := range remoteServiceTaskQueues {
		if err = OnRemoteServiceTaskQueueCallback(v); err != nil {
			return
		}
		//TODO TaskQueueExecution
	}
	return
}

func OnRemoteServiceTaskQueueCallback(remoteServiceTaskQueue *model.RemoteServiceTaskQueue) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		remoteServiceTaskQueue.TransactionState = types.TransactionStateSubmitted
		remoteServiceTaskQueue.InvokeCount += 1

		actionRoute := remoteServiceTaskQueue.RemoteServiceTask.CallbackMethod
		if actionRoute != "" {
			systemConfig := &model.SystemParamsConfig{}
			if err := tx.First(&systemConfig, "`code` = ?", "10010").Error; err == gorm.ErrRecordNotFound {
				return fmt.Errorf("缺少系统参数配置，配置项: 10010")
			} else if err != nil {
				return err
			}

			serviceURI, err := url.Parse(systemConfig.Value)
			if err != nil {
				return err
			}
			requestURL := serviceURI.ResolveReference(&url.URL{Path: actionRoute}).String()

			requestByte, err := json.Marshal(map[string]string{
				"taskNo":       remoteServiceTaskQueue.TaskNo,
				"responseText": remoteServiceTaskQueue.ResponseText,
			})
			if err != nil {
				return err
			}

			resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(requestByte))
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			response := map[string]interface{}{}
			if err = json.Unmarshal(body, &response); err != nil {
				return err
			}

			if code, ok := response["code"].(int); ok && code != types.ServiceResponseCodeSuccess {
				remoteServiceTaskQueue.TransactionState = types.TransactionStateCancelled
				if remoteServiceTaskQueue.RemoteServiceTask.FailureMeasure == types.TaskFailureMeasureRetry && remoteServiceTaskQueue.RemoteServiceTask.MaximumInvokeCount > remoteServiceTaskQueue.InvokeCount {
					remoteServiceTaskQueue.CurrentState = types.RemoteServiceTaskQueueStateWaitExecute
					remoteServiceTaskQueue.TransactionState = ""
				}
			}
		}

		return tx.Save(remoteServiceTaskQueue).Error
	})
}

func RemoteServiceTaskQueueExecute(userID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err = model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/system/remoteservicetaskqueue/execute",
				Level:        types.EventTypeError,
				Source:       "远程服务任务执行",
				Message:      fmt.Sprintf("%v", r),
				StackTrace:   string(debug.Stack()),
				ReportUserID: userID,
			}).Error; err != nil {
				return
			}
		}
	}()

	remoteServiceTaskQueues := []*model.RemoteServiceTaskQueue{}
	if err = model.DB.DB().Preload("RemoteServiceTask.RemoteServiceTaskParameters").Preload("RemoteServiceTaskQueueParameters").Where("`current_state` = ?", types.RemoteServiceTaskQueueStateWaitExecute).Find(&remoteServiceTaskQueues).Error; err != nil {
		return
	}
	remoteServiceTasks := map[string][]*model.RemoteServiceTaskQueue{}
	for _, v := range remoteServiceTaskQueues {
		remoteServiceTasks[v.RemoteServiceTaskID] = append(remoteServiceTasks[v.RemoteServiceTaskID], v)
	}

	for _, remoteServiceTaskQueues := range remoteServiceTasks {
		for _, v := range remoteServiceTaskQueues {
			if err = OnRemoteServiceTaskQueueCallback(v); err != nil {
				return
			}
		}
		//TODO TaskQueueExecution
	}
	return
}

func OnRemoteServiceTaskQueueExecute(remoteServiceTaskQueue *model.RemoteServiceTaskQueue) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		serviceURI, err := url.Parse(remoteServiceTaskQueue.RemoteServiceTask.RemoteService.Address)
		if err != nil {
			return err
		}
		requestURL := serviceURI.ResolveReference(&url.URL{Path: remoteServiceTaskQueue.RemoteServiceTask.InvokeMethod}).String()

		remoteServiceTaskQueue.RequestURL = requestURL

		dictionary := map[string]interface{}{}
		for _, remoteServiceTaskParameter := range remoteServiceTaskQueue.RemoteServiceTask.RemoteServiceTaskParameters {
			if _, ok := dictionary[remoteServiceTaskParameter.Name]; ok {
				delete(dictionary, remoteServiceTaskParameter.Name)
			}

			switch remoteServiceTaskParameter.DataType {
			case types.DataTypeString:
				dictionary[remoteServiceTaskParameter.Name] = remoteServiceTaskParameter.Value
			case types.DataTypeInteger:
				v, err := strconv.Atoi(remoteServiceTaskParameter.Value)
				if err != nil {
					return err
				}
				dictionary[remoteServiceTaskParameter.Name] = v
			case types.DataTypeDouble:
				v, err := strconv.ParseFloat(remoteServiceTaskParameter.Value, 64)
				if err != nil {
					return err
				}
				dictionary[remoteServiceTaskParameter.Name] = v
			case types.DataTypeArray:
			case types.DataTypeObject:
				v := map[string]interface{}{}
				if err := json.Unmarshal([]byte(remoteServiceTaskParameter.Value), &v); err != nil {
					return err
				}
				dictionary[remoteServiceTaskParameter.Name] = v
			default:
			}
		}

		for _, remoteServiceTaskQueueParameter := range remoteServiceTaskQueue.RemoteServiceTaskQueueParameters {
			if _, ok := dictionary[remoteServiceTaskQueueParameter.Name]; ok {
				delete(dictionary, remoteServiceTaskQueueParameter.Name)
			}

			switch remoteServiceTaskQueueParameter.DataType {
			case types.DataTypeString:
				dictionary[remoteServiceTaskQueueParameter.Name] = remoteServiceTaskQueueParameter.Value
			case types.DataTypeInteger:
				v, err := strconv.Atoi(remoteServiceTaskQueueParameter.Value)
				if err != nil {
					return err
				}
				dictionary[remoteServiceTaskQueueParameter.Name] = v
			case types.DataTypeDouble:
				v, err := strconv.ParseFloat(remoteServiceTaskQueueParameter.Value, 64)
				if err != nil {
					return err
				}
				dictionary[remoteServiceTaskQueueParameter.Name] = v
			case types.DataTypeArray:
			case types.DataTypeObject:
				v := map[string]interface{}{}
				if err := json.Unmarshal([]byte(remoteServiceTaskQueueParameter.Value), &v); err != nil {
					return err
				}
				dictionary[remoteServiceTaskQueueParameter.Name] = v
			default:
			}
		}

		requestByte, err := json.Marshal(dictionary)
		if err != nil {
			return err
		}

		remoteServiceTaskQueue.RequestText = string(requestByte)

		var serviceHeaders = remoteServiceTaskQueue.RemoteServiceTask.RemoteService.Headers

		headers := map[string]string{}
		if serviceHeaders != "" {
			matches := regexp.MustCompile("(.*?):(.*?)").FindAllStringSubmatch(serviceHeaders, -1)
			for _, match := range matches {
				var key = match[1]
				var value = match[2]
				headers[key] = value
			}
		}

		useCredential := remoteServiceTaskQueue.RemoteServiceTask.RemoteService.UseCredential
		userName := remoteServiceTaskQueue.RemoteServiceTask.RemoteService.UserName
		password := remoteServiceTaskQueue.RemoteServiceTask.RemoteService.Password
		timeout := remoteServiceTaskQueue.RemoteServiceTask.RemoteService.Timeout
		responseText, err := GetResponse(requestURL, string(requestByte), headers, useCredential, userName, password, timeout)
		if err != nil {
			return err
		}
		remoteServiceTaskQueue.ResponseText = responseText

		remoteServiceTaskQueue.FinishTime = tool.Time2NullTime(time.Now())
		remoteServiceTaskQueue.CurrentState = types.RemoteServiceTaskQueueStateCompleted
		remoteServiceTaskQueue.TransactionState = types.TransactionStateWaitSubmit

		return tx.Save(remoteServiceTaskQueue).Error
	})
}

func GetResponse(url string, jsonData string, headers map[string]string, useCredential bool, userName, password string, timeout int64) (string, error) {
	// 创建一个HTTP客户端，并设置超时
	client := &http.Client{
		Timeout: time.Duration(timeout),
	}

	// 创建请求体
	requestBody := bytes.NewBuffer([]byte(jsonData))

	// 创建请求
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 添加自定义头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 如果需要认证，则添加认证头部
	if useCredential {
		if userName == "" || password == "" {
			return "", fmt.Errorf("")
		}

		// 创建Base64编码的认证字符串
		auth := userName + ":" + password
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", "Basic "+encodedAuth)
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 返回响应体作为字符串
	return string(responseBody), nil
}
