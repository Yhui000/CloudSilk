// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: production_station_alarm.proto

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

type ProductionStationAlarmInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 创建时间
	CreateTime string `protobuf:"bytes,3,opt,name=createTime,proto3" json:"createTime"`
	// 创建人员ID
	CreateUserID   string `protobuf:"bytes,4,opt,name=createUserID,proto3" json:"createUserID"`
	CreateUserName string `protobuf:"bytes,17,opt,name=createUserName,proto3" json:"createUserName"`
	// 生产工站ID
	ProductionStationID string                 `protobuf:"bytes,5,opt,name=productionStationID,proto3" json:"productionStationID"`
	ProductionStation   *ProductionStationInfo `protobuf:"bytes,14,opt,name=productionStation,proto3" json:"productionStation"`
	// 生产工序ID
	ProductionProcessID string                 `protobuf:"bytes,6,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	ProductionProcess   *ProductionProcessInfo `protobuf:"bytes,15,opt,name=productionProcess,proto3" json:"productionProcess"`
	// 产品信息ID
	ProductInfoID string           `protobuf:"bytes,7,opt,name=productInfoID,proto3" json:"productInfoID"`
	ProductInfo   *ProductInfoInfo `protobuf:"bytes,16,opt,name=productInfo,proto3" json:"productInfo"`
	// 报警编号
	AlarmNo string `protobuf:"bytes,8,opt,name=alarmNo,proto3" json:"alarmNo"`
	// 报警信息
	AlarmMessage string `protobuf:"bytes,9,opt,name=alarmMessage,proto3" json:"alarmMessage"`
	// 当前状态
	CurrentState string `protobuf:"bytes,10,opt,name=currentState,proto3" json:"currentState"`
	// 处理方式
	HandleMethod string `protobuf:"bytes,11,opt,name=handleMethod,proto3" json:"handleMethod"`
	// 处理时间
	HandleTime string `protobuf:"bytes,12,opt,name=handleTime,proto3" json:"handleTime"`
	// 备注信息
	Remark string `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark"`
}

func (x *ProductionStationAlarmInfo) Reset() {
	*x = ProductionStationAlarmInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_station_alarm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductionStationAlarmInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductionStationAlarmInfo) ProtoMessage() {}

