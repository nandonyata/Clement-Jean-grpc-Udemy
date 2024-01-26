package main

import (
	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"

	"context"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

	return &pb.BlogId{}, nil
}
