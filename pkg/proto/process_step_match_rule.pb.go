// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: process_step_match_rule.proto

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

type ProcessStepMatchRuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 生产产线ID
	ProductionLineID string `protobuf:"bytes,2,opt,name=productionLineID,proto3" json:"productionLineID"`
	// 优先级
	Priority int32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority"`
	// 采集组
	GroupCode string `protobuf:"bytes,4,opt,name=groupCode,proto3" json:"groupCode"`
	// 是否启用
	Enable bool `protobuf:"varint,5,opt,name=enable,proto3" json:"enable"`
	// 默认匹配
	InitialValue bool `protobuf:"varint,6,opt,name=initialValue,proto3" json:"initialValue"`
	// 描述
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description"`
	// 备注
	Remark                  string                     `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	AttributeExpressions    []*AttributeExpressionInfo `protobuf:"bytes,9,rep,name=AttributeExpressions,proto3" json:"AttributeExpressions"`
	ProductionProcessStepID string                     `protobuf:"bytes,10,opt,name=productionProcessStepID,proto3" json:"productionProcessStepID"`
	ProductionProcessStep   *ProductionProcessStepInfo `protobuf:"bytes,11,opt,name=productionProcessStep,proto3" json:"productionProcessStep"`
}

func (x *ProcessStepMatchRuleInfo) Reset() {
	*x = ProcessStepMatchRuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_step_match_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStepMatchRuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStepMatchRuleInfo) ProtoMessage() {}

func (x *ProcessStepMatchRuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_process_step_match_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStepMatchRuleInfo.ProtoReflect.Descriptor instead.
func (*ProcessStepMatchRuleInfo) Descriptor() ([]byte, []int) {
	return file_process_step_match_rule_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessStepMatchRuleInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ProcessStepMatchRuleInfo) GetGroupCode() string {
	if x != nil {
		return x.GroupCode
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ProcessStepMatchRuleInfo) GetInitialValue() bool {
	if x != nil {
		return x.InitialValue
	}
	return false
}

func (x *ProcessStepMatchRuleInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetAttributeExpressions() []*AttributeExpressionInfo {
	if x != nil {
		return x.AttributeExpressions
	}
	return nil
}

func (x *ProcessStepMatchRuleInfo) GetProductionProcessStepID() string {
	if x != nil {
		return x.ProductionProcessStepID
	}
	return ""
}

func (x *ProcessStepMatchRuleInfo) GetProductionProcessStep() *ProductionProcessStepInfo {
	if x != nil {
		return x.ProductionProcessStep
	}
	return nil
}

type QueryProcessStepMatchRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,4,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// @inject_tag: uri:"enable" form:"enable"
	Enable bool `protobuf:"varint,5,opt,name=enable,proto3" json:"enable" uri:"enable" form:"enable"`
	// @inject_tag: uri:"hasProductionProcessStepID" form:"hasProductionProcessStepID"
	HasProductionProcessStepID bool `protobuf:"varint,6,opt,name=hasProductionProcessStepID,proto3" json:"hasProductionProcessStepID" uri:"hasProductionProcessStepID" form:"hasProductionProcessStepID"`
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,7,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProcessStepMatchRuleRequest) Reset() {
	*x = QueryProcessStepMatchRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_step_match_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProcessStepMatchRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProcessStepMatchRuleRequest) ProtoMessage() {}

func (x *QueryProcessStepMatchRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_step_match_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProcessStepMatchRuleRequest.ProtoReflect.Descriptor instead.
func (*QueryProcessStepMatchRuleRequest) Descriptor() ([]byte, []int) {
	return file_process_step_match_rule_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProcessStepMatchRuleRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProcessStepMatchRuleRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProcessStepMatchRuleRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProcessStepMatchRuleRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryProcessStepMatchRuleRequest) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *QueryProcessStepMatchRuleRequest) GetHasProductionProcessStepID() bool {
	if x != nil {
		return x.HasProductionProcessStepID
	}
	return false
}

func (x *QueryProcessStepMatchRuleRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProcessStepMatchRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProcessStepMatchRuleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                       `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                       `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProcessStepMatchRuleResponse) Reset() {
	*x = QueryProcessStepMatchRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_step_match_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProcessStepMatchRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProcessStepMatchRuleResponse) ProtoMessage() {}

func (x *QueryProcessStepMatchRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_step_match_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProcessStepMatchRuleResponse.ProtoReflect.Descriptor instead.
func (*QueryProcessStepMatchRuleResponse) Descriptor() ([]byte, []int) {
	return file_process_step_match_rule_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProcessStepMatchRuleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProcessStepMatchRuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProcessStepMatchRuleResponse) GetData() []*ProcessStepMatchRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProcessStepMatchRuleResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProcessStepMatchRuleResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProcessStepMatchRuleResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProcessStepMatchRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProcessStepMatchRuleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProcessStepMatchRuleResponse) Reset() {
	*x = GetAllProcessStepMatchRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_step_match_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProcessStepMatchRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProcessStepMatchRuleResponse) ProtoMessage() {}

func (x *GetAllProcessStepMatchRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_step_match_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProcessStepMatchRuleResponse.ProtoReflect.Descriptor instead.
func (*GetAllProcessStepMatchRuleResponse) Descriptor() ([]byte, []int) {
	return file_process_step_match_rule_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProcessStepMatchRuleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProcessStepMatchRuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProcessStepMatchRuleResponse) GetData() []*ProcessStepMatchRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProcessStepMatchRuleDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProcessStepMatchRuleInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProcessStepMatchRuleDetailResponse) Reset() {
	*x = GetProcessStepMatchRuleDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_step_match_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProcessStepMatchRuleDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessStepMatchRuleDetailResponse) ProtoMessage() {}

func (x *GetProcessStepMatchRuleDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_step_match_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessStepMatchRuleDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProcessStepMatchRuleDetailResponse) Descriptor() ([]byte, []int) {
	return file_process_step_match_rule_proto_rawDescGZIP(), []int{4}
}

func (x *GetProcessStepMatchRuleDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProcessStepMatchRuleDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProcessStepMatchRuleDetailResponse) GetData() *ProcessStepMatchRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_process_step_match_rule_proto protoreflect.FileDescriptor

var file_process_step_match_rule_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xec, 0x03, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x52, 0x0a, 0x14,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x38, 0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x44, 0x12, 0x56, 0x0a, 0x15, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x65, 0x70, 0x22, 0x94, 0x02, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x1a, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65,
	0x70, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x68, 0x61, 0x73, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x21, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x94, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a,
	0x25, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x74, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x5c,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_step_match_rule_proto_rawDescOnce sync.Once
	file_process_step_match_rule_proto_rawDescData = file_process_step_match_rule_proto_rawDesc
)

func file_process_step_match_rule_proto_rawDescGZIP() []byte {
	file_process_step_match_rule_proto_rawDescOnce.Do(func() {
		file_process_step_match_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_step_match_rule_proto_rawDescData)
	})
	return file_process_step_match_rule_proto_rawDescData
}

var file_process_step_match_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_process_step_match_rule_proto_goTypes = []interface{}{
	(*ProcessStepMatchRuleInfo)(nil),              // 0: proto.ProcessStepMatchRuleInfo
	(*QueryProcessStepMatchRuleRequest)(nil),      // 1: proto.QueryProcessStepMatchRuleRequest
	(*QueryProcessStepMatchRuleResponse)(nil),     // 2: proto.QueryProcessStepMatchRuleResponse
	(*GetAllProcessStepMatchRuleResponse)(nil),    // 3: proto.GetAllProcessStepMatchRuleResponse
	(*GetProcessStepMatchRuleDetailResponse)(nil), // 4: proto.GetProcessStepMatchRuleDetailResponse
	(*AttributeExpressionInfo)(nil),               // 5: proto.AttributeExpressionInfo
	(*ProductionProcessStepInfo)(nil),             // 6: proto.ProductionProcessStepInfo
	(Code)(0),                                     // 7: proto.Code
}
var file_process_step_match_rule_proto_depIdxs = []int32{
	5, // 0: proto.ProcessStepMatchRuleInfo.AttributeExpressions:type_name -> proto.AttributeExpressionInfo
	6, // 1: proto.ProcessStepMatchRuleInfo.productionProcessStep:type_name -> proto.ProductionProcessStepInfo
	7, // 2: proto.QueryProcessStepMatchRuleResponse.code:type_name -> proto.Code
	0, // 3: proto.QueryProcessStepMatchRuleResponse.data:type_name -> proto.ProcessStepMatchRuleInfo
	7, // 4: proto.GetAllProcessStepMatchRuleResponse.code:type_name -> proto.Code
	0, // 5: proto.GetAllProcessStepMatchRuleResponse.data:type_name -> proto.ProcessStepMatchRuleInfo
	7, // 6: proto.GetProcessStepMatchRuleDetailResponse.code:type_name -> proto.Code
	0, // 7: proto.GetProcessStepMatchRuleDetailResponse.data:type_name -> proto.ProcessStepMatchRuleInfo
	1, // 8: proto.ProcessStepMatchRule.Query:input_type -> proto.QueryProcessStepMatchRuleRequest
	2, // 9: proto.ProcessStepMatchRule.Query:output_type -> proto.QueryProcessStepMatchRuleResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_process_step_match_rule_proto_init() }
func file_process_step_match_rule_proto_init() {
	if File_process_step_match_rule_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_attribute_expression_proto_init()
	file_production_process_step_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_process_step_match_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStepMatchRuleInfo); i {
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
		file_process_step_match_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProcessStepMatchRuleRequest); i {
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
		file_process_step_match_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProcessStepMatchRuleResponse); i {
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
		file_process_step_match_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProcessStepMatchRuleResponse); i {
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
		file_process_step_match_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProcessStepMatchRuleDetailResponse); i {
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
			RawDescriptor: file_process_step_match_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_process_step_match_rule_proto_goTypes,
		DependencyIndexes: file_process_step_match_rule_proto_depIdxs,
		MessageInfos:      file_process_step_match_rule_proto_msgTypes,
	}.Build()
	File_process_step_match_rule_proto = out.File
	file_process_step_match_rule_proto_rawDesc = nil
	file_process_step_match_rule_proto_goTypes = nil
	file_process_step_match_rule_proto_depIdxs = nil
}
