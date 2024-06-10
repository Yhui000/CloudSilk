// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: remote_service.proto

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

type RemoteServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 是否启用
	Enable bool `protobuf:"varint,4,opt,name=enable,proto3" json:"enable"`
	// 服务地址
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	// 请求头
	Headers string `protobuf:"bytes,6,opt,name=headers,proto3" json:"headers"`
	// 超时时间
	Timeout int32 `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout"`
	// 使用凭证
	UseCredential bool `protobuf:"varint,8,opt,name=useCredential,proto3" json:"useCredential"`
	// 用户名
	UserName string `protobuf:"bytes,9,opt,name=userName,proto3" json:"userName"`
	// 密码
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password"`
}

func (x *RemoteServiceInfo) Reset() {
	*x = RemoteServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteServiceInfo) ProtoMessage() {}

func (x *RemoteServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteServiceInfo.ProtoReflect.Descriptor instead.
func (*RemoteServiceInfo) Descriptor() ([]byte, []int) {
	return file_remote_service_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteServiceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteServiceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteServiceInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *RemoteServiceInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RemoteServiceInfo) GetHeaders() string {
	if x != nil {
		return x.Headers
	}
	return ""
}

func (x *RemoteServiceInfo) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *RemoteServiceInfo) GetUseCredential() bool {
	if x != nil {
		return x.UseCredential
	}
	return false
}

func (x *RemoteServiceInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RemoteServiceInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type QueryRemoteServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 名称
	// @inject_tag: uri:"name" form:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" uri:"name" form:"name"`
}

func (x *QueryRemoteServiceRequest) Reset() {
	*x = QueryRemoteServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceRequest) ProtoMessage() {}

func (x *QueryRemoteServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceRequest.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceRequest) Descriptor() ([]byte, []int) {
	return file_remote_service_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRemoteServiceRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryRemoteServiceRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryRemoteServiceRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryRemoteServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryRemoteServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                 `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryRemoteServiceResponse) Reset() {
	*x = QueryRemoteServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRemoteServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRemoteServiceResponse) ProtoMessage() {}

func (x *QueryRemoteServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRemoteServiceResponse.ProtoReflect.Descriptor instead.
func (*QueryRemoteServiceResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRemoteServiceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryRemoteServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryRemoteServiceResponse) GetData() []*RemoteServiceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryRemoteServiceResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryRemoteServiceResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryRemoteServiceResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllRemoteServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                 `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*RemoteServiceInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllRemoteServiceResponse) Reset() {
	*x = GetAllRemoteServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemoteServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemoteServiceResponse) ProtoMessage() {}

func (x *GetAllRemoteServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRemoteServiceResponse.ProtoReflect.Descriptor instead.
func (*GetAllRemoteServiceResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllRemoteServiceResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllRemoteServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllRemoteServiceResponse) GetData() []*RemoteServiceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetRemoteServiceDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code               `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *RemoteServiceInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetRemoteServiceDetailResponse) Reset() {
	*x = GetRemoteServiceDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteServiceDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteServiceDetailResponse) ProtoMessage() {}

func (x *GetRemoteServiceDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteServiceDetailResponse.ProtoReflect.Descriptor instead.
func (*GetRemoteServiceDetailResponse) Descriptor() ([]byte, []int) {
	return file_remote_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRemoteServiceDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetRemoteServiceDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRemoteServiceDetailResponse) GetData() *RemoteServiceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_remote_service_proto protoreflect.FileDescriptor

var file_remote_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d,
	0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfb, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x89, 0x01,
	0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x1a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x86, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x89, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_service_proto_rawDescOnce sync.Once
	file_remote_service_proto_rawDescData = file_remote_service_proto_rawDesc
)

func file_remote_service_proto_rawDescGZIP() []byte {
	file_remote_service_proto_rawDescOnce.Do(func() {
		file_remote_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_service_proto_rawDescData)
	})
	return file_remote_service_proto_rawDescData
}

var file_remote_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_remote_service_proto_goTypes = []interface{}{
	(*RemoteServiceInfo)(nil),              // 0: proto.RemoteServiceInfo
	(*QueryRemoteServiceRequest)(nil),      // 1: proto.QueryRemoteServiceRequest
	(*QueryRemoteServiceResponse)(nil),     // 2: proto.QueryRemoteServiceResponse
	(*GetAllRemoteServiceResponse)(nil),    // 3: proto.GetAllRemoteServiceResponse
	(*GetRemoteServiceDetailResponse)(nil), // 4: proto.GetRemoteServiceDetailResponse
	(Code)(0),                              // 5: proto.Code
}
var file_remote_service_proto_depIdxs = []int32{
	5, // 0: proto.QueryRemoteServiceResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryRemoteServiceResponse.data:type_name -> proto.RemoteServiceInfo
	5, // 2: proto.GetAllRemoteServiceResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllRemoteServiceResponse.data:type_name -> proto.RemoteServiceInfo
	5, // 4: proto.GetRemoteServiceDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetRemoteServiceDetailResponse.data:type_name -> proto.RemoteServiceInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_remote_service_proto_init() }
func file_remote_service_proto_init() {
	if File_remote_service_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remote_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteServiceInfo); i {
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
		file_remote_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceRequest); i {
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
		file_remote_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRemoteServiceResponse); i {
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
		file_remote_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRemoteServiceResponse); i {
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
		file_remote_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteServiceDetailResponse); i {
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
			RawDescriptor: file_remote_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_service_proto_goTypes,
		DependencyIndexes: file_remote_service_proto_depIdxs,
		MessageInfos:      file_remote_service_proto_msgTypes,
	}.Build()
	File_remote_service_proto = out.File
	file_remote_service_proto_rawDesc = nil
	file_remote_service_proto_goTypes = nil
	file_remote_service_proto_depIdxs = nil
}
