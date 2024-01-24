package main

import (
	"log"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed connecting: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	doSum(client)

}
