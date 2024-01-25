package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &pb.SqrtResponse{
		Number: math.Sqrt(float64(number)),
	}, nil
}
