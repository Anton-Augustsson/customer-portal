// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api.proto

package producer

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	SubscriptionLiveStreamAccess(ctx context.Context, in *SubscriptionLiveStreamAccessRequest, opts ...grpc.CallOption) (*SubscriptionLiveStreamAccessReply, error)
	SubscriptionControlAccess(ctx context.Context, in *SubscriptionControlAccessRequest, opts ...grpc.CallOption) (*SubscriptionControlAccessReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SubscriptionLiveStreamAccess(ctx context.Context, in *SubscriptionLiveStreamAccessRequest, opts ...grpc.CallOption) (*SubscriptionLiveStreamAccessReply, error) {
	out := new(SubscriptionLiveStreamAccessReply)
	err := c.cc.Invoke(ctx, "/Greeter/SubscriptionLiveStreamAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SubscriptionControlAccess(ctx context.Context, in *SubscriptionControlAccessRequest, opts ...grpc.CallOption) (*SubscriptionControlAccessReply, error) {
	out := new(SubscriptionControlAccessReply)
	err := c.cc.Invoke(ctx, "/Greeter/SubscriptionControlAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	SubscriptionLiveStreamAccess(context.Context, *SubscriptionLiveStreamAccessRequest) (*SubscriptionLiveStreamAccessReply, error)
	SubscriptionControlAccess(context.Context, *SubscriptionControlAccessRequest) (*SubscriptionControlAccessReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SubscriptionLiveStreamAccess(context.Context, *SubscriptionLiveStreamAccessRequest) (*SubscriptionLiveStreamAccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionLiveStreamAccess not implemented")
}
func (UnimplementedGreeterServer) SubscriptionControlAccess(context.Context, *SubscriptionControlAccessRequest) (*SubscriptionControlAccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionControlAccess not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SubscriptionLiveStreamAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionLiveStreamAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SubscriptionLiveStreamAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SubscriptionLiveStreamAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SubscriptionLiveStreamAccess(ctx, req.(*SubscriptionLiveStreamAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SubscriptionControlAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionControlAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SubscriptionControlAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SubscriptionControlAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SubscriptionControlAccess(ctx, req.(*SubscriptionControlAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscriptionLiveStreamAccess",
			Handler:    _Greeter_SubscriptionLiveStreamAccess_Handler,
		},
		{
			MethodName: "SubscriptionControlAccess",
			Handler:    _Greeter_SubscriptionControlAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
