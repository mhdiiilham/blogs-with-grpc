package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Variables of config
type Variables struct {
	SERVER_NETWORK string
	SERVER_ADDRESS string

	MONGO_DB            string
	MONGO_DB_COLLECTION string
	MONGO_DB_USER       string
	MONGO_DB_PASS       string
}

// LoadVariables function
func LoadVariables() *Variables {
	return &Variables{
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		SERVER_NETWORK: os.Getenv("SERVER_NETWORK"),

		MONGO_DB:            os.Getenv("MONGO_DB"),
		MONGO_DB_COLLECTION: os.Getenv("MONGO_DB_COLLECTION"),
		MONGO_DB_USER:       os.Getenv("MONGO_DB_USER"),
		MONGO_DB_PASS:       os.Getenv("MONGO_DB_PASS"),
	}
}
