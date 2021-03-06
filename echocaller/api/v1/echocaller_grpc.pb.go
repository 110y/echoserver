// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: echocaller/api/v1/echocaller.proto

package v1

import (
	context "context"
	v1 "github.com/110y/echoserver/echoserver/api/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EchoCallerClient is the client API for EchoCaller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoCallerClient interface {
	Echo(ctx context.Context, in *v1.EchoRequest, opts ...grpc.CallOption) (*v1.EchoResponse, error)
}

type echoCallerClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoCallerClient(cc grpc.ClientConnInterface) EchoCallerClient {
	return &echoCallerClient{cc}
}

func (c *echoCallerClient) Echo(ctx context.Context, in *v1.EchoRequest, opts ...grpc.CallOption) (*v1.EchoResponse, error) {
	out := new(v1.EchoResponse)
	err := c.cc.Invoke(ctx, "/labolith.echocaller.v1.EchoCaller/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoCallerServer is the server API for EchoCaller service.
// All implementations should embed UnimplementedEchoCallerServer
// for forward compatibility
type EchoCallerServer interface {
	Echo(context.Context, *v1.EchoRequest) (*v1.EchoResponse, error)
}

// UnimplementedEchoCallerServer should be embedded to have forward compatible implementations.
type UnimplementedEchoCallerServer struct {
}

func (UnimplementedEchoCallerServer) Echo(context.Context, *v1.EchoRequest) (*v1.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

// UnsafeEchoCallerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoCallerServer will
// result in compilation errors.
type UnsafeEchoCallerServer interface {
	mustEmbedUnimplementedEchoCallerServer()
}

func RegisterEchoCallerServer(s grpc.ServiceRegistrar, srv EchoCallerServer) {
	s.RegisterService(&EchoCaller_ServiceDesc, srv)
}

func _EchoCaller_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoCallerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/labolith.echocaller.v1.EchoCaller/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoCallerServer).Echo(ctx, req.(*v1.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoCaller_ServiceDesc is the grpc.ServiceDesc for EchoCaller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoCaller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "labolith.echocaller.v1.EchoCaller",
	HandlerType: (*EchoCallerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoCaller_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echocaller/api/v1/echocaller.proto",
}
