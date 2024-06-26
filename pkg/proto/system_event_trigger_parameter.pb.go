// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: system_event_trigger_parameter.proto

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

type SystemEventTriggerParameterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 数据类型
	DataType string `protobuf:"bytes,3,opt,name=dataType,proto3" json:"dataType"`
	// 名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// 描述
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	// 参数值
	Value string `protobuf:"bytes,6,opt,name=value,proto3" json:"value"`
	// 系统事件触发ID
	SystemEventTriggerID string `protobuf:"bytes,7,opt,name=systemEventTriggerID,proto3" json:"systemEventTriggerID"`
}

func (x *SystemEventTriggerParameterInfo) Reset() {
	*x = SystemEventTriggerParameterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_event_trigger_parameter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemEventTriggerParameterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemEventTriggerParameterInfo) ProtoMessage() {}

func (x *SystemEventTriggerParameterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_system_event_trigger_parameter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemEventTriggerParameterInfo.ProtoReflect.Descriptor instead.
func (*SystemEventTriggerParameterInfo) Descriptor() ([]byte, []int) {
	return file_system_event_trigger_parameter_proto_rawDescGZIP(), []int{0}
}

func (x *SystemEventTriggerParameterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SystemEventTriggerParameterInfo) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *SystemEventTriggerParameterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemEventTriggerParameterInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SystemEventTriggerParameterInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SystemEventTriggerParameterInfo) GetSystemEventTriggerID() string {
	if x != nil {
		return x.SystemEventTriggerID
	}
	return ""
}

