package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "pew-pew!!!",
	})
	if err != nil {
		log.Fatalf("Failed while calling GreetManyTimes: %v\n", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed while reading the stream : %v\n", err)

		}

		log.Printf("Greet many times success : %v\n", resp.Result)
	}
}
