package routes

import (
	"context"
	"fmt"
	"net/http"

	"go-been-to/middleware"
	"go-been-to/models"
	"go-been-to/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var collection *mongo.Collection

// SetClient sets the MongoDB client to be used in the routes
func SetClient(client *mongo.Client) {
	collection = client.Database("myapp").Collection("users")
}

func RegisterUserRoutes(r *gin.Engine) {
	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
	}

	// Apply CORS middleware
	r.Use(cors.New(corsConfig))

	// Register routes
	r.POST("/api/auth/signup", signupHandler)
	r.POST("/api/auth/login", loginHandler)

	// Protected routes (require authentication)
	r.Use(middleware.AuthMiddleware())

	r.GET("/api/user/countries", GetCountriesHandler)       // New /countries route
	r.POST("/api/user/addCountry", AddCountryHandler)       // Add a new country
	r.POST("/api/user/removeCountry", RemoveCountryHandler) // Remove country from the list
}

// Signup handler
func signupHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Insert user into MongoDB
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.Status(http.StatusCreated)
}

// Login handler
func loginHandler(c *gin.Context) {
	var user models.User
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	filter := bson.M{"email": loginReq.Email}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials email"})
		return
	}

	// Compare passwords
	if err := utils.CheckPassword(user.Password, loginReq.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials passw"})
		fmt.Println(err, "err")
		return
	}

	// JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID.Hex()
	tokenString, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
