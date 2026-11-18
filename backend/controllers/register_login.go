package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Data from user
type DevInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Array for Users
var Users = make(map[string]DevInfo)

// Register controller
func Register(c *gin.Context) {
	var input DevInfo

	// Reading data for Error
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data ....!"})
		return
	}

	// NO empty Fields
	if input.Name == "" || input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "All fields are required!"})
		return
	}

	// Existance of Email
	if _, exists := Users[input.Email]; exists {
		c.JSON(400, gin.H{"error": "Email is already registered!"})
		return
	}

	// Hashed the Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		c.JSON(500, gin.H{"error": "Password not encrypted on Server...!"})
		return
	}
	// Hashed Password Saved
	input.Password = string(hashedPassword)

	// Saved Email
	Users[input.Email] = input

	// Successful Message
	c.JSON(http.StatusOK, gin.H{"message": "Registered successfully!"})
}

// Sign in controllers
func Signin(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Reading data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data!"})
		return
	}

	// Check for emptance
	if input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "Email and Password are required!"})
		return
	}

	//    User Exist or not
	savedUser, exists := Users[input.Email]
	if !exists {
		c.JSON(401, gin.H{"error": "Email not found!"})
		return
	}

	// Compare the user vs input email , and password
	err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Wrong password!"})
		return
	}

	// Token  Generating
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": savedUser.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	// Signature
	tokenString, err := token.SignedString([]byte("mySecretKey123"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token!"})
		return
	}

	// Step F: Send token & success response
	c.JSON(200, gin.H{
		"message": "Login successful!",
		"token":   tokenString,
		"user": gin.H{
			"name":  savedUser.Name,
			"email": savedUser.Email,
		},
	})
}
