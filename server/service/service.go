package service

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"

	"blogs/server/entity/blog"
	blogpb "blogs/server/protos"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
)

// Server struct
type Server struct {
	Network    string
	Address    string
	Collection *mongo.Collection
}

// NewService function
// to create new gRPC Service
func NewService(network, address string, mongoDBCollection *mongo.Collection) *Server {
	return &Server{
		Network:    network,
		Address:    address,
		Collection: mongoDBCollection,
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

	res, err := s.Collection.InsertOne(context.Background(), post)

	if err != nil {
		log.Printf("Error when inserting new document. Error: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Error: when converting to OID")
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	return &blogpb.CreatePostResponse{
		Message: "Success create new blog post",
		Post: &blogpb.Post{
			Id:       oid.Hex(),
			AuthorId: data.GetAuthorId(),
			Title:    data.GetTitle(),
			Content:  data.GetContent(),
		},
	}, nil
}
