// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_container.proto

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

type MaterialContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 编号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// 识别码
	Identifier string `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier"`
	// 当前状态
	CurrentState string `protobuf:"bytes,5,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,6,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 备注
	Remark string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	// 容器类型ID
	MaterialContainerTypeID string                     `protobuf:"bytes,8,opt,name=materialContainerTypeID,proto3" json:"materialContainerTypeID"`
	MaterialContainerType   *MaterialContainerTypeInfo `protobuf:"bytes,9,opt,name=materialContainerType,proto3" json:"materialContainerType"`
	// 当前库位ID
	MaterialShelfBinID string                `protobuf:"bytes,10,opt,name=materialShelfBinID,proto3" json:"materialShelfBinID"`
	MaterialShelfBin   *MaterialShelfBinInfo `protobuf:"bytes,11,opt,name=materialShelfBin,proto3" json:"materialShelfBin"`
}

func (x *MaterialContainerInfo) Reset() {
	*x = MaterialContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialContainerInfo) ProtoMessage() {}

func (x *MaterialContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialContainerInfo.ProtoReflect.Descriptor instead.
func (*MaterialContainerInfo) Descriptor() ([]byte, []int) {
	return file_material_container_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialContainerInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialContainerInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaterialContainerInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *MaterialContainerInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *MaterialContainerInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *MaterialContainerInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialContainerInfo) GetMaterialContainerTypeID() string {
	if x != nil {
		return x.MaterialContainerTypeID
	}
	return ""
}

func (x *MaterialContainerInfo) GetMaterialContainerType() *MaterialContainerTypeInfo {
	if x != nil {
		return x.MaterialContainerType
	}
	return nil
}

func (x *MaterialContainerInfo) GetMaterialShelfBinID() string {
	if x != nil {
		return x.MaterialShelfBinID
	}
	return ""
}

func (x *MaterialContainerInfo) GetMaterialShelfBin() *MaterialShelfBinInfo {
	if x != nil {
		return x.MaterialShelfBin
	}
	return nil
}

type QueryMaterialContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 编号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 当前状态
	// @inject_tag: uri:"currentState" form:"currentState"
	CurrentState string `protobuf:"bytes,5,opt,name=currentState,proto3" json:"currentState" uri:"currentState" form:"currentState"`
	// 容器类型ID
	// @inject_tag: uri:"materialContainerTypeID" form:"materialContainerTypeID"
	MaterialContainerTypeID string `protobuf:"bytes,6,opt,name=materialContainerTypeID,proto3" json:"materialContainerTypeID" uri:"materialContainerTypeID" form:"materialContainerTypeID"`
}

func (x *QueryMaterialContainerRequest) Reset() {
	*x = QueryMaterialContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialContainerRequest) ProtoMessage() {}

func (x *QueryMaterialContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialContainerRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialContainerRequest) Descriptor() ([]byte, []int) {
	return file_material_container_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialContainerRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialContainerRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialContainerRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialContainerRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryMaterialContainerRequest) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *QueryMaterialContainerRequest) GetMaterialContainerTypeID() string {
	if x != nil {
		return x.MaterialContainerTypeID
	}
	return ""
}

type QueryMaterialContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialContainerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                    `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                    `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                    `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialContainerResponse) Reset() {
	*x = QueryMaterialContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_container_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialContainerResponse) ProtoMessage() {}

func (x *QueryMaterialContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_container_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialContainerResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialContainerResponse) Descriptor() ([]byte, []int) {
	return file_material_container_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialContainerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialContainerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialContainerResponse) GetData() []*MaterialContainerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialContainerResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialContainerResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialContainerResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialContainerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialContainerResponse) Reset() {
	*x = GetAllMaterialContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_container_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialContainerResponse) ProtoMessage() {}

