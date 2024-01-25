package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, input int32) {
	log.Printf("doSqrt was invoked with: %v/n", input)

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: input,
	})
	if err != nil {
		status, ok := status.FromError(err)

		if ok {
			log.Printf("An Error from server: %v\n", err)

			if status.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Non gRPC err: %v\n", err)
		}
	}

	log.Printf("Nice sqrt: %v\n", res)
}
