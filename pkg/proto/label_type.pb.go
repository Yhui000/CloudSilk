// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.3
// source: label_type.proto

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

type LabelTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 类型
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// 备注
	Remark string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark"`
}

func (x *LabelTypeInfo) Reset() {
	*x = LabelTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelTypeInfo) ProtoMessage() {}

func (x *LabelTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_label_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelTypeInfo.ProtoReflect.Descriptor instead.
func (*LabelTypeInfo) Descriptor() ([]byte, []int) {
	return file_label_type_proto_rawDescGZIP(), []int{0}
}

func (x *LabelTypeInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LabelTypeInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LabelTypeInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LabelTypeInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryLabelTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 类型或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryLabelTypeRequest) Reset() {
	*x = QueryLabelTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLabelTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLabelTypeRequest) ProtoMessage() {}

func (x *QueryLabelTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_label_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLabelTypeRequest.ProtoReflect.Descriptor instead.
func (*QueryLabelTypeRequest) Descriptor() ([]byte, []int) {
	return file_label_type_proto_rawDescGZIP(), []int{1}
}

func (x *QueryLabelTypeRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryLabelTypeRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryLabelTypeRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryLabelTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryLabelTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code             `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*LabelTypeInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64            `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64            `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64            `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryLabelTypeResponse) Reset() {
	*x = QueryLabelTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLabelTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLabelTypeResponse) ProtoMessage() {}

func (x *QueryLabelTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLabelTypeResponse.ProtoReflect.Descriptor instead.
func (*QueryLabelTypeResponse) Descriptor() ([]byte, []int) {
	return file_label_type_proto_rawDescGZIP(), []int{2}
}

func (x *QueryLabelTypeResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryLabelTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryLabelTypeResponse) GetData() []*LabelTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryLabelTypeResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryLabelTypeResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryLabelTypeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllLabelTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code             `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*LabelTypeInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllLabelTypeResponse) Reset() {
	*x = GetAllLabelTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLabelTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLabelTypeResponse) ProtoMessage() {}

func (x *GetAllLabelTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLabelTypeResponse.ProtoReflect.Descriptor instead.
func (*GetAllLabelTypeResponse) Descriptor() ([]byte, []int) {
	return file_label_type_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllLabelTypeResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllLabelTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllLabelTypeResponse) GetData() []*LabelTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetLabelTypeDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code           `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *LabelTypeInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetLabelTypeDetailResponse) Reset() {
	*x = GetLabelTypeDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLabelTypeDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLabelTypeDetailResponse) ProtoMessage() {}

func (x *GetLabelTypeDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLabelTypeDetailResponse.ProtoReflect.Descriptor instead.
func (*GetLabelTypeDetailResponse) Descriptor() ([]byte, []int) {
	return file_label_type_proto_rawDescGZIP(), []int{4}
}

func (x *GetLabelTypeDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetLabelTypeDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetLabelTypeDetailResponse) GetData() *LabelTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_label_type_proto protoreflect.FileDescriptor

var file_label_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0d, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_label_type_proto_rawDescOnce sync.Once
	file_label_type_proto_rawDescData = file_label_type_proto_rawDesc
)

func file_label_type_proto_rawDescGZIP() []byte {
	file_label_type_proto_rawDescOnce.Do(func() {
		file_label_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_label_type_proto_rawDescData)
	})
	return file_label_type_proto_rawDescData
}

var file_label_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_label_type_proto_goTypes = []interface{}{
	(*LabelTypeInfo)(nil),              // 0: proto.LabelTypeInfo
	(*QueryLabelTypeRequest)(nil),      // 1: proto.QueryLabelTypeRequest
	(*QueryLabelTypeResponse)(nil),     // 2: proto.QueryLabelTypeResponse
	(*GetAllLabelTypeResponse)(nil),    // 3: proto.GetAllLabelTypeResponse
	(*GetLabelTypeDetailResponse)(nil), // 4: proto.GetLabelTypeDetailResponse
	(Code)(0),                          // 5: proto.Code
}
var file_label_type_proto_depIdxs = []int32{
	5, // 0: proto.QueryLabelTypeResponse.code:type_name -> proto.Code
	0, // 1: proto.QueryLabelTypeResponse.data:type_name -> proto.LabelTypeInfo
	5, // 2: proto.GetAllLabelTypeResponse.code:type_name -> proto.Code
	0, // 3: proto.GetAllLabelTypeResponse.data:type_name -> proto.LabelTypeInfo
	5, // 4: proto.GetLabelTypeDetailResponse.code:type_name -> proto.Code
	0, // 5: proto.GetLabelTypeDetailResponse.data:type_name -> proto.LabelTypeInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_label_type_proto_init() }
func file_label_type_proto_init() {
	if File_label_type_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_label_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelTypeInfo); i {
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
		file_label_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLabelTypeRequest); i {
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
		file_label_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLabelTypeResponse); i {
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
		file_label_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLabelTypeResponse); i {
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
		file_label_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLabelTypeDetailResponse); i {
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
			RawDescriptor: file_label_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_label_type_proto_goTypes,
		DependencyIndexes: file_label_type_proto_depIdxs,
		MessageInfos:      file_label_type_proto_msgTypes,
	}.Build()
	File_label_type_proto = out.File
	file_label_type_proto_rawDesc = nil
	file_label_type_proto_goTypes = nil
	file_label_type_proto_depIdxs = nil
}
