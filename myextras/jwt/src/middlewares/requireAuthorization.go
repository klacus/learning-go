package middlewares

import (
	"jwt/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RequireAuthorization(c *gin.Context, requiredAccessLevel string, logger *zerolog.Logger) {
	authorizationHeader := c.GetHeader("authorization")
	if authorizationHeader == "" {
		logger.Error().Msg("Authorization header is missing")
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

	// Check if the token has the "accessLevels" claim
	accessLevels, err := token.GetAccessLevelsFromToken(jwtToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get access levels from token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Split the access levels into a slice
	levels := strings.Split(accessLevels, ",")
	// Check if the required access level is in the slice
	for _, level := range levels {
		if strings.TrimSpace(level) == requiredAccessLevel {
			c.Next()
			return
		}
	}

	// The default should be to return a forbidden status to avoid a bug granting access.
	c.AbortWithStatus(http.StatusForbidden)
}
