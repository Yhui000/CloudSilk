// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_issue_record.proto

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

type ProductIssueRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 物料追溯号
	MaterialTraceNo string `protobuf:"bytes,3,opt,name=materialTraceNo,proto3" json:"materialTraceNo"`
	// 发料时间
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime"`
	// 发料人员ID
	CreateUserID string `protobuf:"bytes,5,opt,name=createUserID,proto3" json:"createUserID"`
	// 当前状态
	CurrentState string `protobuf:"bytes,6,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,7,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 绑定产品ID
	ProductInfoID string `protobuf:"bytes,8,opt,name=productInfoID,proto3" json:"productInfoID"`
	// 工单物料ID
	ProductOrderBomID string `protobuf:"bytes,9,opt,name=productOrderBomID,proto3" json:"productOrderBomID"`
	// 发料工序ID
	IssuanceProcessID string `protobuf:"bytes,10,opt,name=issuanceProcessID,proto3" json:"issuanceProcessID"`
	// 产品序列号
	ProductSerialNo string `protobuf:"bytes,11,opt,name=productSerialNo,proto3" json:"productSerialNo"`
	// 生产工单号
	ProductOrderNo string `protobuf:"bytes,12,opt,name=productOrderNo,proto3" json:"productOrderNo"`
	// 物料号
	MaterialNo string `protobuf:"bytes,13,opt,name=materialNo,proto3" json:"materialNo"`
	// 物料描述
	MaterialDescription string `protobuf:"bytes,14,opt,name=materialDescription,proto3" json:"materialDescription"`
	// 发料工序
	ProductionProcess string `protobuf:"bytes,15,opt,name=productionProcess,proto3" json:"productionProcess"`
}

func (x *ProductIssueRecordInfo) Reset() {
	*x = ProductIssueRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_issue_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductIssueRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductIssueRecordInfo) ProtoMessage() {}

func (x *ProductIssueRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_issue_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductIssueRecordInfo.ProtoReflect.Descriptor instead.
func (*ProductIssueRecordInfo) Descriptor() ([]byte, []int) {
	return file_product_issue_record_proto_rawDescGZIP(), []int{0}
}

func (x *ProductIssueRecordInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetMaterialTraceNo() string {
	if x != nil {
		return x.MaterialTraceNo
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetCreateUserID() string {
	if x != nil {
		return x.CreateUserID
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetProductOrderBomID() string {
	if x != nil {
		return x.ProductOrderBomID
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetIssuanceProcessID() string {
	if x != nil {
		return x.IssuanceProcessID
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetMaterialNo() string {
	if x != nil {
		return x.MaterialNo
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetMaterialDescription() string {
	if x != nil {
		return x.MaterialDescription
	}
	return ""
}

func (x *ProductIssueRecordInfo) GetProductionProcess() string {
	if x != nil {
		return x.ProductionProcess
	}
	return ""
}

type QueryProductIssueRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 产品序列号
	// @inject_tag: uri:"productSerialNo" form:"productSerialNo"
	ProductSerialNo string `protobuf:"bytes,4,opt,name=productSerialNo,proto3" json:"productSerialNo" uri:"productSerialNo" form:"productSerialNo"`
	// 生产工单号
	// @inject_tag: uri:"productOrderNo" form:"productOrderNo"
	ProductOrderNo string `protobuf:"bytes,5,opt,name=productOrderNo,proto3" json:"productOrderNo" uri:"productOrderNo" form:"productOrderNo"`
	// @inject_tag: uri:"materialDescription" form:"materialDescription"
	MaterialDescription string `protobuf:"bytes,6,opt,name=materialDescription,proto3" json:"materialDescription" uri:"materialDescription" form:"materialDescription"`
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,7,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,8,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
	// @inject_tag: uri:"productionProcess" form:"productionProcess"
	ProductionProcess string `protobuf:"bytes,9,opt,name=productionProcess,proto3" json:"productionProcess" uri:"productionProcess" form:"productionProcess"`
}

func (x *QueryProductIssueRecordRequest) Reset() {
	*x = QueryProductIssueRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_issue_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductIssueRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductIssueRecordRequest) ProtoMessage() {}

func (x *QueryProductIssueRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_issue_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductIssueRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryProductIssueRecordRequest) Descriptor() ([]byte, []int) {
	return file_product_issue_record_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductIssueRecordRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductIssueRecordRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductIssueRecordRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetMaterialDescription() string {
	if x != nil {
		return x.MaterialDescription
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

func (x *QueryProductIssueRecordRequest) GetProductionProcess() string {
	if x != nil {
		return x.ProductionProcess
	}
	return ""
}

type QueryProductIssueRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductIssueRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                     `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                     `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                     `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductIssueRecordResponse) Reset() {
	*x = QueryProductIssueRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_issue_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductIssueRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductIssueRecordResponse) ProtoMessage() {}

