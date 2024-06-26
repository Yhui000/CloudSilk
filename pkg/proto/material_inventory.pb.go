// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_inventory.proto

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

type MaterialInventoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 物料信息ID
	MaterialInfoID string            `protobuf:"bytes,2,opt,name=materialInfoID,proto3" json:"materialInfoID"`
	MaterialInfo   *MaterialInfoInfo `protobuf:"bytes,3,opt,name=materialInfo,proto3" json:"materialInfo"`
	// 物料仓库ID
	MaterialStoreID string             `protobuf:"bytes,4,opt,name=materialStoreID,proto3" json:"materialStoreID"`
	MaterialStore   *MaterialStoreInfo `protobuf:"bytes,5,opt,name=materialStore,proto3" json:"materialStore"`
	// 库存数量
	StoredQTY int64 `protobuf:"varint,6,opt,name=storedQTY,proto3" json:"storedQTY"`
	// 锁定数量
	IssuedQTY int64 `protobuf:"varint,7,opt,name=issuedQTY,proto3" json:"issuedQTY"`
	// 正在补料数量
	FeedingQTY int64 `protobuf:"varint,8,opt,name=feedingQTY,proto3" json:"feedingQTY"`
	// 正在备库数量
	IssuingQTY int64 `protobuf:"varint,9,opt,name=issuingQTY,proto3" json:"issuingQTY"`
	// 创建时间
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime"`
	// 创建人员
	CreateUserID   string `protobuf:"bytes,11,opt,name=createUserID,proto3" json:"createUserID"`
	CreateUserName string `protobuf:"bytes,15,opt,name=createUserName,proto3" json:"createUserName"`
	// 当前状态
	CurrentState string `protobuf:"bytes,12,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,13,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 备注
	Remark string `protobuf:"bytes,14,opt,name=remark,proto3" json:"remark"`
}

func (x *MaterialInventoryInfo) Reset() {
	*x = MaterialInventoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialInventoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialInventoryInfo) ProtoMessage() {}

func (x *MaterialInventoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialInventoryInfo.ProtoReflect.Descriptor instead.
func (*MaterialInventoryInfo) Descriptor() ([]byte, []int) {
	return file_material_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialInventoryInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialInventoryInfo) GetMaterialInfoID() string {
	if x != nil {
		return x.MaterialInfoID
	}
	return ""
}

func (x *MaterialInventoryInfo) GetMaterialInfo() *MaterialInfoInfo {
	if x != nil {
		return x.MaterialInfo
	}
	return nil
}

func (x *MaterialInventoryInfo) GetMaterialStoreID() string {
	if x != nil {
		return x.MaterialStoreID
	}
	return ""
}

func (x *MaterialInventoryInfo) GetMaterialStore() *MaterialStoreInfo {
	if x != nil {
		return x.MaterialStore
	}
	return nil
}

func (x *MaterialInventoryInfo) GetStoredQTY() int64 {
	if x != nil {
		return x.StoredQTY
	}
	return 0
}

func (x *MaterialInventoryInfo) GetIssuedQTY() int64 {
	if x != nil {
		return x.IssuedQTY
	}
	return 0
}

func (x *MaterialInventoryInfo) GetFeedingQTY() int64 {
	if x != nil {
		return x.FeedingQTY
	}
	return 0
}

func (x *MaterialInventoryInfo) GetIssuingQTY() int64 {
	if x != nil {
		return x.IssuingQTY
	}
	return 0
}

func (x *MaterialInventoryInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *MaterialInventoryInfo) GetCreateUserID() string {
	if x != nil {
		return x.CreateUserID
	}
	return ""
}

func (x *MaterialInventoryInfo) GetCreateUserName() string {
	if x != nil {
		return x.CreateUserName
	}
	return ""
}

func (x *MaterialInventoryInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *MaterialInventoryInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *MaterialInventoryInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryMaterialInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 创建时间开始
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,4,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// 创建时间结束
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,5,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
	// 物料代号或描述
	// @inject_tag: uri:"materialInfo" form:"materialInfo"
	MaterialInfo string `protobuf:"bytes,6,opt,name=materialInfo,proto3" json:"materialInfo" uri:"materialInfo" form:"materialInfo"`
	// 仓库代号或描述
	// @inject_tag: uri:"materialStore" form:"materialStore"
	MaterialStore string `protobuf:"bytes,7,opt,name=materialStore,proto3" json:"materialStore" uri:"materialStore" form:"materialStore"`
}

func (x *QueryMaterialInventoryRequest) Reset() {
	*x = QueryMaterialInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialInventoryRequest) ProtoMessage() {}

func (x *QueryMaterialInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialInventoryRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialInventoryRequest) Descriptor() ([]byte, []int) {
	return file_material_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialInventoryRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialInventoryRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialInventoryRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialInventoryRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryMaterialInventoryRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

func (x *QueryMaterialInventoryRequest) GetMaterialInfo() string {
	if x != nil {
		return x.MaterialInfo
	}
	return ""
}

func (x *QueryMaterialInventoryRequest) GetMaterialStore() string {
	if x != nil {
		return x.MaterialStore
	}
	return ""
}

type QueryMaterialInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialInventoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                    `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                    `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                    `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialInventoryResponse) Reset() {
	*x = QueryMaterialInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialInventoryResponse) ProtoMessage() {}

