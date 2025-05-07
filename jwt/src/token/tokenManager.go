package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	AccessLevels string `json:"accessLevels"`
	jwt.RegisteredClaims
}

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

func GenerateTokenWithClaims(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWTSECRET")))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTSECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	// Validate the token claims
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check if the token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["exp"] != nil {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, fmt.Errorf("token expired")
			}
		}
	} else {
		return nil, fmt.Errorf("token expired")
	}

	return token, nil
}

func GetBearerToken(authorizationHeader string) (string, error) {
	if len(authorizationHeader) <= 0 {
		return "", fmt.Errorf("authorization header is empty")
	}
	bearerToken := (strings.Split(authorizationHeader, "Bearer "))[1]
	if len(bearerToken) <= 0 {
		return "", fmt.Errorf("bearer token is empty")
	}
	return bearerToken, nil
}

func GetUserEmailFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims from token")
	}
	email, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract email from token claims")
	}
	return email, nil
}

func GetAccessLevelsFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims from token")
	}
	accessLevels, ok := claims["accessLevels"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract access levels from token claims")
	}
	return accessLevels, nil
}