func (x *QueryProductIssueRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_issue_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductIssueRecordResponse.ProtoReflect.Descriptor instead.
func (*QueryProductIssueRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_issue_record_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductIssueRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductIssueRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductIssueRecordResponse) GetData() []*ProductIssueRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductIssueRecordResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductIssueRecordResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductIssueRecordResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductIssueRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductIssueRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductIssueRecordResponse) Reset() {
	*x = GetAllProductIssueRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_issue_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductIssueRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductIssueRecordResponse) ProtoMessage() {}

func (x *GetAllProductIssueRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_issue_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductIssueRecordResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductIssueRecordResponse) Descriptor() ([]byte, []int) {
	return file_product_issue_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductIssueRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductIssueRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductIssueRecordResponse) GetData() []*ProductIssueRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductIssueRecordDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductIssueRecordInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductIssueRecordDetailResponse) Reset() {
	*x = GetProductIssueRecordDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_issue_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductIssueRecordDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductIssueRecordDetailResponse) ProtoMessage() {}

func (x *GetProductIssueRecordDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_issue_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductIssueRecordDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductIssueRecordDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_issue_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductIssueRecordDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductIssueRecordDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductIssueRecordDetailResponse) GetData() *ProductIssueRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_issue_record_proto protoreflect.FileDescriptor

var file_product_issue_record_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x04, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x2c,
	0x0a, 0x11, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x73, 0x73, 0x75, 0x61,
	0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x30,
	0x0a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0xf0,
	0x02, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x30, 0x0a,
	0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x30, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x31,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x31, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x22, 0xd5, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a,
	0x23, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_issue_record_proto_rawDescOnce sync.Once
	file_product_issue_record_proto_rawDescData = file_product_issue_record_proto_rawDesc
)

func file_product_issue_record_proto_rawDescGZIP() []byte {
	file_product_issue_record_proto_rawDescOnce.Do(func() {
		file_product_issue_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_issue_record_proto_rawDescData)
	})
	return file_product_issue_record_proto_rawDescData
}

var file_product_issue_record_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_issue_record_proto_goTypes = []interface{}{
	(*ProductIssueRecordInfo)(nil),              // 0: proto.ProductIssueRecordInfo
	(*QueryProductIssueRecordRequest)(nil),      // 1: proto.QueryProductIssueRecordRequest
	(*QueryProductIssueRecordResponse)(nil),     // 2: proto.QueryProductIssueRecordResponse
	(*GetAllProductIssueRecordResponse)(nil),    // 3: proto.GetAllProductIssueRecordResponse
	(*GetProductIssueRecordDetailResponse)(nil), // 4: proto.GetProductIssueRecordDetailResponse
	(Code)(0), // 5: proto.Code
}
var file_product_issue_record_proto_depIdxs = []int32{
	5, // 0: proto.QueryProductIssueRecordResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryProductIssueRecordResponse.data:type_name -> proto.ProductIssueRecordInfo
	5, // 2: proto.GetAllProductIssueRecordResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllProductIssueRecordResponse.data:type_name -> proto.ProductIssueRecordInfo
	5, // 4: proto.GetProductIssueRecordDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetProductIssueRecordDetailResponse.data:type_name -> proto.ProductIssueRecordInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_product_issue_record_proto_init() }
func file_product_issue_record_proto_init() {
	if File_product_issue_record_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_issue_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductIssueRecordInfo); i {
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
		file_product_issue_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductIssueRecordRequest); i {
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
		file_product_issue_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductIssueRecordResponse); i {
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
		file_product_issue_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductIssueRecordResponse); i {
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
		file_product_issue_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductIssueRecordDetailResponse); i {
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
			RawDescriptor: file_product_issue_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_issue_record_proto_goTypes,
		DependencyIndexes: file_product_issue_record_proto_depIdxs,
		MessageInfos:      file_product_issue_record_proto_msgTypes,
	}.Build()
	File_product_issue_record_proto = out.File
	file_product_issue_record_proto_rawDesc = nil
	file_product_issue_record_proto_goTypes = nil
	file_product_issue_record_proto_depIdxs = nil
}
