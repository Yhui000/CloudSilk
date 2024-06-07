// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: operation_trace.proto

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

type OperationTraceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 操作人员ID
	OperateUserID   string `protobuf:"bytes,3,opt,name=operateUserID,proto3" json:"operateUserID"`
	OperateUserName string `protobuf:"bytes,10,opt,name=operateUserName,proto3" json:"operateUserName"`
	// IP地址
	IPAddress string `protobuf:"bytes,4,opt,name=iPAddress,proto3" json:"iPAddress"`
	// 操作时间
	OperateTime string `protobuf:"bytes,5,opt,name=operateTime,proto3" json:"operateTime"`
	// 模块
	ControllerName string `protobuf:"bytes,6,opt,name=controllerName,proto3" json:"controllerName"`
	// 操作
	ActionName string `protobuf:"bytes,7,opt,name=actionName,proto3" json:"actionName"`
	// 提交内容
	RequestContent string `protobuf:"bytes,8,opt,name=requestContent,proto3" json:"requestContent"`
	// 注释
	Annotation string `protobuf:"bytes,9,opt,name=annotation,proto3" json:"annotation"`
}

func (x *OperationTraceInfo) Reset() {
	*x = OperationTraceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationTraceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationTraceInfo) ProtoMessage() {}

func (x *OperationTraceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_operation_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationTraceInfo.ProtoReflect.Descriptor instead.
func (*OperationTraceInfo) Descriptor() ([]byte, []int) {
	return file_operation_trace_proto_rawDescGZIP(), []int{0}
}

func (x *OperationTraceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OperationTraceInfo) GetOperateUserID() string {
	if x != nil {
		return x.OperateUserID
	}
	return ""
}

func (x *OperationTraceInfo) GetOperateUserName() string {
	if x != nil {
		return x.OperateUserName
	}
	return ""
}

func (x *OperationTraceInfo) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

func (x *OperationTraceInfo) GetOperateTime() string {
	if x != nil {
		return x.OperateTime
	}
	return ""
}

func (x *OperationTraceInfo) GetControllerName() string {
	if x != nil {
		return x.ControllerName
	}
	return ""
}

func (x *OperationTraceInfo) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *OperationTraceInfo) GetRequestContent() string {
	if x != nil {
		return x.RequestContent
	}
	return ""
}

func (x *OperationTraceInfo) GetAnnotation() string {
	if x != nil {
		return x.Annotation
	}
	return ""
}

type QueryOperationTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 操作时间开始
	// @inject_tag: uri:"operateTime0" form:"operateTime0"
	OperateTime0 string `protobuf:"bytes,8,opt,name=operateTime0,proto3" json:"operateTime0" uri:"operateTime0" form:"operateTime0"`
	// 操作时间结束
	// @inject_tag: uri:"operateTime1" form:"operateTime1"
	OperateTime1 string `protobuf:"bytes,9,opt,name=operateTime1,proto3" json:"operateTime1" uri:"operateTime1" form:"operateTime1"`
}

func (x *QueryOperationTraceRequest) Reset() {
	*x = QueryOperationTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOperationTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOperationTraceRequest) ProtoMessage() {}

func (x *QueryOperationTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operation_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOperationTraceRequest.ProtoReflect.Descriptor instead.
func (*QueryOperationTraceRequest) Descriptor() ([]byte, []int) {
	return file_operation_trace_proto_rawDescGZIP(), []int{1}
}

func (x *QueryOperationTraceRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryOperationTraceRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryOperationTraceRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryOperationTraceRequest) GetOperateTime0() string {
	if x != nil {
		return x.OperateTime0
	}
	return ""
}

func (x *QueryOperationTraceRequest) GetOperateTime1() string {
	if x != nil {
		return x.OperateTime1
	}
	return ""
}

type QueryOperationTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*OperationTraceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                 `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                 `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryOperationTraceResponse) Reset() {
	*x = QueryOperationTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_trace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOperationTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOperationTraceResponse) ProtoMessage() {}

