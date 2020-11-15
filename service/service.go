package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"blogs/entity/blog"
	blogpb "blogs/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server struct
type Server struct {
	Network string
	Address string
	Manager blog.Manager
}

// NewService function
// to create new gRPC Service
func NewService(network, address string, blogManager blog.Manager) *Server {
	return &Server{
		Network: network,
		Address: address,
		Manager: blogManager,
	}
}

/*
|--------------------------------------------------------------------------
| Server Methods
|--------------------------------------------------------------------------
|
| Here is where you can register gRPC Methods.
|
*/

// CreatePost Handler
func (s *Server) CreatePost(ctx context.Context, req *blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {
	data := req.GetPost()
	post := blog.Post{
		AuthorID: data.GetAuthorId(),
		Title:    data.GetTitle(),
		Content:  data.GetContent(),
	}

	oid := s.Manager.Create(post)
	if oid == "" {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	log.Printf("Created new post with id: %v", oid)
	return &blogpb.CreatePostResponse{
		Message: "Success create new blog post",
		Post: &blogpb.Post{
			Id:       oid,
			AuthorId: data.GetAuthorId(),
			Title:    data.GetTitle(),
			Content:  data.GetContent(),
		},
	}, nil
}

// ReadPost handler
func (s *Server) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	postID := req.GetId()

	oid, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid Post ID",
		)
	}

	post, err := s.Manager.GetOne(oid)

	if err != nil {
		log.Println("Error Occur: ", err.Error())
		return nil, status.Errorf(
			codes.NotFound,
			err.Error(),
		)
	}

	return &blogpb.ReadPostResponse{
		Post: &blogpb.Post{
			Id:       post.ID.Hex(),
			AuthorId: post.AuthorID,
			Title:    post.Title,
			Content:  post.Content,
		},
	}, nil
}
