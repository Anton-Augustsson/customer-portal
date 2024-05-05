package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Anton-Augustsson/customer-portal/subscription-service/api/subscription-service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSubscriptionServiceServer
}

func (s *server) LiveStreamAccess(ctx context.Context, in *pb.LiveStreamAccessRequest) (*pb.LiveStreamAccessReply, error) {
	log.Printf("Received: %v", in.GetSubscriberId())
	return &pb.LiveStreamAccessReply{
		SubscriberId: in.GetSubscriberId(),
		RobotId:      in.GetRobotId(),
		CameraId:     in.GetCameraId(),
		IsAllowed:    false,
	}, nil
}

func (s *server) ControlAccess(ctx context.Context, in *pb.ControlAccessRequest) (*pb.ControlAccessReply, error) {
	log.Printf("Received: %v", in.GetSubscriberId())
	return &pb.ControlAccessReply{
		SubscriberId: in.GetSubscriberId(),
		RobotId:      in.GetRobotId(),
		ActuatorId:   in.GetActuatorId(),
		IsAllowed:    false,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubscriptionServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
