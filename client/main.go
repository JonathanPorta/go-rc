package main

import (
	"fmt"
	"log"

	pb "github.com/jonathanporta/go-rc/remotecontrol"
	termbox "github.com/nsf/termbox-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:9000"
)

var keyToString = map[direction]string{
	UP:    "up",
	DOWN:  "down",
	RIGHT: "right",
	LEFT:  "left",
}

var client pb.RemoteControllerClient

func main() {
	// Establish a connection with the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to %s: %v", address, err)
	}
	defer conn.Close()

	// Instiate the client.
	client = pb.NewRemoteControllerClient(conn)

	// Fire up our termbox for getting input
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	clearTerm()

	keyChan := make(chan keyboardEvent)
	shutdown := make(chan struct{})
	go listenToKeyboard(keyChan)
	go dostuff(keyChan, shutdown)
	<-shutdown
	termbox.Close()
}

func clearTerm() {
	termbox.Clear(termbox.ColorCyan, termbox.ColorBlack)
	// todo handle errors
	termbox.Flush()
	termbox.Sync()
}

func dostuff(events chan keyboardEvent, shutdown chan struct{}) {
	// loop forever waiting for events from the terminal
	for {
		ev := <-events

		if ev.eventType == END {
			shutdown <- struct{}{}
		}

		if ev.eventType == RETRY {
			clearTerm()
			continue
		}

		dir := keyToDirection(ev.key)
		if dir != 0 {
			fmt.Println(keyToString[dir])
			go move(dir)
		}
	}
}

func move(d direction) {
	switch d {
	case LEFT:
		r, err := client.Left(context.Background(), &pb.ControlRequest{Value: "Test"})
		if err != nil {
			log.Fatalf("Unable to send ControlRequest::Left - %v", err)
		}
		log.Printf("ControlResponse: %v", r)
	case RIGHT:
		r, err := client.Right(context.Background(), &pb.ControlRequest{Value: "Test"})
		if err != nil {
			log.Fatalf("Unable to send ControlRequest::Right - %v", err)
		}
		log.Printf("ControlResponse: %v", r)

	case UP:
		r, err := client.Forward(context.Background(), &pb.ControlRequest{Value: "Test"})
		if err != nil {
			log.Fatalf("Unable to send ControlRequest::Forward - %v", err)
		}
		log.Printf("ControlResponse: %v", r)

	case DOWN:
		r, err := client.Backward(context.Background(), &pb.ControlRequest{Value: "Test"})
		if err != nil {
			log.Fatalf("Unable to send ControlRequest::Backward - %v", err)
		}
		log.Printf("ControlResponse: %v", r)

	}
}
