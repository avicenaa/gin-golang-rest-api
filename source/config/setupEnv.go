package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getLoadEnvMongo() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGOURI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in the .env file")
	}

	return mongoURI
}
