// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/pool_service/pool.proto

package pool_service

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
	PoolService_GetPools_FullMethodName       = "/pool_service.PoolService/GetPools"
	PoolService_StreamNewPools_FullMethodName = "/pool_service.PoolService/StreamNewPools"
	PoolService_GetAssets_FullMethodName      = "/pool_service.PoolService/GetAssets"
)

// PoolServiceClient is the client API for PoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolServiceClient interface {
	GetPools(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPoolsResponse, error)
	StreamNewPools(ctx context.Context, in *StreamPoolsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPoolsResponse], error)
	GetAssets(ctx context.Context, in *GetAssetsRequest, opts ...grpc.CallOption) (*GetAssetsResponse, error)
}

type poolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolServiceClient(cc grpc.ClientConnInterface) PoolServiceClient {
	return &poolServiceClient{cc}
}

func (c *poolServiceClient) GetPools(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPoolsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPoolsResponse)
	err := c.cc.Invoke(ctx, PoolService_GetPools_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolServiceClient) StreamNewPools(ctx context.Context, in *StreamPoolsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPoolsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PoolService_ServiceDesc.Streams[0], PoolService_StreamNewPools_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamPoolsRequest, GetPoolsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PoolService_StreamNewPoolsClient = grpc.ServerStreamingClient[GetPoolsResponse]

func (c *poolServiceClient) GetAssets(ctx context.Context, in *GetAssetsRequest, opts ...grpc.CallOption) (*GetAssetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAssetsResponse)
	err := c.cc.Invoke(ctx, PoolService_GetAssets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolServiceServer is the server API for PoolService service.
// All implementations must embed UnimplementedPoolServiceServer
// for forward compatibility.
type PoolServiceServer interface {
	GetPools(context.Context, *Empty) (*GetPoolsResponse, error)
	StreamNewPools(*StreamPoolsRequest, grpc.ServerStreamingServer[GetPoolsResponse]) error
	GetAssets(context.Context, *GetAssetsRequest) (*GetAssetsResponse, error)
	mustEmbedUnimplementedPoolServiceServer()
}

// UnimplementedPoolServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPoolServiceServer struct{}

func (UnimplementedPoolServiceServer) GetPools(context.Context, *Empty) (*GetPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPools not implemented")
}
func (UnimplementedPoolServiceServer) StreamNewPools(*StreamPoolsRequest, grpc.ServerStreamingServer[GetPoolsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamNewPools not implemented")
}
func (UnimplementedPoolServiceServer) GetAssets(context.Context, *GetAssetsRequest) (*GetAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}
func (UnimplementedPoolServiceServer) mustEmbedUnimplementedPoolServiceServer() {}
func (UnimplementedPoolServiceServer) testEmbeddedByValue()                     {}

// UnsafePoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolServiceServer will
// result in compilation errors.
type UnsafePoolServiceServer interface {
	mustEmbedUnimplementedPoolServiceServer()
}

func RegisterPoolServiceServer(s grpc.ServiceRegistrar, srv PoolServiceServer) {
	// If the following call pancis, it indicates UnimplementedPoolServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PoolService_ServiceDesc, srv)
}

func _PoolService_GetPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).GetPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoolService_GetPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).GetPools(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolService_StreamNewPools_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPoolsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PoolServiceServer).StreamNewPools(m, &grpc.GenericServerStream[StreamPoolsRequest, GetPoolsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PoolService_StreamNewPoolsServer = grpc.ServerStreamingServer[GetPoolsResponse]

func _PoolService_GetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).GetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoolService_GetAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).GetAssets(ctx, req.(*GetAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoolService_ServiceDesc is the grpc.ServiceDesc for PoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pool_service.PoolService",
	HandlerType: (*PoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPools",
			Handler:    _PoolService_GetPools_Handler,
		},
		{
			MethodName: "GetAssets",
			Handler:    _PoolService_GetAssets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNewPools",
			Handler:       _PoolService_StreamNewPools_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/pool_service/pool.proto",
}
