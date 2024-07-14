// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_channel_proposing_record.proto

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

type MaterialChannelProposingRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 料道
	MaterialChannelID string               `protobuf:"bytes,2,opt,name=materialChannelID,proto3" json:"materialChannelID"`
	MaterialChannel   *MaterialChannelInfo `protobuf:"bytes,3,opt,name=materialChannel,proto3" json:"materialChannel"`
	// 叫料时间
	RequestTime string `protobuf:"bytes,4,opt,name=requestTime,proto3" json:"requestTime"`
	// 补料时间
	ResponseTime string `protobuf:"bytes,5,opt,name=responseTime,proto3" json:"responseTime"`
	// 候料时长
	Duration int64 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration"`
	// 当前状态
	CurrentState string `protobuf:"bytes,7,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,8,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
}

func (x *MaterialChannelProposingRecordInfo) Reset() {
	*x = MaterialChannelProposingRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_proposing_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialChannelProposingRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialChannelProposingRecordInfo) ProtoMessage() {}

func (x *MaterialChannelProposingRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_proposing_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialChannelProposingRecordInfo.ProtoReflect.Descriptor instead.
func (*MaterialChannelProposingRecordInfo) Descriptor() ([]byte, []int) {
	return file_material_channel_proposing_record_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialChannelProposingRecordInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialChannelProposingRecordInfo) GetMaterialChannelID() string {
	if x != nil {
		return x.MaterialChannelID
	}
	return ""
}

func (x *MaterialChannelProposingRecordInfo) GetMaterialChannel() *MaterialChannelInfo {
	if x != nil {
		return x.MaterialChannel
	}
	return nil
}

func (x *MaterialChannelProposingRecordInfo) GetRequestTime() string {
	if x != nil {
		return x.RequestTime
	}
	return ""
}

func (x *MaterialChannelProposingRecordInfo) GetResponseTime() string {
	if x != nil {
		return x.ResponseTime
	}
	return ""
}

func (x *MaterialChannelProposingRecordInfo) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *MaterialChannelProposingRecordInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *MaterialChannelProposingRecordInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

type QueryMaterialChannelProposingRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 叫料时间开始
	// @inject_tag: uri:"requestTime0" form:"requestTime0"
	RequestTime0 string `protobuf:"bytes,4,opt,name=requestTime0,proto3" json:"requestTime0" uri:"requestTime0" form:"requestTime0"`
	// 叫料时间结束
	// @inject_tag: uri:"requestTime1" form:"requestTime1"
	RequestTime1 string `protobuf:"bytes,5,opt,name=requestTime1,proto3" json:"requestTime1" uri:"requestTime1" form:"requestTime1"`
	// 代号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,6,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,7,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
}

func (x *QueryMaterialChannelProposingRecordRequest) Reset() {
	*x = QueryMaterialChannelProposingRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_proposing_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialChannelProposingRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialChannelProposingRecordRequest) ProtoMessage() {}

func (x *QueryMaterialChannelProposingRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_proposing_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialChannelProposingRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialChannelProposingRecordRequest) Descriptor() ([]byte, []int) {
	return file_material_channel_proposing_record_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialChannelProposingRecordRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialChannelProposingRecordRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialChannelProposingRecordRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialChannelProposingRecordRequest) GetRequestTime0() string {
	if x != nil {
		return x.RequestTime0
	}
	return ""
}

func (x *QueryMaterialChannelProposingRecordRequest) GetRequestTime1() string {
	if x != nil {
		return x.RequestTime1
	}
	return ""
}

func (x *QueryMaterialChannelProposingRecordRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryMaterialChannelProposingRecordRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

type QueryMaterialChannelProposingRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialChannelProposingRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                                 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                                 `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                                 `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialChannelProposingRecordResponse) Reset() {
	*x = QueryMaterialChannelProposingRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_proposing_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialChannelProposingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialChannelProposingRecordResponse) ProtoMessage() {}

func (x *QueryMaterialChannelProposingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_proposing_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialChannelProposingRecordResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialChannelProposingRecordResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_proposing_record_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialChannelProposingRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialChannelProposingRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialChannelProposingRecordResponse) GetData() []*MaterialChannelProposingRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialChannelProposingRecordResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialChannelProposingRecordResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialChannelProposingRecordResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialChannelProposingRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                                  `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                                `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialChannelProposingRecordInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialChannelProposingRecordResponse) Reset() {
	*x = GetAllMaterialChannelProposingRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_proposing_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialChannelProposingRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialChannelProposingRecordResponse) ProtoMessage() {}

func (x *GetAllMaterialChannelProposingRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_proposing_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialChannelProposingRecordResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialChannelProposingRecordResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_proposing_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialChannelProposingRecordResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialChannelProposingRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialChannelProposingRecordResponse) GetData() []*MaterialChannelProposingRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialChannelProposingRecordDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialChannelProposingRecordInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialChannelProposingRecordDetailResponse) Reset() {
	*x = GetMaterialChannelProposingRecordDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_proposing_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialChannelProposingRecordDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialChannelProposingRecordDetailResponse) ProtoMessage() {}

func (x *GetMaterialChannelProposingRecordDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_proposing_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialChannelProposingRecordDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialChannelProposingRecordDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_proposing_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialChannelProposingRecordDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialChannelProposingRecordDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialChannelProposingRecordDetailResponse) GetData() *MaterialChannelProposingRecordInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_channel_proposing_record_proto protoreflect.FileDescriptor

var file_material_channel_proposing_record_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd6, 0x02, 0x0a, 0x22, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8e, 0x02, 0x0a, 0x2a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x30, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x22, 0xed, 0x01, 0x0a, 0x2b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa8, 0x01, 0x0a, 0x2c, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xab, 0x01, 0x0a, 0x2f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_channel_proposing_record_proto_rawDescOnce sync.Once
	file_material_channel_proposing_record_proto_rawDescData = file_material_channel_proposing_record_proto_rawDesc
)

func file_material_channel_proposing_record_proto_rawDescGZIP() []byte {
	file_material_channel_proposing_record_proto_rawDescOnce.Do(func() {
		file_material_channel_proposing_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_channel_proposing_record_proto_rawDescData)
	})
	return file_material_channel_proposing_record_proto_rawDescData
}

var file_material_channel_proposing_record_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_material_channel_proposing_record_proto_goTypes = []interface{}{
	(*MaterialChannelProposingRecordInfo)(nil),              // 0: proto.MaterialChannelProposingRecordInfo
	(*QueryMaterialChannelProposingRecordRequest)(nil),      // 1: proto.QueryMaterialChannelProposingRecordRequest
	(*QueryMaterialChannelProposingRecordResponse)(nil),     // 2: proto.QueryMaterialChannelProposingRecordResponse
	(*GetAllMaterialChannelProposingRecordResponse)(nil),    // 3: proto.GetAllMaterialChannelProposingRecordResponse
	(*GetMaterialChannelProposingRecordDetailResponse)(nil), // 4: proto.GetMaterialChannelProposingRecordDetailResponse
	(*MaterialChannelInfo)(nil),                             // 5: proto.MaterialChannelInfo
	(Code)(0),                                               // 6: proto.Code
}
var file_material_channel_proposing_record_proto_depIdxs = []int32{
	5, // 0: proto.MaterialChannelProposingRecordInfo.materialChannel:type_name -> proto.MaterialChannelInfo
	6, // 1: proto.QueryMaterialChannelProposingRecordResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryMaterialChannelProposingRecordResponse.data:type_name -> proto.MaterialChannelProposingRecordInfo
	6, // 3: proto.GetAllMaterialChannelProposingRecordResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllMaterialChannelProposingRecordResponse.data:type_name -> proto.MaterialChannelProposingRecordInfo
	6, // 5: proto.GetMaterialChannelProposingRecordDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetMaterialChannelProposingRecordDetailResponse.data:type_name -> proto.MaterialChannelProposingRecordInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_material_channel_proposing_record_proto_init() }
func file_material_channel_proposing_record_proto_init() {
	if File_material_channel_proposing_record_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_channel_layer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_channel_proposing_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialChannelProposingRecordInfo); i {
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
		file_material_channel_proposing_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialChannelProposingRecordRequest); i {
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
		file_material_channel_proposing_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialChannelProposingRecordResponse); i {
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
		file_material_channel_proposing_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialChannelProposingRecordResponse); i {
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
		file_material_channel_proposing_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialChannelProposingRecordDetailResponse); i {
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
			RawDescriptor: file_material_channel_proposing_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_channel_proposing_record_proto_goTypes,
		DependencyIndexes: file_material_channel_proposing_record_proto_depIdxs,
		MessageInfos:      file_material_channel_proposing_record_proto_msgTypes,
	}.Build()
	File_material_channel_proposing_record_proto = out.File
	file_material_channel_proposing_record_proto_rawDesc = nil
	file_material_channel_proposing_record_proto_goTypes = nil
	file_material_channel_proposing_record_proto_depIdxs = nil
}
