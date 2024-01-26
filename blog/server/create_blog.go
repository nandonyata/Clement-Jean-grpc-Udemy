package main

import (
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create was invoked with: %v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatalf("Failed insert: %v\n", err)

		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	_id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Error(
			codes.Internal,
			"Cant conver to objectId",
		)
	}

	return &pb.BlogId{Id: _id.Hex()}, nil
}
