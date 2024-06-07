// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: remote_service_task_queue.proto

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

type RemoteServiceTaskQueueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 任务编号
	TaskNo string `protobuf:"bytes,3,opt,name=taskNo,proto3" json:"taskNo"`
	// 创建时间
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime"`
	// 请求路径
	RequestURL string `protobuf:"bytes,5,opt,name=requestURL,proto3" json:"requestURL"`
	// 请求内容
	RequestText string `protobuf:"bytes,6,opt,name=requestText,proto3" json:"requestText"`
	// 响应内容
	ResponseText string `protobuf:"bytes,7,opt,name=responseText,proto3" json:"responseText"`
	// 完成时间
	FinishTime string `protobuf:"bytes,8,opt,name=finishTime,proto3" json:"finishTime"`
	// 调用计数
	InvokeCount int32 `protobuf:"varint,9,opt,name=invokeCount,proto3" json:"invokeCount"`
	// 当前状态
	CurrentState string `protobuf:"bytes,10,opt,name=currentState,proto3" json:"currentState"`
	// 事务状态
	TransactionState string `protobuf:"bytes,11,opt,name=transactionState,proto3" json:"transactionState"`
	// 远程任务ID
	RemoteServiceTaskID string `protobuf:"bytes,12,opt,name=remoteServiceTaskID,proto3" json:"remoteServiceTaskID"`
	// 远程任务
	RemoteServiceTaskName string `protobuf:"bytes,13,opt,name=remoteServiceTaskName,proto3" json:"remoteServiceTaskName"`
	// 远程服务
	RemoteServiceName string `protobuf:"bytes,14,opt,name=remoteServiceName,proto3" json:"remoteServiceName"`
	// 远程任务队列参数
	RemoteServiceTaskQueueParameters []*RemoteServiceTaskQueueParameterInfo `protobuf:"bytes,15,rep,name=remoteServiceTaskQueueParameters,proto3" json:"remoteServiceTaskQueueParameters"`
}

func (x *RemoteServiceTaskQueueInfo) Reset() {
	*x = RemoteServiceTaskQueueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteServiceTaskQueueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteServiceTaskQueueInfo) ProtoMessage() {}

func (x *RemoteServiceTaskQueueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteServiceTaskQueueInfo.ProtoReflect.Descriptor instead.
func (*RemoteServiceTaskQueueInfo) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteServiceTaskQueueInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetTaskNo() string {
	if x != nil {
		return x.TaskNo
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRequestURL() string {
	if x != nil {
		return x.RequestURL
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRequestText() string {
	if x != nil {
		return x.RequestText
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetResponseText() string {
	if x != nil {
		return x.ResponseText
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetFinishTime() string {
	if x != nil {
		return x.FinishTime
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetInvokeCount() int32 {
	if x != nil {
		return x.InvokeCount
	}
	return 0
}

func (x *RemoteServiceTaskQueueInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetTransactionState() string {
	if x != nil {
		return x.TransactionState
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRemoteServiceTaskID() string {
	if x != nil {
		return x.RemoteServiceTaskID
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRemoteServiceTaskName() string {
	if x != nil {
		return x.RemoteServiceTaskName
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRemoteServiceName() string {
	if x != nil {
		return x.RemoteServiceName
	}
	return ""
}

func (x *RemoteServiceTaskQueueInfo) GetRemoteServiceTaskQueueParameters() []*RemoteServiceTaskQueueParameterInfo {
	if x != nil {
		return x.RemoteServiceTaskQueueParameters
	}
	return nil
}

type RemoteServiceTaskQueueParameterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	DataType                 string `protobuf:"bytes,2,opt,name=dataType,proto3" json:"dataType"`
	Name                     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Description              string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	Value                    string `protobuf:"bytes,5,opt,name=value,proto3" json:"value"`
	RemoteServiceTaskQueueID string `protobuf:"bytes,6,opt,name=remoteServiceTaskQueueID,proto3" json:"remoteServiceTaskQueueID"`
}

func (x *RemoteServiceTaskQueueParameterInfo) Reset() {
	*x = RemoteServiceTaskQueueParameterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteServiceTaskQueueParameterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteServiceTaskQueueParameterInfo) ProtoMessage() {}

func (x *RemoteServiceTaskQueueParameterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteServiceTaskQueueParameterInfo.ProtoReflect.Descriptor instead.
func (*RemoteServiceTaskQueueParameterInfo) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteServiceTaskQueueParameterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteServiceTaskQueueParameterInfo) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *RemoteServiceTaskQueueParameterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteServiceTaskQueueParameterInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RemoteServiceTaskQueueParameterInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RemoteServiceTaskQueueParameterInfo) GetRemoteServiceTaskQueueID() string {
	if x != nil {
		return x.RemoteServiceTaskQueueID
	}
	return ""
}

type QueryRemoteServiceTaskQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 任务编号或请求内容或响应内容
	// @inject_tag: uri:"taskNo" form:"taskNo"
	TaskNo string `protobuf:"bytes,4,opt,name=taskNo,proto3" json:"taskNo" uri:"taskNo" form:"taskNo"`
	// 创建时间开始
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,5,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// 创建时间结束
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,6,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
}

func (x *QueryRemoteServiceTaskQueueRequest) Reset() {
	*x = QueryRemoteServiceTaskQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceTaskQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceTaskQueueRequest) ProtoMessage() {}

func (x *QueryRemoteServiceTaskQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceTaskQueueRequest.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceTaskQueueRequest) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRemoteServiceTaskQueueRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryRemoteServiceTaskQueueRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryRemoteServiceTaskQueueRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryRemoteServiceTaskQueueRequest) GetTaskNo() string {
	if x != nil {
		return x.TaskNo
	}
	return ""
}

func (x *QueryRemoteServiceTaskQueueRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryRemoteServiceTaskQueueRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

type QueryRemoteServiceTaskQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceTaskQueueInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                         `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                         `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                         `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryRemoteServiceTaskQueueResponse) Reset() {
	*x = QueryRemoteServiceTaskQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceTaskQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceTaskQueueResponse) ProtoMessage() {}

func (x *QueryRemoteServiceTaskQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceTaskQueueResponse.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceTaskQueueResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRemoteServiceTaskQueueResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryRemoteServiceTaskQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryRemoteServiceTaskQueueResponse) GetData() []*RemoteServiceTaskQueueInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryRemoteServiceTaskQueueResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryRemoteServiceTaskQueueResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryRemoteServiceTaskQueueResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllRemoteServiceTaskQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceTaskQueueInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllRemoteServiceTaskQueueResponse) Reset() {
	*x = GetAllRemoteServiceTaskQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemoteServiceTaskQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemoteServiceTaskQueueResponse) ProtoMessage() {}

func (x *GetAllRemoteServiceTaskQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRemoteServiceTaskQueueResponse.ProtoReflect.Descriptor instead.
func (*GetAllRemoteServiceTaskQueueResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllRemoteServiceTaskQueueResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllRemoteServiceTaskQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllRemoteServiceTaskQueueResponse) GetData() []*RemoteServiceTaskQueueInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetRemoteServiceTaskQueueDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *RemoteServiceTaskQueueInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetRemoteServiceTaskQueueDetailResponse) Reset() {
	*x = GetRemoteServiceTaskQueueDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_queue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteServiceTaskQueueDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteServiceTaskQueueDetailResponse) ProtoMessage() {}

func (x *GetRemoteServiceTaskQueueDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_queue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteServiceTaskQueueDetailResponse.ProtoReflect.Descriptor instead.
func (*GetRemoteServiceTaskQueueDetailResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_queue_proto_rawDescGZIP(), []int{5}
}

func (x *GetRemoteServiceTaskQueueDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetRemoteServiceTaskQueueDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRemoteServiceTaskQueueDetailResponse) GetData() *RemoteServiceTaskQueueInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_remote_service_task_queue_proto protoreflect.FileDescriptor

var file_remote_service_task_queue_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a, 0x1a, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x4e,
	0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x15, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x76, 0x0a, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x23, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x49, 0x44, 0x22, 0xda, 0x01, 0x0a, 0x22, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31,
	0x22, 0xdd, 0x01, 0x0a, 0x23, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x98, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9b, 0x01, 0x0a, 0x27,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_service_task_queue_proto_rawDescOnce sync.Once
	file_remote_service_task_queue_proto_rawDescData = file_remote_service_task_queue_proto_rawDesc
)

func file_remote_service_task_queue_proto_rawDescGZIP() []byte {
	file_remote_service_task_queue_proto_rawDescOnce.Do(func() {
		file_remote_service_task_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_service_task_queue_proto_rawDescData)
	})
	return file_remote_service_task_queue_proto_rawDescData
}

var file_remote_service_task_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_remote_service_task_queue_proto_goTypes = []interface{}{
	(*RemoteServiceTaskQueueInfo)(nil),              // 0: proto.RemoteServiceTaskQueueInfo
	(*RemoteServiceTaskQueueParameterInfo)(nil),     // 1: proto.RemoteServiceTaskQueueParameterInfo
	(*QueryRemoteServiceTaskQueueRequest)(nil),      // 2: proto.QueryRemoteServiceTaskQueueRequest
	(*QueryRemoteServiceTaskQueueResponse)(nil),     // 3: proto.QueryRemoteServiceTaskQueueResponse
	(*GetAllRemoteServiceTaskQueueResponse)(nil),    // 4: proto.GetAllRemoteServiceTaskQueueResponse
	(*GetRemoteServiceTaskQueueDetailResponse)(nil), // 5: proto.GetRemoteServiceTaskQueueDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_remote_service_task_queue_proto_depIdxs = []int32{
	1, // 0: proto.RemoteServiceTaskQueueInfo.remoteServiceTaskQueueParameters:type_name -> proto.RemoteServiceTaskQueueParameterInfo
	6, // 1: proto.QueryRemoteServiceTaskQueueResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryRemoteServiceTaskQueueResponse.data:type_name -> proto.RemoteServiceTaskQueueInfo
	6, // 3: proto.GetAllRemoteServiceTaskQueueResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllRemoteServiceTaskQueueResponse.data:type_name -> proto.RemoteServiceTaskQueueInfo
	6, // 5: proto.GetRemoteServiceTaskQueueDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetRemoteServiceTaskQueueDetailResponse.data:type_name -> proto.RemoteServiceTaskQueueInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_remote_service_task_queue_proto_init() }
func file_remote_service_task_queue_proto_init() {
	if File_remote_service_task_queue_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remote_service_task_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteServiceTaskQueueInfo); i {
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
		file_remote_service_task_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteServiceTaskQueueParameterInfo); i {
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
		file_remote_service_task_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceTaskQueueRequest); i {
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
		file_remote_service_task_queue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceTaskQueueResponse); i {
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
		file_remote_service_task_queue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRemoteServiceTaskQueueResponse); i {
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
		file_remote_service_task_queue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteServiceTaskQueueDetailResponse); i {
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
			RawDescriptor: file_remote_service_task_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_service_task_queue_proto_goTypes,
		DependencyIndexes: file_remote_service_task_queue_proto_depIdxs,
		MessageInfos:      file_remote_service_task_queue_proto_msgTypes,
	}.Build()
	File_remote_service_task_queue_proto = out.File
	file_remote_service_task_queue_proto_rawDesc = nil
	file_remote_service_task_queue_proto_goTypes = nil
	file_remote_service_task_queue_proto_depIdxs = nil
}
