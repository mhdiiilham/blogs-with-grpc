package blog

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post struct
type Post struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"title"`
}
