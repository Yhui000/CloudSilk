// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: invocation_authentication.proto

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

type InvocationAuthenticationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// IP地址
	IPAddress string `protobuf:"bytes,4,opt,name=iPAddress,proto3" json:"iPAddress"`
	// API密钥
	APIKey string `protobuf:"bytes,5,opt,name=aPIKey,proto3" json:"aPIKey"`
	// 备注
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
}

func (x *InvocationAuthenticationInfo) Reset() {
	*x = InvocationAuthenticationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invocation_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvocationAuthenticationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvocationAuthenticationInfo) ProtoMessage() {}

func (x *InvocationAuthenticationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_invocation_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvocationAuthenticationInfo.ProtoReflect.Descriptor instead.
func (*InvocationAuthenticationInfo) Descriptor() ([]byte, []int) {
	return file_invocation_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *InvocationAuthenticationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvocationAuthenticationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InvocationAuthenticationInfo) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

func (x *InvocationAuthenticationInfo) GetAPIKey() string {
	if x != nil {
		return x.APIKey
	}
	return ""
}

func (x *InvocationAuthenticationInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryInvocationAuthenticationRequest struct {
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
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" uri:"name" form:"name"`
}

func (x *QueryInvocationAuthenticationRequest) Reset() {
	*x = QueryInvocationAuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invocation_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInvocationAuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInvocationAuthenticationRequest) ProtoMessage() {}

func (x *QueryInvocationAuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invocation_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInvocationAuthenticationRequest.ProtoReflect.Descriptor instead.
func (*QueryInvocationAuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_invocation_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *QueryInvocationAuthenticationRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryInvocationAuthenticationRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryInvocationAuthenticationRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryInvocationAuthenticationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryInvocationAuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                            `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*InvocationAuthenticationInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                           `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                           `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                           `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryInvocationAuthenticationResponse) Reset() {
	*x = QueryInvocationAuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invocation_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInvocationAuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInvocationAuthenticationResponse) ProtoMessage() {}

func (x *QueryInvocationAuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invocation_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInvocationAuthenticationResponse.ProtoReflect.Descriptor instead.
func (*QueryInvocationAuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_invocation_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *QueryInvocationAuthenticationResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryInvocationAuthenticationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryInvocationAuthenticationResponse) GetData() []*InvocationAuthenticationInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryInvocationAuthenticationResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryInvocationAuthenticationResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryInvocationAuthenticationResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllInvocationAuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                            `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*InvocationAuthenticationInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllInvocationAuthenticationResponse) Reset() {
	*x = GetAllInvocationAuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invocation_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllInvocationAuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllInvocationAuthenticationResponse) ProtoMessage() {}

func (x *GetAllInvocationAuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invocation_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllInvocationAuthenticationResponse.ProtoReflect.Descriptor instead.
func (*GetAllInvocationAuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_invocation_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllInvocationAuthenticationResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllInvocationAuthenticationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllInvocationAuthenticationResponse) GetData() []*InvocationAuthenticationInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetInvocationAuthenticationDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *InvocationAuthenticationInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetInvocationAuthenticationDetailResponse) Reset() {
	*x = GetInvocationAuthenticationDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invocation_authentication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvocationAuthenticationDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvocationAuthenticationDetailResponse) ProtoMessage() {}

func (x *GetInvocationAuthenticationDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invocation_authentication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvocationAuthenticationDetailResponse.ProtoReflect.Descriptor instead.
func (*GetInvocationAuthenticationDetailResponse) Descriptor() ([]byte, []int) {
	return file_invocation_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *GetInvocationAuthenticationDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetInvocationAuthenticationDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetInvocationAuthenticationDetailResponse) GetData() *InvocationAuthenticationInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_invocation_authentication_proto protoreflect.FileDescriptor

var file_invocation_authentication_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x1c, 0x49,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x94, 0x01,
	0x0a, 0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x25, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x9c, 0x01, 0x0a, 0x26, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invocation_authentication_proto_rawDescOnce sync.Once
	file_invocation_authentication_proto_rawDescData = file_invocation_authentication_proto_rawDesc
)

func file_invocation_authentication_proto_rawDescGZIP() []byte {
	file_invocation_authentication_proto_rawDescOnce.Do(func() {
		file_invocation_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_invocation_authentication_proto_rawDescData)
	})
	return file_invocation_authentication_proto_rawDescData
}

var file_invocation_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_invocation_authentication_proto_goTypes = []interface{}{
	(*InvocationAuthenticationInfo)(nil),              // 0: proto.InvocationAuthenticationInfo
	(*QueryInvocationAuthenticationRequest)(nil),      // 1: proto.QueryInvocationAuthenticationRequest
	(*QueryInvocationAuthenticationResponse)(nil),     // 2: proto.QueryInvocationAuthenticationResponse
	(*GetAllInvocationAuthenticationResponse)(nil),    // 3: proto.GetAllInvocationAuthenticationResponse
	(*GetInvocationAuthenticationDetailResponse)(nil), // 4: proto.GetInvocationAuthenticationDetailResponse
	(Code)(0), // 5: proto.Code
}
var file_invocation_authentication_proto_depIdxs = []int32{
	5, // 0: proto.QueryInvocationAuthenticationResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryInvocationAuthenticationResponse.data:type_name -> proto.InvocationAuthenticationInfo
	5, // 2: proto.GetAllInvocationAuthenticationResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllInvocationAuthenticationResponse.data:type_name -> proto.InvocationAuthenticationInfo
	5, // 4: proto.GetInvocationAuthenticationDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetInvocationAuthenticationDetailResponse.data:type_name -> proto.InvocationAuthenticationInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_invocation_authentication_proto_init() }
func file_invocation_authentication_proto_init() {
	if File_invocation_authentication_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_invocation_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvocationAuthenticationInfo); i {
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
		file_invocation_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInvocationAuthenticationRequest); i {
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
		file_invocation_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInvocationAuthenticationResponse); i {
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
		file_invocation_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllInvocationAuthenticationResponse); i {
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
		file_invocation_authentication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvocationAuthenticationDetailResponse); i {
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
			RawDescriptor: file_invocation_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_invocation_authentication_proto_goTypes,
		DependencyIndexes: file_invocation_authentication_proto_depIdxs,
		MessageInfos:      file_invocation_authentication_proto_msgTypes,
	}.Build()
	File_invocation_authentication_proto = out.File
	file_invocation_authentication_proto_rawDesc = nil
	file_invocation_authentication_proto_goTypes = nil
	file_invocation_authentication_proto_depIdxs = nil
}
