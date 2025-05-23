package controllers

import (
	"jwt/models"
	"jwt/token"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	token, err := NewUserToken(user.Email, database, logger)
	if err != nil {
		logger.Error().Msgf("Error generating token: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.Header("Authorization", "Bearer "+token) // Set the Authorization header with the token
	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully."})
}

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
	existingUser, err := GetUserByEmail(user.Email, database)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Email or password invalid."}) // Do not return what components of the credential was invalid to avoid leaking information.
			return
		}
		logger.Error().Msgf("Error getting user by email: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	// Check if the user is active
	if !existingUser.Active {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not active."})
		return
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password invalid."}) // Do not return what components of the credential was invalid to avoid leaking information.
		return
	}

	// Generate a JWT token
	token, err := NewUserToken(user.Email, database, logger)
	if err != nil {
		logger.Error().Msgf("Error generating token: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	// Set the Authorization header with the token
	c.Header("Authorization", "Bearer "+token) // Set the Authorization header with the token

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"message": "User login successful."})
}

func GetUserAccessLevelsString(accessLevels []models.AccessLevel) string {
	accessLevelsString := ""
	for _, accessLevel := range accessLevels {
		accessLevelsString += accessLevel.Name + ","
	}
	if len(accessLevelsString) > 0 {
		accessLevelsString = accessLevelsString[:len(accessLevelsString)-1] // Remove the last comma
	}
	return accessLevelsString
}

// Logout function to remove the Authorization header
// This is a placeholder function. In a real application, you would handle logout differently.
// For example, you might want to invalidate the token on the server side or remove it from the client.
// In this case, we are just removing the Authorization header from the response.
// In a real application, you would also want to handle token revocation.
func Logout(c *gin.Context) {
	c.Header("Authorization", "") // Remove the Authorization header
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully."})
}

func GetUserByEmail(email string, database *gorm.DB) (*models.User, error) {
	var user models.User
	if err := database.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserAccessLevels(email string, database *gorm.DB) ([]models.AccessLevel, error) {
	var user models.User
	if err := database.Preload("AccessLevels").Find(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return user.AccessLevels, nil
}

func NewUserToken(email string, database *gorm.DB, logger *zerolog.Logger) (string, error) {
	// Get user by email
	user, err := GetUserByEmail(email, database)
	if err != nil {
		logger.Error().Msgf("Error getting user by email: %s", err)
		return "", err
	}

	// Get user access levels
	accessLevels, err := GetUserAccessLevels(user.Email, database)
	if err != nil {
		logger.Error().Msgf("Error getting user access levels: %s", err)
		return "", err
	}

	// Set up the claims for the token
	claims := token.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWTISSUER"),
			Subject:   user.Email,
			Audience:  []string{os.Getenv("JWTAUDIENCE")},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
		AccessLevels: GetUserAccessLevelsString(accessLevels),
	}

	// Generate the token with claims
	token, err := token.GenerateTokenWithClaims(claims)
	if err != nil {
		logger.Error().Msgf("Error generating token: %s", err)
		return "", err
	}

	return token, nil
}
