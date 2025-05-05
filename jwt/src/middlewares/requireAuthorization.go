package middleware

import (
	"jwt/auth"
	"jwt/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func RequireAuthentication(c *gin.Context, database *gorm.DB) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := auth.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// find user
	user, err := controllers.GetUserByEmail(token.Claims.(jwt.MapClaims)["sub"].(string), database)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Next()
}
