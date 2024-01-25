package main

import (
	"log"
	"net"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.CalculatorServiceServer
}

var address string = "localhost:3001"

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed listen: %v\n", err)
	}

	log.Print("Listening ", address)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed serving: %v\n", err)
	}
}
