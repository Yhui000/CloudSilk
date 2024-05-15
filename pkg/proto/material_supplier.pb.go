// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_supplier.proto

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

type MaterialSupplierInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// 备注
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark"`
	// 持有物料信息
	AvailableMaterials []*AvailableMaterialInfo `protobuf:"bytes,5,rep,name=availableMaterials,proto3" json:"availableMaterials"`
}

func (x *MaterialSupplierInfo) Reset() {
	*x = MaterialSupplierInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialSupplierInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialSupplierInfo) ProtoMessage() {}

func (x *MaterialSupplierInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialSupplierInfo.ProtoReflect.Descriptor instead.
func (*MaterialSupplierInfo) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialSupplierInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialSupplierInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialSupplierInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaterialSupplierInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialSupplierInfo) GetAvailableMaterials() []*AvailableMaterialInfo {
	if x != nil {
		return x.AvailableMaterials
	}
	return nil
}

type AvailableMaterialInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 物料供应商ID
	MaterialSupplierID string `protobuf:"bytes,2,opt,name=materialSupplierID,proto3" json:"materialSupplierID"`
	// 物料信息ID
	MaterialInfoID string `protobuf:"bytes,3,opt,name=materialInfoID,proto3" json:"materialInfoID"`
}

func (x *AvailableMaterialInfo) Reset() {
	*x = AvailableMaterialInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableMaterialInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableMaterialInfo) ProtoMessage() {}

func (x *AvailableMaterialInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableMaterialInfo.ProtoReflect.Descriptor instead.
func (*AvailableMaterialInfo) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *AvailableMaterialInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AvailableMaterialInfo) GetMaterialSupplierID() string {
	if x != nil {
		return x.MaterialSupplierID
	}
	return ""
}

func (x *AvailableMaterialInfo) GetMaterialInfoID() string {
	if x != nil {
		return x.MaterialInfoID
	}
	return ""
}

type QueryMaterialSupplierRequest struct {
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
}

func (x *QueryMaterialSupplierRequest) Reset() {
	*x = QueryMaterialSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialSupplierRequest) ProtoMessage() {}

func (x *QueryMaterialSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialSupplierRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialSupplierRequest) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialSupplierRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialSupplierRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialSupplierRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialSupplierRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type QueryMaterialSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialSupplierInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                   `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                   `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                   `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialSupplierResponse) Reset() {
	*x = QueryMaterialSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialSupplierResponse) ProtoMessage() {}

func (x *QueryMaterialSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialSupplierResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialSupplierResponse) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMaterialSupplierResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialSupplierResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialSupplierResponse) GetData() []*MaterialSupplierInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialSupplierResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialSupplierResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialSupplierResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialSupplierInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialSupplierResponse) Reset() {
	*x = GetAllMaterialSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialSupplierResponse) ProtoMessage() {}

func (x *GetAllMaterialSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialSupplierResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialSupplierResponse) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllMaterialSupplierResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialSupplierResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialSupplierResponse) GetData() []*MaterialSupplierInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialSupplierDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialSupplierInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialSupplierDetailResponse) Reset() {
	*x = GetMaterialSupplierDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_supplier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialSupplierDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialSupplierDetailResponse) ProtoMessage() {}

func (x *GetMaterialSupplierDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_supplier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialSupplierDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialSupplierDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_supplier_proto_rawDescGZIP(), []int{5}
}

func (x *GetMaterialSupplierDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialSupplierDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialSupplierDetailResponse) GetData() *MaterialSupplierInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_supplier_proto protoreflect.FileDescriptor

var file_material_supplier_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x4c, 0x0a, 0x12, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x7f, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x22, 0x98, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_supplier_proto_rawDescOnce sync.Once
	file_material_supplier_proto_rawDescData = file_material_supplier_proto_rawDesc
)

func file_material_supplier_proto_rawDescGZIP() []byte {
	file_material_supplier_proto_rawDescOnce.Do(func() {
		file_material_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_supplier_proto_rawDescData)
	})
	return file_material_supplier_proto_rawDescData
}

var file_material_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_material_supplier_proto_goTypes = []interface{}{
	(*MaterialSupplierInfo)(nil),              // 0: proto.MaterialSupplierInfo
	(*AvailableMaterialInfo)(nil),             // 1: proto.AvailableMaterialInfo
	(*QueryMaterialSupplierRequest)(nil),      // 2: proto.QueryMaterialSupplierRequest
	(*QueryMaterialSupplierResponse)(nil),     // 3: proto.QueryMaterialSupplierResponse
	(*GetAllMaterialSupplierResponse)(nil),    // 4: proto.GetAllMaterialSupplierResponse
	(*GetMaterialSupplierDetailResponse)(nil), // 5: proto.GetMaterialSupplierDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_material_supplier_proto_depIdxs = []int32{
	1, // 0: proto.MaterialSupplierInfo.availableMaterials:type_name -> proto.AvailableMaterialInfo
	6, // 1: proto.QueryMaterialSupplierResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryMaterialSupplierResponse.data:type_name -> proto.MaterialSupplierInfo
	6, // 3: proto.GetAllMaterialSupplierResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllMaterialSupplierResponse.data:type_name -> proto.MaterialSupplierInfo
	6, // 5: proto.GetMaterialSupplierDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetMaterialSupplierDetailResponse.data:type_name -> proto.MaterialSupplierInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_material_supplier_proto_init() }
func file_material_supplier_proto_init() {
	if File_material_supplier_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialSupplierInfo); i {
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
		file_material_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableMaterialInfo); i {
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
		file_material_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialSupplierRequest); i {
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
		file_material_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialSupplierResponse); i {
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
		file_material_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialSupplierResponse); i {
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
		file_material_supplier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialSupplierDetailResponse); i {
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
			RawDescriptor: file_material_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_supplier_proto_goTypes,
		DependencyIndexes: file_material_supplier_proto_depIdxs,
		MessageInfos:      file_material_supplier_proto_msgTypes,
	}.Build()
	File_material_supplier_proto = out.File
	file_material_supplier_proto_rawDesc = nil
	file_material_supplier_proto_goTypes = nil
	file_material_supplier_proto_depIdxs = nil
}
