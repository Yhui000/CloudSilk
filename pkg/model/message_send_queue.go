package model

import "time"

// 消息发送队列
type MessageSendQueue struct {
	ModelID
	TaskNo              string                       `gorm:"size:50;comment:任务编号"`
	MessageTypeID       string                       `gorm:"size:36;comment:消息类型ID"`
	MessageType         *MessageType                 `gorm:"constraint:OnDelete:CASCADE"`
	To                  string                       `gorm:"size:-1;comment:收件地址"`
	Cc                  string                       `gorm:"size:-1;comment:抄送地址"`
	Title               string                       `gorm:"size:500;comment:标题"`
	Content             string                       `gorm:"size:-1;comment:内容"`
	SendTimes           int                          `gorm:"comment:发送次数"`
	CreateTime          time.Time                    `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        int                          `gorm:"comment:创建人员"`
	CurrentState        string                       `gorm:"size:50;comment:当前状态"`
	LastUpdateTime      time.Time                    `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark              string                       `gorm:"size:500;comment:备注"`
	RemoteServiceTaskID *string                      `gorm:"size:36;comment:远程任务ID"`
	RemoteServiceTask   *RemoteServiceTask           `gorm:"constraint:OnDelete:SET NULL"`
	MessageParameters   []*MessageSendQueueParameter `gorm:"constraint:OnDelete:CASCADE"`
	SendExecutions      []*MessageSendQueueExecution `gorm:"constraint:OnDelete:CASCADE"`
}

// 消息发送队列参数
type MessageSendQueueParameter struct {
	ModelID
	MessageSendQueueID string `gorm:"index;size:36;comment:消息发送队列ID"`
	Name               string `gorm:"size:50;comment:名称"`
	FixedValue         string `gorm:"size:5000;comment:设定值"`
	Remark             string `gorm:"size:500;comment:备注"`
}

// 消息发送队列执行
type MessageSendQueueExecution struct {
	ModelID
	MessageSendQueueID string    `gorm:"index;size:36;comment:消息发送队列ID"`
	Success            bool      `gorm:"comment:是否成功"`
	FailureReason      string    `gorm:"size:5000;comment:失败原因"`
	CreateTime         time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
}
