// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: remote_service_task.proto

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

type RemoteServiceTaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// 远程服务ID
	RemoteServiceID   string `protobuf:"bytes,5,opt,name=remoteServiceID,proto3" json:"remoteServiceID"`
	RemoteServiceName string `protobuf:"bytes,17,opt,name=remoteServiceName,proto3" json:"remoteServiceName"`
	// 调用方法
	InvokeMethod string `protobuf:"bytes,6,opt,name=invokeMethod,proto3" json:"invokeMethod"`
	// 回调方法
	CallbackMethod string `protobuf:"bytes,7,opt,name=callbackMethod,proto3" json:"callbackMethod"`
	// 失败措施
	FailureMeasure int32 `protobuf:"varint,8,opt,name=failureMeasure,proto3" json:"failureMeasure"`
	// 调用计数
	MaximumInvokeCount int32 `protobuf:"varint,9,opt,name=maximumInvokeCount,proto3" json:"maximumInvokeCount"`
	// 常规调用
	RoutineInvoke bool `protobuf:"varint,10,opt,name=routineInvoke,proto3" json:"routineInvoke"`
	// 定期调用
	RegularInvoke bool `protobuf:"varint,11,opt,name=regularInvoke,proto3" json:"regularInvoke"`
	// 起始时间
	StartTime string `protobuf:"bytes,12,opt,name=startTime,proto3" json:"startTime"`
	// 运行间隔
	Interval int32 `protobuf:"varint,13,opt,name=interval,proto3" json:"interval"`
	// 结束时间
	FinishTime string `protobuf:"bytes,14,opt,name=finishTime,proto3" json:"finishTime"`
	// 最后调用时间
	LastInvokeTime string `protobuf:"bytes,15,opt,name=lastInvokeTime,proto3" json:"lastInvokeTime"`
	// 远程服务任务参数
	RemoteServiceTaskParameters []*RemoteServiceTaskParameterInfo `protobuf:"bytes,16,rep,name=remoteServiceTaskParameters,proto3" json:"remoteServiceTaskParameters"`
}

func (x *RemoteServiceTaskInfo) Reset() {
	*x = RemoteServiceTaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteServiceTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteServiceTaskInfo) ProtoMessage() {}

