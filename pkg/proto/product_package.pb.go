// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: product_package.proto

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

type ProductPackageInfo struct {
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
	// 包装数量
	PackageQuantity int32 `protobuf:"varint,6,opt,name=packageQuantity,proto3" json:"packageQuantity"`
	// 数量单位
	QuantityUnit string `protobuf:"bytes,7,opt,name=quantityUnit,proto3" json:"quantityUnit"`
	// 毛重
	GrossWeight float32 `protobuf:"fixed32,8,opt,name=grossWeight,proto3" json:"grossWeight"`
	// 净重
	NetWeight float32 `protobuf:"fixed32,9,opt,name=netWeight,proto3" json:"netWeight"`
	// 重量单位
	WeightUnit string `protobuf:"bytes,10,opt,name=weightUnit,proto3" json:"weightUnit"`
	// 尺寸
	Measure string `protobuf:"bytes,11,opt,name=measure,proto3" json:"measure"`
	// 尺寸单位
	MeasureUnit string `protobuf:"bytes,12,opt,name=measureUnit,proto3" json:"measureUnit"`
	// 备注
	Remark string `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark"`
	// 包装类型ID
	ProductPackageTypeID string                  `protobuf:"bytes,14,opt,name=productPackageTypeID,proto3" json:"productPackageTypeID"`
	ProductPackageType   *ProductPackageTypeInfo `protobuf:"bytes,15,opt,name=productPackageType,proto3" json:"productPackageType"`
}

func (x *ProductPackageInfo) Reset() {
	*x = ProductPackageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPackageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPackageInfo) ProtoMessage() {}

func (x *ProductPackageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPackageInfo.ProtoReflect.Descriptor instead.
func (*ProductPackageInfo) Descriptor() ([]byte, []int) {
	return file_product_package_proto_rawDescGZIP(), []int{0}
}

func (x *ProductPackageInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductPackageInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductPackageInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductPackageInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProductPackageInfo) GetPackageQuantity() int32 {
	if x != nil {
		return x.PackageQuantity
	}
	return 0
}

func (x *ProductPackageInfo) GetQuantityUnit() string {
	if x != nil {
		return x.QuantityUnit
	}
	return ""
}

func (x *ProductPackageInfo) GetGrossWeight() float32 {
	if x != nil {
		return x.GrossWeight
	}
	return 0
}

func (x *ProductPackageInfo) GetNetWeight() float32 {
	if x != nil {
		return x.NetWeight
	}
	return 0
}

func (x *ProductPackageInfo) GetWeightUnit() string {
	if x != nil {
		return x.WeightUnit
	}
	return ""
}

func (x *ProductPackageInfo) GetMeasure() string {
	if x != nil {
		return x.Measure
	}
	return ""
}

func (x *ProductPackageInfo) GetMeasureUnit() string {
	if x != nil {
		return x.MeasureUnit
	}
	return ""
}

func (x *ProductPackageInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductPackageInfo) GetProductPackageTypeID() string {
	if x != nil {
		return x.ProductPackageTypeID
	}
	return ""
}

func (x *ProductPackageInfo) GetProductPackageType() *ProductPackageTypeInfo {
	if x != nil {
		return x.ProductPackageType
	}
	return nil
}

type QueryProductPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 代号
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProductPackageRequest) Reset() {
	*x = QueryProductPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductPackageRequest) ProtoMessage() {}

func (x *QueryProductPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductPackageRequest.ProtoReflect.Descriptor instead.
func (*QueryProductPackageRequest) Descriptor() ([]byte, []int) {
	return file_product_package_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductPackageRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductPackageRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductPackageRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductPackageRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductPackageInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                 `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                 `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductPackageResponse) Reset() {
	*x = QueryProductPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductPackageResponse) ProtoMessage() {}

func (x *QueryProductPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductPackageResponse.ProtoReflect.Descriptor instead.
func (*QueryProductPackageResponse) Descriptor() ([]byte, []int) {
	return file_product_package_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductPackageResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductPackageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductPackageResponse) GetData() []*ProductPackageInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductPackageResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductPackageResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductPackageResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductPackageInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductPackageResponse) Reset() {
	*x = GetAllProductPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductPackageResponse) ProtoMessage() {}

func (x *GetAllProductPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductPackageResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductPackageResponse) Descriptor() ([]byte, []int) {
	return file_product_package_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductPackageResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductPackageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductPackageResponse) GetData() []*ProductPackageInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductPackageDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductPackageInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductPackageDetailResponse) Reset() {
	*x = GetProductPackageDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductPackageDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductPackageDetailResponse) ProtoMessage() {}

func (x *GetProductPackageDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductPackageDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductPackageDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_package_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductPackageDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductPackageDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductPackageDetailResponse) GetData() *ProductPackageInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_package_proto protoreflect.FileDescriptor

var file_product_package_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x03, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55,
	0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x73, 0x73,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x67, 0x72,
	0x6f, 0x73, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6e, 0x65,
	0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x32, 0x0a, 0x14, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x4d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8a,
	0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x1b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_package_proto_rawDescOnce sync.Once
	file_product_package_proto_rawDescData = file_product_package_proto_rawDesc
)

func file_product_package_proto_rawDescGZIP() []byte {
	file_product_package_proto_rawDescOnce.Do(func() {
		file_product_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_package_proto_rawDescData)
	})
	return file_product_package_proto_rawDescData
}

var file_product_package_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_package_proto_goTypes = []interface{}{
	(*ProductPackageInfo)(nil),              // 0: proto.ProductPackageInfo
	(*QueryProductPackageRequest)(nil),      // 1: proto.QueryProductPackageRequest
	(*QueryProductPackageResponse)(nil),     // 2: proto.QueryProductPackageResponse
	(*GetAllProductPackageResponse)(nil),    // 3: proto.GetAllProductPackageResponse
	(*GetProductPackageDetailResponse)(nil), // 4: proto.GetProductPackageDetailResponse
	(*ProductPackageTypeInfo)(nil),          // 5: proto.ProductPackageTypeInfo
	(Code)(0),                               // 6: proto.Code
}
var file_product_package_proto_depIdxs = []int32{
	5, // 0: proto.ProductPackageInfo.productPackageType:type_name -> proto.ProductPackageTypeInfo
	6, // 1: proto.QueryProductPackageResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductPackageResponse.data:type_name -> proto.ProductPackageInfo
	6, // 3: proto.GetAllProductPackageResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductPackageResponse.data:type_name -> proto.ProductPackageInfo
	6, // 5: proto.GetProductPackageDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductPackageDetailResponse.data:type_name -> proto.ProductPackageInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_product_package_proto_init() }
func file_product_package_proto_init() {
	if File_product_package_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_product_package_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPackageInfo); i {
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
		file_product_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductPackageRequest); i {
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
		file_product_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductPackageResponse); i {
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
		file_product_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductPackageResponse); i {
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
		file_product_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductPackageDetailResponse); i {
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
			RawDescriptor: file_product_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_package_proto_goTypes,
		DependencyIndexes: file_product_package_proto_depIdxs,
		MessageInfos:      file_product_package_proto_msgTypes,
	}.Build()
	File_product_package_proto = out.File
	file_product_package_proto_rawDesc = nil
	file_product_package_proto_goTypes = nil
	file_product_package_proto_depIdxs = nil
}
