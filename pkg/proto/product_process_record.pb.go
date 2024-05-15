// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_process_record.proto

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

type ProductProcessRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 作业类型
	ProcessStepType string `protobuf:"bytes,3,opt,name=processStepType,proto3" json:"processStepType"`
	// 作业描述
	WorkDescription string `protobuf:"bytes,4,opt,name=workDescription,proto3" json:"workDescription"`
	// 作业数据
	WorkData string `protobuf:"bytes,5,opt,name=workData,proto3" json:"workData"`
	// 作业结果
	WorkResult string `protobuf:"bytes,6,opt,name=workResult,proto3" json:"workResult"`
	// 作业时间
	WorkTime string `protobuf:"bytes,7,opt,name=workTime,proto3" json:"workTime"`
	// 作业人员ID
	WorkUserID string `protobuf:"bytes,8,opt,name=workUserID,proto3" json:"workUserID"`
	// 备注
	Remark string `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark"`
	// 生产工站ID
	ProductionStationID string `protobuf:"bytes,10,opt,name=productionStationID,proto3" json:"productionStationID"`
	// 生产工序ID
	ProductionProcessID string `protobuf:"bytes,11,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	// 产品信息ID
	ProductInfoID string `protobuf:"bytes,12,opt,name=productInfoID,proto3" json:"productInfoID"`
	// 产品序列号
	ProductSerialNo string `protobuf:"bytes,13,opt,name=productSerialNo,proto3" json:"productSerialNo"`
	// 工站
	ProductionStation string `protobuf:"bytes,14,opt,name=productionStation,proto3" json:"productionStation"`
	// 生产工单号
	ProductOrderNo string `protobuf:"bytes,15,opt,name=productOrderNo,proto3" json:"productOrderNo"`
}

func (x *ProductProcessRecordInfo) Reset() {
	*x = ProductProcessRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductProcessRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductProcessRecordInfo) ProtoMessage() {}

func (x *ProductProcessRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductProcessRecordInfo.ProtoReflect.Descriptor instead.
func (*ProductProcessRecordInfo) Descriptor() ([]byte, []int) {
	return file_product_process_record_proto_rawDescGZIP(), []int{0}
}

func (x *ProductProcessRecordInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProcessStepType() string {
	if x != nil {
		return x.ProcessStepType
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetWorkDescription() string {
	if x != nil {
		return x.WorkDescription
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetWorkData() string {
	if x != nil {
		return x.WorkData
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetWorkResult() string {
	if x != nil {
		return x.WorkResult
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetWorkTime() string {
	if x != nil {
		return x.WorkTime
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetWorkUserID() string {
	if x != nil {
		return x.WorkUserID
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductionStation() string {
	if x != nil {
		return x.ProductionStation
	}
	return ""
}

func (x *ProductProcessRecordInfo) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

type QueryProductProcessRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 产品序列号
	// @inject_tag: uri:"productSerialNo" form:"productSerialNo"
	ProductSerialNo string `protobuf:"bytes,4,opt,name=productSerialNo,proto3" json:"productSerialNo" uri:"productSerialNo" form:"productSerialNo"`
	// 生产工单号
	// @inject_tag: uri:"productOrderNo" form:"productOrderNo"
	ProductOrderNo string `protobuf:"bytes,5,opt,name=productOrderNo,proto3" json:"productOrderNo" uri:"productOrderNo" form:"productOrderNo"`
	// @inject_tag: uri:"workTime0" form:"workTime0"
	WorkTime0 string `protobuf:"bytes,6,opt,name=workTime0,proto3" json:"workTime0" uri:"workTime0" form:"workTime0"`
	// @inject_tag: uri:"workTime1" form:"workTime1"
	WorkTime1 string `protobuf:"bytes,7,opt,name=workTime1,proto3" json:"workTime1" uri:"workTime1" form:"workTime1"`
	// @inject_tag: uri:"workDescription" form:"workDescription"
	WorkDescription string `protobuf:"bytes,8,opt,name=workDescription,proto3" json:"workDescription" uri:"workDescription" form:"workDescription"`
}

func (x *QueryProductProcessRecordRequest) Reset() {
	*x = QueryProductProcessRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductProcessRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductProcessRecordRequest) ProtoMessage() {}

func (x *QueryProductProcessRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductProcessRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryProductProcessRecordRequest) Descriptor() ([]byte, []int) {
	return file_product_process_record_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductProcessRecordRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductProcessRecordRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductProcessRecordRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductProcessRecordRequest) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *QueryProductProcessRecordRequest) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *QueryProductProcessRecordRequest) GetWorkTime0() string {
	if x != nil {
		return x.WorkTime0
	}
	return ""
}

func (x *QueryProductProcessRecordRequest) GetWorkTime1() string {
	if x != nil {
		return x.WorkTime1
	}
	return ""
}

func (x *QueryProductProcessRecordRequest) GetWorkDescription() string {
	if x != nil {
		return x.WorkDescription
	}
	return ""
}

type QueryProductProcessRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductProcessRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                       `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                       `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductProcessRecordResponse) Reset() {
	*x = QueryProductProcessRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductProcessRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductProcessRecordResponse) ProtoMessage() {}

func (x *QueryProductProcessRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductProcessRecordResponse.ProtoReflect.Descriptor instead.
func (*QueryProductProcessRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_process_record_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductProcessRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductProcessRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductProcessRecordResponse) GetData() []*ProductProcessRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductProcessRecordResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductProcessRecordResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductProcessRecordResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductProcessRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductProcessRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductProcessRecordResponse) Reset() {
	*x = GetAllProductProcessRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductProcessRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductProcessRecordResponse) ProtoMessage() {}

func (x *GetAllProductProcessRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductProcessRecordResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductProcessRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_process_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductProcessRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductProcessRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductProcessRecordResponse) GetData() []*ProductProcessRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductProcessRecordDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductProcessRecordInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductProcessRecordDetailResponse) Reset() {
	*x = GetProductProcessRecordDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductProcessRecordDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductProcessRecordDetailResponse) ProtoMessage() {}

func (x *GetProductProcessRecordDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductProcessRecordDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductProcessRecordDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_process_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductProcessRecordDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductProcessRecordDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductProcessRecordDetailResponse) GetData() *ProductProcessRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_process_record_proto protoreflect.FileDescriptor

var file_product_process_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x04, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x2c, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x30,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x30, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x12,
	0x28, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x21, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x94, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a,
	0x25, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_process_record_proto_rawDescOnce sync.Once
	file_product_process_record_proto_rawDescData = file_product_process_record_proto_rawDesc
)

func file_product_process_record_proto_rawDescGZIP() []byte {
	file_product_process_record_proto_rawDescOnce.Do(func() {
		file_product_process_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_process_record_proto_rawDescData)
	})
	return file_product_process_record_proto_rawDescData
}

var file_product_process_record_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_process_record_proto_goTypes = []interface{}{
	(*ProductProcessRecordInfo)(nil),              // 0: proto.ProductProcessRecordInfo
	(*QueryProductProcessRecordRequest)(nil),      // 1: proto.QueryProductProcessRecordRequest
	(*QueryProductProcessRecordResponse)(nil),     // 2: proto.QueryProductProcessRecordResponse
	(*GetAllProductProcessRecordResponse)(nil),    // 3: proto.GetAllProductProcessRecordResponse
	(*GetProductProcessRecordDetailResponse)(nil), // 4: proto.GetProductProcessRecordDetailResponse
	(Code)(0), // 5: proto.Code
}
var file_product_process_record_proto_depIdxs = []int32{
	5, // 0: proto.QueryProductProcessRecordResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryProductProcessRecordResponse.data:type_name -> proto.ProductProcessRecordInfo
	5, // 2: proto.GetAllProductProcessRecordResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllProductProcessRecordResponse.data:type_name -> proto.ProductProcessRecordInfo
	5, // 4: proto.GetProductProcessRecordDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetProductProcessRecordDetailResponse.data:type_name -> proto.ProductProcessRecordInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_product_process_record_proto_init() }
func file_product_process_record_proto_init() {
	if File_product_process_record_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_process_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductProcessRecordInfo); i {
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
		file_product_process_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductProcessRecordRequest); i {
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
		file_product_process_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductProcessRecordResponse); i {
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
		file_product_process_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductProcessRecordResponse); i {
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
		file_product_process_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductProcessRecordDetailResponse); i {
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
			RawDescriptor: file_product_process_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_process_record_proto_goTypes,
		DependencyIndexes: file_product_process_record_proto_depIdxs,
		MessageInfos:      file_product_process_record_proto_msgTypes,
	}.Build()
	File_product_process_record_proto = out.File
	file_product_process_record_proto_rawDesc = nil
	file_product_process_record_proto_goTypes = nil
	file_product_process_record_proto_depIdxs = nil
}