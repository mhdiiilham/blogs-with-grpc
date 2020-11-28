package main

import (
	"blogs/entity/blog"
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"blogs/config"
	mongodbConn "blogs/mongodb"
	blogpb "blogs/protos"
	"blogs/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// if we crash the go code,
	// we get the file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// config variables
	cfg := config.LoadVariables()

	// Connection to MongoDB
	client, collection, mongoErr := mongodbConn.NewMongoDBConnection(
		cfg.MONGO_DB_USER,
		cfg.MONGO_DB_PASS,
		cfg.MONGO_DB,
		cfg.MONGO_DB_COLLECTION,
	)

	if mongoErr != nil {
		panic("MongoDB Connection Error")
	}

	blogRepo := blog.NewMongoDBRepository(collection)
	blogManager := blog.NewManager(blogRepo)

	// Blogs server instance
	blogServer := service.NewService(
		cfg.SERVER_NETWORK,
		cfg.SERVER_ADDRESS,
		blogManager,
	)

	lis, err := net.Listen(blogServer.Network, blogServer.Address)
	if err != nil {
		log.Fatalf("Failed to listen to port 50051. Error: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	blogpb.RegisterBlogServiceServer(s, blogServer)
	go func() {
		log.Println("Starting server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Stopping the server")
	s.Stop()
	log.Println("Closing listener")
	lis.Close()
	log.Println("Disconnection from MongoDB")
	client.Disconnect(context.TODO())
}
