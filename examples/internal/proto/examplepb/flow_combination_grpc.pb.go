// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/flow_combination.proto

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FlowCombination_RpcEmptyRpc_FullMethodName               = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcEmptyRpc"
	FlowCombination_RpcEmptyStream_FullMethodName            = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcEmptyStream"
	FlowCombination_StreamEmptyRpc_FullMethodName            = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/StreamEmptyRpc"
	FlowCombination_StreamEmptyStream_FullMethodName         = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/StreamEmptyStream"
	FlowCombination_RpcBodyRpc_FullMethodName                = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcBodyRpc"
	FlowCombination_RpcPathSingleNestedRpc_FullMethodName    = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcPathSingleNestedRpc"
	FlowCombination_RpcPathNestedRpc_FullMethodName          = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcPathNestedRpc"
	FlowCombination_RpcBodyStream_FullMethodName             = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcBodyStream"
	FlowCombination_RpcPathSingleNestedStream_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcPathSingleNestedStream"
	FlowCombination_RpcPathNestedStream_FullMethodName       = "/grpc.gateway.examples.internal.proto.examplepb.FlowCombination/RpcPathNestedStream"
)

// FlowCombinationClient is the client API for FlowCombination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowCombinationClient interface {
	RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error)
	StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EmptyProto, EmptyProto], error)
	StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EmptyProto, EmptyProto], error)
	RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error)
	RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error)
	RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error)
}

type flowCombinationClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowCombinationClient(cc grpc.ClientConnInterface) FlowCombinationClient {
	return &flowCombinationClient{cc}
}

func (c *flowCombinationClient) RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyProto)
	err := c.cc.Invoke(ctx, FlowCombination_RpcEmptyRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[0], FlowCombination_RpcEmptyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EmptyProto, EmptyProto]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcEmptyStreamClient = grpc.ServerStreamingClient[EmptyProto]

func (c *flowCombinationClient) StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EmptyProto, EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[1], FlowCombination_StreamEmptyRpc_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EmptyProto, EmptyProto]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_StreamEmptyRpcClient = grpc.ClientStreamingClient[EmptyProto, EmptyProto]

func (c *flowCombinationClient) StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EmptyProto, EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[2], FlowCombination_StreamEmptyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EmptyProto, EmptyProto]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_StreamEmptyStreamClient = grpc.BidiStreamingClient[EmptyProto, EmptyProto]

func (c *flowCombinationClient) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyProto)
	err := c.cc.Invoke(ctx, FlowCombination_RpcBodyRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyProto)
	err := c.cc.Invoke(ctx, FlowCombination_RpcPathSingleNestedRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyProto)
	err := c.cc.Invoke(ctx, FlowCombination_RpcPathNestedRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[3], FlowCombination_RpcBodyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NonEmptyProto, EmptyProto]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcBodyStreamClient = grpc.ServerStreamingClient[EmptyProto]

func (c *flowCombinationClient) RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[4], FlowCombination_RpcPathSingleNestedStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SingleNestedProto, EmptyProto]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcPathSingleNestedStreamClient = grpc.ServerStreamingClient[EmptyProto]

func (c *flowCombinationClient) RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EmptyProto], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlowCombination_ServiceDesc.Streams[5], FlowCombination_RpcPathNestedStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NestedProto, EmptyProto]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcPathNestedStreamClient = grpc.ServerStreamingClient[EmptyProto]

// FlowCombinationServer is the server API for FlowCombination service.
// All implementations should embed UnimplementedFlowCombinationServer
// for forward compatibility.
type FlowCombinationServer interface {
	RpcEmptyRpc(context.Context, *EmptyProto) (*EmptyProto, error)
	RpcEmptyStream(*EmptyProto, grpc.ServerStreamingServer[EmptyProto]) error
	StreamEmptyRpc(grpc.ClientStreamingServer[EmptyProto, EmptyProto]) error
	StreamEmptyStream(grpc.BidiStreamingServer[EmptyProto, EmptyProto]) error
	RpcBodyRpc(context.Context, *NonEmptyProto) (*EmptyProto, error)
	RpcPathSingleNestedRpc(context.Context, *SingleNestedProto) (*EmptyProto, error)
	RpcPathNestedRpc(context.Context, *NestedProto) (*EmptyProto, error)
	RpcBodyStream(*NonEmptyProto, grpc.ServerStreamingServer[EmptyProto]) error
	RpcPathSingleNestedStream(*SingleNestedProto, grpc.ServerStreamingServer[EmptyProto]) error
	RpcPathNestedStream(*NestedProto, grpc.ServerStreamingServer[EmptyProto]) error
}

// UnimplementedFlowCombinationServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFlowCombinationServer struct{}

