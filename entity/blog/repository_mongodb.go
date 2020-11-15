package blog

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepo struct
type MongoDBRepo struct {
	Collection *mongo.Collection
}

// NewMongoDBRepository function
func NewMongoDBRepository(collection *mongo.Collection) *MongoDBRepo {
	return &MongoDBRepo{
		Collection: collection,
	}
}

// Create function
func (r *MongoDBRepo) Create(post Post) string {
	res, err := r.Collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Printf("Error when inserting new document. Error: %v", err)
		return ""
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Error: when converting to OID")
		return ""
	}

	return oid.Hex()
}
