// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_shelf_bin.proto

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

type MaterialShelfBinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 编号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 行
	RowIndex int32 `protobuf:"varint,3,opt,name=rowIndex,proto3" json:"rowIndex"`
	// 列
	ColumnIndex int32 `protobuf:"varint,4,opt,name=columnIndex,proto3" json:"columnIndex"`
	// 识别码
	Identifier string `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier"`
	// 兼容容器类型ID
	MaterialContainerTypeID string                     `protobuf:"bytes,6,opt,name=materialContainerTypeID,proto3" json:"materialContainerTypeID"`
	MaterialContainerType   *MaterialContainerTypeInfo `protobuf:"bytes,7,opt,name=materialContainerType,proto3" json:"materialContainerType"`
	// 缺料调度
	EnableDispatch bool `protobuf:"varint,8,opt,name=enableDispatch,proto3" json:"enableDispatch"`
	// 状态监控
	EnableMonitor bool `protobuf:"varint,9,opt,name=enableMonitor,proto3" json:"enableMonitor"`
	// 当前状态
	CurrentState string `protobuf:"bytes,10,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,11,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 备注
	Remark string `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark"`
	// 隶属料架ID
	MaterialShelfID string             `protobuf:"bytes,13,opt,name=materialShelfID,proto3" json:"materialShelfID"`
	MaterialShelf   *MaterialShelfInfo `protobuf:"bytes,14,opt,name=materialShelf,proto3" json:"materialShelf"`
	// 当前物料ID
	MaterialInfoID string            `protobuf:"bytes,15,opt,name=materialInfoID,proto3" json:"materialInfoID"`
	MaterialInfo   *MaterialInfoInfo `protobuf:"bytes,16,opt,name=materialInfo,proto3" json:"materialInfo"`
}

func (x *MaterialShelfBinInfo) Reset() {
	*x = MaterialShelfBinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_shelf_bin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialShelfBinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialShelfBinInfo) ProtoMessage() {}

