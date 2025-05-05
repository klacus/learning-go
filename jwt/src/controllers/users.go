package controllers

import (
	"jwt/auth"
	"jwt/models"
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

	// Get access levels
	accessLevels := []models.AccessLevel{}
	if err := database.Model(&existingUser).Association("AccessLevels").Find(&accessLevels); err != nil {
		logger.Error().Msgf("Error getting access levels: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	// ADD ACCESS LEVELS TO THE JWT TOKEN
	// Generate a JWT token with access levels
	accessLevelsString := ""
	for _, accessLevel := range accessLevels {
		accessLevelsString += accessLevel.Name + ","
	}
	if len(accessLevelsString) > 0 {
		accessLevelsString = accessLevelsString[:len(accessLevelsString)-1] // Remove the last comma
	}
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  os.Getenv("JWTISSUER"),
			Subject: user.Email,
			// Audience:  []string{"your-app-audience"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
		AccessLevels: accessLevelsString,
	}
	token, err := auth.GenerateTokenWithClaims(claims)
	if err != nil {
		logger.Error().Msgf("Error generating token: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	// // Generate a JWT token
	// token, err := auth.GenerateToken(user.Email)
	// if err != nil {
	// 	logger.Error().Msgf("Error generating token: %s", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
	// 	return
	// }

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true) // Set the cookie with the token
	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"message": "User login successful."})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true) // Clear the cookie by setting MaxAge to negative value, works with non synchronized clocks also
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully."})
}

func GetUserByEmail(email string, database *gorm.DB) (*models.User, error) {
	var user models.User
	if err := database.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
