// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_category.proto

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

type ProductCategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// 识别码
	Identifier string `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier"`
	// 是否授权
	IsAuthorized bool `protobuf:"varint,6,opt,name=isAuthorized,proto3" json:"isAuthorized"`
	// 特性表达式
	AttributeExpression string `protobuf:"bytes,7,opt,name=attributeExpression,proto3" json:"attributeExpression"`
	// 备注
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	// 产品品牌ID
	ProductBrandID   string `protobuf:"bytes,9,opt,name=productBrandID,proto3" json:"productBrandID"`
	ProductBrandName string `protobuf:"bytes,10,opt,name=productBrandName,proto3" json:"productBrandName"`
}

func (x *ProductCategoryInfo) Reset() {
	*x = ProductCategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategoryInfo) ProtoMessage() {}

func (x *ProductCategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategoryInfo.ProtoReflect.Descriptor instead.
func (*ProductCategoryInfo) Descriptor() ([]byte, []int) {
	return file_product_category_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCategoryInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductCategoryInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductCategoryInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductCategoryInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProductCategoryInfo) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

func (x *ProductCategoryInfo) GetAttributeExpression() string {
	if x != nil {
		return x.AttributeExpression
	}
	return ""
}

func (x *ProductCategoryInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductCategoryInfo) GetProductBrandID() string {
	if x != nil {
		return x.ProductBrandID
	}
	return ""
}

func (x *ProductCategoryInfo) GetProductBrandName() string {
	if x != nil {
		return x.ProductBrandName
	}
	return ""
}

type QueryProductCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 代号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 产品品牌ID
	// @inject_tag: uri:"productBrandID" form:"productBrandID"
	ProductBrandID string `protobuf:"bytes,11,opt,name=productBrandID,proto3" json:"productBrandID" uri:"productBrandID" form:"productBrandID"`
}

func (x *QueryProductCategoryRequest) Reset() {
	*x = QueryProductCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductCategoryRequest) ProtoMessage() {}

func (x *QueryProductCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductCategoryRequest.ProtoReflect.Descriptor instead.
func (*QueryProductCategoryRequest) Descriptor() ([]byte, []int) {
	return file_product_category_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductCategoryRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductCategoryRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductCategoryRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductCategoryRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryProductCategoryRequest) GetProductBrandID() string {
	if x != nil {
		return x.ProductBrandID
	}
	return ""
}

type QueryProductCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductCategoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                  `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                  `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                  `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductCategoryResponse) Reset() {
	*x = QueryProductCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductCategoryResponse) ProtoMessage() {}

func (x *QueryProductCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductCategoryResponse.ProtoReflect.Descriptor instead.
func (*QueryProductCategoryResponse) Descriptor() ([]byte, []int) {
	return file_product_category_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductCategoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductCategoryResponse) GetData() []*ProductCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductCategoryResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductCategoryResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductCategoryResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductCategoryInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductCategoryResponse) Reset() {
	*x = GetAllProductCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductCategoryResponse) ProtoMessage() {}

func (x *GetAllProductCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductCategoryResponse) Descriptor() ([]byte, []int) {
	return file_product_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductCategoryResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductCategoryResponse) GetData() []*ProductCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductCategoryDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                 `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductCategoryInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductCategoryDetailResponse) Reset() {
	*x = GetProductCategoryDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategoryDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategoryDetailResponse) ProtoMessage() {}

func (x *GetProductCategoryDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategoryDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductCategoryDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductCategoryDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductCategoryDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductCategoryDetailResponse) GetData() *ProductCategoryInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_category_proto protoreflect.FileDescriptor

var file_product_category_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xb3, 0x01, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49,
	0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x22, 0xcf, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_category_proto_rawDescOnce sync.Once
	file_product_category_proto_rawDescData = file_product_category_proto_rawDesc
)

func file_product_category_proto_rawDescGZIP() []byte {
	file_product_category_proto_rawDescOnce.Do(func() {
		file_product_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_category_proto_rawDescData)
	})
	return file_product_category_proto_rawDescData
}

var file_product_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_category_proto_goTypes = []interface{}{
	(*ProductCategoryInfo)(nil),              // 0: proto.ProductCategoryInfo
	(*QueryProductCategoryRequest)(nil),      // 1: proto.QueryProductCategoryRequest
	(*QueryProductCategoryResponse)(nil),     // 2: proto.QueryProductCategoryResponse
	(*GetAllProductCategoryResponse)(nil),    // 3: proto.GetAllProductCategoryResponse
	(*GetProductCategoryDetailResponse)(nil), // 4: proto.GetProductCategoryDetailResponse
	(Code)(0),                                // 5: proto.Code
}
var file_product_category_proto_depIdxs = []int32{
	5, // 0: proto.QueryProductCategoryResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryProductCategoryResponse.data:type_name -> proto.ProductCategoryInfo
	5, // 2: proto.GetAllProductCategoryResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllProductCategoryResponse.data:type_name -> proto.ProductCategoryInfo
	5, // 4: proto.GetProductCategoryDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetProductCategoryDetailResponse.data:type_name -> proto.ProductCategoryInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_product_category_proto_init() }
func file_product_category_proto_init() {
	if File_product_category_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategoryInfo); i {
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
		file_product_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductCategoryRequest); i {
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
		file_product_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductCategoryResponse); i {
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
		file_product_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductCategoryResponse); i {
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
		file_product_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategoryDetailResponse); i {
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
			RawDescriptor: file_product_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_category_proto_goTypes,
		DependencyIndexes: file_product_category_proto_depIdxs,
		MessageInfos:      file_product_category_proto_msgTypes,
	}.Build()
	File_product_category_proto = out.File
	file_product_category_proto_rawDesc = nil
	file_product_category_proto_goTypes = nil
	file_product_category_proto_depIdxs = nil
}
