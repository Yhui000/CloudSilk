// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: production_process.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductionProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
	// 顺序
	SortIndex int32 `protobuf:"varint,4,opt,name=sortIndex,proto3" json:"sortIndex"`
	// 代号
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
	// 识别码
	Identifier string `protobuf:"bytes,7,opt,name=identifier,proto3" json:"identifier"`
	// 是否启用
	Enable bool `protobuf:"varint,8,opt,name=enable,proto3" json:"enable"`
	// 默认匹配
	InitialValue bool `protobuf:"varint,9,opt,name=initialValue,proto3" json:"initialValue"`
	// 是否报工
	EnableReport bool `protobuf:"varint,10,opt,name=enableReport,proto3" json:"enableReport"`
	// 是否管控
	EnableControl bool `protobuf:"varint,11,opt,name=enableControl,proto3" json:"enableControl"`
	// 模拟执行
	SimulateExecution bool `protobuf:"varint,12,opt,name=simulateExecution,proto3" json:"simulateExecution"`
	// 工序类型
	ProcessType int32 `protobuf:"varint,13,opt,name=processType,proto3" json:"processType"`
	// 载具类型
	VehicleType int32 `protobuf:"varint,14,opt,name=vehicleType,proto3" json:"vehicleType"`
	// 产品状态
	ProductState string `protobuf:"bytes,15,opt,name=productState,proto3" json:"productState"`
	// 备注
	Remark string `protobuf:"bytes,16,opt,name=remark,proto3" json:"remark"`
	// 生产产线ID
	ProductionLineID string              `protobuf:"bytes,17,opt,name=productionLineID,proto3" json:"productionLineID"`
	ProductionLine   *ProductionLineInfo `protobuf:"bytes,18,opt,name=productionLine,proto3" json:"productionLine"`
	// 支持工站
	ProductionProcessAvailableStations []*ProductionProcessAvailableStationInfo `protobuf:"bytes,22,rep,name=productionProcessAvailableStations,proto3" json:"productionProcessAvailableStations"`
	// 支援工序
	// repeated ProductionProcessInfo productionProcesses=20;
	AttributeExpressions []*AttributeExpressionInfo `protobuf:"bytes,21,rep,name=attributeExpressions,proto3" json:"attributeExpressions"`
}

func (x *ProductionProcessInfo) Reset() {
	*x = ProductionProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductionProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductionProcessInfo) ProtoMessage() {}

func (x *ProductionProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductionProcessInfo.ProtoReflect.Descriptor instead.
func (*ProductionProcessInfo) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{0}
}

func (x *ProductionProcessInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductionProcessInfo) GetSortIndex() int32 {
	if x != nil {
		return x.SortIndex
	}
	return 0
}

func (x *ProductionProcessInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductionProcessInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductionProcessInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProductionProcessInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ProductionProcessInfo) GetInitialValue() bool {
	if x != nil {
		return x.InitialValue
	}
	return false
}

func (x *ProductionProcessInfo) GetEnableReport() bool {
	if x != nil {
		return x.EnableReport
	}
	return false
}

func (x *ProductionProcessInfo) GetEnableControl() bool {
	if x != nil {
		return x.EnableControl
	}
	return false
}

func (x *ProductionProcessInfo) GetSimulateExecution() bool {
	if x != nil {
		return x.SimulateExecution
	}
	return false
}

func (x *ProductionProcessInfo) GetProcessType() int32 {
	if x != nil {
		return x.ProcessType
	}
	return 0
}

func (x *ProductionProcessInfo) GetVehicleType() int32 {
	if x != nil {
		return x.VehicleType
	}
	return 0
}

func (x *ProductionProcessInfo) GetProductState() string {
	if x != nil {
		return x.ProductState
	}
	return ""
}

func (x *ProductionProcessInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductionProcessInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *ProductionProcessInfo) GetProductionLine() *ProductionLineInfo {
	if x != nil {
		return x.ProductionLine
	}
	return nil
}

func (x *ProductionProcessInfo) GetProductionProcessAvailableStations() []*ProductionProcessAvailableStationInfo {
	if x != nil {
		return x.ProductionProcessAvailableStations
	}
	return nil
}

func (x *ProductionProcessInfo) GetAttributeExpressions() []*AttributeExpressionInfo {
	if x != nil {
		return x.AttributeExpressions
	}
	return nil
}

type ProductionProcessAvailableStationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string                 `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	ProductionProcessID string                 `protobuf:"bytes,1,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	ProductionStationID string                 `protobuf:"bytes,2,opt,name=productionStationID,proto3" json:"productionStationID"`
	ProductionStation   *ProductionStationInfo `protobuf:"bytes,3,opt,name=productionStation,proto3" json:"productionStation"`
}

func (x *ProductionProcessAvailableStationInfo) Reset() {
	*x = ProductionProcessAvailableStationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductionProcessAvailableStationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductionProcessAvailableStationInfo) ProtoMessage() {}

func (x *ProductionProcessAvailableStationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductionProcessAvailableStationInfo.ProtoReflect.Descriptor instead.
func (*ProductionProcessAvailableStationInfo) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProductionProcessAvailableStationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductionProcessAvailableStationInfo) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *ProductionProcessAvailableStationInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *ProductionProcessAvailableStationInfo) GetProductionStation() *ProductionStationInfo {
	if x != nil {
		return x.ProductionStation
	}
	return nil
}

type QueryProductionProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 生产产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,4,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// @inject_tag: uri:"productionStationID" form:"productionStationID"
	ProductionStationID string `protobuf:"bytes,5,opt,name=productionStationID,proto3" json:"productionStationID" uri:"productionStationID" form:"productionStationID"`
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,6,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProductionProcessRequest) Reset() {
	*x = QueryProductionProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionProcessRequest) ProtoMessage() {}

