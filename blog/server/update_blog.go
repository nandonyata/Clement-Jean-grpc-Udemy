package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("Update blog was invoked with: %v\n", in)

	objectId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"Failed to parse id",
		)
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"authorid": in.AuthorId,
			"title":    in.Title,
			"content":  in.Content,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Failed updating: %v", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			fmt.Sprintf("Data with id:%v not found", in.Id),
		)
	}

	log.Printf("RES UPDATE: %v", res)

	return &emptypb.Empty{}, nil
}
