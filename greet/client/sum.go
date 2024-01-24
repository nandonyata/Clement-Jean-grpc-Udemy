package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Print("doSum was invoked")

	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		First:  10,
		Second: 3,
	})

	if err != nil {
		log.Fatalf("Failed sum: %v\n", err)
	}

	log.Printf("Result sum: %v\n", resp)
}
