package blog

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type manager struct {
	Repo Repository
}

// NewManager create new repository
func NewManager(r Repository) *manager {
	return &manager{
		Repo: r,
	}
}

// Create new blog post
func (s *manager) Create(post Post) string {
	return s.Repo.Create(post)
}

// GetOne post
func (s *manager) GetOne(id primitive.ObjectID) (*Post, error) {
	return s.Repo.GetOne(id)
}

// UpdateTitle of a post
func (s *manager) UpdateTitle(id primitive.ObjectID, newTitle string) (*Post, error) {
	return s.Repo.UpdateTitle(id, newTitle)
}

// Find all posts
func (s *manager) Find() ([]*Post, error) {
	return s.Repo.Find()
}