type QuerySystemEventTriggerParameterRequest struct {
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

func (x *QuerySystemEventTriggerParameterRequest) Reset() {
	*x = QuerySystemEventTriggerParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_event_trigger_parameter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySystemEventTriggerParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySystemEventTriggerParameterRequest) ProtoMessage() {}

func (x *QuerySystemEventTriggerParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_event_trigger_parameter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySystemEventTriggerParameterRequest.ProtoReflect.Descriptor instead.
func (*QuerySystemEventTriggerParameterRequest) Descriptor() ([]byte, []int) {
	return file_system_event_trigger_parameter_proto_rawDescGZIP(), []int{1}
}

func (x *QuerySystemEventTriggerParameterRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QuerySystemEventTriggerParameterRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QuerySystemEventTriggerParameterRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

type QuerySystemEventTriggerParameterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                               `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                             `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*SystemEventTriggerParameterInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                              `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                              `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                              `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QuerySystemEventTriggerParameterResponse) Reset() {
	*x = QuerySystemEventTriggerParameterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_event_trigger_parameter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySystemEventTriggerParameterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySystemEventTriggerParameterResponse) ProtoMessage() {}

func (x *QuerySystemEventTriggerParameterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_event_trigger_parameter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySystemEventTriggerParameterResponse.ProtoReflect.Descriptor instead.
func (*QuerySystemEventTriggerParameterResponse) Descriptor() ([]byte, []int) {
	return file_system_event_trigger_parameter_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySystemEventTriggerParameterResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QuerySystemEventTriggerParameterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QuerySystemEventTriggerParameterResponse) GetData() []*SystemEventTriggerParameterInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QuerySystemEventTriggerParameterResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QuerySystemEventTriggerParameterResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QuerySystemEventTriggerParameterResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllSystemEventTriggerParameterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                               `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                             `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*SystemEventTriggerParameterInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllSystemEventTriggerParameterResponse) Reset() {
	*x = GetAllSystemEventTriggerParameterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_event_trigger_parameter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSystemEventTriggerParameterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSystemEventTriggerParameterResponse) ProtoMessage() {}

func (x *GetAllSystemEventTriggerParameterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_event_trigger_parameter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSystemEventTriggerParameterResponse.ProtoReflect.Descriptor instead.
func (*GetAllSystemEventTriggerParameterResponse) Descriptor() ([]byte, []int) {
	return file_system_event_trigger_parameter_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllSystemEventTriggerParameterResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllSystemEventTriggerParameterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllSystemEventTriggerParameterResponse) GetData() []*SystemEventTriggerParameterInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSystemEventTriggerParameterDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                             `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                           `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *SystemEventTriggerParameterInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetSystemEventTriggerParameterDetailResponse) Reset() {
	*x = GetSystemEventTriggerParameterDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_event_trigger_parameter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemEventTriggerParameterDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemEventTriggerParameterDetailResponse) ProtoMessage() {}

func (x *GetSystemEventTriggerParameterDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_event_trigger_parameter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemEventTriggerParameterDetailResponse.ProtoReflect.Descriptor instead.
func (*GetSystemEventTriggerParameterDetailResponse) Descriptor() ([]byte, []int) {
	return file_system_event_trigger_parameter_proto_rawDescGZIP(), []int{4}
}

func (x *GetSystemEventTriggerParameterDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetSystemEventTriggerParameterDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSystemEventTriggerParameterDetailResponse) GetData() *SystemEventTriggerParameterInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_system_event_trigger_parameter_proto protoreflect.FileDescriptor

var file_system_event_trigger_parameter_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d,
	0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcd, 0x01, 0x0a, 0x1f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x83, 0x01, 0x0a, 0x27, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe7, 0x01, 0x0a, 0x28, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0xa2, 0x01, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xa5, 0x01, 0x0a, 0x2c, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x65, 0x0a, 0x1b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x03, 0x41,
	0x64, 0x64, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_event_trigger_parameter_proto_rawDescOnce sync.Once
	file_system_event_trigger_parameter_proto_rawDescData = file_system_event_trigger_parameter_proto_rawDesc
)

func file_system_event_trigger_parameter_proto_rawDescGZIP() []byte {
	file_system_event_trigger_parameter_proto_rawDescOnce.Do(func() {
		file_system_event_trigger_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_event_trigger_parameter_proto_rawDescData)
	})
	return file_system_event_trigger_parameter_proto_rawDescData
}

var file_system_event_trigger_parameter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_system_event_trigger_parameter_proto_goTypes = []interface{}{
	(*SystemEventTriggerParameterInfo)(nil),              // 0: proto.SystemEventTriggerParameterInfo
	(*QuerySystemEventTriggerParameterRequest)(nil),      // 1: proto.QuerySystemEventTriggerParameterRequest
	(*QuerySystemEventTriggerParameterResponse)(nil),     // 2: proto.QuerySystemEventTriggerParameterResponse
	(*GetAllSystemEventTriggerParameterResponse)(nil),    // 3: proto.GetAllSystemEventTriggerParameterResponse
	(*GetSystemEventTriggerParameterDetailResponse)(nil), // 4: proto.GetSystemEventTriggerParameterDetailResponse
	(Code)(0),              // 5: proto.Code
	(*CommonResponse)(nil), // 6: proto.CommonResponse
}
var file_system_event_trigger_parameter_proto_depIdxs = []int32{
	5, // 0: proto.QuerySystemEventTriggerParameterResponse.code:type_name -> proto.Code
	0, // 1: proto.QuerySystemEventTriggerParameterResponse.data:type_name -> proto.SystemEventTriggerParameterInfo
	5, // 2: proto.GetAllSystemEventTriggerParameterResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllSystemEventTriggerParameterResponse.data:type_name -> proto.SystemEventTriggerParameterInfo
	5, // 4: proto.GetSystemEventTriggerParameterDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetSystemEventTriggerParameterDetailResponse.data:type_name -> proto.SystemEventTriggerParameterInfo
	0, // 6: proto.SystemEventTriggerParameter.Add:input_type -> proto.SystemEventTriggerParameterInfo
	6, // 7: proto.SystemEventTriggerParameter.Add:output_type -> proto.CommonResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_system_event_trigger_parameter_proto_init() }
func file_system_event_trigger_parameter_proto_init() {
	if File_system_event_trigger_parameter_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_system_event_trigger_parameter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemEventTriggerParameterInfo); i {
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
		file_system_event_trigger_parameter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySystemEventTriggerParameterRequest); i {
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
		file_system_event_trigger_parameter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySystemEventTriggerParameterResponse); i {
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
		file_system_event_trigger_parameter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSystemEventTriggerParameterResponse); i {
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
		file_system_event_trigger_parameter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemEventTriggerParameterDetailResponse); i {
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
			RawDescriptor: file_system_event_trigger_parameter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_event_trigger_parameter_proto_goTypes,
		DependencyIndexes: file_system_event_trigger_parameter_proto_depIdxs,
		MessageInfos:      file_system_event_trigger_parameter_proto_msgTypes,
	}.Build()
	File_system_event_trigger_parameter_proto = out.File
	file_system_event_trigger_parameter_proto_rawDesc = nil
	file_system_event_trigger_parameter_proto_goTypes = nil
	file_system_event_trigger_parameter_proto_depIdxs = nil
}