func (x *RemoteServiceTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteServiceTaskInfo.ProtoReflect.Descriptor instead.
func (*RemoteServiceTaskInfo) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteServiceTaskInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetRemoteServiceID() string {
	if x != nil {
		return x.RemoteServiceID
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetRemoteServiceName() string {
	if x != nil {
		return x.RemoteServiceName
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetInvokeMethod() string {
	if x != nil {
		return x.InvokeMethod
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetCallbackMethod() string {
	if x != nil {
		return x.CallbackMethod
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetFailureMeasure() int32 {
	if x != nil {
		return x.FailureMeasure
	}
	return 0
}

func (x *RemoteServiceTaskInfo) GetMaximumInvokeCount() int32 {
	if x != nil {
		return x.MaximumInvokeCount
	}
	return 0
}

func (x *RemoteServiceTaskInfo) GetRoutineInvoke() bool {
	if x != nil {
		return x.RoutineInvoke
	}
	return false
}

func (x *RemoteServiceTaskInfo) GetRegularInvoke() bool {
	if x != nil {
		return x.RegularInvoke
	}
	return false
}

func (x *RemoteServiceTaskInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *RemoteServiceTaskInfo) GetFinishTime() string {
	if x != nil {
		return x.FinishTime
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetLastInvokeTime() string {
	if x != nil {
		return x.LastInvokeTime
	}
	return ""
}

func (x *RemoteServiceTaskInfo) GetRemoteServiceTaskParameters() []*RemoteServiceTaskParameterInfo {
	if x != nil {
		return x.RemoteServiceTaskParameters
	}
	return nil
}

type RemoteServiceTaskParameterInfo struct {
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
	// 远程服务任务ID
	RemoteServiceTaskID string `protobuf:"bytes,7,opt,name=remoteServiceTaskID,proto3" json:"remoteServiceTaskID"`
}

func (x *RemoteServiceTaskParameterInfo) Reset() {
	*x = RemoteServiceTaskParameterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteServiceTaskParameterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteServiceTaskParameterInfo) ProtoMessage() {}

func (x *RemoteServiceTaskParameterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteServiceTaskParameterInfo.ProtoReflect.Descriptor instead.
func (*RemoteServiceTaskParameterInfo) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteServiceTaskParameterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteServiceTaskParameterInfo) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *RemoteServiceTaskParameterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteServiceTaskParameterInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RemoteServiceTaskParameterInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RemoteServiceTaskParameterInfo) GetRemoteServiceTaskID() string {
	if x != nil {
		return x.RemoteServiceTaskID
	}
	return ""
}

type QueryRemoteServiceTaskRequest struct {
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

func (x *QueryRemoteServiceTaskRequest) Reset() {
	*x = QueryRemoteServiceTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceTaskRequest) ProtoMessage() {}

func (x *QueryRemoteServiceTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceTaskRequest) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRemoteServiceTaskRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryRemoteServiceTaskRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryRemoteServiceTaskRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

type QueryRemoteServiceTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceTaskInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                    `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                    `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                    `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryRemoteServiceTaskResponse) Reset() {
	*x = QueryRemoteServiceTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceTaskResponse) ProtoMessage() {}

func (x *QueryRemoteServiceTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceTaskResponse.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceTaskResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRemoteServiceTaskResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryRemoteServiceTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryRemoteServiceTaskResponse) GetData() []*RemoteServiceTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryRemoteServiceTaskResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryRemoteServiceTaskResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryRemoteServiceTaskResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllRemoteServiceTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceTaskInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllRemoteServiceTaskResponse) Reset() {
	*x = GetAllRemoteServiceTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemoteServiceTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemoteServiceTaskResponse) ProtoMessage() {}

func (x *GetAllRemoteServiceTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRemoteServiceTaskResponse.ProtoReflect.Descriptor instead.
func (*GetAllRemoteServiceTaskResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllRemoteServiceTaskResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllRemoteServiceTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllRemoteServiceTaskResponse) GetData() []*RemoteServiceTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetRemoteServiceTaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *RemoteServiceTaskInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetRemoteServiceTaskDetailResponse) Reset() {
	*x = GetRemoteServiceTaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteServiceTaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteServiceTaskDetailResponse) ProtoMessage() {}

func (x *GetRemoteServiceTaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteServiceTaskDetailResponse.ProtoReflect.Descriptor instead.
func (*GetRemoteServiceTaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_task_proto_rawDescGZIP(), []int{5}
}

func (x *GetRemoteServiceTaskDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetRemoteServiceTaskDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRemoteServiceTaskDetailResponse) GetData() *RemoteServiceTaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_remote_service_task_proto protoreflect.FileDescriptor

var file_remote_service_task_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x05, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2c,
	0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x67,
	0x0a, 0x1b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x1b, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x22, 0x79, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0xd3, 0x01, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8e, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_service_task_proto_rawDescOnce sync.Once
	file_remote_service_task_proto_rawDescData = file_remote_service_task_proto_rawDesc
)

func file_remote_service_task_proto_rawDescGZIP() []byte {
	file_remote_service_task_proto_rawDescOnce.Do(func() {
		file_remote_service_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_service_task_proto_rawDescData)
	})
	return file_remote_service_task_proto_rawDescData
}

var file_remote_service_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_remote_service_task_proto_goTypes = []interface{}{
	(*RemoteServiceTaskInfo)(nil),              // 0: proto.RemoteServiceTaskInfo
	(*RemoteServiceTaskParameterInfo)(nil),     // 1: proto.RemoteServiceTaskParameterInfo
	(*QueryRemoteServiceTaskRequest)(nil),      // 2: proto.QueryRemoteServiceTaskRequest
	(*QueryRemoteServiceTaskResponse)(nil),     // 3: proto.QueryRemoteServiceTaskResponse
	(*GetAllRemoteServiceTaskResponse)(nil),    // 4: proto.GetAllRemoteServiceTaskResponse
	(*GetRemoteServiceTaskDetailResponse)(nil), // 5: proto.GetRemoteServiceTaskDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_remote_service_task_proto_depIdxs = []int32{
	1, // 0: proto.RemoteServiceTaskInfo.remoteServiceTaskParameters:type_name -> proto.RemoteServiceTaskParameterInfo
	6, // 1: proto.QueryRemoteServiceTaskResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryRemoteServiceTaskResponse.data:type_name -> proto.RemoteServiceTaskInfo
	6, // 3: proto.GetAllRemoteServiceTaskResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllRemoteServiceTaskResponse.data:type_name -> proto.RemoteServiceTaskInfo
	6, // 5: proto.GetRemoteServiceTaskDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetRemoteServiceTaskDetailResponse.data:type_name -> proto.RemoteServiceTaskInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_remote_service_task_proto_init() }
func file_remote_service_task_proto_init() {
	if File_remote_service_task_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remote_service_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteServiceTaskInfo); i {
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
		file_remote_service_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteServiceTaskParameterInfo); i {
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
		file_remote_service_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceTaskRequest); i {
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
		file_remote_service_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceTaskResponse); i {
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
		file_remote_service_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRemoteServiceTaskResponse); i {
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
		file_remote_service_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteServiceTaskDetailResponse); i {
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
			RawDescriptor: file_remote_service_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_service_task_proto_goTypes,
		DependencyIndexes: file_remote_service_task_proto_depIdxs,
		MessageInfos:      file_remote_service_task_proto_msgTypes,
	}.Build()
	File_remote_service_task_proto = out.File
	file_remote_service_task_proto_rawDesc = nil
	file_remote_service_task_proto_goTypes = nil
	file_remote_service_task_proto_depIdxs = nil
}