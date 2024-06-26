// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: print_server.proto

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

type PrintServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 打印机
	Printers []*PrinterInfo `protobuf:"bytes,1,rep,name=printers,proto3" json:"printers"`
	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 身份标识
	Identity string `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity"`
	// 运行状态
	RunningState string `protobuf:"bytes,5,opt,name=runningState,proto3" json:"runningState"`
}

func (x *PrintServerInfo) Reset() {
	*x = PrintServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintServerInfo) ProtoMessage() {}

func (x *PrintServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintServerInfo.ProtoReflect.Descriptor instead.
func (*PrintServerInfo) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{0}
}

func (x *PrintServerInfo) GetPrinters() []*PrinterInfo {
	if x != nil {
		return x.Printers
	}
	return nil
}

func (x *PrintServerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrintServerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrintServerInfo) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *PrintServerInfo) GetRunningState() string {
	if x != nil {
		return x.RunningState
	}
	return ""
}

type PrinterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 是否启用
	Enable bool `protobuf:"varint,4,opt,name=enable,proto3" json:"enable"`
	// 打印服务器ID
	PrintServerID string           `protobuf:"bytes,5,opt,name=printServerID,proto3" json:"printServerID"`
	PrintServer   *PrintServerInfo `protobuf:"bytes,6,opt,name=printServer,proto3" json:"printServer"`
}

func (x *PrinterInfo) Reset() {
	*x = PrinterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrinterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterInfo) ProtoMessage() {}

func (x *PrinterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterInfo.ProtoReflect.Descriptor instead.
func (*PrinterInfo) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{1}
}

func (x *PrinterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrinterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrinterInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *PrinterInfo) GetPrintServerID() string {
	if x != nil {
		return x.PrintServerID
	}
	return ""
}

func (x *PrinterInfo) GetPrintServer() *PrintServerInfo {
	if x != nil {
		return x.PrintServer
	}
	return nil
}

type QueryPrintServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
}

func (x *QueryPrintServerRequest) Reset() {
	*x = QueryPrintServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPrintServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPrintServerRequest) ProtoMessage() {}

func (x *QueryPrintServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPrintServerRequest.ProtoReflect.Descriptor instead.
func (*QueryPrintServerRequest) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPrintServerRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryPrintServerRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryPrintServerRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

type QueryPrintServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code               `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*PrintServerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64              `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64              `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64              `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryPrintServerResponse) Reset() {
	*x = QueryPrintServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPrintServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPrintServerResponse) ProtoMessage() {}

func (x *QueryPrintServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPrintServerResponse.ProtoReflect.Descriptor instead.
func (*QueryPrintServerResponse) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{3}
}

func (x *QueryPrintServerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryPrintServerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryPrintServerResponse) GetData() []*PrintServerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryPrintServerResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryPrintServerResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryPrintServerResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllPrintServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code               `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*PrintServerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllPrintServerResponse) Reset() {
	*x = GetAllPrintServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPrintServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPrintServerResponse) ProtoMessage() {}

func (x *GetAllPrintServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPrintServerResponse.ProtoReflect.Descriptor instead.
func (*GetAllPrintServerResponse) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllPrintServerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllPrintServerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllPrintServerResponse) GetData() []*PrintServerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllPrinterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code           `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*PrinterInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllPrinterResponse) Reset() {
	*x = GetAllPrinterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPrinterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPrinterResponse) ProtoMessage() {}

func (x *GetAllPrinterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPrinterResponse.ProtoReflect.Descriptor instead.
func (*GetAllPrinterResponse) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllPrinterResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllPrinterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllPrinterResponse) GetData() []*PrinterInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPrintServerDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code             `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *PrintServerInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetPrintServerDetailResponse) Reset() {
	*x = GetPrintServerDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrintServerDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrintServerDetailResponse) ProtoMessage() {}

func (x *GetPrintServerDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_print_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrintServerDetailResponse.ProtoReflect.Descriptor instead.
func (*GetPrintServerDetailResponse) Descriptor() ([]byte, []int) {
	return file_print_server_proto_rawDescGZIP(), []int{6}
}

func (x *GetPrintServerDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetPrintServerDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPrintServerDetailResponse) GetData() *PrintServerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_print_server_proto protoreflect.FileDescriptor

var file_print_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01,
	0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x73, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xc7, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x82, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x85, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_print_server_proto_rawDescOnce sync.Once
	file_print_server_proto_rawDescData = file_print_server_proto_rawDesc
)

func file_print_server_proto_rawDescGZIP() []byte {
	file_print_server_proto_rawDescOnce.Do(func() {
		file_print_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_print_server_proto_rawDescData)
	})
	return file_print_server_proto_rawDescData
}

var file_print_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_print_server_proto_goTypes = []interface{}{
	(*PrintServerInfo)(nil),              // 0: proto.PrintServerInfo
	(*PrinterInfo)(nil),                  // 1: proto.PrinterInfo
	(*QueryPrintServerRequest)(nil),      // 2: proto.QueryPrintServerRequest
	(*QueryPrintServerResponse)(nil),     // 3: proto.QueryPrintServerResponse
	(*GetAllPrintServerResponse)(nil),    // 4: proto.GetAllPrintServerResponse
	(*GetAllPrinterResponse)(nil),        // 5: proto.GetAllPrinterResponse
	(*GetPrintServerDetailResponse)(nil), // 6: proto.GetPrintServerDetailResponse
	(Code)(0),                            // 7: proto.Code
}
var file_print_server_proto_depIdxs = []int32{
	1,  // 0: proto.PrintServerInfo.printers:type_name -> proto.PrinterInfo
	0,  // 1: proto.PrinterInfo.printServer:type_name -> proto.PrintServerInfo
	7,  // 2: proto.QueryPrintServerResponse.code:type_name -> proto.Code
	0,  // 3: proto.QueryPrintServerResponse.data:type_name -> proto.PrintServerInfo
	7,  // 4: proto.GetAllPrintServerResponse.code:type_name -> proto.Code
	0,  // 5: proto.GetAllPrintServerResponse.data:type_name -> proto.PrintServerInfo
	7,  // 6: proto.GetAllPrinterResponse.code:type_name -> proto.Code
	1,  // 7: proto.GetAllPrinterResponse.data:type_name -> proto.PrinterInfo
	7,  // 8: proto.GetPrintServerDetailResponse.code:type_name -> proto.Code
	0,  // 9: proto.GetPrintServerDetailResponse.data:type_name -> proto.PrintServerInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_print_server_proto_init() }
func file_print_server_proto_init() {
	if File_print_server_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_print_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintServerInfo); i {
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
		file_print_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrinterInfo); i {
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
		file_print_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPrintServerRequest); i {
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
		file_print_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPrintServerResponse); i {
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
		file_print_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPrintServerResponse); i {
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
		file_print_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPrinterResponse); i {
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
		file_print_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrintServerDetailResponse); i {
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
			RawDescriptor: file_print_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_print_server_proto_goTypes,
		DependencyIndexes: file_print_server_proto_depIdxs,
		MessageInfos:      file_print_server_proto_msgTypes,
	}.Build()
	File_print_server_proto = out.File
	file_print_server_proto_rawDesc = nil
	file_print_server_proto_goTypes = nil
	file_print_server_proto_depIdxs = nil
}
