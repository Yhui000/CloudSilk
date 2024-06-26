// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.21.12
// source: production_process_sop.proto

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

// ProductionProcessSopClient is the client API for ProductionProcessSop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductionProcessSopClient interface {
	Get(ctx context.Context, in *GetProductionProcessSopRequest, opts ...grpc_go.CallOption) (*GetProductionProcessSopDetailResponse, common.ErrorWithAttachment)
}

type productionProcessSopClient struct {
	cc *triple.TripleConn
}

type ProductionProcessSopClientImpl struct {
	Get func(ctx context.Context, in *GetProductionProcessSopRequest) (*GetProductionProcessSopDetailResponse, error)
}

func (c *ProductionProcessSopClientImpl) GetDubboStub(cc *triple.TripleConn) ProductionProcessSopClient {
	return NewProductionProcessSopClient(cc)
}

func (c *ProductionProcessSopClientImpl) XXX_InterfaceName() string {
	return "proto.ProductionProcessSop"
}

func NewProductionProcessSopClient(cc *triple.TripleConn) ProductionProcessSopClient {
	return &productionProcessSopClient{cc}
}

func (c *productionProcessSopClient) Get(ctx context.Context, in *GetProductionProcessSopRequest, opts ...grpc_go.CallOption) (*GetProductionProcessSopDetailResponse, common.ErrorWithAttachment) {
	out := new(GetProductionProcessSopDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Get", in, out)
}

// ProductionProcessSopServer is the server API for ProductionProcessSop service.
// All implementations must embed UnimplementedProductionProcessSopServer
// for forward compatibility
type ProductionProcessSopServer interface {
	Get(context.Context, *GetProductionProcessSopRequest) (*GetProductionProcessSopDetailResponse, error)
	mustEmbedUnimplementedProductionProcessSopServer()
}

// UnimplementedProductionProcessSopServer must be embedded to have forward compatible implementations.
type UnimplementedProductionProcessSopServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedProductionProcessSopServer) Get(context.Context, *GetProductionProcessSopRequest) (*GetProductionProcessSopDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *UnimplementedProductionProcessSopServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedProductionProcessSopServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedProductionProcessSopServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &ProductionProcessSop_ServiceDesc
}
func (s *UnimplementedProductionProcessSopServer) XXX_InterfaceName() string {
	return "proto.ProductionProcessSop"
}

func (UnimplementedProductionProcessSopServer) mustEmbedUnimplementedProductionProcessSopServer() {}

// UnsafeProductionProcessSopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductionProcessSopServer will
// result in compilation errors.
type UnsafeProductionProcessSopServer interface {
	mustEmbedUnimplementedProductionProcessSopServer()
}

func RegisterProductionProcessSopServer(s grpc_go.ServiceRegistrar, srv ProductionProcessSopServer) {
	s.RegisterService(&ProductionProcessSop_ServiceDesc, srv)
}

func _ProductionProcessSop_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductionProcessSopRequest)
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

// ProductionProcessSop_ServiceDesc is the grpc_go.ServiceDesc for ProductionProcessSop service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductionProcessSop_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.ProductionProcessSop",
	HandlerType: (*ProductionProcessSopServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProductionProcessSop_Get_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "production_process_sop.proto",
}