func (x *QueryMaterialInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialInventoryResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialInventoryResponse) Descriptor() ([]byte, []int) {
	return file_material_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialInventoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialInventoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialInventoryResponse) GetData() []*MaterialInventoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialInventoryResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialInventoryResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialInventoryResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialInventoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialInventoryResponse) Reset() {
	*x = GetAllMaterialInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialInventoryResponse) ProtoMessage() {}

func (x *GetAllMaterialInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialInventoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialInventoryResponse) Descriptor() ([]byte, []int) {
	return file_material_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialInventoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialInventoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialInventoryResponse) GetData() []*MaterialInventoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialInventoryDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialInventoryInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialInventoryDetailResponse) Reset() {
	*x = GetMaterialInventoryDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialInventoryDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialInventoryDetailResponse) ProtoMessage() {}

func (x *GetMaterialInventoryDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialInventoryDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialInventoryDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialInventoryDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialInventoryDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialInventoryDetailResponse) GetData() *MaterialInventoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_inventory_proto protoreflect.FileDescriptor

var file_material_inventory_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2,
	0x04, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44,
	0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a,
	0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x51, 0x54, 0x59, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x51, 0x54, 0x59, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x51,
	0x54, 0x59, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x51, 0x54, 0x59, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x54,
	0x59, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67,
	0x51, 0x54, 0x59, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x51, 0x54,
	0x59, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67,
	0x51, 0x54, 0x59, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x22, 0x87, 0x02, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x30, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xd3, 0x01,
	0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x8e, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_inventory_proto_rawDescOnce sync.Once
	file_material_inventory_proto_rawDescData = file_material_inventory_proto_rawDesc
)

func file_material_inventory_proto_rawDescGZIP() []byte {
	file_material_inventory_proto_rawDescOnce.Do(func() {
		file_material_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_inventory_proto_rawDescData)
	})
	return file_material_inventory_proto_rawDescData
}

var file_material_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_material_inventory_proto_goTypes = []interface{}{
	(*MaterialInventoryInfo)(nil),              // 0: proto.MaterialInventoryInfo
	(*QueryMaterialInventoryRequest)(nil),      // 1: proto.QueryMaterialInventoryRequest
	(*QueryMaterialInventoryResponse)(nil),     // 2: proto.QueryMaterialInventoryResponse
	(*GetAllMaterialInventoryResponse)(nil),    // 3: proto.GetAllMaterialInventoryResponse
	(*GetMaterialInventoryDetailResponse)(nil), // 4: proto.GetMaterialInventoryDetailResponse
	(*MaterialInfoInfo)(nil),                   // 5: proto.MaterialInfoInfo
	(*MaterialStoreInfo)(nil),                  // 6: proto.MaterialStoreInfo
	(Code)(0),                                  // 7: proto.Code
}
var file_material_inventory_proto_depIdxs = []int32{
	5, // 0: proto.MaterialInventoryInfo.materialInfo:type_name -> proto.MaterialInfoInfo
	6, // 1: proto.MaterialInventoryInfo.materialStore:type_name -> proto.MaterialStoreInfo
	7, // 2: proto.QueryMaterialInventoryResponse.code:type_name -> proto.Code
	0, // 3: proto.QueryMaterialInventoryResponse.data:type_name -> proto.MaterialInventoryInfo
	7, // 4: proto.GetAllMaterialInventoryResponse.code:type_name -> proto.Code
	0, // 5: proto.GetAllMaterialInventoryResponse.data:type_name -> proto.MaterialInventoryInfo
	7, // 6: proto.GetMaterialInventoryDetailResponse.code:type_name -> proto.Code
	0, // 7: proto.GetMaterialInventoryDetailResponse.data:type_name -> proto.MaterialInventoryInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_material_inventory_proto_init() }
func file_material_inventory_proto_init() {
	if File_material_inventory_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_info_proto_init()
	file_material_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialInventoryInfo); i {
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
		file_material_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialInventoryRequest); i {
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
		file_material_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialInventoryResponse); i {
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
		file_material_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialInventoryResponse); i {
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
		file_material_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialInventoryDetailResponse); i {
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
			RawDescriptor: file_material_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_inventory_proto_goTypes,
		DependencyIndexes: file_material_inventory_proto_depIdxs,
		MessageInfos:      file_material_inventory_proto_msgTypes,
	}.Build()
	File_material_inventory_proto = out.File
	file_material_inventory_proto_rawDesc = nil
	file_material_inventory_proto_goTypes = nil
	file_material_inventory_proto_depIdxs = nil
}
