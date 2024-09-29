package config

import (
	"context"
	"log"
	"os"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

// ConnectDB - Connect to MongoDB
func ConnectDB() *mongo.Client {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	mongoURI := os.Getenv("MONGO_URI")
    fmt.Println("MONGO_URI:", mongoURI)

	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return client
}
