// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_rhythm_record.proto

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

type GetProductRhythmRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductionProcessID string `protobuf:"bytes,1,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	ProductInfoID       string `protobuf:"bytes,2,opt,name=productInfoID,proto3" json:"productInfoID"`
	ProductionStationID string `protobuf:"bytes,3,opt,name=productionStationID,proto3" json:"productionStationID"`
	HasWorkEndTime      bool   `protobuf:"varint,4,opt,name=hasWorkEndTime,proto3" json:"hasWorkEndTime"`
}

func (x *GetProductRhythmRecordRequest) Reset() {
	*x = GetProductRhythmRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRhythmRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRhythmRecordRequest) ProtoMessage() {}

func (x *GetProductRhythmRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRhythmRecordRequest.ProtoReflect.Descriptor instead.
func (*GetProductRhythmRecordRequest) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRhythmRecordRequest) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *GetProductRhythmRecordRequest) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *GetProductRhythmRecordRequest) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *GetProductRhythmRecordRequest) GetHasWorkEndTime() bool {
	if x != nil {
		return x.HasWorkEndTime
	}
	return false
}

type ProductRhythmRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 创建时间
	CreateTime string `protobuf:"bytes,3,opt,name=createTime,proto3" json:"createTime"`
	// 标准时长(秒)
	StandardWorkTime int32 `protobuf:"varint,4,opt,name=standardWorkTime,proto3" json:"standardWorkTime"`
	// 作业人员ID
	WorkUserID string `protobuf:"bytes,5,opt,name=workUserID,proto3" json:"workUserID"`
	// 开始作业时间
	WorkStartTime string `protobuf:"bytes,6,opt,name=workStartTime,proto3" json:"workStartTime"`
	// 等待时长(秒)
	WaitTime int32 `protobuf:"varint,7,opt,name=waitTime,proto3" json:"waitTime"`
	// 作业时长(秒)
	WorkTime int32 `protobuf:"varint,8,opt,name=workTime,proto3" json:"workTime"`
	// 超时时长(秒)
	OverTime int32 `protobuf:"varint,9,opt,name=overTime,proto3" json:"overTime"`
	// 结束作业时间
	WorkEndTime string `protobuf:"bytes,10,opt,name=workEndTime,proto3" json:"workEndTime"`
	// 是否返工
	IsRework bool `protobuf:"varint,11,opt,name=isRework,proto3" json:"isRework"`
	// 备注
	Remark string `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark"`
	// 生产工站ID
	ProductionStationID string `protobuf:"bytes,13,opt,name=productionStationID,proto3" json:"productionStationID"`
	// 生产工序ID
	ProductionProcessID string `protobuf:"bytes,14,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	// 产品信息ID
	ProductInfoID string `protobuf:"bytes,15,opt,name=productInfoID,proto3" json:"productInfoID"`
	// 生产工单号
	ProductOrderNo string `protobuf:"bytes,16,opt,name=productOrderNo,proto3" json:"productOrderNo"`
	// 产品序列号
	ProductSerialNo string `protobuf:"bytes,17,opt,name=productSerialNo,proto3" json:"productSerialNo"`
	// 生产工站
	ProductionStationDescription string `protobuf:"bytes,18,opt,name=productionStationDescription,proto3" json:"productionStationDescription"`
	// 生产工序
	ProductionProcessDescription string `protobuf:"bytes,19,opt,name=productionProcessDescription,proto3" json:"productionProcessDescription"`
}

func (x *ProductRhythmRecordInfo) Reset() {
	*x = ProductRhythmRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRhythmRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRhythmRecordInfo) ProtoMessage() {}

func (x *ProductRhythmRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRhythmRecordInfo.ProtoReflect.Descriptor instead.
func (*ProductRhythmRecordInfo) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{1}
}

func (x *ProductRhythmRecordInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetStandardWorkTime() int32 {
	if x != nil {
		return x.StandardWorkTime
	}
	return 0
}

func (x *ProductRhythmRecordInfo) GetWorkUserID() string {
	if x != nil {
		return x.WorkUserID
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetWorkStartTime() string {
	if x != nil {
		return x.WorkStartTime
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetWaitTime() int32 {
	if x != nil {
		return x.WaitTime
	}
	return 0
}

func (x *ProductRhythmRecordInfo) GetWorkTime() int32 {
	if x != nil {
		return x.WorkTime
	}
	return 0
}

func (x *ProductRhythmRecordInfo) GetOverTime() int32 {
	if x != nil {
		return x.OverTime
	}
	return 0
}

func (x *ProductRhythmRecordInfo) GetWorkEndTime() string {
	if x != nil {
		return x.WorkEndTime
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetIsRework() bool {
	if x != nil {
		return x.IsRework
	}
	return false
}

func (x *ProductRhythmRecordInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductionStationDescription() string {
	if x != nil {
		return x.ProductionStationDescription
	}
	return ""
}

func (x *ProductRhythmRecordInfo) GetProductionProcessDescription() string {
	if x != nil {
		return x.ProductionProcessDescription
	}
	return ""
}

type QueryProductRhythmRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// @inject_tag: uri:"productOrderNo" form:"productOrderNo"
	ProductOrderNo string `protobuf:"bytes,4,opt,name=productOrderNo,proto3" json:"productOrderNo" uri:"productOrderNo" form:"productOrderNo"`
	// @inject_tag: uri:"productSerialNo" form:"productSerialNo"
	ProductSerialNo string `protobuf:"bytes,5,opt,name=productSerialNo,proto3" json:"productSerialNo" uri:"productSerialNo" form:"productSerialNo"`
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,6,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,7,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
}

func (x *QueryProductRhythmRecordRequest) Reset() {
	*x = QueryProductRhythmRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductRhythmRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductRhythmRecordRequest) ProtoMessage() {}

func (x *QueryProductRhythmRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductRhythmRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryProductRhythmRecordRequest) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductRhythmRecordRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductRhythmRecordRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductRhythmRecordRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductRhythmRecordRequest) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *QueryProductRhythmRecordRequest) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *QueryProductRhythmRecordRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryProductRhythmRecordRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

type QueryProductRhythmRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductRhythmRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                      `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                      `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                      `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductRhythmRecordResponse) Reset() {
	*x = QueryProductRhythmRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductRhythmRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductRhythmRecordResponse) ProtoMessage() {}

