package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDBConnection function
// to create new connection (WTF??)
func NewMongoDBConnection() (*mongo.Client, *mongo.Collection) {
	log.Println("Connection to MongoDB")

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
	return client, collection
}
