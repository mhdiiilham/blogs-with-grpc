package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
func LoadVariables(file ...string) (*Variables, error) {
	var err error
	if len(file) > 0 {
		err = godotenv.Load(file[0])
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &Variables{
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		SERVER_NETWORK: os.Getenv("SERVER_NETWORK"),

		MONGO_DB:            os.Getenv("MONGO_DB"),
		MONGO_DB_COLLECTION: os.Getenv("MONGO_DB_COLLECTION"),
		MONGO_DB_USER:       os.Getenv("MONGO_DB_USER"),
		MONGO_DB_PASS:       os.Getenv("MONGO_DB_PASS"),
	}, nil
}
