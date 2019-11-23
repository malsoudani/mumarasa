package main

import (
	"fmt"
	"google.golang.org/grpc"
	"greetpb"
	"log"
	"net"
)

const (
	PROTO = "tcp"
	PORT  = "0.0.0.0:50051"
)

type server struct{}

func main() {
	fmt.Println("Hello World!")

	listener, err := net.Listen(PROTO, PORT)
	if err != nil {
		log.Fatalf("something is wrong opening listening to the tcp server %v", err)
	}

	greetpb.RegisterGreetServiceServer(s, &server{})
	s := grpc.NewServer()

	if err := s.Serve(listener); err != nil {
		log.Fatalf("falied to serve: %v", err)
	}
}
