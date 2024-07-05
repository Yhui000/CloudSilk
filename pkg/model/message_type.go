package model

// 消息类型
type MessageType struct {
	ModelID
	Code        string `gorm:"size:50;comment:代号"`
	Description string `gorm:"size:500;comment:描述"`
	Remark      string `gorm:"size:500;comment:备注"`
}
