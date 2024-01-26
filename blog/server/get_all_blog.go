package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetAllBlog(_ *emptypb.Empty, stream pb.BlogService_GetAllBlogServer) error {
	log.Print("GetAllBlog was invoked !!!")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("Err internal: %v\n", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}

		if err = cur.Decode(&data); err != nil {
			return status.Error(
				codes.Internal,
				fmt.Sprintf("Err Decoding data: %v\n", err),
			)
		}

		stream.Send(documentToBlog(data))

	}

	if err = cur.Err(); err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("Err internal: %v\n", err),
		)
	}

	return nil
}
