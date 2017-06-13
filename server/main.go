package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/configor"
	pb "github.com/jonathanporta/go-rc/remotecontrol"
	"github.com/jonathanporta/go-rc/server/controller"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Configuration struct {
	Motors []controller.MotorConfiguration
	Port   string `default:":9000"`
}

var car controller.CarController

var Config = Configuration{}

type server struct {
	// car    controller.CarController
	Config Configuration
}

func (s *server) Left(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	car.Left()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Right(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	car.Right()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Forward(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	fmt.Printf("\n\nfunc (s *server) Forward(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {:\n\n%v\n\n", car)
	car.Forward()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Backward(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	car.Backward()
	return &pb.ControlReply{Success: true}, nil
}
func (s *server) Stop(ctx context.Context, in *pb.ControlRequest) (*pb.ControlReply, error) {
	car.Stop()
	return &pb.ControlReply{Success: true}, nil
}

func (s *server) Listen() {
	lis, err := net.Listen("tcp", s.Config.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		fmt.Printf("Listening on '%v'\n", s.Config.Port)
	}

	instance := grpc.NewServer()
	pb.RegisterRemoteControllerServer(instance, &server{})

	reflection.Register(instance)
	if err := instance.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func newServer(c Configuration) server {
	car = controller.NewCarController(c.Motors)
	fmt.Printf("ctrl: \n=================================\n\n%v\n\n===========================", car)
	return server{c}
}

func main() {
	fmt.Printf("Starting up server...\n")
	configor.Load(&Config, "config.yml")
	fmt.Printf("Loaded configuration:\n%#v\n", Config)

	s := newServer(Config)
	s.Listen()

	fmt.Printf("Exiting!\n")
}
