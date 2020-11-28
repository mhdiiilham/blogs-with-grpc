package service

import (
	"blogs/config"
	"blogs/entity/blog"
	blogpb "blogs/protos"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	mongodbConn "blogs/mongodb"

	"google.golang.org/grpc/codes"
)

type testCases struct {
	Name string
	Post *blogpb.Post
	Code codes.Code
}

func TestCreatePost(t *testing.T) {

	// config variables
	cfg, envErr := config.LoadVariables("../.env")
	if envErr != nil {
		panic(envErr)
	}

	// Connection to MongoDB
	client, collection, mongoErr := mongodbConn.NewMongoDBConnection(
		cfg.MONGO_DB_USER,
		cfg.MONGO_DB_PASS,
		cfg.MONGO_DB,
		cfg.MONGO_DB_COLLECTION,
	)
	defer client.Disconnect(context.TODO())

	if mongoErr != nil {
		panic("MongoDB Connection Error")
	}

	blogRepo := blog.NewMongoDBRepository(collection)
	blogManager := blog.NewManager(blogRepo)

	// Blogs server instance
	blogServer := NewService(
		cfg.SERVER_NETWORK,
		cfg.SERVER_ADDRESS,
		blogManager,
	)

	testCases := []testCases{
		testCases{
			Name: "Success Posted A Post",
			Post: &blogpb.Post{
				AuthorId: "1",
				Title:    "Test Post",
				Content:  "This Post Was Created From Tests",
			},
			Code: codes.OK,
		},
		testCases{
			Name: "Failed Posted A Post",
			Post: &blogpb.Post{
				AuthorId: "",
				Title:    "",
				Content:  "",
			},
			Code: codes.InvalidArgument,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {

			req := &blogpb.CreatePostRequest{
				Post: testCase.Post,
			}

			fmt.Println(testCase.Name)
			res, err := blogServer.CreatePost(context.TODO(), req)
			if testCase.Code == codes.OK {
				assert.Nil(t, err, "Error should be nil")
				assert.NotEmpty(t, res.GetPost(), "Response should not be empty")
				assert.Equal(t, testCase.Post.GetTitle(), res.GetPost().GetTitle())
			} else {
				assert.NotNil(t, err, "Error should be not nill")
				assert.Empty(t, res, "Response should be empty")
			}

		})
	}
}
