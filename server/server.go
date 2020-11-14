package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	blogpb "blogs/server/protos"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// if we crash the go code,
	// we get the file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog server")

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen to port 50051. Error: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing listener")
	lis.Close()
}
