// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package iot

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

// IOTClient is the client API for IOT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IOTClient interface {
	UpdateDevice(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type iOTClient struct {
	cc grpc.ClientConnInterface
}

func NewIOTClient(cc grpc.ClientConnInterface) IOTClient {
	return &iOTClient{cc}
}

func (c *iOTClient) UpdateDevice(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/iot.IOT/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IOTServer is the server API for IOT service.
// All implementations must embed UnimplementedIOTServer
// for forward compatibility
type IOTServer interface {
	UpdateDevice(context.Context, *UpdateRequest) (*Empty, error)
	mustEmbedUnimplementedIOTServer()
}

// UnimplementedIOTServer must be embedded to have forward compatible implementations.
type UnimplementedIOTServer struct {
}

func (UnimplementedIOTServer) UpdateDevice(context.Context, *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedIOTServer) mustEmbedUnimplementedIOTServer() {}

// UnsafeIOTServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IOTServer will
// result in compilation errors.
type UnsafeIOTServer interface {
	mustEmbedUnimplementedIOTServer()
}

func RegisterIOTServer(s grpc.ServiceRegistrar, srv IOTServer) {
	s.RegisterService(&IOT_ServiceDesc, srv)
}

func _IOT_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOTServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iot.IOT/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOTServer).UpdateDevice(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IOT_ServiceDesc is the grpc.ServiceDesc for IOT service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IOT_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iot.IOT",
	HandlerType: (*IOTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDevice",
			Handler:    _IOT_UpdateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/iot/iot.proto",
}
