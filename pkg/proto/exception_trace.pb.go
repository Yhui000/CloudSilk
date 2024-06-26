// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: exception_trace.proto

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

type ExceptionTraceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 主机
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host"`
	// 来源
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
	// 等级
	Level string `protobuf:"bytes,5,opt,name=level,proto3" json:"level"`
	// 异常信息
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message"`
	// 堆栈跟踪
	StackTrace string `protobuf:"bytes,7,opt,name=stackTrace,proto3" json:"stackTrace"`
	// 屏幕截图
	ScreenCapture string `protobuf:"bytes,8,opt,name=screenCapture,proto3" json:"screenCapture"`
	// 上报时间
	TimeReported string `protobuf:"bytes,9,opt,name=timeReported,proto3" json:"timeReported"`
	// 上报人员ID
	ReportUserID   string `protobuf:"bytes,10,opt,name=reportUserID,proto3" json:"reportUserID"`
	ReportUserName string `protobuf:"bytes,11,opt,name=reportUserName,proto3" json:"reportUserName"`
}

func (x *ExceptionTraceInfo) Reset() {
	*x = ExceptionTraceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exception_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExceptionTraceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExceptionTraceInfo) ProtoMessage() {}

func (x *ExceptionTraceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_exception_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExceptionTraceInfo.ProtoReflect.Descriptor instead.
func (*ExceptionTraceInfo) Descriptor() ([]byte, []int) {
	return file_exception_trace_proto_rawDescGZIP(), []int{0}
}

func (x *ExceptionTraceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExceptionTraceInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ExceptionTraceInfo) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ExceptionTraceInfo) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *ExceptionTraceInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExceptionTraceInfo) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *ExceptionTraceInfo) GetScreenCapture() string {
	if x != nil {
		return x.ScreenCapture
	}
	return ""
}

func (x *ExceptionTraceInfo) GetTimeReported() string {
	if x != nil {
		return x.TimeReported
	}
	return ""
}

func (x *ExceptionTraceInfo) GetReportUserID() string {
	if x != nil {
		return x.ReportUserID
	}
	return ""
}

func (x *ExceptionTraceInfo) GetReportUserName() string {
	if x != nil {
		return x.ReportUserName
	}
	return ""
}

type QueryExceptionTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 异常信息
	// @inject_tag: uri:"message" form:"message"
	Message string `protobuf:"bytes,8,opt,name=message,proto3" json:"message" uri:"message" form:"message"`
	// 上报时间开始
	// @inject_tag: uri:"timeReported0" form:"timeReported0"
	TimeReported0 string `protobuf:"bytes,11,opt,name=TimeReported0,proto3" json:"TimeReported0" uri:"timeReported0" form:"timeReported0"`
	// 上报时间结束
	// @inject_tag: uri:"timeReported1" form:"timeReported1"
	TimeReported1 string `protobuf:"bytes,12,opt,name=TimeReported1,proto3" json:"TimeReported1" uri:"timeReported1" form:"timeReported1"`
}

func (x *QueryExceptionTraceRequest) Reset() {
	*x = QueryExceptionTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exception_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryExceptionTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryExceptionTraceRequest) ProtoMessage() {}

func (x *QueryExceptionTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exception_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryExceptionTraceRequest.ProtoReflect.Descriptor instead.
func (*QueryExceptionTraceRequest) Descriptor() ([]byte, []int) {
	return file_exception_trace_proto_rawDescGZIP(), []int{1}
}

func (x *QueryExceptionTraceRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryExceptionTraceRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryExceptionTraceRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryExceptionTraceRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryExceptionTraceRequest) GetTimeReported0() string {
	if x != nil {
		return x.TimeReported0
	}
	return ""
}

func (x *QueryExceptionTraceRequest) GetTimeReported1() string {
	if x != nil {
		return x.TimeReported1
	}
	return ""
}

type QueryExceptionTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ExceptionTraceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                 `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                 `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryExceptionTraceResponse) Reset() {
	*x = QueryExceptionTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exception_trace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryExceptionTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryExceptionTraceResponse) ProtoMessage() {}

