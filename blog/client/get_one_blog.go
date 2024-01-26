package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/grpc/status"
)

func GetOneBlog(c pb.BlogServiceClient) {
	log.Print("GetOneBlog was invoked ,,,")

	blog, err := c.GetOneBlog(context.Background(), &pb.BlogId{Id: "65b358957665261755c89c2d"})
	if err != nil {
		_, ok := status.FromError(err)

		if ok {
			log.Printf("Error from server: %v\n", err)
		} else {
			log.Printf("Non gRPC err: %v", err)
		}
	}

	log.Printf("Success get one: %v", blog)
}
