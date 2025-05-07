package middlewares

import (
	"jwt/controllers"
	"jwt/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequireAuthentication(c *gin.Context, database *gorm.DB) {
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

	email, err := token.GetUserEmailFromToken(jwtToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// find user
	user, err := controllers.GetUserByEmail(email, database)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Next()
}