func (x *MaterialShelfBinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_shelf_bin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialShelfBinInfo.ProtoReflect.Descriptor instead.
func (*MaterialShelfBinInfo) Descriptor() ([]byte, []int) {
	return file_material_shelf_bin_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialShelfBinInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetRowIndex() int32 {
	if x != nil {
		return x.RowIndex
	}
	return 0
}

func (x *MaterialShelfBinInfo) GetColumnIndex() int32 {
	if x != nil {
		return x.ColumnIndex
	}
	return 0
}

func (x *MaterialShelfBinInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetMaterialContainerTypeID() string {
	if x != nil {
		return x.MaterialContainerTypeID
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetMaterialContainerType() *MaterialContainerTypeInfo {
	if x != nil {
		return x.MaterialContainerType
	}
	return nil
}

func (x *MaterialShelfBinInfo) GetEnableDispatch() bool {
	if x != nil {
		return x.EnableDispatch
	}
	return false
}

func (x *MaterialShelfBinInfo) GetEnableMonitor() bool {
	if x != nil {
		return x.EnableMonitor
	}
	return false
}

func (x *MaterialShelfBinInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetMaterialShelfID() string {
	if x != nil {
		return x.MaterialShelfID
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetMaterialShelf() *MaterialShelfInfo {
	if x != nil {
		return x.MaterialShelf
	}
	return nil
}

func (x *MaterialShelfBinInfo) GetMaterialInfoID() string {
	if x != nil {
		return x.MaterialInfoID
	}
	return ""
}

func (x *MaterialShelfBinInfo) GetMaterialInfo() *MaterialInfoInfo {
	if x != nil {
		return x.MaterialInfo
	}
	return nil
}

type QueryMaterialShelfBinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 编号
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 当前状态
	// @inject_tag: uri:"currentState" form:"currentState"
	CurrentState string `protobuf:"bytes,5,opt,name=currentState,proto3" json:"currentState" uri:"currentState" form:"currentState"`
}

func (x *QueryMaterialShelfBinRequest) Reset() {
	*x = QueryMaterialShelfBinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_shelf_bin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialShelfBinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialShelfBinRequest) ProtoMessage() {}

func (x *QueryMaterialShelfBinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_shelf_bin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialShelfBinRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialShelfBinRequest) Descriptor() ([]byte, []int) {
	return file_material_shelf_bin_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialShelfBinRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialShelfBinRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialShelfBinRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialShelfBinRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryMaterialShelfBinRequest) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

type QueryMaterialShelfBinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialShelfBinInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                   `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                   `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                   `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialShelfBinResponse) Reset() {
	*x = QueryMaterialShelfBinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_shelf_bin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialShelfBinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialShelfBinResponse) ProtoMessage() {}

func (x *QueryMaterialShelfBinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_shelf_bin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialShelfBinResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialShelfBinResponse) Descriptor() ([]byte, []int) {
	return file_material_shelf_bin_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialShelfBinResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialShelfBinResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialShelfBinResponse) GetData() []*MaterialShelfBinInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialShelfBinResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialShelfBinResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialShelfBinResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialShelfBinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialShelfBinInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialShelfBinResponse) Reset() {
	*x = GetAllMaterialShelfBinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_shelf_bin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialShelfBinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialShelfBinResponse) ProtoMessage() {}

func (x *GetAllMaterialShelfBinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_shelf_bin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialShelfBinResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialShelfBinResponse) Descriptor() ([]byte, []int) {
	return file_material_shelf_bin_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialShelfBinResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialShelfBinResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialShelfBinResponse) GetData() []*MaterialShelfBinInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialShelfBinDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialShelfBinInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialShelfBinDetailResponse) Reset() {
	*x = GetMaterialShelfBinDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_shelf_bin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialShelfBinDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialShelfBinDetailResponse) ProtoMessage() {}

func (x *GetMaterialShelfBinDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_shelf_bin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialShelfBinDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialShelfBinDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_shelf_bin_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialShelfBinDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialShelfBinDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialShelfBinDetailResponse) GetData() *MaterialShelfBinInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_shelf_bin_proto protoreflect.FileDescriptor

var file_material_shelf_bin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x5f, 0x62, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x05,
	0x0a, 0x14, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x12, 0x56, 0x0a, 0x15, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x15, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x28, 0x0a, 0x0f,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x49, 0x44, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x3b,
	0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x1c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xd1,
	0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x8f, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_shelf_bin_proto_rawDescOnce sync.Once
	file_material_shelf_bin_proto_rawDescData = file_material_shelf_bin_proto_rawDesc
)

func file_material_shelf_bin_proto_rawDescGZIP() []byte {
	file_material_shelf_bin_proto_rawDescOnce.Do(func() {
		file_material_shelf_bin_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_shelf_bin_proto_rawDescData)
	})
	return file_material_shelf_bin_proto_rawDescData
}

var file_material_shelf_bin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_material_shelf_bin_proto_goTypes = []interface{}{
	(*MaterialShelfBinInfo)(nil),              // 0: proto.MaterialShelfBinInfo
	(*QueryMaterialShelfBinRequest)(nil),      // 1: proto.QueryMaterialShelfBinRequest
	(*QueryMaterialShelfBinResponse)(nil),     // 2: proto.QueryMaterialShelfBinResponse
	(*GetAllMaterialShelfBinResponse)(nil),    // 3: proto.GetAllMaterialShelfBinResponse
	(*GetMaterialShelfBinDetailResponse)(nil), // 4: proto.GetMaterialShelfBinDetailResponse
	(*MaterialContainerTypeInfo)(nil),         // 5: proto.MaterialContainerTypeInfo
	(*MaterialShelfInfo)(nil),                 // 6: proto.MaterialShelfInfo
	(*MaterialInfoInfo)(nil),                  // 7: proto.MaterialInfoInfo
	(Code)(0),                                 // 8: proto.Code
}
var file_material_shelf_bin_proto_depIdxs = []int32{
	5, // 0: proto.MaterialShelfBinInfo.materialContainerType:type_name -> proto.MaterialContainerTypeInfo
	6, // 1: proto.MaterialShelfBinInfo.materialShelf:type_name -> proto.MaterialShelfInfo
	7, // 2: proto.MaterialShelfBinInfo.materialInfo:type_name -> proto.MaterialInfoInfo
	8, // 3: proto.QueryMaterialShelfBinResponse.code:type_name -> proto.Code
	0, // 4: proto.QueryMaterialShelfBinResponse.data:type_name -> proto.MaterialShelfBinInfo
	8, // 5: proto.GetAllMaterialShelfBinResponse.code:type_name -> proto.Code
	0, // 6: proto.GetAllMaterialShelfBinResponse.data:type_name -> proto.MaterialShelfBinInfo
	8, // 7: proto.GetMaterialShelfBinDetailResponse.code:type_name -> proto.Code
	0, // 8: proto.GetMaterialShelfBinDetailResponse.data:type_name -> proto.MaterialShelfBinInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_material_shelf_bin_proto_init() }
func file_material_shelf_bin_proto_init() {
	if File_material_shelf_bin_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_container_type_proto_init()
	file_material_shelf_proto_init()
	file_material_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_shelf_bin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialShelfBinInfo); i {
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
		file_material_shelf_bin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialShelfBinRequest); i {
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
		file_material_shelf_bin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialShelfBinResponse); i {
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
		file_material_shelf_bin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialShelfBinResponse); i {
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
		file_material_shelf_bin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialShelfBinDetailResponse); i {
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
			RawDescriptor: file_material_shelf_bin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_shelf_bin_proto_goTypes,
		DependencyIndexes: file_material_shelf_bin_proto_depIdxs,
		MessageInfos:      file_material_shelf_bin_proto_msgTypes,
	}.Build()
	File_material_shelf_bin_proto = out.File
	file_material_shelf_bin_proto_rawDesc = nil
	file_material_shelf_bin_proto_goTypes = nil
	file_material_shelf_bin_proto_depIdxs = nil
}
