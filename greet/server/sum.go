package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.First + in.Second,
	}, nil
}
