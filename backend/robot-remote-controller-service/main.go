package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Anton-Augustsson/customer-portal/robot-remote-controller-service/api/robot-remote-controller-service"
	pb_subscription "github.com/Anton-Augustsson/customer-portal/robot-remote-controller-service/api/subscription-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port                      = flag.Int("port", 50052, "The server port")
	subscription_service_addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type server struct {
	pb.UnimplementedRobotRemoteControllerServiceServer
}

func (s *server) Speed(ctx context.Context, in *pb.SpeedRequest) (*pb.SpeedReply, error) {
	log.Printf("Received: %v", in.GetSpeedOpt())

	conn, err := grpc.Dial(*subscription_service_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_subscription.NewSubscriptionServiceClient(conn)

	// Contact the server and print out its response.
	ctx_subscription, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ControlAccess(ctx_subscription, &pb_subscription.ControlAccessRequest{
		SubscriberId: 123,
		RobotId:      312,
		ActuatorId:   21,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %t", r.GetIsAllowed())

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
	pb.RegisterRobotRemoteControllerServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
