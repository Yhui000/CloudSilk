// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_order_process_step.proto

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

type ProductOrderProcessStepInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 排序
	SortIndex int32 `protobuf:"varint,3,opt,name=sortIndex,proto3" json:"sortIndex"`
	// 作业描述
	WorkDescription string `protobuf:"bytes,4,opt,name=workDescription,proto3" json:"workDescription"`
	// 作业图示
	WorkGraphic string `protobuf:"bytes,5,opt,name=workGraphic,proto3" json:"workGraphic"`
	// 创建时间
	CreateTime string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime"`
	// 创建人员ID
	CreateUserID   string `protobuf:"bytes,7,opt,name=createUserID,proto3" json:"createUserID"`
	CreateUserName string `protobuf:"bytes,14,opt,name=createUserName,proto3" json:"createUserName"`
	// 备注
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	// 作业类型ID
	ProcessStepTypeID string `protobuf:"bytes,9,opt,name=processStepTypeID,proto3" json:"processStepTypeID"`
	// 工单工序ID
	ProductOrderProcessID string `protobuf:"bytes,10,opt,name=productOrderProcessID,proto3" json:"productOrderProcessID"`
	// 工步附件
	ProductOrderProcessStepAttachments []*ProductOrderProcessStepAttachmentInfo `protobuf:"bytes,11,rep,name=productOrderProcessStepAttachments,proto3" json:"productOrderProcessStepAttachments"`
	// 工布参数
	ProductOrderProcessStepTypeParameters []*ProductOrderProcessStepTypeParameterInfo `protobuf:"bytes,12,rep,name=productOrderProcessStepTypeParameters,proto3" json:"productOrderProcessStepTypeParameters"`
	// 作业类型
	ProcessStepType string `protobuf:"bytes,13,opt,name=processStepType,proto3" json:"processStepType"`
}

func (x *ProductOrderProcessStepInfo) Reset() {
	*x = ProductOrderProcessStepInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOrderProcessStepInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOrderProcessStepInfo) ProtoMessage() {}

func (x *ProductOrderProcessStepInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOrderProcessStepInfo.ProtoReflect.Descriptor instead.
func (*ProductOrderProcessStepInfo) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{0}
}

