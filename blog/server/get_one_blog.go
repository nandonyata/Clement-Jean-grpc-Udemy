package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOneBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("GetOneBlog was invoked with: %v\n", in)

	objectId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("cant parse id: %v, err: %v", in.Id, err))
	}

	data := &BlogItem{}
	filter := bson.M{"_id": objectId}

	err = collection.FindOne(ctx, filter).Decode(data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Not Found: %v", err))

		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("Err Internal: %v", err))
	}

	return documentToBlog(data), nil
}
