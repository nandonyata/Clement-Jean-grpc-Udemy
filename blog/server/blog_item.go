package main

import (
	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id", omitempty`
	AuthorId string             `bson:"author_id`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(b *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       b.ID.Hex(),
		AuthorId: b.AuthorId,
		Title:    b.Title,
		Content:  b.Content,
	}
}