func (x *ProductionStationAlarmInfo) ProtoReflect() protoreflect.Message {
	mi := &file_production_station_alarm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductionStationAlarmInfo.ProtoReflect.Descriptor instead.
func (*ProductionStationAlarmInfo) Descriptor() ([]byte, []int) {
	return file_production_station_alarm_proto_rawDescGZIP(), []int{0}
}

func (x *ProductionStationAlarmInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetCreateUserID() string {
	if x != nil {
		return x.CreateUserID
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetCreateUserName() string {
	if x != nil {
		return x.CreateUserName
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetProductionStation() *ProductionStationInfo {
	if x != nil {
		return x.ProductionStation
	}
	return nil
}

func (x *ProductionStationAlarmInfo) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetProductionProcess() *ProductionProcessInfo {
	if x != nil {
		return x.ProductionProcess
	}
	return nil
}

func (x *ProductionStationAlarmInfo) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetProductInfo() *ProductInfoInfo {
	if x != nil {
		return x.ProductInfo
	}
	return nil
}

func (x *ProductionStationAlarmInfo) GetAlarmNo() string {
	if x != nil {
		return x.AlarmNo
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetAlarmMessage() string {
	if x != nil {
		return x.AlarmMessage
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetHandleMethod() string {
	if x != nil {
		return x.HandleMethod
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetHandleTime() string {
	if x != nil {
		return x.HandleTime
	}
	return ""
}

func (x *ProductionStationAlarmInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryProductionStationAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 创建时间开始
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,5,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// 创建时间结束
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,6,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
	// 生产产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,14,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// 生产工单号
	// @inject_tag: uri:"productOrderNo" form:"productOrderNo"
	ProductOrderNo string `protobuf:"bytes,15,opt,name=productOrderNo,proto3" json:"productOrderNo" uri:"productOrderNo" form:"productOrderNo"`
	// 产品序列号
	// @inject_tag: uri:"productSerialNo" form:"productSerialNo"
	ProductSerialNo string `protobuf:"bytes,16,opt,name=productSerialNo,proto3" json:"productSerialNo" uri:"productSerialNo" form:"productSerialNo"`
	// 报警信息
	// @inject_tag: uri:"alarmMessage" form:"alarmMessage"
	AlarmMessage string `protobuf:"bytes,17,opt,name=alarmMessage,proto3" json:"alarmMessage" uri:"alarmMessage" form:"alarmMessage"`
}

func (x *QueryProductionStationAlarmRequest) Reset() {
	*x = QueryProductionStationAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_station_alarm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionStationAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionStationAlarmRequest) ProtoMessage() {}

func (x *QueryProductionStationAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_station_alarm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionStationAlarmRequest.ProtoReflect.Descriptor instead.
func (*QueryProductionStationAlarmRequest) Descriptor() ([]byte, []int) {
	return file_production_station_alarm_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductionStationAlarmRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductionStationAlarmRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductionStationAlarmRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *QueryProductionStationAlarmRequest) GetAlarmMessage() string {
	if x != nil {
		return x.AlarmMessage
	}
	return ""
}

type QueryProductionStationAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionStationAlarmInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                         `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                         `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                         `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductionStationAlarmResponse) Reset() {
	*x = QueryProductionStationAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_station_alarm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionStationAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionStationAlarmResponse) ProtoMessage() {}

func (x *QueryProductionStationAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_station_alarm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionStationAlarmResponse.ProtoReflect.Descriptor instead.
func (*QueryProductionStationAlarmResponse) Descriptor() ([]byte, []int) {
	return file_production_station_alarm_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductionStationAlarmResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductionStationAlarmResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductionStationAlarmResponse) GetData() []*ProductionStationAlarmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductionStationAlarmResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductionStationAlarmResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductionStationAlarmResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductionStationAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionStationAlarmInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductionStationAlarmResponse) Reset() {
	*x = GetAllProductionStationAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_station_alarm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductionStationAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductionStationAlarmResponse) ProtoMessage() {}

func (x *GetAllProductionStationAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_station_alarm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductionStationAlarmResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductionStationAlarmResponse) Descriptor() ([]byte, []int) {
	return file_production_station_alarm_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductionStationAlarmResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductionStationAlarmResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductionStationAlarmResponse) GetData() []*ProductionStationAlarmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductionStationAlarmDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductionStationAlarmInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductionStationAlarmDetailResponse) Reset() {
	*x = GetProductionStationAlarmDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_station_alarm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductionStationAlarmDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionStationAlarmDetailResponse) ProtoMessage() {}

func (x *GetProductionStationAlarmDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_station_alarm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionStationAlarmDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductionStationAlarmDetailResponse) Descriptor() ([]byte, []int) {
	return file_production_station_alarm_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductionStationAlarmDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductionStationAlarmDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductionStationAlarmDetailResponse) GetData() *ProductionStationAlarmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_production_station_alarm_proto protoreflect.FileDescriptor

var file_production_station_alarm_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x05, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4e,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4e, 0x6f,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0xe4, 0x02, 0x0a, 0x22, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x30, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x28, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x23,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x98, 0x01, 0x0a, 0x24,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9b, 0x01, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_production_station_alarm_proto_rawDescOnce sync.Once
	file_production_station_alarm_proto_rawDescData = file_production_station_alarm_proto_rawDesc
)

func file_production_station_alarm_proto_rawDescGZIP() []byte {
	file_production_station_alarm_proto_rawDescOnce.Do(func() {
		file_production_station_alarm_proto_rawDescData = protoimpl.X.CompressGZIP(file_production_station_alarm_proto_rawDescData)
	})
	return file_production_station_alarm_proto_rawDescData
}

var file_production_station_alarm_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_production_station_alarm_proto_goTypes = []interface{}{
	(*ProductionStationAlarmInfo)(nil),              // 0: proto.ProductionStationAlarmInfo
	(*QueryProductionStationAlarmRequest)(nil),      // 1: proto.QueryProductionStationAlarmRequest
	(*QueryProductionStationAlarmResponse)(nil),     // 2: proto.QueryProductionStationAlarmResponse
	(*GetAllProductionStationAlarmResponse)(nil),    // 3: proto.GetAllProductionStationAlarmResponse
	(*GetProductionStationAlarmDetailResponse)(nil), // 4: proto.GetProductionStationAlarmDetailResponse
	(*ProductionStationInfo)(nil),                   // 5: proto.ProductionStationInfo
	(*ProductionProcessInfo)(nil),                   // 6: proto.ProductionProcessInfo
	(*ProductInfoInfo)(nil),                         // 7: proto.ProductInfoInfo
	(Code)(0),                                       // 8: proto.Code
}
var file_production_station_alarm_proto_depIdxs = []int32{
	5, // 0: proto.ProductionStationAlarmInfo.productionStation:type_name -> proto.ProductionStationInfo
	6, // 1: proto.ProductionStationAlarmInfo.productionProcess:type_name -> proto.ProductionProcessInfo
	7, // 2: proto.ProductionStationAlarmInfo.productInfo:type_name -> proto.ProductInfoInfo
	8, // 3: proto.QueryProductionStationAlarmResponse.code:type_name -> proto.Code
	0, // 4: proto.QueryProductionStationAlarmResponse.data:type_name -> proto.ProductionStationAlarmInfo
	8, // 5: proto.GetAllProductionStationAlarmResponse.code:type_name -> proto.Code
	0, // 6: proto.GetAllProductionStationAlarmResponse.data:type_name -> proto.ProductionStationAlarmInfo
	8, // 7: proto.GetProductionStationAlarmDetailResponse.code:type_name -> proto.Code
	0, // 8: proto.GetProductionStationAlarmDetailResponse.data:type_name -> proto.ProductionStationAlarmInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_production_station_alarm_proto_init() }
func file_production_station_alarm_proto_init() {
	if File_production_station_alarm_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_production_line_proto_init()
	file_production_process_proto_init()
	file_product_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_production_station_alarm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductionStationAlarmInfo); i {
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
		file_production_station_alarm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionStationAlarmRequest); i {
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
		file_production_station_alarm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionStationAlarmResponse); i {
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
		file_production_station_alarm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductionStationAlarmResponse); i {
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
		file_production_station_alarm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductionStationAlarmDetailResponse); i {
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
			RawDescriptor: file_production_station_alarm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_production_station_alarm_proto_goTypes,
		DependencyIndexes: file_production_station_alarm_proto_depIdxs,
		MessageInfos:      file_production_station_alarm_proto_msgTypes,
	}.Build()
	File_production_station_alarm_proto = out.File
	file_production_station_alarm_proto_rawDesc = nil
	file_production_station_alarm_proto_goTypes = nil
	file_production_station_alarm_proto_depIdxs = nil
}
