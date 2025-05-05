package middleware

import (
	"net/http"
	"strings"

	"jwt/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuthorization(c *gin.Context, requiredAccessLevel string) {
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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["accessLevels"] != nil {
		if accessLevels, ok := claims["accessLevels"].(string); ok {
			// Split the access levels into a slice
			levels := strings.Split(accessLevels, ",")
			// Check if the required access level is in the slice
			for _, level := range levels {
				if strings.TrimSpace(level) == requiredAccessLevel {
					c.Next()
				}
			}
		}
	}
	c.AbortWithStatus(http.StatusForbidden)

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Retrieve the authentication cookie
	// 	cookie, err := r.Cookie("auth_token")
	// 	if err != nil {
	// 		http.Error(w, "authentication cookie not found", http.StatusUnauthorized)
	// 		return
	// 	}

	// 	// Validate the token
	// 	token, err := auth.ValidateToken(cookie.Value)
	// 	if err != nil {
	// 		http.Error(w, "invalid token", http.StatusUnauthorized)
	// 		return
	// 	}

	// 	// Check if the token has the "access_levels" claim
	// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["access_levels"] != nil {
	// 		if accessLevels, ok := claims["access_levels"].(string); ok {
	// 			// Split the access levels into a slice
	// 			levels := strings.Split(accessLevels, ",")
	// 			// Check if the required access level is in the slice
	// 			for _, level := range levels {
	// 				if strings.TrimSpace(level) == requiredAccessLevel {
	// 					// Call the next handler if the access level is found
	// 					next.ServeHTTP(w, r)
	// 					return
	// 				}
	// 			}
	// 		}
	// 	}

	// 	// If the required access level is not found, return an error
	// 	http.Error(w, fmt.Sprintf("access denied: required access level '%s' not found", requiredAccessLevel), http.StatusForbidden)
	// })
}
