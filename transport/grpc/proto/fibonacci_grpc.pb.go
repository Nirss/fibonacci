// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FibonacciServiceClient is the client API for FibonacciService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibonacciServiceClient interface {
	GetRange(ctx context.Context, in *GetRangeRequest, opts ...grpc.CallOption) (*GetRangeResponse, error)
}

type fibonacciServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFibonacciServiceClient(cc grpc.ClientConnInterface) FibonacciServiceClient {
	return &fibonacciServiceClient{cc}
}

func (c *fibonacciServiceClient) GetRange(ctx context.Context, in *GetRangeRequest, opts ...grpc.CallOption) (*GetRangeResponse, error) {
	out := new(GetRangeResponse)
	err := c.cc.Invoke(ctx, "/FibonacciService/GetRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibonacciServiceServer is the server API for FibonacciService service.
// All implementations must embed UnimplementedFibonacciServiceServer
// for forward compatibility
type FibonacciServiceServer interface {
	GetRange(context.Context, *GetRangeRequest) (*GetRangeResponse, error)
	mustEmbedUnimplementedFibonacciServiceServer()
}

// UnimplementedFibonacciServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFibonacciServiceServer struct {
}

func (UnimplementedFibonacciServiceServer) GetRange(context.Context, *GetRangeRequest) (*GetRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRange not implemented")
}
func (UnimplementedFibonacciServiceServer) mustEmbedUnimplementedFibonacciServiceServer() {}

// UnsafeFibonacciServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibonacciServiceServer will
// result in compilation errors.
type UnsafeFibonacciServiceServer interface {
	mustEmbedUnimplementedFibonacciServiceServer()
}

func RegisterFibonacciServiceServer(s grpc.ServiceRegistrar, srv FibonacciServiceServer) {
	s.RegisterService(&FibonacciService_ServiceDesc, srv)
}

func _FibonacciService_GetRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServiceServer).GetRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FibonacciService/GetRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServiceServer).GetRange(ctx, req.(*GetRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FibonacciService_ServiceDesc is the grpc.ServiceDesc for FibonacciService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FibonacciService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FibonacciService",
	HandlerType: (*FibonacciServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRange",
			Handler:    _FibonacciService_GetRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fibonacci.proto",
}
