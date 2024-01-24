package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Sum was invoked with: %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.First + in.Second,
	}, nil
}