func (x *QueryProductionProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionProcessRequest.ProtoReflect.Descriptor instead.
func (*QueryProductionProcessRequest) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductionProcessRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductionProcessRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductionProcessRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductionProcessRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryProductionProcessRequest) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *QueryProductionProcessRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductionProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionProcessInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                    `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                    `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                    `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductionProcessResponse) Reset() {
	*x = QueryProductionProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionProcessResponse) ProtoMessage() {}

func (x *QueryProductionProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionProcessResponse.ProtoReflect.Descriptor instead.
func (*QueryProductionProcessResponse) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductionProcessResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductionProcessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductionProcessResponse) GetData() []*ProductionProcessInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductionProcessResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductionProcessResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductionProcessResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductionProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionProcessInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductionProcessResponse) Reset() {
	*x = GetAllProductionProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductionProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductionProcessResponse) ProtoMessage() {}

func (x *GetAllProductionProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductionProcessResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductionProcessResponse) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductionProcessResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductionProcessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductionProcessResponse) GetData() []*ProductionProcessInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductionProcessDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductionProcessInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductionProcessDetailResponse) Reset() {
	*x = GetProductionProcessDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_process_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductionProcessDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionProcessDetailResponse) ProtoMessage() {}

func (x *GetProductionProcessDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_process_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionProcessDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductionProcessDetailResponse) Descriptor() ([]byte, []int) {
	return file_production_process_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductionProcessDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductionProcessDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductionProcessDetailResponse) GetData() *ProductionProcessInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_production_process_proto protoreflect.FileDescriptor

var file_production_process_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x06, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x7c, 0x0a, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x15, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x25, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xeb, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0xd3, 0x01, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8e, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xbe, 0x01, 0x0a,
	0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x56, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_production_process_proto_rawDescOnce sync.Once
	file_production_process_proto_rawDescData = file_production_process_proto_rawDesc
)

func file_production_process_proto_rawDescGZIP() []byte {
	file_production_process_proto_rawDescOnce.Do(func() {
		file_production_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_production_process_proto_rawDescData)
	})
	return file_production_process_proto_rawDescData
}

var file_production_process_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_production_process_proto_goTypes = []interface{}{
	(*ProductionProcessInfo)(nil),                 // 0: proto.ProductionProcessInfo
	(*ProductionProcessAvailableStationInfo)(nil), // 1: proto.ProductionProcessAvailableStationInfo
	(*QueryProductionProcessRequest)(nil),         // 2: proto.QueryProductionProcessRequest
	(*QueryProductionProcessResponse)(nil),        // 3: proto.QueryProductionProcessResponse
	(*GetAllProductionProcessResponse)(nil),       // 4: proto.GetAllProductionProcessResponse
	(*GetProductionProcessDetailResponse)(nil),    // 5: proto.GetProductionProcessDetailResponse
	(*ProductionLineInfo)(nil),                    // 6: proto.ProductionLineInfo
	(*AttributeExpressionInfo)(nil),               // 7: proto.AttributeExpressionInfo
	(*ProductionStationInfo)(nil),                 // 8: proto.ProductionStationInfo
	(Code)(0),                                     // 9: proto.Code
	(*GetDetailRequest)(nil),                      // 10: proto.GetDetailRequest
}
var file_production_process_proto_depIdxs = []int32{
	6,  // 0: proto.ProductionProcessInfo.productionLine:type_name -> proto.ProductionLineInfo
	1,  // 1: proto.ProductionProcessInfo.productionProcessAvailableStations:type_name -> proto.ProductionProcessAvailableStationInfo
	7,  // 2: proto.ProductionProcessInfo.attributeExpressions:type_name -> proto.AttributeExpressionInfo
	8,  // 3: proto.ProductionProcessAvailableStationInfo.productionStation:type_name -> proto.ProductionStationInfo
	9,  // 4: proto.QueryProductionProcessResponse.code:type_name -> proto.Code
	0,  // 5: proto.QueryProductionProcessResponse.data:type_name -> proto.ProductionProcessInfo
	9,  // 6: proto.GetAllProductionProcessResponse.code:type_name -> proto.Code
	0,  // 7: proto.GetAllProductionProcessResponse.data:type_name -> proto.ProductionProcessInfo
	9,  // 8: proto.GetProductionProcessDetailResponse.code:type_name -> proto.Code
	0,  // 9: proto.GetProductionProcessDetailResponse.data:type_name -> proto.ProductionProcessInfo
	2,  // 10: proto.ProductionProcess.Query:input_type -> proto.QueryProductionProcessRequest
	10, // 11: proto.ProductionProcess.GetDetail:input_type -> proto.GetDetailRequest
	3,  // 12: proto.ProductionProcess.Query:output_type -> proto.QueryProductionProcessResponse
	5,  // 13: proto.ProductionProcess.GetDetail:output_type -> proto.GetProductionProcessDetailResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_production_process_proto_init() }
func file_production_process_proto_init() {
	if File_production_process_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_production_line_proto_init()
	file_attribute_expression_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_production_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductionProcessInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_production_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductionProcessAvailableStationInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_production_process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionProcessRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_production_process_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionProcessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_production_process_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductionProcessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_production_process_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductionProcessDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_production_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_production_process_proto_goTypes,
		DependencyIndexes: file_production_process_proto_depIdxs,
		MessageInfos:      file_production_process_proto_msgTypes,
	}.Build()
	File_production_process_proto = out.File
	file_production_process_proto_rawDesc = nil
	file_production_process_proto_goTypes = nil
	file_production_process_proto_depIdxs = nil
}
