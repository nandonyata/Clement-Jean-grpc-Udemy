package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/grpc/status"
)

func createBlog(client pb.BlogServiceClient) {
	fmt.Println("CreateBlog was invoked..")

	res, err := client.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "pewpew",
		Title:    "SLEBEW",
		Content:  "WHO KNOWS THE CONTENT???",
	})
	if err != nil {
		_, ok := status.FromError(err)

		if ok {
			log.Printf("An error occured from server: %v", err)
			return
		} else {
			log.Printf("A non gRPC err: %v", err)
			return
		}
	}

	log.Printf("Success create blog, result id: %v\n", res)

}
