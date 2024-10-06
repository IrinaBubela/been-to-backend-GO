package main

import (
	"context"
	"log"
	"os"

	"go-been-to/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Connect to MongoDB and check for errors
	if err := connectToDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return
	}

	// Register routes
	routes.RegisterUserRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Connect to MongoDB
func connectToDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/myapp")
	var err error

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err // Return the error if connection fails
	}

	// Ping the MongoDB server to verify connection
	if err = client.Ping(context.Background(), nil); err != nil {
		return err // Return the error if ping fails
	}

	// Set global client in routes package
	routes.SetClient(client)

	log.Println("Connected to MongoDB!")
	return nil // Return nil if everything is fine
}
