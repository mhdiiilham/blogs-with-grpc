package blog

// Writer interface
type Writer interface {
	Create(post Post) string
}

// Repository interface
type Repository interface {
	Writer
}

// Manager interface
type Manager interface {
	Repository
}
