package main

import (
	"fmt"
	pb "github.com/jonathanporta/go-rc/remotecontrol"
	"github.com/jonathanporta/go-rc/server/controller"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":9000"
)

type server struct{}

func (s *server) Left(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	go controller.MoveLeft()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Right(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	go controller.MoveRight()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Forward(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	go controller.MoveForward()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Backward(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	go controller.MoveBackward()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Stop(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	go controller.Stop()
	return &pb.ControlReply{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRemoteControllerServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	} else {
		fmt.Printf("Listening on '%v'\n", port)
	}
}
