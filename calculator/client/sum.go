package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Print("Do sum was invoked")

	resp, err := client.Sum(context.Background(), &pb.CalculatorRequest{
		First:  12,
		Second: 3,
	})

	if err != nil {
		log.Fatalf("Failed doing sum: %v\n", err)
	}

	log.Printf("Result sum: %v\n", resp)
}
