// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.21.12
// source: production_station_signup.proto

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

// ProductionStationSignupClient is the client API for ProductionStationSignup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductionStationSignupClient interface {
	Add(ctx context.Context, in *ProductionStationSignupInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Update(ctx context.Context, in *ProductionStationSignupInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Get(ctx context.Context, in *GetProductionStationSignupRequest, opts ...grpc_go.CallOption) (*GetProductionStationSignupDetailResponse, common.ErrorWithAttachment)
}

type productionStationSignupClient struct {
	cc *triple.TripleConn
}

type ProductionStationSignupClientImpl struct {
	Add    func(ctx context.Context, in *ProductionStationSignupInfo) (*CommonResponse, error)
	Update func(ctx context.Context, in *ProductionStationSignupInfo) (*CommonResponse, error)
	Get    func(ctx context.Context, in *GetProductionStationSignupRequest) (*GetProductionStationSignupDetailResponse, error)
}

func (c *ProductionStationSignupClientImpl) GetDubboStub(cc *triple.TripleConn) ProductionStationSignupClient {
	return NewProductionStationSignupClient(cc)
}

func (c *ProductionStationSignupClientImpl) XXX_InterfaceName() string {
	return "proto.ProductionStationSignup"
}

func NewProductionStationSignupClient(cc *triple.TripleConn) ProductionStationSignupClient {
	return &productionStationSignupClient{cc}
}

func (c *productionStationSignupClient) Add(ctx context.Context, in *ProductionStationSignupInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Add", in, out)
}

func (c *productionStationSignupClient) Update(ctx context.Context, in *ProductionStationSignupInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Update", in, out)
}

func (c *productionStationSignupClient) Get(ctx context.Context, in *GetProductionStationSignupRequest, opts ...grpc_go.CallOption) (*GetProductionStationSignupDetailResponse, common.ErrorWithAttachment) {
	out := new(GetProductionStationSignupDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Get", in, out)
}

// ProductionStationSignupServer is the server API for ProductionStationSignup service.
// All implementations must embed UnimplementedProductionStationSignupServer
// for forward compatibility
type ProductionStationSignupServer interface {
	Add(context.Context, *ProductionStationSignupInfo) (*CommonResponse, error)
	Update(context.Context, *ProductionStationSignupInfo) (*CommonResponse, error)
	Get(context.Context, *GetProductionStationSignupRequest) (*GetProductionStationSignupDetailResponse, error)
	mustEmbedUnimplementedProductionStationSignupServer()
}

// UnimplementedProductionStationSignupServer must be embedded to have forward compatible implementations.
type UnimplementedProductionStationSignupServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedProductionStationSignupServer) Add(context.Context, *ProductionStationSignupInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedProductionStationSignupServer) Update(context.Context, *ProductionStationSignupInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductionStationSignupServer) Get(context.Context, *GetProductionStationSignupRequest) (*GetProductionStationSignupDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *UnimplementedProductionStationSignupServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedProductionStationSignupServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedProductionStationSignupServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &ProductionStationSignup_ServiceDesc
}
func (s *UnimplementedProductionStationSignupServer) XXX_InterfaceName() string {
	return "proto.ProductionStationSignup"
}

func (UnimplementedProductionStationSignupServer) mustEmbedUnimplementedProductionStationSignupServer() {
}

// UnsafeProductionStationSignupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductionStationSignupServer will
// result in compilation errors.
type UnsafeProductionStationSignupServer interface {
	mustEmbedUnimplementedProductionStationSignupServer()
}

func RegisterProductionStationSignupServer(s grpc_go.ServiceRegistrar, srv ProductionStationSignupServer) {
	s.RegisterService(&ProductionStationSignup_ServiceDesc, srv)
}

func _ProductionStationSignup_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductionStationSignupInfo)
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
	invo := invocation.NewRPCInvocation("Add", args, invAttachment)
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

func _ProductionStationSignup_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductionStationSignupInfo)
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
	invo := invocation.NewRPCInvocation("Update", args, invAttachment)
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

func _ProductionStationSignup_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductionStationSignupRequest)
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

// ProductionStationSignup_ServiceDesc is the grpc_go.ServiceDesc for ProductionStationSignup service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductionStationSignup_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.ProductionStationSignup",
	HandlerType: (*ProductionStationSignupServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ProductionStationSignup_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductionStationSignup_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProductionStationSignup_Get_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "production_station_signup.proto",
}
