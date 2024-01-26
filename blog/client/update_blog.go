package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/grpc/status"
)

func UpdateBlog(c pb.BlogServiceClient) {
	log.Print("udpate was invoked bro..")

	res, err := c.UpdateBlog(context.Background(), &pb.Blog{
		Id:       "65b358957665261755c89c2d",
		AuthorId: "NANDO NYATA",
		Title:    "NEW TAITEL BRO",
		Content:  "HAHAHAH",
	})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			log.Printf("Error from server %v\n", err)
		} else {
			log.Printf("Non grpc error %v", err)
		}
	}

	log.Printf("NOICE... %v\n", res)
}
