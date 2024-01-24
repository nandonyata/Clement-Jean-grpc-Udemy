package main

import (
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %v, number %v", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
