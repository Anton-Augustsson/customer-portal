package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Anton-Augustsson/customer-portal/robot-remote-controller-service/api/robot-remote-controller-service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Speed(ctx context.Context, in *pb.SpeedRequest) (*pb.SpeedReply, error) {
	log.Printf("Received: %v", in.GetSpeedOpt())
	return &pb.SpeedReply{
		SubscriberId: in.GetSubscriberId(),
		RobotId:      in.GetRobotId(),
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
