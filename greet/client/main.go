package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nandonyata/Clement-Jean-grpc-Udemy/greet/proto"
)

var address string = "localhost:3001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed connection client: %v\n", err)
	}

	defer conn.Close()

	ccGreet := pb.NewGreetServiceClient(conn)

	// doGreet(ccGreet)
	// doGreetManyTimes(ccGreet)
	// doLongGreet(ccGreet)
	doGreetEveryOne(ccGreet)
}
