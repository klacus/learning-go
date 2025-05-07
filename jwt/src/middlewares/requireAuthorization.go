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

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Retrieve the authentication cookie
	// 	cookie, err := r.Cookie("auth_token")
	// 	if err != nil {
	// 		http.Error(w, "authentication cookie not found", http.StatusUnauthorized)
	// 		return
	// 	}

	// 	// Validate the token
	// 	token, err := middleware.ValidateToken(cookie.Value)
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
