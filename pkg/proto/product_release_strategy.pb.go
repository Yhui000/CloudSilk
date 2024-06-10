// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_release_strategy.proto

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

type ProductReleaseStrategyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 工厂产线ID
	ProductionLineID string `protobuf:"bytes,3,opt,name=productionLineID,proto3" json:"productionLineID"`
	// 产品类别ID
	ProductCategoryID   string `protobuf:"bytes,4,opt,name=productCategoryID,proto3" json:"productCategoryID"`
	ProductCategoryName string `protobuf:"bytes,7,opt,name=productCategoryName,proto3" json:"productCategoryName"`
	// 投料方式
	ReleaseMethod int32 `protobuf:"varint,5,opt,name=releaseMethod,proto3" json:"releaseMethod"`
	// 备注
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
	// 产品类别特性
	ProductReleaseStrategyComparableAttributes []*ProductReleaseStrategyComparableAttributeInfo `protobuf:"bytes,8,rep,name=ProductReleaseStrategyComparableAttributes,proto3" json:"ProductReleaseStrategyComparableAttributes"`
}

func (x *ProductReleaseStrategyInfo) Reset() {
	*x = ProductReleaseStrategyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReleaseStrategyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReleaseStrategyInfo) ProtoMessage() {}

func (x *ProductReleaseStrategyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReleaseStrategyInfo.ProtoReflect.Descriptor instead.
func (*ProductReleaseStrategyInfo) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *ProductReleaseStrategyInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReleaseStrategyInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *ProductReleaseStrategyInfo) GetProductCategoryID() string {
	if x != nil {
		return x.ProductCategoryID
	}
	return ""
}

func (x *ProductReleaseStrategyInfo) GetProductCategoryName() string {
	if x != nil {
		return x.ProductCategoryName
	}
	return ""
}

func (x *ProductReleaseStrategyInfo) GetReleaseMethod() int32 {
	if x != nil {
		return x.ReleaseMethod
	}
	return 0
}

func (x *ProductReleaseStrategyInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductReleaseStrategyInfo) GetProductReleaseStrategyComparableAttributes() []*ProductReleaseStrategyComparableAttributeInfo {
	if x != nil {
		return x.ProductReleaseStrategyComparableAttributes
	}
	return nil
}

type ProductReleaseStrategyComparableAttributeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	// 产品发布策略ID
	ProductReleaseStrategyID string `protobuf:"bytes,2,opt,name=productReleaseStrategyID,proto3" json:"productReleaseStrategyID"`
	// 产品类别特性ID
	ProductCategoryAttributeID string `protobuf:"bytes,3,opt,name=productCategoryAttributeID,proto3" json:"productCategoryAttributeID"`
}

func (x *ProductReleaseStrategyComparableAttributeInfo) Reset() {
	*x = ProductReleaseStrategyComparableAttributeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReleaseStrategyComparableAttributeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReleaseStrategyComparableAttributeInfo) ProtoMessage() {}

func (x *ProductReleaseStrategyComparableAttributeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReleaseStrategyComparableAttributeInfo.ProtoReflect.Descriptor instead.
func (*ProductReleaseStrategyComparableAttributeInfo) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *ProductReleaseStrategyComparableAttributeInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReleaseStrategyComparableAttributeInfo) GetProductReleaseStrategyID() string {
	if x != nil {
		return x.ProductReleaseStrategyID
	}
	return ""
}

func (x *ProductReleaseStrategyComparableAttributeInfo) GetProductCategoryAttributeID() string {
	if x != nil {
		return x.ProductCategoryAttributeID
	}
	return ""
}

type QueryProductReleaseStrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,4,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// 类别或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProductReleaseStrategyRequest) Reset() {
	*x = QueryProductReleaseStrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReleaseStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReleaseStrategyRequest) ProtoMessage() {}