func (x *QueryExceptionTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exception_trace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryExceptionTraceResponse.ProtoReflect.Descriptor instead.
func (*QueryExceptionTraceResponse) Descriptor() ([]byte, []int) {
	return file_exception_trace_proto_rawDescGZIP(), []int{2}
}

func (x *QueryExceptionTraceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryExceptionTraceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryExceptionTraceResponse) GetData() []*ExceptionTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryExceptionTraceResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryExceptionTraceResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryExceptionTraceResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllExceptionTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ExceptionTraceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllExceptionTraceResponse) Reset() {
	*x = GetAllExceptionTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exception_trace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllExceptionTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExceptionTraceResponse) ProtoMessage() {}

func (x *GetAllExceptionTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exception_trace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExceptionTraceResponse.ProtoReflect.Descriptor instead.
func (*GetAllExceptionTraceResponse) Descriptor() ([]byte, []int) {
	return file_exception_trace_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllExceptionTraceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllExceptionTraceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllExceptionTraceResponse) GetData() []*ExceptionTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetExceptionTraceDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ExceptionTraceInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetExceptionTraceDetailResponse) Reset() {
	*x = GetExceptionTraceDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exception_trace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExceptionTraceDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExceptionTraceDetailResponse) ProtoMessage() {}

func (x *GetExceptionTraceDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exception_trace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExceptionTraceDetailResponse.ProtoReflect.Descriptor instead.
func (*GetExceptionTraceDetailResponse) Descriptor() ([]byte, []int) {
	return file_exception_trace_proto_rawDescGZIP(), []int{4}
}

func (x *GetExceptionTraceDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetExceptionTraceDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetExceptionTraceDetailResponse) GetData() *ExceptionTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_exception_trace_proto protoreflect.FileDescriptor

var file_exception_trace_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x02, 0x0a, 0x12, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x1a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x30, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x30, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x31, 0x22, 0xcd, 0x01, 0x0a, 0x1b, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exception_trace_proto_rawDescOnce sync.Once
	file_exception_trace_proto_rawDescData = file_exception_trace_proto_rawDesc
)

func file_exception_trace_proto_rawDescGZIP() []byte {
	file_exception_trace_proto_rawDescOnce.Do(func() {
		file_exception_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_exception_trace_proto_rawDescData)
	})
	return file_exception_trace_proto_rawDescData
}

var file_exception_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_exception_trace_proto_goTypes = []interface{}{
	(*ExceptionTraceInfo)(nil),              // 0: proto.ExceptionTraceInfo
	(*QueryExceptionTraceRequest)(nil),      // 1: proto.QueryExceptionTraceRequest
	(*QueryExceptionTraceResponse)(nil),     // 2: proto.QueryExceptionTraceResponse
	(*GetAllExceptionTraceResponse)(nil),    // 3: proto.GetAllExceptionTraceResponse
	(*GetExceptionTraceDetailResponse)(nil), // 4: proto.GetExceptionTraceDetailResponse
	(Code)(0),                               // 5: proto.Code
}
var file_exception_trace_proto_depIdxs = []int32{
	5, // 0: proto.QueryExceptionTraceResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryExceptionTraceResponse.data:type_name -> proto.ExceptionTraceInfo
	5, // 2: proto.GetAllExceptionTraceResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllExceptionTraceResponse.data:type_name -> proto.ExceptionTraceInfo
	5, // 4: proto.GetExceptionTraceDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetExceptionTraceDetailResponse.data:type_name -> proto.ExceptionTraceInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_exception_trace_proto_init() }
func file_exception_trace_proto_init() {
	if File_exception_trace_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exception_trace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExceptionTraceInfo); i {
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
		file_exception_trace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryExceptionTraceRequest); i {
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
		file_exception_trace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryExceptionTraceResponse); i {
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
		file_exception_trace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllExceptionTraceResponse); i {
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
		file_exception_trace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExceptionTraceDetailResponse); i {
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
			RawDescriptor: file_exception_trace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exception_trace_proto_goTypes,
		DependencyIndexes: file_exception_trace_proto_depIdxs,
		MessageInfos:      file_exception_trace_proto_msgTypes,
	}.Build()
	File_exception_trace_proto = out.File
	file_exception_trace_proto_rawDesc = nil
	file_exception_trace_proto_goTypes = nil
	file_exception_trace_proto_depIdxs = nil
}