func (x *QueryProductRhythmRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductRhythmRecordResponse.ProtoReflect.Descriptor instead.
func (*QueryProductRhythmRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductRhythmRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductRhythmRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductRhythmRecordResponse) GetData() []*ProductRhythmRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductRhythmRecordResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductRhythmRecordResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductRhythmRecordResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductRhythmRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductRhythmRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductRhythmRecordResponse) Reset() {
	*x = GetAllProductRhythmRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductRhythmRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductRhythmRecordResponse) ProtoMessage() {}

func (x *GetAllProductRhythmRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductRhythmRecordResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductRhythmRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductRhythmRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductRhythmRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductRhythmRecordResponse) GetData() []*ProductRhythmRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductRhythmRecordDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductRhythmRecordInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductRhythmRecordDetailResponse) Reset() {
	*x = GetProductRhythmRecordDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rhythm_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRhythmRecordDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRhythmRecordDetailResponse) ProtoMessage() {}

func (x *GetProductRhythmRecordDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rhythm_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRhythmRecordDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductRhythmRecordDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_rhythm_record_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductRhythmRecordDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductRhythmRecordDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductRhythmRecordDetailResponse) GetData() *ProductRhythmRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_rhythm_record_proto protoreflect.FileDescriptor

