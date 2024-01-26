package main

import (
	"context"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("Delete blog was invoked with : %v\n", in)

	objectId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"Failed decoding id",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"Err deleting",
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Data not foun",
		)
	}

	return &emptypb.Empty{}, nil
}
