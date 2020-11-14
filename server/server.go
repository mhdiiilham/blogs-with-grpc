package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	mongodbConn "blogs/server/mongodb"
	blogpb "blogs/server/protos"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type server struct {
	Network    string
	Address    string
	Collection *mongo.Collection
}

func main() {
	// if we crash the go code,
	// we get the file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connection to MongoDB
	client, collection := mongodbConn.NewMongoDBConnection()

	// Blogs server instance
	blogServer := newServer(collection)

	lis, err := net.Listen(blogServer.Network, blogServer.Address)
	if err != nil {
		log.Fatalf("Failed to listen to port 50051. Error: %v", err)
	}
	s := grpc.NewServer()
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

func newServer(mongoDBCollection *mongo.Collection) *server {
	log.Println("Function newServer invoked")

	return &server{
		Network:    os.Getenv("SERVER_NETWORK"),
		Address:    os.Getenv("SERVER_ADDRESS"),
		Collection: mongoDBCollection,
	}
}
