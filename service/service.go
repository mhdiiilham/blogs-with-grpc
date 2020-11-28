package service

import (
	"context"
	"log"
	"time"

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

	if data.GetAuthorId() == "" || data.GetTitle() == "" || data.GetContent() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"You Must Input All Required Fields",
		)
	}

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

// UpdatePostTitle handler
func (s *Server) UpdatePostTitle(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid Post ID",
		)
	}

	newPost, err := s.Manager.UpdateTitle(oid, req.GetNewTitle())
	if err != nil {
		log.Println("Error Occur: ", err.Error())
		return nil, status.Errorf(
			codes.NotFound,
			err.Error(),
		)
	}

	return &blogpb.UpdatePostResponse{
		Post: &blogpb.Post{
			Id:       newPost.ID.Hex(),
			AuthorId: newPost.AuthorID,
			Title:    req.GetNewTitle(),
			Content:  newPost.Content,
		},
	}, nil
}

// Find handler
func (s *Server) Find(ctx context.Context, req *blogpb.FindRequest) (*blogpb.FindResponse, error) {
	var res []*blogpb.Post

	posts, err := s.Manager.Find()

	if err != nil {
		log.Println("Error: ", err)
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	for _, post := range posts {
		res = append(res, &blogpb.Post{
			Id:       post.ID.Hex(),
			AuthorId: post.AuthorID,
			Title:    post.Title,
			Content:  post.Content,
		})
	}

	return &blogpb.FindResponse{
		Post: res,
	}, nil
}

// List handler
func (s *Server) List(req *blogpb.ListRequest, stream blogpb.BlogService_ListServer) error {
	log.Println("List all blog collections. Req: ", req)

	posts, err := s.Manager.Find()
	if err != nil {
		log.Println("Error: ", err)
		return status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	for _, post := range posts {
		log.Println("Sending:", post.ID.Hex())
		stream.Send(&blogpb.ListResponse{
			Post: &blogpb.Post{
				Id:       post.ID.Hex(),
				AuthorId: post.AuthorID,
				Title:    post.Title,
				Content:  post.Content,
			},
		})
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
