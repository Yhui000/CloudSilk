package model

import "time"

type MaterialContainer struct {
	ModelID
	Code                            string                 `gorm:"size:50;comment:编号"`
	Description                     string                 `gorm:"size:100;comment:描述"`
	Identifier                      string                 `gorm:"size:50;comment:识别码"`
	CurrentState                    string                 `gorm:"size:50;comment:当前状态"`
	LastUpdateTime                  time.Time              `gorm:"comment:状态变更时间"`
	Remark                          string                 `gorm:"size:500;comment:备注"`
	MaterialContainerTypeID         *string                `gorm:"size:36;comment:容器类型ID"`
	MaterialContainerType           *MaterialContainerType `gorm:"constraint:OnDelete:SET NULL"`
	MaterialShelfBinID              *string                `gorm:"size:36;comment:当前库位ID"`
	MaterialShelfBin                *MaterialShelfBin      `gorm:"constraint:OnDelete:SET NULL"`
	MaterialContainerBindingRecords []*MaterialContainerBindingRecord
}

type MaterialContainerBindingRecord struct {
	ModelID
	MaterialContainerID string            `gorm:"size:36;comment:物料容器ID"`
	MaterialContainer   MaterialContainer `` //物料容器
	MaterialInfoID      string            `gorm:"size:36;comment:绑定物料ID"`
	MaterialInfo        MaterialInfo      `` //绑定物料
	BindQTY             int               `gorm:"comment:装入数量"`
	MaterialTraceNo     string            `gorm:"size:100;comment:物料追溯号"`
	CreateTime          time.Time         `gorm:"autoCreateTime:nano;comment:绑定时间"`
	CreateUserID        string            `gorm:"size:36;comment:操作人员"`
	CurrentState        string            `gorm:"size:50;comment:当前状态"`
	LastUpdateTime      time.Time         `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark              string            `gorm:"size:500;comment:备注"`
}

type MaterialContainerType struct {
	ModelID
	Code        string  `grom:"size:50;comment:编号"`
	Description string  `grom:"size:100;comment:描述"`
	Length      int64   `grom:"comment:长(mm)"`
	Width       int64   `grom:"comment:宽(mm)"`
	Height      int64   `grom:"comment:高(mm)"`
	WeightLimit float64 `grom:"comment:限重(kg)"`
	HeightLimit int64   `grom:"comment:限高(mm)"`
	Remark      string  `grom:"size:500;comment:备注"`
}

type MaterialShelfBin struct {
	ModelID
	Code                           string                          `gorm:"size:50;comment:编号"`
	RowIndex                       int32                           `gorm:"comment:行"`
	ColumnIndex                    int32                           `gorm:"comment:列"`
	Identifier                     string                          `gorm:"size:50;comment:识别码"`
	MaterialContainerTypeID        string                          `gorm:"size:36;comment:兼容容器ID"`
	MaterialContainerType          *MaterialContainerType          `` //兼容容器
	EnableDispatch                 bool                            `gorm:"comment:缺料调度"`
	EnableMonitor                  bool                            `gorm:"comment:状态监控"`
	CurrentState                   string                          `gorm:"size:50;comment:当前状态"`
	LastUpdateTime                 time.Time                       `gorm:"comment:状态变更时间"`
	Remark                         string                          `gorm:"size:500;comment:备注"`
	MaterialShelfID                string                          `gorm:"size:36;comment:隶属料架ID"`
	MaterialShelf                  *MaterialShelf                  `` //隶属料架
	MaterialInfoID                 string                          `gorm:"size:36;comment:当前物料ID"`
	MaterialInfo                   *MaterialInfo                   `` //当前物料
	MaterialContainers             []*MaterialContainer            `` //当前容器
	MaterialShelfBinDispatchInfoes []*MaterialShelfBinDispatchInfo `` //调度信息
}

type MaterialShelf struct {
	ModelID
	Code              string              `gorm:"size:50;comment:代号"`
	Description       string              `gorm:"size:500;comment:描述"`
	Identifier        string              `gorm:"size:50;comment:识别码"`
	ShelfType         int                 `gorm:"comment:货架类型"`
	Enable            bool                `gorm:"comment:是否启用"`
	Remark            string              `gorm:"size:500;comment:备注"`
	MaterialStoreID   string              `gorm:"size:36;comment:隶属仓库ID"`
	MaterialStore     *MaterialStore      `` //隶属仓库
	MaterialShelfBins []*MaterialShelfBin `` //下属库位
	ProductionLines   []*ProductionLine   `` //支援产线
}

type MaterialStore struct {
	ModelID
	Code                string               `gorm:"size:50;comment:代号"`
	Description         string               `gorm:"size:500;comment:描述"`
	Remark              string               `gorm:"size:500;comment:备注"`
	MaterialShelves     []*MaterialShelf     //物料货架
	MaterialInventories []*MaterialInventory //物料库存
}

type MaterialShelfBinDispatchInfo struct {
	ModelID
	MaterialShelfBinID string           `gorm:"size:36;comment:隶属库位ID"`
	MaterialShelfBin   MaterialShelfBin `` //隶属库位
	Priority           int32            `gorm:"comment:优先级"`
	MaterialInfoID     string           `gorm:"size:36;comment:调度物料ID"`
	MaterialInfo       MaterialInfo     `` //调度物料
	Enable             bool             `gorm:"comment:是否启用"`
	Remark             string           `gorm:"size:500;comment:备注"`
}

type MaterialInventory struct {
	ModelID
	MaterialInfoID  string        `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo    MaterialInfo  `` //物料信息
	MaterialStoreID string        `gorm:"size:36;comment:物料仓库ID"`
	MaterialStore   MaterialStore `` //物料仓库
	StoredQTY       int64         `gorm:"comment:库存数量"`
	IssuedQTY       int64         `gorm:"comment:锁定数量"`
	FeedingQTY      int64         `gorm:"comment:正在补料数量"`
	IssuingQTY      int64         `gorm:"comment:正在备库数量"`
	CreateTime      time.Time     `gorm:"comment:创建时间"`
	CreateUserID    string        `gorm:"size:36;comment:创建人员"`
	CurrentState    string        `gorm:"size:50;comment:当前状态"`
	LastUpdateTime  time.Time     `gorm:"comment:状态变更时间"`
	Remark          string        `gorm:"size:500;comment:备注"`
}
