package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Print("doGreetEveryone was invoked")

	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "nandoW",
		},
		{
			FirstName: "pewpew",
		},
		{
			FirstName: "wow",
		},
		{
			FirstName: "gge",
		},
	}

	waitC := make(chan struct{})
	go func() {
		for _, e := range reqs {
			log.Printf("Sending req: %v\n", e)
			stream.Send(e)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed receiving: %v\n", err)
				break
			}

			log.Printf("Received nice: %v\n", res)
		}
		close(waitC)
	}()

	<-waitC
}
