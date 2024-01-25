package main

import (
	"io"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Print("GreetEveryOne was invoked!")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed reading client streaming: %v\n", err)
		}

		res := "Helo Bro " + req.FirstName
		if err = stream.Send(&pb.GreetResponse{
			Result: res,
		}); err != nil {
			log.Fatalf("Failed sending stream: %v\n", err)
		}

	}
}
