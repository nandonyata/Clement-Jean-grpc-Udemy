package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	log.Print("SUCCESS DELETED")
}