func (x *GetAllMaterialContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_container_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialContainerResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialContainerResponse) Descriptor() ([]byte, []int) {
	return file_material_container_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialContainerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialContainerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialContainerResponse) GetData() []*MaterialContainerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialContainerDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialContainerInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialContainerDetailResponse) Reset() {
	*x = GetMaterialContainerDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_container_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialContainerDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialContainerDetailResponse) ProtoMessage() {}

func (x *GetMaterialContainerDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_container_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialContainerDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialContainerDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_container_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialContainerDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialContainerDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialContainerDetailResponse) GetData() *MaterialContainerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_container_proto protoreflect.FileDescriptor

var file_material_container_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x5f, 0x62, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x03, 0x0a,
	0x15, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x38, 0x0a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x56, 0x0a, 0x15, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e,
	0x49, 0x44, 0x12, 0x47, 0x0a, 0x10, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x42, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x22, 0xeb, 0x01, 0x0a, 0x1d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x0a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0xd3, 0x01, 0x0a, 0x1e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x8e, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x91, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_container_proto_rawDescOnce sync.Once
	file_material_container_proto_rawDescData = file_material_container_proto_rawDesc
)

func file_material_container_proto_rawDescGZIP() []byte {
	file_material_container_proto_rawDescOnce.Do(func() {
		file_material_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_container_proto_rawDescData)
	})
	return file_material_container_proto_rawDescData
}

var file_material_container_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_material_container_proto_goTypes = []interface{}{
	(*MaterialContainerInfo)(nil),              // 0: proto.MaterialContainerInfo
	(*QueryMaterialContainerRequest)(nil),      // 1: proto.QueryMaterialContainerRequest
	(*QueryMaterialContainerResponse)(nil),     // 2: proto.QueryMaterialContainerResponse
	(*GetAllMaterialContainerResponse)(nil),    // 3: proto.GetAllMaterialContainerResponse
	(*GetMaterialContainerDetailResponse)(nil), // 4: proto.GetMaterialContainerDetailResponse
	(*MaterialContainerTypeInfo)(nil),          // 5: proto.MaterialContainerTypeInfo
	(*MaterialShelfBinInfo)(nil),               // 6: proto.MaterialShelfBinInfo
	(Code)(0),                                  // 7: proto.Code
}
var file_material_container_proto_depIdxs = []int32{
	5, // 0: proto.MaterialContainerInfo.materialContainerType:type_name -> proto.MaterialContainerTypeInfo
	6, // 1: proto.MaterialContainerInfo.materialShelfBin:type_name -> proto.MaterialShelfBinInfo
	7, // 2: proto.QueryMaterialContainerResponse.code:type_name -> proto.Code
	0, // 3: proto.QueryMaterialContainerResponse.data:type_name -> proto.MaterialContainerInfo
	7, // 4: proto.GetAllMaterialContainerResponse.code:type_name -> proto.Code
	0, // 5: proto.GetAllMaterialContainerResponse.data:type_name -> proto.MaterialContainerInfo
	7, // 6: proto.GetMaterialContainerDetailResponse.code:type_name -> proto.Code
	0, // 7: proto.GetMaterialContainerDetailResponse.data:type_name -> proto.MaterialContainerInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_material_container_proto_init() }
func file_material_container_proto_init() {
	if File_material_container_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_container_type_proto_init()
	file_material_shelf_bin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialContainerInfo); i {
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
		file_material_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialContainerRequest); i {
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
		file_material_container_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialContainerResponse); i {
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
		file_material_container_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialContainerResponse); i {
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
		file_material_container_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialContainerDetailResponse); i {
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
			RawDescriptor: file_material_container_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_container_proto_goTypes,
		DependencyIndexes: file_material_container_proto_depIdxs,
		MessageInfos:      file_material_container_proto_msgTypes,
	}.Build()
	File_material_container_proto = out.File
	file_material_container_proto_rawDesc = nil
	file_material_container_proto_goTypes = nil
	file_material_container_proto_depIdxs = nil
}