var file_product_rhythm_record_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x68, 0x79, 0x74, 0x68, 0x6d,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44,
	0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x57,
	0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc9, 0x05, 0x0a, 0x17, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x69, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x61, 0x69, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x42,
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x02, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x28, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x30, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x22, 0xd7, 0x01, 0x0a, 0x20, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68,
	0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x92, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x24, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68,
	0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xb1, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79,
	0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_rhythm_record_proto_rawDescOnce sync.Once
	file_product_rhythm_record_proto_rawDescData = file_product_rhythm_record_proto_rawDesc
)

func file_product_rhythm_record_proto_rawDescGZIP() []byte {
	file_product_rhythm_record_proto_rawDescOnce.Do(func() {
		file_product_rhythm_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_rhythm_record_proto_rawDescData)
	})
	return file_product_rhythm_record_proto_rawDescData
}

var file_product_rhythm_record_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_rhythm_record_proto_goTypes = []interface{}{
	(*GetProductRhythmRecordRequest)(nil),        // 0: proto.GetProductRhythmRecordRequest
	(*ProductRhythmRecordInfo)(nil),              // 1: proto.ProductRhythmRecordInfo
	(*QueryProductRhythmRecordRequest)(nil),      // 2: proto.QueryProductRhythmRecordRequest
	(*QueryProductRhythmRecordResponse)(nil),     // 3: proto.QueryProductRhythmRecordResponse
	(*GetAllProductRhythmRecordResponse)(nil),    // 4: proto.GetAllProductRhythmRecordResponse
	(*GetProductRhythmRecordDetailResponse)(nil), // 5: proto.GetProductRhythmRecordDetailResponse
	(Code)(0),              // 6: proto.Code
	(*CommonResponse)(nil), // 7: proto.CommonResponse
}
var file_product_rhythm_record_proto_depIdxs = []int32{
	6, // 0: proto.QueryProductRhythmRecordResponse.code:type_name -> proto.Code
	1, // 1: proto.QueryProductRhythmRecordResponse.data:type_name -> proto.ProductRhythmRecordInfo
	6, // 2: proto.GetAllProductRhythmRecordResponse.code:type_name -> proto.Code
	1, // 3: proto.GetAllProductRhythmRecordResponse.data:type_name -> proto.ProductRhythmRecordInfo
	6, // 4: proto.GetProductRhythmRecordDetailResponse.code:type_name -> proto.Code
	1, // 5: proto.GetProductRhythmRecordDetailResponse.data:type_name -> proto.ProductRhythmRecordInfo
	1, // 6: proto.ProductRhythmRecord.Add:input_type -> proto.ProductRhythmRecordInfo
	0, // 7: proto.ProductRhythmRecord.Get:input_type -> proto.GetProductRhythmRecordRequest
	7, // 8: proto.ProductRhythmRecord.Add:output_type -> proto.CommonResponse
	5, // 9: proto.ProductRhythmRecord.Get:output_type -> proto.GetProductRhythmRecordDetailResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_product_rhythm_record_proto_init() }
func file_product_rhythm_record_proto_init() {
	if File_product_rhythm_record_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_rhythm_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRhythmRecordRequest); i {
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
		file_product_rhythm_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRhythmRecordInfo); i {
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
		file_product_rhythm_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductRhythmRecordRequest); i {
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
		file_product_rhythm_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductRhythmRecordResponse); i {
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
		file_product_rhythm_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductRhythmRecordResponse); i {
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
		file_product_rhythm_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRhythmRecordDetailResponse); i {
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
			RawDescriptor: file_product_rhythm_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_rhythm_record_proto_goTypes,
		DependencyIndexes: file_product_rhythm_record_proto_depIdxs,
		MessageInfos:      file_product_rhythm_record_proto_msgTypes,
	}.Build()
	File_product_rhythm_record_proto = out.File
	file_product_rhythm_record_proto_rawDesc = nil
	file_product_rhythm_record_proto_goTypes = nil
	file_product_rhythm_record_proto_depIdxs = nil
}