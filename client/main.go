package main

import (
	"context"
	"log"

	pb "github.com/jonathanporta/go-rc/remotecontrol"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to %s: %v", address, err)
	}
	defer conn.Close()
	c := pb.NewRemoteControllerClient(conn)

	r, err := c.Left(context.Background(), &pb.ControlRequest{Value: "Test"})
	if err != nil {
		log.Fatalf("Unable to send ControlRequest::Left - %v", err)
	}
	log.Printf("ControlResponse: %v", r)
}
