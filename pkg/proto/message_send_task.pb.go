// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: message_send_task.proto

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

type MessageSendTaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// 是否启用
	Enable bool `protobuf:"varint,4,opt,name=enable,proto3" json:"enable"`
	// 工厂产线ID
	ProductionLineID string              `protobuf:"bytes,5,opt,name=productionLineID,proto3" json:"productionLineID"`
	ProductionLine   *ProductionLineInfo `protobuf:"bytes,6,opt,name=productionLine,proto3" json:"productionLine"`
	// 消息模板ID
	MessageTemplateID string               `protobuf:"bytes,7,opt,name=messageTemplateID,proto3" json:"messageTemplateID"`
	MessageTemplate   *MessageTemplateInfo `protobuf:"bytes,8,opt,name=messageTemplate,proto3" json:"messageTemplate"`
	// 收件地址
	To string `protobuf:"bytes,9,opt,name=To,proto3" json:"To"`
	// 抄送地址
	Cc string `protobuf:"bytes,10,opt,name=Cc,proto3" json:"Cc"`
	// 远程服务ID
	RemoteServiceTaskID string                 `protobuf:"bytes,11,opt,name=remoteServiceTaskID,proto3" json:"remoteServiceTaskID"`
	RemoteServiceTask   *RemoteServiceTaskInfo `protobuf:"bytes,12,opt,name=remoteServiceTask,proto3" json:"remoteServiceTask"`
	// 备注
	Remark string `protobuf:"bytes,13,opt,name=Remark,proto3" json:"Remark"`
}

func (x *MessageSendTaskInfo) Reset() {
	*x = MessageSendTaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendTaskInfo) ProtoMessage() {}

func (x *MessageSendTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendTaskInfo.ProtoReflect.Descriptor instead.
func (*MessageSendTaskInfo) Descriptor() ([]byte, []int) {
	return file_message_send_task_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSendTaskInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageSendTaskInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MessageSendTaskInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MessageSendTaskInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *MessageSendTaskInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *MessageSendTaskInfo) GetProductionLine() *ProductionLineInfo {
	if x != nil {
		return x.ProductionLine
	}
	return nil
}

func (x *MessageSendTaskInfo) GetMessageTemplateID() string {
	if x != nil {
		return x.MessageTemplateID
	}
	return ""
}

func (x *MessageSendTaskInfo) GetMessageTemplate() *MessageTemplateInfo {
	if x != nil {
		return x.MessageTemplate
	}
	return nil
}

func (x *MessageSendTaskInfo) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *MessageSendTaskInfo) GetCc() string {
	if x != nil {
		return x.Cc
	}
	return ""
}

func (x *MessageSendTaskInfo) GetRemoteServiceTaskID() string {
	if x != nil {
		return x.RemoteServiceTaskID
	}
	return ""
}

func (x *MessageSendTaskInfo) GetRemoteServiceTask() *RemoteServiceTaskInfo {
	if x != nil {
		return x.RemoteServiceTask
	}
	return nil
}

func (x *MessageSendTaskInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryMessageSendTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// @inject_tag: uri:"identifier" form:"identifier"
	Identifier string `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier" uri:"identifier" form:"identifier"`
	// 代号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 工厂产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,6,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// 消息类型ID
	// @inject_tag: uri:"messageTypeID" form:"messageTypeID"
	MessageTypeID string `protobuf:"bytes,7,opt,name=messageTypeID,proto3" json:"messageTypeID" uri:"messageTypeID" form:"messageTypeID"`
}

func (x *QueryMessageSendTaskRequest) Reset() {
	*x = QueryMessageSendTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessageSendTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessageSendTaskRequest) ProtoMessage() {}

func (x *QueryMessageSendTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessageSendTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryMessageSendTaskRequest) Descriptor() ([]byte, []int) {
	return file_message_send_task_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMessageSendTaskRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMessageSendTaskRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMessageSendTaskRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMessageSendTaskRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *QueryMessageSendTaskRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryMessageSendTaskRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryMessageSendTaskRequest) GetMessageTypeID() string {
	if x != nil {
		return x.MessageTypeID
	}
	return ""
}

type QueryMessageSendTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MessageSendTaskInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                  `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                  `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                  `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMessageSendTaskResponse) Reset() {
	*x = QueryMessageSendTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessageSendTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessageSendTaskResponse) ProtoMessage() {}

func (x *QueryMessageSendTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessageSendTaskResponse.ProtoReflect.Descriptor instead.
func (*QueryMessageSendTaskResponse) Descriptor() ([]byte, []int) {
	return file_message_send_task_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMessageSendTaskResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMessageSendTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMessageSendTaskResponse) GetData() []*MessageSendTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMessageSendTaskResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMessageSendTaskResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMessageSendTaskResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMessageSendTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MessageSendTaskInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMessageSendTaskResponse) Reset() {
	*x = GetAllMessageSendTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMessageSendTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMessageSendTaskResponse) ProtoMessage() {}

func (x *GetAllMessageSendTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMessageSendTaskResponse.ProtoReflect.Descriptor instead.
func (*GetAllMessageSendTaskResponse) Descriptor() ([]byte, []int) {
	return file_message_send_task_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMessageSendTaskResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMessageSendTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMessageSendTaskResponse) GetData() []*MessageSendTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMessageSendTaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                 `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MessageSendTaskInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMessageSendTaskDetailResponse) Reset() {
	*x = GetMessageSendTaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageSendTaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageSendTaskDetailResponse) ProtoMessage() {}

func (x *GetMessageSendTaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageSendTaskDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMessageSendTaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_message_send_task_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageSendTaskDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMessageSendTaskDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMessageSendTaskDetailResponse) GetData() *MessageSendTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_message_send_task_proto protoreflect.FileDescriptor

var file_message_send_task_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x04, 0x0a,
	0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x41,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12,
	0x44, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x43, 0x63, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xfd, 0x01, 0x0a, 0x1b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0xcf, 0x01, 0x0a, 0x1c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8a, 0x01,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_send_task_proto_rawDescOnce sync.Once
	file_message_send_task_proto_rawDescData = file_message_send_task_proto_rawDesc
)

func file_message_send_task_proto_rawDescGZIP() []byte {
	file_message_send_task_proto_rawDescOnce.Do(func() {
		file_message_send_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_send_task_proto_rawDescData)
	})
	return file_message_send_task_proto_rawDescData
}

var file_message_send_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_send_task_proto_goTypes = []interface{}{
	(*MessageSendTaskInfo)(nil),              // 0: proto.MessageSendTaskInfo
	(*QueryMessageSendTaskRequest)(nil),      // 1: proto.QueryMessageSendTaskRequest
	(*QueryMessageSendTaskResponse)(nil),     // 2: proto.QueryMessageSendTaskResponse
	(*GetAllMessageSendTaskResponse)(nil),    // 3: proto.GetAllMessageSendTaskResponse
	(*GetMessageSendTaskDetailResponse)(nil), // 4: proto.GetMessageSendTaskDetailResponse
	(*ProductionLineInfo)(nil),               // 5: proto.ProductionLineInfo
	(*MessageTemplateInfo)(nil),              // 6: proto.MessageTemplateInfo
	(*RemoteServiceTaskInfo)(nil),            // 7: proto.RemoteServiceTaskInfo
	(Code)(0),                                // 8: proto.Code
}
var file_message_send_task_proto_depIdxs = []int32{
	5, // 0: proto.MessageSendTaskInfo.productionLine:type_name -> proto.ProductionLineInfo
	6, // 1: proto.MessageSendTaskInfo.messageTemplate:type_name -> proto.MessageTemplateInfo
	7, // 2: proto.MessageSendTaskInfo.remoteServiceTask:type_name -> proto.RemoteServiceTaskInfo
	8, // 3: proto.QueryMessageSendTaskResponse.code:type_name -> proto.Code
	0, // 4: proto.QueryMessageSendTaskResponse.data:type_name -> proto.MessageSendTaskInfo
	8, // 5: proto.GetAllMessageSendTaskResponse.code:type_name -> proto.Code
	0, // 6: proto.GetAllMessageSendTaskResponse.data:type_name -> proto.MessageSendTaskInfo
	8, // 7: proto.GetMessageSendTaskDetailResponse.code:type_name -> proto.Code
	0, // 8: proto.GetMessageSendTaskDetailResponse.data:type_name -> proto.MessageSendTaskInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_message_send_task_proto_init() }
func file_message_send_task_proto_init() {
	if File_message_send_task_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_production_line_proto_init()
	file_message_template_proto_init()
	file_remote_service_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_send_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendTaskInfo); i {
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
		file_message_send_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessageSendTaskRequest); i {
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
		file_message_send_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessageSendTaskResponse); i {
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
		file_message_send_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMessageSendTaskResponse); i {
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
		file_message_send_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageSendTaskDetailResponse); i {
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
			RawDescriptor: file_message_send_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_send_task_proto_goTypes,
		DependencyIndexes: file_message_send_task_proto_depIdxs,
		MessageInfos:      file_message_send_task_proto_msgTypes,
	}.Build()
	File_message_send_task_proto = out.File
	file_message_send_task_proto_rawDesc = nil
	file_message_send_task_proto_goTypes = nil
	file_message_send_task_proto_depIdxs = nil
}
