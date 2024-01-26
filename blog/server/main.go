package main

import (
	"log"
	"net"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
}

var address string = "localhost:3002"

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed listening address: %v\n", err)
	}

	log.Printf("Listening on address: %v\n", address)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed serving: %v\n", err)
	}

}
