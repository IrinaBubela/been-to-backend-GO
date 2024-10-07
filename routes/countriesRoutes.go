package routes

import (
	"context"
	"net/http"

	"go-been-to/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddCountryHandler allows the authenticated user to add a country to their list
func AddCountryHandler(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse the country from the request body
	var body struct {
		Country string `json:"country"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert userId to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch the user from MongoDB
	var user models.User
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Add the new country to the user's list
	user.Countries = append(user.Countries, body.Country)

	// Update the user in MongoDB
	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": bson.M{"countries": user.Countries}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add country"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// RemoveCountryHandler allows the authenticated user to remove a country from their list
func RemoveCountryHandler(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse the country from the request body
	var body struct {
		Country string `json:"country"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert userId to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch the user from MongoDB
	var user models.User
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if the country exists in the user's list
	var found bool
	for _, c := range user.Countries {
		if c == body.Country {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country not found in user's list"})
		return
	}

	// Remove the country from the user's list
	var updatedCountries []string
	for _, c := range user.Countries {
		if c != body.Country {
			updatedCountries = append(updatedCountries, c)
		}
	}

	// Update the user in MongoDB
	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": bson.M{"countries": updatedCountries}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove country"})
		return
	}

	user.Countries = updatedCountries
	c.JSON(http.StatusOK, user)
}

// GetCountriesHandler returns the list of countries for the authenticated user
func GetCountriesHandler(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert userId to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch the user from MongoDB
	var user models.User
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user's countries
	c.JSON(http.StatusOK, gin.H{"countries": user.Countries})
}
