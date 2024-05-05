package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Anton-Augustsson/customer-portal/subscription-service/api/producer"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SubscriptionLiveStreamAccess(ctx context.Context, in *pb.SubscriptionLiveStreamAccessRequest) (*pb.SubscriptionLiveStreamAccessReply, error) {
	log.Printf("Received: %v", in.GetSubscriberId())
	return &pb.SubscriptionLiveStreamAccessReply{
		SubscriberId: in.GetSubscriberId(),
		RobotId:      in.GetRobotId(),
		CameraId:     in.GetCameraId(),
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
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
