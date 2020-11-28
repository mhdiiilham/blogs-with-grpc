package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvSuccess(t *testing.T) {
	c, err := LoadVariables("../.env.test")

	assert.Nil(t, err, "Error Expected To Be A Nil")
	assert.IsType(t, &Variables{}, c, "config should be type of `variables`")
	assert.Equal(t, "tcp", c.SERVER_NETWORK, "SERVER_NETWORK should be equal to `tcp`")
	assert.Equal(t, "0.0.0.0:50051", c.SERVER_ADDRESS, "SERVER_ADDRESS should be equal to `0.0.0.0:50051`")
	assert.Equal(t, "blogs_test", c.MONGO_DB, "MONGO_DB should be equal to `blogs_test`")
	assert.Equal(t, "blogs_test", c.MONGO_DB_COLLECTION, "MONGO_DB_COLLECTION should be equal to `blogs_test`")
	assert.Equal(t, "blogs_test", c.MONGO_DB_USER, "MONGO_DB should be equal to `blogs_test`")
	assert.Equal(t, "password_test", c.MONGO_DB_PASS, "MONGO_DB should be equal to `password_test`")
}

func TestLoadEnvNotExist(t *testing.T) {
	_, err := LoadVariables("../.env.failed")

	assert.NotNil(t, err, "Error Expected To Be Not A Nil")
	assert.Equal(t, err.Error(), "open ../.env.failed: no such file or directory")
}
