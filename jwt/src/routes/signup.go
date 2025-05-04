package routes

import (
	"jwt/auth"
	"jwt/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context, database *gorm.DB, logger *zerolog.Logger) {
	// Parse the request body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the user input
	if user.Name == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, password, and email are required."})
		return
	}

	// Check if the email is already registered
	var existingUser models.User
	if err := database.First(&existingUser, "email = ?", user.Email).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered."})
		return
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error().Msgf("Error hashing password: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := database.Create(&user).Error; err != nil {
		logger.Error().Msgf("Error saving user to database: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	// Generate a JWT token
	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		logger.Error().Msgf("Error generating token: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true) // Set the cookie with the token
	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully."})
}

// func generateToken(userEmail string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(userEmail), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "Failed to hash password.", err
// 	}
// 	return string(hash), nil
// }

func RegisterSignupRoute(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.POST("/signup", func(c *gin.Context) {
		Signup(c, database, logger)
	})
}
