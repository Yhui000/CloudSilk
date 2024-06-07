// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: coding_serial.proto

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

type CodingSerialInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 前缀
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix"`
	// 种子
	Seed int32 `protobuf:"varint,3,opt,name=seed,proto3" json:"seed"`
	// 创建时间
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime"`
	// 最后增长时间
	LastIncreaseTime string `protobuf:"bytes,5,opt,name=lastIncreaseTime,proto3" json:"lastIncreaseTime"`
}

func (x *CodingSerialInfo) Reset() {
	*x = CodingSerialInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coding_serial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodingSerialInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodingSerialInfo) ProtoMessage() {}

func (x *CodingSerialInfo) ProtoReflect() protoreflect.Message {
	mi := &file_coding_serial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodingSerialInfo.ProtoReflect.Descriptor instead.
func (*CodingSerialInfo) Descriptor() ([]byte, []int) {
	return file_coding_serial_proto_rawDescGZIP(), []int{0}
}

func (x *CodingSerialInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CodingSerialInfo) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *CodingSerialInfo) GetSeed() int32 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *CodingSerialInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *CodingSerialInfo) GetLastIncreaseTime() string {
	if x != nil {
		return x.LastIncreaseTime
	}
	return ""
}

type QueryCodingSerialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 前缀
	// @inject_tag: uri:"prefix" form:"prefix"
	Prefix string `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix" uri:"prefix" form:"prefix"`
	// 种子
	// @inject_tag: uri:"seed" form:"seed"
	Seed int32 `protobuf:"varint,5,opt,name=seed,proto3" json:"seed" uri:"seed" form:"seed"`
	// 创建时间开始
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,6,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// 创建时间结束
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,7,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
}

func (x *QueryCodingSerialRequest) Reset() {
	*x = QueryCodingSerialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coding_serial_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCodingSerialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCodingSerialRequest) ProtoMessage() {}

func (x *QueryCodingSerialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coding_serial_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCodingSerialRequest.ProtoReflect.Descriptor instead.
func (*QueryCodingSerialRequest) Descriptor() ([]byte, []int) {
	return file_coding_serial_proto_rawDescGZIP(), []int{1}
}

func (x *QueryCodingSerialRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryCodingSerialRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryCodingSerialRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryCodingSerialRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *QueryCodingSerialRequest) GetSeed() int32 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *QueryCodingSerialRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryCodingSerialRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

type QueryCodingSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*CodingSerialInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64               `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64               `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64               `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryCodingSerialResponse) Reset() {
	*x = QueryCodingSerialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coding_serial_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCodingSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCodingSerialResponse) ProtoMessage() {}

func (x *QueryCodingSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coding_serial_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCodingSerialResponse.ProtoReflect.Descriptor instead.
func (*QueryCodingSerialResponse) Descriptor() ([]byte, []int) {
	return file_coding_serial_proto_rawDescGZIP(), []int{2}
}

func (x *QueryCodingSerialResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryCodingSerialResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryCodingSerialResponse) GetData() []*CodingSerialInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryCodingSerialResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryCodingSerialResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryCodingSerialResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllCodingSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*CodingSerialInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllCodingSerialResponse) Reset() {
	*x = GetAllCodingSerialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coding_serial_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCodingSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCodingSerialResponse) ProtoMessage() {}

func (x *GetAllCodingSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coding_serial_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCodingSerialResponse.ProtoReflect.Descriptor instead.
func (*GetAllCodingSerialResponse) Descriptor() ([]byte, []int) {
	return file_coding_serial_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCodingSerialResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllCodingSerialResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllCodingSerialResponse) GetData() []*CodingSerialInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCodingSerialDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code              `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *CodingSerialInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetCodingSerialDetailResponse) Reset() {
	*x = GetCodingSerialDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coding_serial_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodingSerialDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodingSerialDetailResponse) ProtoMessage() {}

func (x *GetCodingSerialDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coding_serial_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodingSerialDetailResponse.ProtoReflect.Descriptor instead.
func (*GetCodingSerialDetailResponse) Descriptor() ([]byte, []int) {
	return file_coding_serial_proto_rawDescGZIP(), []int{4}
}

func (x *GetCodingSerialDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetCodingSerialDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetCodingSerialDetailResponse) GetData() *CodingSerialInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_coding_serial_proto protoreflect.FileDescriptor

var file_coding_serial_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f,
	0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x01, 0x0a, 0x10, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x18,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x31, 0x22, 0xc9, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x84,
	0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_coding_serial_proto_rawDescOnce sync.Once
	file_coding_serial_proto_rawDescData = file_coding_serial_proto_rawDesc
)

func file_coding_serial_proto_rawDescGZIP() []byte {
	file_coding_serial_proto_rawDescOnce.Do(func() {
		file_coding_serial_proto_rawDescData = protoimpl.X.CompressGZIP(file_coding_serial_proto_rawDescData)
	})
	return file_coding_serial_proto_rawDescData
}

var file_coding_serial_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_coding_serial_proto_goTypes = []interface{}{
	(*CodingSerialInfo)(nil),              // 0: proto.CodingSerialInfo
	(*QueryCodingSerialRequest)(nil),      // 1: proto.QueryCodingSerialRequest
	(*QueryCodingSerialResponse)(nil),     // 2: proto.QueryCodingSerialResponse
	(*GetAllCodingSerialResponse)(nil),    // 3: proto.GetAllCodingSerialResponse
	(*GetCodingSerialDetailResponse)(nil), // 4: proto.GetCodingSerialDetailResponse
	(Code)(0),                             // 5: proto.Code
}
var file_coding_serial_proto_depIdxs = []int32{
	5, // 0: proto.QueryCodingSerialResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryCodingSerialResponse.data:type_name -> proto.CodingSerialInfo
	5, // 2: proto.GetAllCodingSerialResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllCodingSerialResponse.data:type_name -> proto.CodingSerialInfo
	5, // 4: proto.GetCodingSerialDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetCodingSerialDetailResponse.data:type_name -> proto.CodingSerialInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_coding_serial_proto_init() }
func file_coding_serial_proto_init() {
	if File_coding_serial_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coding_serial_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodingSerialInfo); i {
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
		file_coding_serial_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCodingSerialRequest); i {
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
		file_coding_serial_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCodingSerialResponse); i {
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
		file_coding_serial_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCodingSerialResponse); i {
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
		file_coding_serial_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodingSerialDetailResponse); i {
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
			RawDescriptor: file_coding_serial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coding_serial_proto_goTypes,
		DependencyIndexes: file_coding_serial_proto_depIdxs,
		MessageInfos:      file_coding_serial_proto_msgTypes,
	}.Build()
	File_coding_serial_proto = out.File
	file_coding_serial_proto_rawDesc = nil
	file_coding_serial_proto_goTypes = nil
	file_coding_serial_proto_depIdxs = nil
}
