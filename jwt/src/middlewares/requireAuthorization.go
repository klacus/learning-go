package middlewares

import (
	"jwt/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuthorization(c *gin.Context, requiredAccessLevel string) {
	authorizationHeader := c.GetHeader("authorization")
	if authorizationHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := token.GetBearerToken(authorizationHeader)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtToken, err := token.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if the token has the "accessLevels" claim
	accessLevels, err := token.GetAccessLevelsFromToken(jwtToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Split the access levels into a slice
	levels := strings.Split(accessLevels, ",")
	// Check if the required access level is in the slice
	for _, level := range levels {
		if strings.TrimSpace(level) == requiredAccessLevel {
			c.Next()
		}
	}
	c.AbortWithStatus(http.StatusForbidden)

}
