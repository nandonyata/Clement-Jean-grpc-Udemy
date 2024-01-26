package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
}

var (
	address    string = "localhost:3002"
	collection *mongo.Collection
)

func main() {
	// mongo ->
	mc, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Err mongo: %v\n", err)
	}

	if err = mc.Connect(context.Background()); err != nil {
		log.Fatalf("Failed connect to mongo: %v\n", err)
	}

	collection = mc.Database("nando-gRPC-blogDB").Collection("blog")
	// <-

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed listening address: %v\n", err)
	}

	log.Printf("Listening on address: %v\n", address)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed serving: %v\n", err)
	}

}
