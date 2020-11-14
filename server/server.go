package main

import (
	"fmt"
	"log"
	"net"

	blogpb "blogs/server/protos"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Blog server")

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen to port 50051. Error: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
