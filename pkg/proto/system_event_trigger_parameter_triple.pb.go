// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: system_event_trigger_parameter.proto

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

// SystemEventTriggerParameterClient is the client API for SystemEventTriggerParameter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemEventTriggerParameterClient interface {
	Add(ctx context.Context, in *SystemEventTriggerParameterInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
}

type systemEventTriggerParameterClient struct {
	cc *triple.TripleConn
}

type SystemEventTriggerParameterClientImpl struct {
	Add func(ctx context.Context, in *SystemEventTriggerParameterInfo) (*CommonResponse, error)
}

func (c *SystemEventTriggerParameterClientImpl) GetDubboStub(cc *triple.TripleConn) SystemEventTriggerParameterClient {
	return NewSystemEventTriggerParameterClient(cc)
}

func (c *SystemEventTriggerParameterClientImpl) XXX_InterfaceName() string {
	return "proto.SystemEventTriggerParameter"
}

func NewSystemEventTriggerParameterClient(cc *triple.TripleConn) SystemEventTriggerParameterClient {
	return &systemEventTriggerParameterClient{cc}
}

func (c *systemEventTriggerParameterClient) Add(ctx context.Context, in *SystemEventTriggerParameterInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Add", in, out)
}

// SystemEventTriggerParameterServer is the server API for SystemEventTriggerParameter service.
// All implementations must embed UnimplementedSystemEventTriggerParameterServer
// for forward compatibility
type SystemEventTriggerParameterServer interface {
	Add(context.Context, *SystemEventTriggerParameterInfo) (*CommonResponse, error)
	mustEmbedUnimplementedSystemEventTriggerParameterServer()
}

// UnimplementedSystemEventTriggerParameterServer must be embedded to have forward compatible implementations.
type UnimplementedSystemEventTriggerParameterServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedSystemEventTriggerParameterServer) Add(context.Context, *SystemEventTriggerParameterInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (s *UnimplementedSystemEventTriggerParameterServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedSystemEventTriggerParameterServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedSystemEventTriggerParameterServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &SystemEventTriggerParameter_ServiceDesc
}
func (s *UnimplementedSystemEventTriggerParameterServer) XXX_InterfaceName() string {
	return "proto.SystemEventTriggerParameter"
}

func (UnimplementedSystemEventTriggerParameterServer) mustEmbedUnimplementedSystemEventTriggerParameterServer() {
}

// UnsafeSystemEventTriggerParameterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemEventTriggerParameterServer will
// result in compilation errors.
type UnsafeSystemEventTriggerParameterServer interface {
	mustEmbedUnimplementedSystemEventTriggerParameterServer()
}

func RegisterSystemEventTriggerParameterServer(s grpc_go.ServiceRegistrar, srv SystemEventTriggerParameterServer) {
	s.RegisterService(&SystemEventTriggerParameter_ServiceDesc, srv)
}

func _SystemEventTriggerParameter_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemEventTriggerParameterInfo)
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

// SystemEventTriggerParameter_ServiceDesc is the grpc_go.ServiceDesc for SystemEventTriggerParameter service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemEventTriggerParameter_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.SystemEventTriggerParameter",
	HandlerType: (*SystemEventTriggerParameterServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _SystemEventTriggerParameter_Add_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "system_event_trigger_parameter.proto",
}
