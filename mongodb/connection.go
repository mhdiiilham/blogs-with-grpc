package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDBConnection function
// to create new connection (WTF??)
func NewMongoDBConnection(user, pass, db, collection string) (*mongo.Client, *mongo.Collection) {
	log.Println("Connection to MongoDB")

	mongoURI := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.ub9ns.mongodb.net/%s?retryWrites=true&w=majority",
		user,
		pass,
		db,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Cannot connect to MongoDB: %v", err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	mongoCollection := client.Database(db).Collection(collection)
	return client, mongoCollection
}
