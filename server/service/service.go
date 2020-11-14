package service

import "go.mongodb.org/mongo-driver/mongo"

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