func (x *ProductOrderProcessStepInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetSortIndex() int32 {
	if x != nil {
		return x.SortIndex
	}
	return 0
}

func (x *ProductOrderProcessStepInfo) GetWorkDescription() string {
	if x != nil {
		return x.WorkDescription
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetWorkGraphic() string {
	if x != nil {
		return x.WorkGraphic
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetCreateUserID() string {
	if x != nil {
		return x.CreateUserID
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetCreateUserName() string {
	if x != nil {
		return x.CreateUserName
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetProcessStepTypeID() string {
	if x != nil {
		return x.ProcessStepTypeID
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetProductOrderProcessID() string {
	if x != nil {
		return x.ProductOrderProcessID
	}
	return ""
}

func (x *ProductOrderProcessStepInfo) GetProductOrderProcessStepAttachments() []*ProductOrderProcessStepAttachmentInfo {
	if x != nil {
		return x.ProductOrderProcessStepAttachments
	}
	return nil
}

func (x *ProductOrderProcessStepInfo) GetProductOrderProcessStepTypeParameters() []*ProductOrderProcessStepTypeParameterInfo {
	if x != nil {
		return x.ProductOrderProcessStepTypeParameters
	}
	return nil
}

func (x *ProductOrderProcessStepInfo) GetProcessStepType() string {
	if x != nil {
		return x.ProcessStepType
	}
	return ""
}

type ProductOrderProcessStepAttachmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 文件名
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName"`
	// 文件类型
	FileType string `protobuf:"bytes,4,opt,name=fileType,proto3" json:"fileType"`
	// 文件大小
	FileSize float32 `protobuf:"fixed32,5,opt,name=fileSize,proto3" json:"fileSize"`
	// 保存文件名
	SavedFileName string `protobuf:"bytes,6,opt,name=savedFileName,proto3" json:"savedFileName"`
	// 文件状态
	FileStatus string `protobuf:"bytes,7,opt,name=fileStatus,proto3" json:"fileStatus"`
	// 创建时间
	CreateTime string `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime"`
	// 创建人员ID
	CreateUserID string `protobuf:"bytes,9,opt,name=createUserID,proto3" json:"createUserID"`
	// 隶属工步ID
	ProductOrderProcessStepID string `protobuf:"bytes,10,opt,name=productOrderProcessStepID,proto3" json:"productOrderProcessStepID"`
}

func (x *ProductOrderProcessStepAttachmentInfo) Reset() {
	*x = ProductOrderProcessStepAttachmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOrderProcessStepAttachmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOrderProcessStepAttachmentInfo) ProtoMessage() {}

func (x *ProductOrderProcessStepAttachmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOrderProcessStepAttachmentInfo.ProtoReflect.Descriptor instead.
func (*ProductOrderProcessStepAttachmentInfo) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{1}
}

func (x *ProductOrderProcessStepAttachmentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetFileSize() float32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ProductOrderProcessStepAttachmentInfo) GetSavedFileName() string {
	if x != nil {
		return x.SavedFileName
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetFileStatus() string {
	if x != nil {
		return x.FileStatus
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetCreateUserID() string {
	if x != nil {
		return x.CreateUserID
	}
	return ""
}

func (x *ProductOrderProcessStepAttachmentInfo) GetProductOrderProcessStepID() string {
	if x != nil {
		return x.ProductOrderProcessStepID
	}
	return ""
}

type ProductOrderProcessStepTypeParameterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 参数值
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	// 备注
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark"`
	// 工单工步ID
	ProductOrderProcessStepID string `protobuf:"bytes,5,opt,name=productOrderProcessStepID,proto3" json:"productOrderProcessStepID"`
	// 作业参数ID
	ProcessStepTypeParameterID string `protobuf:"bytes,6,opt,name=processStepTypeParameterID,proto3" json:"processStepTypeParameterID"`
}

func (x *ProductOrderProcessStepTypeParameterInfo) Reset() {
	*x = ProductOrderProcessStepTypeParameterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOrderProcessStepTypeParameterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOrderProcessStepTypeParameterInfo) ProtoMessage() {}

func (x *ProductOrderProcessStepTypeParameterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOrderProcessStepTypeParameterInfo.ProtoReflect.Descriptor instead.
func (*ProductOrderProcessStepTypeParameterInfo) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{2}
}

func (x *ProductOrderProcessStepTypeParameterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductOrderProcessStepTypeParameterInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ProductOrderProcessStepTypeParameterInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductOrderProcessStepTypeParameterInfo) GetProductOrderProcessStepID() string {
	if x != nil {
		return x.ProductOrderProcessStepID
	}
	return ""
}

func (x *ProductOrderProcessStepTypeParameterInfo) GetProcessStepTypeParameterID() string {
	if x != nil {
		return x.ProcessStepTypeParameterID
	}
	return ""
}

type QueryProductOrderProcessStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 工单工序ID
	// @inject_tag: uri:"productOrderProcessID" form:"productOrderProcessID"
	ProductOrderProcessID string `protobuf:"bytes,4,opt,name=productOrderProcessID,proto3" json:"productOrderProcessID" uri:"productOrderProcessID" form:"productOrderProcessID"`
}

func (x *QueryProductOrderProcessStepRequest) Reset() {
	*x = QueryProductOrderProcessStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductOrderProcessStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductOrderProcessStepRequest) ProtoMessage() {}

func (x *QueryProductOrderProcessStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductOrderProcessStepRequest.ProtoReflect.Descriptor instead.
func (*QueryProductOrderProcessStepRequest) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductOrderProcessStepRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductOrderProcessStepRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductOrderProcessStepRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductOrderProcessStepRequest) GetProductOrderProcessID() string {
	if x != nil {
		return x.ProductOrderProcessID
	}
	return ""
}

type QueryProductOrderProcessStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                           `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                         `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductOrderProcessStepInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                          `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                          `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                          `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductOrderProcessStepResponse) Reset() {
	*x = QueryProductOrderProcessStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductOrderProcessStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductOrderProcessStepResponse) ProtoMessage() {}

func (x *QueryProductOrderProcessStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductOrderProcessStepResponse.ProtoReflect.Descriptor instead.
func (*QueryProductOrderProcessStepResponse) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{4}
}

func (x *QueryProductOrderProcessStepResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductOrderProcessStepResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductOrderProcessStepResponse) GetData() []*ProductOrderProcessStepInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductOrderProcessStepResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductOrderProcessStepResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductOrderProcessStepResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductOrderProcessStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                           `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                         `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductOrderProcessStepInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductOrderProcessStepResponse) Reset() {
	*x = GetAllProductOrderProcessStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductOrderProcessStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductOrderProcessStepResponse) ProtoMessage() {}

func (x *GetAllProductOrderProcessStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductOrderProcessStepResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductOrderProcessStepResponse) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllProductOrderProcessStepResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductOrderProcessStepResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductOrderProcessStepResponse) GetData() []*ProductOrderProcessStepInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductOrderProcessStepDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                         `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                       `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductOrderProcessStepInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductOrderProcessStepDetailResponse) Reset() {
	*x = GetProductOrderProcessStepDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_process_step_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductOrderProcessStepDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductOrderProcessStepDetailResponse) ProtoMessage() {}

func (x *GetProductOrderProcessStepDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_process_step_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductOrderProcessStepDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductOrderProcessStepDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_order_process_step_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductOrderProcessStepDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductOrderProcessStepDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductOrderProcessStepDetailResponse) GetData() *ProductOrderProcessStepInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_order_process_step_proto protoreflect.FileDescriptor

var file_product_order_process_step_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x05, 0x0a, 0x1b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x6f, 0x72,
	0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x7c, 0x0a,
	0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x25,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x25, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd3, 0x02,
	0x0a, 0x25, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65,
	0x70, 0x49, 0x44, 0x22, 0xe6, 0x01, 0x0a, 0x28, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x3c,
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x1a,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb5, 0x01, 0x0a,
	0x23, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34,
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x44, 0x22, 0xdf, 0x01, 0x0a, 0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x65, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_order_process_step_proto_rawDescOnce sync.Once
	file_product_order_process_step_proto_rawDescData = file_product_order_process_step_proto_rawDesc
)

func file_product_order_process_step_proto_rawDescGZIP() []byte {
	file_product_order_process_step_proto_rawDescOnce.Do(func() {
		file_product_order_process_step_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_order_process_step_proto_rawDescData)
	})
	return file_product_order_process_step_proto_rawDescData
}

var file_product_order_process_step_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_order_process_step_proto_goTypes = []interface{}{
	(*ProductOrderProcessStepInfo)(nil),              // 0: proto.ProductOrderProcessStepInfo
	(*ProductOrderProcessStepAttachmentInfo)(nil),    // 1: proto.ProductOrderProcessStepAttachmentInfo
	(*ProductOrderProcessStepTypeParameterInfo)(nil), // 2: proto.ProductOrderProcessStepTypeParameterInfo
	(*QueryProductOrderProcessStepRequest)(nil),      // 3: proto.QueryProductOrderProcessStepRequest
	(*QueryProductOrderProcessStepResponse)(nil),     // 4: proto.QueryProductOrderProcessStepResponse
	(*GetAllProductOrderProcessStepResponse)(nil),    // 5: proto.GetAllProductOrderProcessStepResponse
	(*GetProductOrderProcessStepDetailResponse)(nil), // 6: proto.GetProductOrderProcessStepDetailResponse
	(Code)(0), // 7: proto.Code
}
var file_product_order_process_step_proto_depIdxs = []int32{
	1, // 0: proto.ProductOrderProcessStepInfo.productOrderProcessStepAttachments:type_name -> proto.ProductOrderProcessStepAttachmentInfo
	2, // 1: proto.ProductOrderProcessStepInfo.productOrderProcessStepTypeParameters:type_name -> proto.ProductOrderProcessStepTypeParameterInfo
	7, // 2: proto.QueryProductOrderProcessStepResponse.code:type_name -> proto.Code
	0, // 3: proto.QueryProductOrderProcessStepResponse.data:type_name -> proto.ProductOrderProcessStepInfo
	7, // 4: proto.GetAllProductOrderProcessStepResponse.code:type_name -> proto.Code
	0, // 5: proto.GetAllProductOrderProcessStepResponse.data:type_name -> proto.ProductOrderProcessStepInfo
	7, // 6: proto.GetProductOrderProcessStepDetailResponse.code:type_name -> proto.Code
	0, // 7: proto.GetProductOrderProcessStepDetailResponse.data:type_name -> proto.ProductOrderProcessStepInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_product_order_process_step_proto_init() }
func file_product_order_process_step_proto_init() {
	if File_product_order_process_step_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_order_process_step_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOrderProcessStepInfo); i {
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
		file_product_order_process_step_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOrderProcessStepAttachmentInfo); i {
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
		file_product_order_process_step_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOrderProcessStepTypeParameterInfo); i {
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
		file_product_order_process_step_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductOrderProcessStepRequest); i {
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
		file_product_order_process_step_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductOrderProcessStepResponse); i {
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
		file_product_order_process_step_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductOrderProcessStepResponse); i {
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
		file_product_order_process_step_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductOrderProcessStepDetailResponse); i {
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
			RawDescriptor: file_product_order_process_step_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_order_process_step_proto_goTypes,
		DependencyIndexes: file_product_order_process_step_proto_depIdxs,
		MessageInfos:      file_product_order_process_step_proto_msgTypes,
	}.Build()
	File_product_order_process_step_proto = out.File
	file_product_order_process_step_proto_rawDesc = nil
	file_product_order_process_step_proto_goTypes = nil
	file_product_order_process_step_proto_depIdxs = nil
}