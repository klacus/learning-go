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

func Login(c *gin.Context, database *gorm.DB, logger *zerolog.Logger) {
	// Parse the request body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the user input
	if user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password, and email are required."})
		return
	}

	// Get user by email
	var existingUser models.User
	if err := database.First(&existingUser, "email = ?", user.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email or password invalid."}) // Do not return what components of the credential was invalid to avoid leaking information.
		return
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password invalid."}) // Do not return what components of the credential was invalid to avoid leaking information.
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
	c.JSON(http.StatusOK, gin.H{"message": "User login successful."})
}

func RegisterLoginRoute(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.POST("/login", func(c *gin.Context) {
		Login(c, database, logger)
	})
}
