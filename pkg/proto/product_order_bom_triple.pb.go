// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: product_order_bom.proto

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

// ProductOrderBomClient is the client API for ProductOrderBom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductOrderBomClient interface {
	Query(ctx context.Context, in *QueryProductOrderBomRequest, opts ...grpc_go.CallOption) (*QueryProductOrderBomResponse, common.ErrorWithAttachment)
}

type productOrderBomClient struct {
	cc *triple.TripleConn
}

type ProductOrderBomClientImpl struct {
	Query func(ctx context.Context, in *QueryProductOrderBomRequest) (*QueryProductOrderBomResponse, error)
}

func (c *ProductOrderBomClientImpl) GetDubboStub(cc *triple.TripleConn) ProductOrderBomClient {
	return NewProductOrderBomClient(cc)
}

func (c *ProductOrderBomClientImpl) XXX_InterfaceName() string {
	return "proto.ProductOrderBom"
}

func NewProductOrderBomClient(cc *triple.TripleConn) ProductOrderBomClient {
	return &productOrderBomClient{cc}
}

func (c *productOrderBomClient) Query(ctx context.Context, in *QueryProductOrderBomRequest, opts ...grpc_go.CallOption) (*QueryProductOrderBomResponse, common.ErrorWithAttachment) {
	out := new(QueryProductOrderBomResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

// ProductOrderBomServer is the server API for ProductOrderBom service.
// All implementations must embed UnimplementedProductOrderBomServer
// for forward compatibility
type ProductOrderBomServer interface {
	Query(context.Context, *QueryProductOrderBomRequest) (*QueryProductOrderBomResponse, error)
	mustEmbedUnimplementedProductOrderBomServer()
}

// UnimplementedProductOrderBomServer must be embedded to have forward compatible implementations.
type UnimplementedProductOrderBomServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedProductOrderBomServer) Query(context.Context, *QueryProductOrderBomRequest) (*QueryProductOrderBomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (s *UnimplementedProductOrderBomServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedProductOrderBomServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedProductOrderBomServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &ProductOrderBom_ServiceDesc
}
func (s *UnimplementedProductOrderBomServer) XXX_InterfaceName() string {
	return "proto.ProductOrderBom"
}

func (UnimplementedProductOrderBomServer) mustEmbedUnimplementedProductOrderBomServer() {}

// UnsafeProductOrderBomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductOrderBomServer will
// result in compilation errors.
type UnsafeProductOrderBomServer interface {
	mustEmbedUnimplementedProductOrderBomServer()
}

func RegisterProductOrderBomServer(s grpc_go.ServiceRegistrar, srv ProductOrderBomServer) {
	s.RegisterService(&ProductOrderBom_ServiceDesc, srv)
}

func _ProductOrderBom_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProductOrderBomRequest)
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

// ProductOrderBom_ServiceDesc is the grpc_go.ServiceDesc for ProductOrderBom service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductOrderBom_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.ProductOrderBom",
	HandlerType: (*ProductOrderBomServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _ProductOrderBom_Query_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "product_order_bom.proto",
}