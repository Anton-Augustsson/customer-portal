// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api/robot-remote-controller-service/robot-remote-controller-service.proto

package robot_remote_controller_service

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

// RobotRemoteControllerServiceClient is the client API for RobotRemoteControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotRemoteControllerServiceClient interface {
	Speed(ctx context.Context, in *SpeedRequest, opts ...grpc.CallOption) (*SpeedReply, error)
	Steer(ctx context.Context, in *SteerRequest, opts ...grpc.CallOption) (*SteerReply, error)
}

type robotRemoteControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotRemoteControllerServiceClient(cc grpc.ClientConnInterface) RobotRemoteControllerServiceClient {
	return &robotRemoteControllerServiceClient{cc}
}

func (c *robotRemoteControllerServiceClient) Speed(ctx context.Context, in *SpeedRequest, opts ...grpc.CallOption) (*SpeedReply, error) {
	out := new(SpeedReply)
	err := c.cc.Invoke(ctx, "/RobotRemoteControllerService/Speed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotRemoteControllerServiceClient) Steer(ctx context.Context, in *SteerRequest, opts ...grpc.CallOption) (*SteerReply, error) {
	out := new(SteerReply)
	err := c.cc.Invoke(ctx, "/RobotRemoteControllerService/Steer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotRemoteControllerServiceServer is the server API for RobotRemoteControllerService service.
// All implementations must embed UnimplementedRobotRemoteControllerServiceServer
// for forward compatibility
type RobotRemoteControllerServiceServer interface {
	Speed(context.Context, *SpeedRequest) (*SpeedReply, error)
	Steer(context.Context, *SteerRequest) (*SteerReply, error)
	mustEmbedUnimplementedRobotRemoteControllerServiceServer()
}

// UnimplementedRobotRemoteControllerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotRemoteControllerServiceServer struct {
}

func (UnimplementedRobotRemoteControllerServiceServer) Speed(context.Context, *SpeedRequest) (*SpeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Speed not implemented")
}
func (UnimplementedRobotRemoteControllerServiceServer) Steer(context.Context, *SteerRequest) (*SteerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Steer not implemented")
}
func (UnimplementedRobotRemoteControllerServiceServer) mustEmbedUnimplementedRobotRemoteControllerServiceServer() {
}

// UnsafeRobotRemoteControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotRemoteControllerServiceServer will
// result in compilation errors.
type UnsafeRobotRemoteControllerServiceServer interface {
	mustEmbedUnimplementedRobotRemoteControllerServiceServer()
}

func RegisterRobotRemoteControllerServiceServer(s grpc.ServiceRegistrar, srv RobotRemoteControllerServiceServer) {
	s.RegisterService(&RobotRemoteControllerService_ServiceDesc, srv)
}

func _RobotRemoteControllerService_Speed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotRemoteControllerServiceServer).Speed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RobotRemoteControllerService/Speed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotRemoteControllerServiceServer).Speed(ctx, req.(*SpeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotRemoteControllerService_Steer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotRemoteControllerServiceServer).Steer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RobotRemoteControllerService/Steer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotRemoteControllerServiceServer).Steer(ctx, req.(*SteerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotRemoteControllerService_ServiceDesc is the grpc.ServiceDesc for RobotRemoteControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotRemoteControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RobotRemoteControllerService",
	HandlerType: (*RobotRemoteControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Speed",
			Handler:    _RobotRemoteControllerService_Speed_Handler,
		},
		{
			MethodName: "Steer",
			Handler:    _RobotRemoteControllerService_Steer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/robot-remote-controller-service/robot-remote-controller-service.proto",
}