func (UnimplementedFlowCombinationServer) RpcEmptyRpc(context.Context, *EmptyProto) (*EmptyProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcEmptyRpc not implemented")
}
func (UnimplementedFlowCombinationServer) RpcEmptyStream(*EmptyProto, grpc.ServerStreamingServer[EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method RpcEmptyStream not implemented")
}
func (UnimplementedFlowCombinationServer) StreamEmptyRpc(grpc.ClientStreamingServer[EmptyProto, EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method StreamEmptyRpc not implemented")
}
func (UnimplementedFlowCombinationServer) StreamEmptyStream(grpc.BidiStreamingServer[EmptyProto, EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method StreamEmptyStream not implemented")
}
func (UnimplementedFlowCombinationServer) RpcBodyRpc(context.Context, *NonEmptyProto) (*EmptyProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcBodyRpc not implemented")
}
func (UnimplementedFlowCombinationServer) RpcPathSingleNestedRpc(context.Context, *SingleNestedProto) (*EmptyProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcPathSingleNestedRpc not implemented")
}
func (UnimplementedFlowCombinationServer) RpcPathNestedRpc(context.Context, *NestedProto) (*EmptyProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcPathNestedRpc not implemented")
}
func (UnimplementedFlowCombinationServer) RpcBodyStream(*NonEmptyProto, grpc.ServerStreamingServer[EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method RpcBodyStream not implemented")
}
func (UnimplementedFlowCombinationServer) RpcPathSingleNestedStream(*SingleNestedProto, grpc.ServerStreamingServer[EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method RpcPathSingleNestedStream not implemented")
}
func (UnimplementedFlowCombinationServer) RpcPathNestedStream(*NestedProto, grpc.ServerStreamingServer[EmptyProto]) error {
	return status.Errorf(codes.Unimplemented, "method RpcPathNestedStream not implemented")
}
func (UnimplementedFlowCombinationServer) testEmbeddedByValue() {}

// UnsafeFlowCombinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowCombinationServer will
// result in compilation errors.
type UnsafeFlowCombinationServer interface {
	mustEmbedUnimplementedFlowCombinationServer()
}

func RegisterFlowCombinationServer(s grpc.ServiceRegistrar, srv FlowCombinationServer) {
	// If the following call pancis, it indicates UnimplementedFlowCombinationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FlowCombination_ServiceDesc, srv)
}

func _FlowCombination_RpcEmptyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcEmptyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlowCombination_RpcEmptyRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcEmptyRpc(ctx, req.(*EmptyProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcEmptyStream(m, &grpc.GenericServerStream[EmptyProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcEmptyStreamServer = grpc.ServerStreamingServer[EmptyProto]

func _FlowCombination_StreamEmptyRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyRpc(&grpc.GenericServerStream[EmptyProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_StreamEmptyRpcServer = grpc.ClientStreamingServer[EmptyProto, EmptyProto]

func _FlowCombination_StreamEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyStream(&grpc.GenericServerStream[EmptyProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_StreamEmptyStreamServer = grpc.BidiStreamingServer[EmptyProto, EmptyProto]

func _FlowCombination_RpcBodyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonEmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcBodyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlowCombination_RpcBodyRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcBodyRpc(ctx, req.(*NonEmptyProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcPathSingleNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlowCombination_RpcPathSingleNestedRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, req.(*SingleNestedProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcPathNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlowCombination_RpcPathNestedRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, req.(*NestedProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcBodyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NonEmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcBodyStream(m, &grpc.GenericServerStream[NonEmptyProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcBodyStreamServer = grpc.ServerStreamingServer[EmptyProto]

func _FlowCombination_RpcPathSingleNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SingleNestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathSingleNestedStream(m, &grpc.GenericServerStream[SingleNestedProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcPathSingleNestedStreamServer = grpc.ServerStreamingServer[EmptyProto]

func _FlowCombination_RpcPathNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathNestedStream(m, &grpc.GenericServerStream[NestedProto, EmptyProto]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlowCombination_RpcPathNestedStreamServer = grpc.ServerStreamingServer[EmptyProto]

// FlowCombination_ServiceDesc is the grpc.ServiceDesc for FlowCombination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowCombination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.FlowCombination",
	HandlerType: (*FlowCombinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcEmptyRpc",
			Handler:    _FlowCombination_RpcEmptyRpc_Handler,
		},
		{
			MethodName: "RpcBodyRpc",
			Handler:    _FlowCombination_RpcBodyRpc_Handler,
		},
		{
			MethodName: "RpcPathSingleNestedRpc",
			Handler:    _FlowCombination_RpcPathSingleNestedRpc_Handler,
		},
		{
			MethodName: "RpcPathNestedRpc",
			Handler:    _FlowCombination_RpcPathNestedRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RpcEmptyStream",
			Handler:       _FlowCombination_RpcEmptyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEmptyRpc",
			Handler:       _FlowCombination_StreamEmptyRpc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamEmptyStream",
			Handler:       _FlowCombination_StreamEmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RpcBodyStream",
			Handler:       _FlowCombination_RpcBodyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathSingleNestedStream",
			Handler:       _FlowCombination_RpcPathSingleNestedStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathNestedStream",
			Handler:       _FlowCombination_RpcPathNestedStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/flow_combination.proto",
}
