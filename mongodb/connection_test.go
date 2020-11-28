package mongodb

import (
	"blogs/config"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stretchr/testify/assert"
)

func TestSuccessCreateMongoDBConnection(t *testing.T) {
	cfg, _ := config.LoadVariables("../.env")
	client, collection, err := NewMongoDBConnection(
		cfg.MONGO_DB_USER,
		cfg.MONGO_DB_PASS,
		cfg.MONGO_DB,
		cfg.MONGO_DB_COLLECTION,
	)

	assert.Nil(t, err, "Error Expected To Be A Nil")
	assert.NotEqual(t, nil, client, "Client Expected To Be Not A Nil")
	assert.NotEqual(t, nil, collection, "Collection Expected To Be Not A Nil")
	assert.IsType(t, &mongo.Client{}, client, "Client Expected To Be Type of A MongoDB Client")
	assert.IsType(t, &mongo.Collection{}, collection, "Collection Expected To Be Type of A MongoDB Collection")
}
