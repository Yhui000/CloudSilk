package model

// 消息模版
type MessageTemplate struct {
	ModelID
	Code              string              `gorm:"size:50;comment:代号"`
	Description       string              `gorm:"size:500;comment:描述"`
	Title             string              `gorm:"size:500;comment:标题"`
	Content           string              `gorm:"size:-1;comment:内容"`
	Remark            string              `gorm:"size:500;comment:备注"`
	MessageTypeID     string              `gorm:"size:36;comment:消息类型ID"`
	MessageType       *MessageType        `gorm:"constraint:OnDelete:CASCADE"`
	MessageParameters []*MessageParameter `gorm:"constraint:OnDelete:CASCADE"`
}

// 消息参数
type MessageParameter struct {
	ModelID
	MessageTemplateID string `gorm:"index;size:36;comment:消息模版ID"`
	Name              string `gorm:"size:50;comment:名称"`
	DefaultValue      string `gorm:"size:1000;comment:预设值"`
}
