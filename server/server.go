package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	blogpb "blogs/server/protos"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	collection := mongoDBConnection()

	// Blogs server instance
	blogServer := newServer(collection)

	lis, err := net.Listen(blogServer.Network, blogServer.Address)
	if err != nil {
		log.Fatalf("Failed to listen to port 50051. Error: %v", err)
	}
	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, blogServer)
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

func newServer(mongoDBCollection *mongo.Collection) *server {
	return &server{
		Network:    os.Getenv("SERVER_NETWORK"),
		Address:    os.Getenv("SERVER_ADDRESS"),
		Collection: mongoDBCollection,
	}
}

func mongoDBConnection() *mongo.Collection {
	mongoURI := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.ub9ns.mongodb.net/%s?retryWrites=true&w=majority",
		os.Getenv("MONGO_DB_USER"),
		os.Getenv("MONGO_DB_PASS"),
		os.Getenv("MONGO_DB_DB"),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Cannot connect to MongoDB: %v", err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_DB_COLLECTION"))
	return collection
}
