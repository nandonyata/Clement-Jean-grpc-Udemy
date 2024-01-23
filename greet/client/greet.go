package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Print("doGreet was invoked")

	resp, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "pewpew",
	})

	if err != nil {
		log.Fatalf("Failed greeting from doGreet %v\n", err)
	}

	log.Printf("Response from doGreet: %v\n", resp.Result)
}
