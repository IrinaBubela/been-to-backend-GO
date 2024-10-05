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

	// Connect to MongoDB
	connectToDB()

	// Register routes
	routes.RegisterUserRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Connect to MongoDB
func connectToDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Set global client
	routes.SetClient(client)
}
