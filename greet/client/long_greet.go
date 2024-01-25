package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Print("doLongGreet was invoked")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("Failed calling long greet: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "nando"},
		{FirstName: "pewpew"},
		{FirstName: "rido"},
		{FirstName: "luna"},
	}

	for _, r := range reqs {
		stream.Send(r)
		log.Printf("Sending req: %v\n", r)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Err receiving response longGreet: %v\n", err)
	}
	log.Printf("LongGreetDone: %v\n", res)

}
