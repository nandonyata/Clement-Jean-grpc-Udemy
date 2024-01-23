package main

import (
	"log"
	"net"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
	"google.golang.org/grpc"
)

var address string = "localhost:3001"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed listening on server: %v\n", err)
	}

	log.Printf("Listening on: %v\n", address)

	s := grpc.NewServer()

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed serve: %v\n", err)
	}
}