func (x *QueryProductReleaseStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReleaseStrategyRequest.ProtoReflect.Descriptor instead.
func (*QueryProductReleaseStrategyRequest) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductReleaseStrategyRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductReleaseStrategyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductReleaseStrategyRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductReleaseStrategyRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryProductReleaseStrategyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductReleaseStrategyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReleaseStrategyInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                         `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                         `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                         `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductReleaseStrategyResponse) Reset() {
	*x = QueryProductReleaseStrategyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReleaseStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReleaseStrategyResponse) ProtoMessage() {}

func (x *QueryProductReleaseStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReleaseStrategyResponse.ProtoReflect.Descriptor instead.
func (*QueryProductReleaseStrategyResponse) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductReleaseStrategyResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductReleaseStrategyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductReleaseStrategyResponse) GetData() []*ProductReleaseStrategyInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductReleaseStrategyResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductReleaseStrategyResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductReleaseStrategyResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductReleaseStrategyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReleaseStrategyInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductReleaseStrategyResponse) Reset() {
	*x = GetAllProductReleaseStrategyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReleaseStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReleaseStrategyResponse) ProtoMessage() {}

func (x *GetAllProductReleaseStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReleaseStrategyResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductReleaseStrategyResponse) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductReleaseStrategyResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductReleaseStrategyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductReleaseStrategyResponse) GetData() []*ProductReleaseStrategyInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductReleaseStrategyDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductReleaseStrategyInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductReleaseStrategyDetailResponse) Reset() {
	*x = GetProductReleaseStrategyDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_release_strategy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReleaseStrategyDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReleaseStrategyDetailResponse) ProtoMessage() {}

func (x *GetProductReleaseStrategyDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_release_strategy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReleaseStrategyDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductReleaseStrategyDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_release_strategy_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductReleaseStrategyDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductReleaseStrategyDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductReleaseStrategyDetailResponse) GetData() *ProductReleaseStrategyInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_release_strategy_proto protoreflect.FileDescriptor

var file_product_release_strategy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x1a, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x94, 0x01, 0x0a, 0x2a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x2a, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x2d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x18, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x44, 0x22, 0xbe, 0x01, 0x0a, 0x22, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x23, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x98, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x9b, 0x01, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_release_strategy_proto_rawDescOnce sync.Once
	file_product_release_strategy_proto_rawDescData = file_product_release_strategy_proto_rawDesc
)

func file_product_release_strategy_proto_rawDescGZIP() []byte {
	file_product_release_strategy_proto_rawDescOnce.Do(func() {
		file_product_release_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_release_strategy_proto_rawDescData)
	})
	return file_product_release_strategy_proto_rawDescData
}

var file_product_release_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_release_strategy_proto_goTypes = []interface{}{
	(*ProductReleaseStrategyInfo)(nil),                    // 0: proto.ProductReleaseStrategyInfo
	(*ProductReleaseStrategyComparableAttributeInfo)(nil), // 1: proto.ProductReleaseStrategyComparableAttributeInfo
	(*QueryProductReleaseStrategyRequest)(nil),            // 2: proto.QueryProductReleaseStrategyRequest
	(*QueryProductReleaseStrategyResponse)(nil),           // 3: proto.QueryProductReleaseStrategyResponse
	(*GetAllProductReleaseStrategyResponse)(nil),          // 4: proto.GetAllProductReleaseStrategyResponse
	(*GetProductReleaseStrategyDetailResponse)(nil),       // 5: proto.GetProductReleaseStrategyDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_product_release_strategy_proto_depIdxs = []int32{
	1, // 0: proto.ProductReleaseStrategyInfo.ProductReleaseStrategyComparableAttributes:type_name -> proto.ProductReleaseStrategyComparableAttributeInfo
	6, // 1: proto.QueryProductReleaseStrategyResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductReleaseStrategyResponse.data:type_name -> proto.ProductReleaseStrategyInfo
	6, // 3: proto.GetAllProductReleaseStrategyResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductReleaseStrategyResponse.data:type_name -> proto.ProductReleaseStrategyInfo
	6, // 5: proto.GetProductReleaseStrategyDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductReleaseStrategyDetailResponse.data:type_name -> proto.ProductReleaseStrategyInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_product_release_strategy_proto_init() }
func file_product_release_strategy_proto_init() {
	if File_product_release_strategy_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_release_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReleaseStrategyInfo); i {
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
		file_product_release_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReleaseStrategyComparableAttributeInfo); i {
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
		file_product_release_strategy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReleaseStrategyRequest); i {
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
		file_product_release_strategy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReleaseStrategyResponse); i {
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
		file_product_release_strategy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductReleaseStrategyResponse); i {
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
		file_product_release_strategy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReleaseStrategyDetailResponse); i {
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
			RawDescriptor: file_product_release_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_release_strategy_proto_goTypes,
		DependencyIndexes: file_product_release_strategy_proto_depIdxs,
		MessageInfos:      file_product_release_strategy_proto_msgTypes,
	}.Build()
	File_product_release_strategy_proto = out.File
	file_product_release_strategy_proto_rawDesc = nil
	file_product_release_strategy_proto_goTypes = nil
	file_product_release_strategy_proto_depIdxs = nil
}
