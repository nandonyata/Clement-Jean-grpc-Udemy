package main

import (
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3002"

func main() {
	c, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed conecting address: %v\n", err)
	}

	defer c.Close()
	client := pb.NewBlogServiceClient(c)

	// createBlog(client)
	// GetOneBlog(client)
	UpdateBlog(client)
}
