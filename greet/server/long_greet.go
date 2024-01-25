package main

import (
	"io"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Print("LongGreet was invoked")

	res := "WOI APA LU"

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Failed streaming, : %v\n", err)
		}

		log.Printf("Receiving req long greet: %v\n", req)
		res += " " + req.FirstName
	}

}
