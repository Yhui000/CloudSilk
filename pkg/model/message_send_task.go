package model

type MessageSendTask struct {
	ModelID
	Code                string             `gorm:"size:50;comment:代号"`
	Description         string             `gorm:"size:500;comment:描述"`
	Enable              bool               `gorm:"comment:是否启用"`
	ProductionLineID    *string            `gorm:"size:36;comment:工厂产线ID"`
	ProductionLine      *ProductionLine    `gorm:"constraint:OnDelete:SET NULL"`
	MessageTemplateID   string             `gorm:"size:36;comment:消息模板ID"`
	MessageTemplate     *MessageTemplate   `gorm:"constraint:OnDelete:CASCADE"`
	To                  string             `gorm:"size:-1;comment:收件地址"`
	Cc                  string             `gorm:"size:-1;comment:抄送地址"`
	RemoteServiceTaskID *string            `gorm:"constraint:OnDelete:SET NULL"`
	RemoteServiceTask   *RemoteServiceTask `gorm:"constraint:OnDelete:CASCADE"`
	Remark              string             `gorm:"size:500;comment:备注"`
}