func (x *QueryOperationTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operation_trace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOperationTraceResponse.ProtoReflect.Descriptor instead.
func (*QueryOperationTraceResponse) Descriptor() ([]byte, []int) {
	return file_operation_trace_proto_rawDescGZIP(), []int{2}
}

func (x *QueryOperationTraceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryOperationTraceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryOperationTraceResponse) GetData() []*OperationTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryOperationTraceResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryOperationTraceResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryOperationTraceResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllOperationTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*OperationTraceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllOperationTraceResponse) Reset() {
	*x = GetAllOperationTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_trace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOperationTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOperationTraceResponse) ProtoMessage() {}

func (x *GetAllOperationTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operation_trace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOperationTraceResponse.ProtoReflect.Descriptor instead.
func (*GetAllOperationTraceResponse) Descriptor() ([]byte, []int) {
	return file_operation_trace_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllOperationTraceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllOperationTraceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllOperationTraceResponse) GetData() []*OperationTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetOperationTraceDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *OperationTraceInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetOperationTraceDetailResponse) Reset() {
	*x = GetOperationTraceDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_trace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationTraceDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationTraceDetailResponse) ProtoMessage() {}

func (x *GetOperationTraceDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operation_trace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationTraceDetailResponse.ProtoReflect.Descriptor instead.
func (*GetOperationTraceDetailResponse) Descriptor() ([]byte, []int) {
	return file_operation_trace_proto_rawDescGZIP(), []int{4}
}

func (x *GetOperationTraceDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetOperationTraceDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetOperationTraceDetailResponse) GetData() *OperationTraceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_operation_trace_proto protoreflect.FileDescriptor

var file_operation_trace_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc4, 0x02, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a,
	0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x30, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x22, 0xcd, 0x01, 0x0a, 0x1b, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operation_trace_proto_rawDescOnce sync.Once
	file_operation_trace_proto_rawDescData = file_operation_trace_proto_rawDesc
)

func file_operation_trace_proto_rawDescGZIP() []byte {
	file_operation_trace_proto_rawDescOnce.Do(func() {
		file_operation_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_operation_trace_proto_rawDescData)
	})
	return file_operation_trace_proto_rawDescData
}

var file_operation_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operation_trace_proto_goTypes = []interface{}{
	(*OperationTraceInfo)(nil),              // 0: proto.OperationTraceInfo
	(*QueryOperationTraceRequest)(nil),      // 1: proto.QueryOperationTraceRequest
	(*QueryOperationTraceResponse)(nil),     // 2: proto.QueryOperationTraceResponse
	(*GetAllOperationTraceResponse)(nil),    // 3: proto.GetAllOperationTraceResponse
	(*GetOperationTraceDetailResponse)(nil), // 4: proto.GetOperationTraceDetailResponse
	(Code)(0),                               // 5: proto.Code
}
var file_operation_trace_proto_depIdxs = []int32{
	5, // 0: proto.QueryOperationTraceResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryOperationTraceResponse.data:type_name -> proto.OperationTraceInfo
	5, // 2: proto.GetAllOperationTraceResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllOperationTraceResponse.data:type_name -> proto.OperationTraceInfo
	5, // 4: proto.GetOperationTraceDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetOperationTraceDetailResponse.data:type_name -> proto.OperationTraceInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_operation_trace_proto_init() }
func file_operation_trace_proto_init() {
	if File_operation_trace_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_operation_trace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationTraceInfo); i {
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
		file_operation_trace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOperationTraceRequest); i {
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
		file_operation_trace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOperationTraceResponse); i {
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
		file_operation_trace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOperationTraceResponse); i {
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
		file_operation_trace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationTraceDetailResponse); i {
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
			RawDescriptor: file_operation_trace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_operation_trace_proto_goTypes,
		DependencyIndexes: file_operation_trace_proto_depIdxs,
		MessageInfos:      file_operation_trace_proto_msgTypes,
	}.Build()
	File_operation_trace_proto = out.File
	file_operation_trace_proto_rawDesc = nil
	file_operation_trace_proto_goTypes = nil
	file_operation_trace_proto_depIdxs = nil
}
