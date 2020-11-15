package blog

import "go.mongodb.org/mongo-driver/bson/primitive"

// Writer interface
type Writer interface {
	Create(post Post) string
}

// Reader Interface
type Reader interface {
	GetOne(id primitive.ObjectID) (*Post, error)
}

// Repository interface
type Repository interface {
	Writer
	Reader
}

// Manager interface
type Manager interface {
	Repository
}
