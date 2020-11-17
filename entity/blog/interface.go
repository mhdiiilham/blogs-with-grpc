package blog

import "go.mongodb.org/mongo-driver/bson/primitive"

// Writer interface
type Writer interface {
	Create(post Post) string
	UpdateTitle(id primitive.ObjectID, newTitle string) (*Post, error)
}

// Reader Interface
type Reader interface {
	GetOne(id primitive.ObjectID) (*Post, error)
	Find() ([]*Post, error)
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
