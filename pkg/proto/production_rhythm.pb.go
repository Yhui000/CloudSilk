// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: production_rhythm.proto

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

type ProductionRhythmInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
	// 优先级
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority"`
	// 标准时长(秒)
	StandardTime int32 `protobuf:"varint,5,opt,name=standardTime,proto3" json:"standardTime"`
	// 是否启用
	Enable bool `protobuf:"varint,6,opt,name=enable,proto3" json:"enable"`
	// 默认匹配
	InitialValue bool `protobuf:"varint,7,opt,name=initialValue,proto3" json:"initialValue"`
	// 备注
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	// 生产产线ID
	ProductionLineID string `protobuf:"bytes,9,opt,name=productionLineID,proto3" json:"productionLineID"`
	// 特性表达式
	AttributeExpressions []*AttributeExpressionInfo `protobuf:"bytes,2,rep,name=attributeExpressions,proto3" json:"attributeExpressions"`
}

func (x *ProductionRhythmInfo) Reset() {
	*x = ProductionRhythmInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_rhythm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductionRhythmInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductionRhythmInfo) ProtoMessage() {}

func (x *ProductionRhythmInfo) ProtoReflect() protoreflect.Message {
	mi := &file_production_rhythm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductionRhythmInfo.ProtoReflect.Descriptor instead.
func (*ProductionRhythmInfo) Descriptor() ([]byte, []int) {
	return file_production_rhythm_proto_rawDescGZIP(), []int{0}
}

func (x *ProductionRhythmInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductionRhythmInfo) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ProductionRhythmInfo) GetStandardTime() int32 {
	if x != nil {
		return x.StandardTime
	}
	return 0
}

func (x *ProductionRhythmInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ProductionRhythmInfo) GetInitialValue() bool {
	if x != nil {
		return x.InitialValue
	}
	return false
}

func (x *ProductionRhythmInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductionRhythmInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *ProductionRhythmInfo) GetAttributeExpressions() []*AttributeExpressionInfo {
	if x != nil {
		return x.AttributeExpressions
	}
	return nil
}

type QueryProductionRhythmRequest struct {
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
	ProductionLineID string `protobuf:"bytes,11,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
}

func (x *QueryProductionRhythmRequest) Reset() {
	*x = QueryProductionRhythmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_rhythm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionRhythmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionRhythmRequest) ProtoMessage() {}

func (x *QueryProductionRhythmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_production_rhythm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionRhythmRequest.ProtoReflect.Descriptor instead.
func (*QueryProductionRhythmRequest) Descriptor() ([]byte, []int) {
	return file_production_rhythm_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductionRhythmRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductionRhythmRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductionRhythmRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductionRhythmRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

type QueryProductionRhythmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionRhythmInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                   `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                   `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                   `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductionRhythmResponse) Reset() {
	*x = QueryProductionRhythmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_rhythm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductionRhythmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductionRhythmResponse) ProtoMessage() {}

func (x *QueryProductionRhythmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_rhythm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductionRhythmResponse.ProtoReflect.Descriptor instead.
func (*QueryProductionRhythmResponse) Descriptor() ([]byte, []int) {
	return file_production_rhythm_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductionRhythmResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductionRhythmResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductionRhythmResponse) GetData() []*ProductionRhythmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductionRhythmResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductionRhythmResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductionRhythmResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductionRhythmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductionRhythmInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductionRhythmResponse) Reset() {
	*x = GetAllProductionRhythmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_rhythm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductionRhythmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductionRhythmResponse) ProtoMessage() {}

func (x *GetAllProductionRhythmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_rhythm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductionRhythmResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductionRhythmResponse) Descriptor() ([]byte, []int) {
	return file_production_rhythm_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductionRhythmResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductionRhythmResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductionRhythmResponse) GetData() []*ProductionRhythmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductionRhythmDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductionRhythmInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductionRhythmDetailResponse) Reset() {
	*x = GetProductionRhythmDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_production_rhythm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductionRhythmDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductionRhythmDetailResponse) ProtoMessage() {}

func (x *GetProductionRhythmDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_production_rhythm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductionRhythmDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductionRhythmDetailResponse) Descriptor() ([]byte, []int) {
	return file_production_rhythm_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductionRhythmDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductionRhythmDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductionRhythmDetailResponse) GetData() *ProductionRhythmInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_production_rhythm_proto protoreflect.FileDescriptor

var file_production_rhythm_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x68, 0x79,
	0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x02, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79,
	0x74, 0x68, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x52, 0x0a, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x1c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x44, 0x22, 0xd1, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79, 0x74, 0x68,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x68, 0x79, 0x74, 0x68, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_production_rhythm_proto_rawDescOnce sync.Once
	file_production_rhythm_proto_rawDescData = file_production_rhythm_proto_rawDesc
)

func file_production_rhythm_proto_rawDescGZIP() []byte {
	file_production_rhythm_proto_rawDescOnce.Do(func() {
		file_production_rhythm_proto_rawDescData = protoimpl.X.CompressGZIP(file_production_rhythm_proto_rawDescData)
	})
	return file_production_rhythm_proto_rawDescData
}

var file_production_rhythm_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_production_rhythm_proto_goTypes = []interface{}{
	(*ProductionRhythmInfo)(nil),              // 0: proto.ProductionRhythmInfo
	(*QueryProductionRhythmRequest)(nil),      // 1: proto.QueryProductionRhythmRequest
	(*QueryProductionRhythmResponse)(nil),     // 2: proto.QueryProductionRhythmResponse
	(*GetAllProductionRhythmResponse)(nil),    // 3: proto.GetAllProductionRhythmResponse
	(*GetProductionRhythmDetailResponse)(nil), // 4: proto.GetProductionRhythmDetailResponse
	(*AttributeExpressionInfo)(nil),           // 5: proto.AttributeExpressionInfo
	(Code)(0),                                 // 6: proto.Code
}
var file_production_rhythm_proto_depIdxs = []int32{
	5, // 0: proto.ProductionRhythmInfo.attributeExpressions:type_name -> proto.AttributeExpressionInfo
	6, // 1: proto.QueryProductionRhythmResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductionRhythmResponse.data:type_name -> proto.ProductionRhythmInfo
	6, // 3: proto.GetAllProductionRhythmResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductionRhythmResponse.data:type_name -> proto.ProductionRhythmInfo
	6, // 5: proto.GetProductionRhythmDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductionRhythmDetailResponse.data:type_name -> proto.ProductionRhythmInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_production_rhythm_proto_init() }
func file_production_rhythm_proto_init() {
	if File_production_rhythm_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_attribute_expression_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_production_rhythm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductionRhythmInfo); i {
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
		file_production_rhythm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionRhythmRequest); i {
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
		file_production_rhythm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductionRhythmResponse); i {
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
		file_production_rhythm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductionRhythmResponse); i {
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
		file_production_rhythm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductionRhythmDetailResponse); i {
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
			RawDescriptor: file_production_rhythm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_production_rhythm_proto_goTypes,
		DependencyIndexes: file_production_rhythm_proto_depIdxs,
		MessageInfos:      file_production_rhythm_proto_msgTypes,
	}.Build()
	File_production_rhythm_proto = out.File
	file_production_rhythm_proto_rawDesc = nil
	file_production_rhythm_proto_goTypes = nil
	file_production_rhythm_proto_depIdxs = nil
}
