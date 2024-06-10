// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_model_bom.proto

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

type ProductModelBomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 项目号
	ItemNo string `protobuf:"bytes,3,opt,name=itemNo,proto3" json:"itemNo"`
	// 物料号
	MaterialNo string `protobuf:"bytes,4,opt,name=materialNo,proto3" json:"materialNo"`
	// 物料描述
	MaterialDescription string `protobuf:"bytes,5,opt,name=materialDescription,proto3" json:"materialDescription"`
	// 需求数量
	RequireQTY float32 `protobuf:"fixed32,6,opt,name=requireQTY,proto3" json:"requireQTY"`
	// 单位
	Unit string `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit"`
	// 需求工序
	ProductionProcess string `protobuf:"bytes,8,opt,name=productionProcess,proto3" json:"productionProcess"`
	// 备注
	Remark string `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark"`
	// 产品型号ID
	ProductModelID string `protobuf:"bytes,10,opt,name=productModelID,proto3" json:"productModelID"`
	// 产品类别ID
	ProductCategoryID string `protobuf:"bytes,11,opt,name=productCategoryID,proto3" json:"productCategoryID"`
}

func (x *ProductModelBomInfo) Reset() {
	*x = ProductModelBomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_bom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductModelBomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductModelBomInfo) ProtoMessage() {}

func (x *ProductModelBomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_bom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductModelBomInfo.ProtoReflect.Descriptor instead.
func (*ProductModelBomInfo) Descriptor() ([]byte, []int) {
	return file_product_model_bom_proto_rawDescGZIP(), []int{0}
}

func (x *ProductModelBomInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductModelBomInfo) GetItemNo() string {
	if x != nil {
		return x.ItemNo
	}
	return ""
}

func (x *ProductModelBomInfo) GetMaterialNo() string {
	if x != nil {
		return x.MaterialNo
	}
	return ""
}

func (x *ProductModelBomInfo) GetMaterialDescription() string {
	if x != nil {
		return x.MaterialDescription
	}
	return ""
}

func (x *ProductModelBomInfo) GetRequireQTY() float32 {
	if x != nil {
		return x.RequireQTY
	}
	return 0
}

func (x *ProductModelBomInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ProductModelBomInfo) GetProductionProcess() string {
	if x != nil {
		return x.ProductionProcess
	}
	return ""
}

func (x *ProductModelBomInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductModelBomInfo) GetProductModelID() string {
	if x != nil {
		return x.ProductModelID
	}
	return ""
}

func (x *ProductModelBomInfo) GetProductCategoryID() string {
	if x != nil {
		return x.ProductCategoryID
	}
	return ""
}

type QueryProductModelBomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 物料号
	// @inject_tag: uri:"materialNo" form:"materialNo"
	MaterialNo string `protobuf:"bytes,6,opt,name=materialNo,proto3" json:"materialNo" uri:"materialNo" form:"materialNo"`
	// 产品类别ID
	// @inject_tag: uri:"productCategoryID" form:"productCategoryID"
	ProductCategoryID string `protobuf:"bytes,7,opt,name=productCategoryID,proto3" json:"productCategoryID" uri:"productCategoryID" form:"productCategoryID"`
	// 产品型号ID
	// @inject_tag: uri:"productModelID" form:"productModelID"
	ProductModelID string `protobuf:"bytes,8,opt,name=productModelID,proto3" json:"productModelID" uri:"productModelID" form:"productModelID"`
}

func (x *QueryProductModelBomRequest) Reset() {
	*x = QueryProductModelBomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_bom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductModelBomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductModelBomRequest) ProtoMessage() {}

