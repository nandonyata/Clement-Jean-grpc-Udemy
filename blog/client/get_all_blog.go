package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetAllBlog(c pb.BlogServiceClient) {
	log.Print("GetAllBlog was invoked ..")

	res, err := c.GetAllBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Err from server: %v\n", err)
	}

	for {
		stream, err := res.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Err receiving server: %v\n", err)
		}

		log.Printf("Success: %v", stream)
	}

}
