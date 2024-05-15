// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.21.12
// source: personnel_qualification.proto

package proto

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// PersonnelQualificationClient is the client API for PersonnelQualification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonnelQualificationClient interface {
	Query(ctx context.Context, in *QueryPersonnelQualificationRequest, opts ...grpc_go.CallOption) (*QueryPersonnelQualificationResponse, common.ErrorWithAttachment)
	Get(ctx context.Context, in *GetPersonnelQualificationRequest, opts ...grpc_go.CallOption) (*GetPersonnelQualificationDetailResponse, common.ErrorWithAttachment)
}

type personnelQualificationClient struct {
	cc *triple.TripleConn
}

type PersonnelQualificationClientImpl struct {
	Query func(ctx context.Context, in *QueryPersonnelQualificationRequest) (*QueryPersonnelQualificationResponse, error)
	Get   func(ctx context.Context, in *GetPersonnelQualificationRequest) (*GetPersonnelQualificationDetailResponse, error)
}

func (c *PersonnelQualificationClientImpl) GetDubboStub(cc *triple.TripleConn) PersonnelQualificationClient {
	return NewPersonnelQualificationClient(cc)
}

func (c *PersonnelQualificationClientImpl) XXX_InterfaceName() string {
	return "proto.PersonnelQualification"
}

func NewPersonnelQualificationClient(cc *triple.TripleConn) PersonnelQualificationClient {
	return &personnelQualificationClient{cc}
}

func (c *personnelQualificationClient) Query(ctx context.Context, in *QueryPersonnelQualificationRequest, opts ...grpc_go.CallOption) (*QueryPersonnelQualificationResponse, common.ErrorWithAttachment) {
	out := new(QueryPersonnelQualificationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

func (c *personnelQualificationClient) Get(ctx context.Context, in *GetPersonnelQualificationRequest, opts ...grpc_go.CallOption) (*GetPersonnelQualificationDetailResponse, common.ErrorWithAttachment) {
	out := new(GetPersonnelQualificationDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Get", in, out)
}

// PersonnelQualificationServer is the server API for PersonnelQualification service.
// All implementations must embed UnimplementedPersonnelQualificationServer
// for forward compatibility
type PersonnelQualificationServer interface {
	Query(context.Context, *QueryPersonnelQualificationRequest) (*QueryPersonnelQualificationResponse, error)
	Get(context.Context, *GetPersonnelQualificationRequest) (*GetPersonnelQualificationDetailResponse, error)
	mustEmbedUnimplementedPersonnelQualificationServer()
}

// UnimplementedPersonnelQualificationServer must be embedded to have forward compatible implementations.
type UnimplementedPersonnelQualificationServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedPersonnelQualificationServer) Query(context.Context, *QueryPersonnelQualificationRequest) (*QueryPersonnelQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedPersonnelQualificationServer) Get(context.Context, *GetPersonnelQualificationRequest) (*GetPersonnelQualificationDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *UnimplementedPersonnelQualificationServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedPersonnelQualificationServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedPersonnelQualificationServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &PersonnelQualification_ServiceDesc
}
func (s *UnimplementedPersonnelQualificationServer) XXX_InterfaceName() string {
	return "proto.PersonnelQualification"
}

func (UnimplementedPersonnelQualificationServer) mustEmbedUnimplementedPersonnelQualificationServer() {
}

// UnsafePersonnelQualificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonnelQualificationServer will
// result in compilation errors.
type UnsafePersonnelQualificationServer interface {
	mustEmbedUnimplementedPersonnelQualificationServer()
}

func RegisterPersonnelQualificationServer(s grpc_go.ServiceRegistrar, srv PersonnelQualificationServer) {
	s.RegisterService(&PersonnelQualification_ServiceDesc, srv)
}

func _PersonnelQualification_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPersonnelQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Query", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonnelQualification_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonnelQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Get", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// PersonnelQualification_ServiceDesc is the grpc_go.ServiceDesc for PersonnelQualification service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonnelQualification_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.PersonnelQualification",
	HandlerType: (*PersonnelQualificationServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _PersonnelQualification_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PersonnelQualification_Get_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "personnel_qualification.proto",
}