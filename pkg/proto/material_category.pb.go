// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_category.proto

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

type MaterialCategoryInfo struct {
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
}

func (x *MaterialCategoryInfo) Reset() {
	*x = MaterialCategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialCategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialCategoryInfo) ProtoMessage() {}

func (x *MaterialCategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialCategoryInfo.ProtoReflect.Descriptor instead.
func (*MaterialCategoryInfo) Descriptor() ([]byte, []int) {
	return file_material_category_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialCategoryInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialCategoryInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialCategoryInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaterialCategoryInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryMaterialCategoryRequest struct {
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

func (x *QueryMaterialCategoryRequest) Reset() {
	*x = QueryMaterialCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialCategoryRequest) ProtoMessage() {}

func (x *QueryMaterialCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialCategoryRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialCategoryRequest) Descriptor() ([]byte, []int) {
	return file_material_category_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialCategoryRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialCategoryRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialCategoryRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialCategoryRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type QueryMaterialCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialCategoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                   `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                   `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                   `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialCategoryResponse) Reset() {
	*x = QueryMaterialCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialCategoryResponse) ProtoMessage() {}

func (x *QueryMaterialCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialCategoryResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialCategoryResponse) Descriptor() ([]byte, []int) {
	return file_material_category_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialCategoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialCategoryResponse) GetData() []*MaterialCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialCategoryResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialCategoryResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialCategoryResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialCategoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialCategoryResponse) Reset() {
	*x = GetAllMaterialCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialCategoryResponse) ProtoMessage() {}

func (x *GetAllMaterialCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialCategoryResponse) Descriptor() ([]byte, []int) {
	return file_material_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialCategoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialCategoryResponse) GetData() []*MaterialCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialCategoryDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialCategoryInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialCategoryDetailResponse) Reset() {
	*x = GetMaterialCategoryDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialCategoryDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialCategoryDetailResponse) ProtoMessage() {}

func (x *GetMaterialCategoryDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialCategoryDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialCategoryDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialCategoryDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialCategoryDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialCategoryDetailResponse) GetData() *MaterialCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_category_proto protoreflect.FileDescriptor

var file_material_category_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x98, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_category_proto_rawDescOnce sync.Once
	file_material_category_proto_rawDescData = file_material_category_proto_rawDesc
)

func file_material_category_proto_rawDescGZIP() []byte {
	file_material_category_proto_rawDescOnce.Do(func() {
		file_material_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_category_proto_rawDescData)
	})
	return file_material_category_proto_rawDescData
}

var file_material_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_material_category_proto_goTypes = []interface{}{
	(*MaterialCategoryInfo)(nil),              // 0: proto.MaterialCategoryInfo
	(*QueryMaterialCategoryRequest)(nil),      // 1: proto.QueryMaterialCategoryRequest
	(*QueryMaterialCategoryResponse)(nil),     // 2: proto.QueryMaterialCategoryResponse
	(*GetAllMaterialCategoryResponse)(nil),    // 3: proto.GetAllMaterialCategoryResponse
	(*GetMaterialCategoryDetailResponse)(nil), // 4: proto.GetMaterialCategoryDetailResponse
	(Code)(0), // 5: proto.Code
}
var file_material_category_proto_depIdxs = []int32{
	5, // 0: proto.QueryMaterialCategoryResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryMaterialCategoryResponse.data:type_name -> proto.MaterialCategoryInfo
	5, // 2: proto.GetAllMaterialCategoryResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllMaterialCategoryResponse.data:type_name -> proto.MaterialCategoryInfo
	5, // 4: proto.GetMaterialCategoryDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetMaterialCategoryDetailResponse.data:type_name -> proto.MaterialCategoryInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_material_category_proto_init() }
func file_material_category_proto_init() {
	if File_material_category_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialCategoryInfo); i {
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
		file_material_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialCategoryRequest); i {
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
		file_material_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialCategoryResponse); i {
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
		file_material_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialCategoryResponse); i {
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
		file_material_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialCategoryDetailResponse); i {
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
			RawDescriptor: file_material_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_category_proto_goTypes,
		DependencyIndexes: file_material_category_proto_depIdxs,
		MessageInfos:      file_material_category_proto_msgTypes,
	}.Build()
	File_material_category_proto = out.File
	file_material_category_proto_rawDesc = nil
	file_material_category_proto_goTypes = nil
	file_material_category_proto_depIdxs = nil
}
