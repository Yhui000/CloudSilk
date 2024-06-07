// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: product_attribute.proto

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

// ProductAttributeClient is the client API for ProductAttribute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductAttributeClient interface {
	Query(ctx context.Context, in *QueryProductAttributeRequest, opts ...grpc_go.CallOption) (*QueryProductAttributeResponse, common.ErrorWithAttachment)
}

type productAttributeClient struct {
	cc *triple.TripleConn
}

type ProductAttributeClientImpl struct {
	Query func(ctx context.Context, in *QueryProductAttributeRequest) (*QueryProductAttributeResponse, error)
}

func (c *ProductAttributeClientImpl) GetDubboStub(cc *triple.TripleConn) ProductAttributeClient {
	return NewProductAttributeClient(cc)
}

func (c *ProductAttributeClientImpl) XXX_InterfaceName() string {
	return "proto.ProductAttribute"
}

func NewProductAttributeClient(cc *triple.TripleConn) ProductAttributeClient {
	return &productAttributeClient{cc}
}

func (c *productAttributeClient) Query(ctx context.Context, in *QueryProductAttributeRequest, opts ...grpc_go.CallOption) (*QueryProductAttributeResponse, common.ErrorWithAttachment) {
	out := new(QueryProductAttributeResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

// ProductAttributeServer is the server API for ProductAttribute service.
// All implementations must embed UnimplementedProductAttributeServer
// for forward compatibility
type ProductAttributeServer interface {
	Query(context.Context, *QueryProductAttributeRequest) (*QueryProductAttributeResponse, error)
	mustEmbedUnimplementedProductAttributeServer()
}

// UnimplementedProductAttributeServer must be embedded to have forward compatible implementations.
type UnimplementedProductAttributeServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedProductAttributeServer) Query(context.Context, *QueryProductAttributeRequest) (*QueryProductAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (s *UnimplementedProductAttributeServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedProductAttributeServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedProductAttributeServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &ProductAttribute_ServiceDesc
}
func (s *UnimplementedProductAttributeServer) XXX_InterfaceName() string {
	return "proto.ProductAttribute"
}

func (UnimplementedProductAttributeServer) mustEmbedUnimplementedProductAttributeServer() {}

// UnsafeProductAttributeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductAttributeServer will
// result in compilation errors.
type UnsafeProductAttributeServer interface {
	mustEmbedUnimplementedProductAttributeServer()
}

func RegisterProductAttributeServer(s grpc_go.ServiceRegistrar, srv ProductAttributeServer) {
	s.RegisterService(&ProductAttribute_ServiceDesc, srv)
}

func _ProductAttribute_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProductAttributeRequest)
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

// ProductAttribute_ServiceDesc is the grpc_go.ServiceDesc for ProductAttribute service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductAttribute_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.ProductAttribute",
	HandlerType: (*ProductAttributeServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _ProductAttribute_Query_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "product_attribute.proto",
}
