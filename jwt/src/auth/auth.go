package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userEmail,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "Failed to create token.", err
	}
	return tokenString, nil
}