func (x *QueryProductModelBomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_bom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductModelBomRequest.ProtoReflect.Descriptor instead.
func (*QueryProductModelBomRequest) Descriptor() ([]byte, []int) {
	return file_product_model_bom_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductModelBomRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductModelBomRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductModelBomRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductModelBomRequest) GetMaterialNo() string {
	if x != nil {
		return x.MaterialNo
	}
	return ""
}

func (x *QueryProductModelBomRequest) GetProductCategoryID() string {
	if x != nil {
		return x.ProductCategoryID
	}
	return ""
}

func (x *QueryProductModelBomRequest) GetProductModelID() string {
	if x != nil {
		return x.ProductModelID
	}
	return ""
}

type QueryProductModelBomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductModelBomInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                  `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                  `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                  `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductModelBomResponse) Reset() {
	*x = QueryProductModelBomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_bom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductModelBomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductModelBomResponse) ProtoMessage() {}

func (x *QueryProductModelBomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_bom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductModelBomResponse.ProtoReflect.Descriptor instead.
func (*QueryProductModelBomResponse) Descriptor() ([]byte, []int) {
	return file_product_model_bom_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductModelBomResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductModelBomResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductModelBomResponse) GetData() []*ProductModelBomInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductModelBomResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductModelBomResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductModelBomResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductModelBomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductModelBomInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductModelBomResponse) Reset() {
	*x = GetAllProductModelBomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_bom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductModelBomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductModelBomResponse) ProtoMessage() {}

func (x *GetAllProductModelBomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_bom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductModelBomResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductModelBomResponse) Descriptor() ([]byte, []int) {
	return file_product_model_bom_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductModelBomResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductModelBomResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductModelBomResponse) GetData() []*ProductModelBomInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductModelBomDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                 `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductModelBomInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductModelBomDetailResponse) Reset() {
	*x = GetProductModelBomDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_bom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductModelBomDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductModelBomDetailResponse) ProtoMessage() {}

func (x *GetProductModelBomDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_bom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductModelBomDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductModelBomDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_model_bom_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductModelBomDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductModelBomDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductModelBomDetailResponse) GetData() *ProductModelBomInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_model_bom_proto protoreflect.FileDescriptor

var file_product_model_bom_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x62, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x4e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x6f, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x51,
	0x54, 0x59, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x51, 0x54, 0x59, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x22, 0xed, 0x01, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x2c,
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x44, 0x22, 0xcf, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_model_bom_proto_rawDescOnce sync.Once
	file_product_model_bom_proto_rawDescData = file_product_model_bom_proto_rawDesc
)

func file_product_model_bom_proto_rawDescGZIP() []byte {
	file_product_model_bom_proto_rawDescOnce.Do(func() {
		file_product_model_bom_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_model_bom_proto_rawDescData)
	})
	return file_product_model_bom_proto_rawDescData
}

var file_product_model_bom_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_model_bom_proto_goTypes = []interface{}{
	(*ProductModelBomInfo)(nil),              // 0: proto.ProductModelBomInfo
	(*QueryProductModelBomRequest)(nil),      // 1: proto.QueryProductModelBomRequest
	(*QueryProductModelBomResponse)(nil),     // 2: proto.QueryProductModelBomResponse
	(*GetAllProductModelBomResponse)(nil),    // 3: proto.GetAllProductModelBomResponse
	(*GetProductModelBomDetailResponse)(nil), // 4: proto.GetProductModelBomDetailResponse
	(Code)(0),                                // 5: proto.Code
}
var file_product_model_bom_proto_depIdxs = []int32{
	5, // 0: proto.QueryProductModelBomResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryProductModelBomResponse.data:type_name -> proto.ProductModelBomInfo
	5, // 2: proto.GetAllProductModelBomResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllProductModelBomResponse.data:type_name -> proto.ProductModelBomInfo
	5, // 4: proto.GetProductModelBomDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetProductModelBomDetailResponse.data:type_name -> proto.ProductModelBomInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_product_model_bom_proto_init() }
func file_product_model_bom_proto_init() {
	if File_product_model_bom_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_model_bom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductModelBomInfo); i {
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
		file_product_model_bom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductModelBomRequest); i {
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
		file_product_model_bom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductModelBomResponse); i {
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
		file_product_model_bom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductModelBomResponse); i {
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
		file_product_model_bom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductModelBomDetailResponse); i {
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
			RawDescriptor: file_product_model_bom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_model_bom_proto_goTypes,
		DependencyIndexes: file_product_model_bom_proto_depIdxs,
		MessageInfos:      file_product_model_bom_proto_msgTypes,
	}.Build()
	File_product_model_bom_proto = out.File
	file_product_model_bom_proto_rawDesc = nil
	file_product_model_bom_proto_goTypes = nil
	file_product_model_bom_proto_depIdxs = nil
}
