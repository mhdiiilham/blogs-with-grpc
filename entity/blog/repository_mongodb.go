package blog

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepo struct
type MongoDBRepo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

// NewMongoDBRepository function
func NewMongoDBRepository(collection *mongo.Collection) *MongoDBRepo {
	ctx := context.Background()
	return &MongoDBRepo{
		Collection: collection,
		Ctx:        ctx,
	}
}

// Create function
func (r *MongoDBRepo) Create(post Post) string {
	res, err := r.Collection.InsertOne(r.Ctx, post)
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

// GetOne Function
func (r *MongoDBRepo) GetOne(id primitive.ObjectID) (*Post, error) {
	var post Post

	if err := r.Collection.FindOne(r.Ctx, bson.M{"_id": id}).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// UpdateTitle of a post
func (r *MongoDBRepo) UpdateTitle(id primitive.ObjectID, newTitle string) (*Post, error) {
	var post Post
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"title": newTitle}}
	if err := r.Collection.FindOneAndUpdate(r.Ctx, filter, update).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}
