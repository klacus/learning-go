package middlewares

import (
	"jwt/controllers"
	"jwt/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func RequireAuthentication(c *gin.Context, database *gorm.DB, logger *zerolog.Logger) {
	authorizationHeader := c.GetHeader("authorization")
	if authorizationHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := token.GetBearerToken(authorizationHeader)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get bearer token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtToken, err := token.ValidateToken(tokenString)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	email, err := token.GetUserEmailFromToken(jwtToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get user email from token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// find user
	user, err := controllers.GetUserByEmail(email, database)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to find user by email")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user != nil {
		c.Set("user", user)
		c.Next()
		return
	}

	// The default should be to return a forbidden status to avoid a bug granting access.
	logger.Error().Msg("unexpected error: user not found")
	c.AbortWithStatus(http.StatusInternalServerError)
}
