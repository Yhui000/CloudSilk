// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: material_channel_layer.proto

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

// MaterialChannelLayerClient is the client API for MaterialChannelLayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaterialChannelLayerClient interface {
	Get(ctx context.Context, in *GetMaterialChannelLayerRequest, opts ...grpc_go.CallOption) (*GetAllMaterialChannelLayerResponse, common.ErrorWithAttachment)
}

type materialChannelLayerClient struct {
	cc *triple.TripleConn
}

type MaterialChannelLayerClientImpl struct {
	Get func(ctx context.Context, in *GetMaterialChannelLayerRequest) (*GetAllMaterialChannelLayerResponse, error)
}

func (c *MaterialChannelLayerClientImpl) GetDubboStub(cc *triple.TripleConn) MaterialChannelLayerClient {
	return NewMaterialChannelLayerClient(cc)
}

func (c *MaterialChannelLayerClientImpl) XXX_InterfaceName() string {
	return "proto.MaterialChannelLayer"
}

func NewMaterialChannelLayerClient(cc *triple.TripleConn) MaterialChannelLayerClient {
	return &materialChannelLayerClient{cc}
}

func (c *materialChannelLayerClient) Get(ctx context.Context, in *GetMaterialChannelLayerRequest, opts ...grpc_go.CallOption) (*GetAllMaterialChannelLayerResponse, common.ErrorWithAttachment) {
	out := new(GetAllMaterialChannelLayerResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Get", in, out)
}

// MaterialChannelLayerServer is the server API for MaterialChannelLayer service.
// All implementations must embed UnimplementedMaterialChannelLayerServer
// for forward compatibility
type MaterialChannelLayerServer interface {
	Get(context.Context, *GetMaterialChannelLayerRequest) (*GetAllMaterialChannelLayerResponse, error)
	mustEmbedUnimplementedMaterialChannelLayerServer()
}

// UnimplementedMaterialChannelLayerServer must be embedded to have forward compatible implementations.
type UnimplementedMaterialChannelLayerServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedMaterialChannelLayerServer) Get(context.Context, *GetMaterialChannelLayerRequest) (*GetAllMaterialChannelLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *UnimplementedMaterialChannelLayerServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedMaterialChannelLayerServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedMaterialChannelLayerServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &MaterialChannelLayer_ServiceDesc
}
func (s *UnimplementedMaterialChannelLayerServer) XXX_InterfaceName() string {
	return "proto.MaterialChannelLayer"
}

func (UnimplementedMaterialChannelLayerServer) mustEmbedUnimplementedMaterialChannelLayerServer() {}

// UnsafeMaterialChannelLayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaterialChannelLayerServer will
// result in compilation errors.
type UnsafeMaterialChannelLayerServer interface {
	mustEmbedUnimplementedMaterialChannelLayerServer()
}

func RegisterMaterialChannelLayerServer(s grpc_go.ServiceRegistrar, srv MaterialChannelLayerServer) {
	s.RegisterService(&MaterialChannelLayer_ServiceDesc, srv)
}

func _MaterialChannelLayer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaterialChannelLayerRequest)
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

// MaterialChannelLayer_ServiceDesc is the grpc_go.ServiceDesc for MaterialChannelLayer service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaterialChannelLayer_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.MaterialChannelLayer",
	HandlerType: (*MaterialChannelLayerServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MaterialChannelLayer_Get_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "material_channel_layer.proto",
}
