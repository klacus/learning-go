package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	AccessLevels string `json:"accessLevels"`
}

func GenerateTokenFromEmail(userEmail string) (string, error) {
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

	// Check if the token is signed with the correct algorithm
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	// Check if the token is not before the current time
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["nbf"] != nil {
		if nbf, ok := claims["nbf"].(float64); ok {
			if time.Unix(int64(nbf), 0).After(time.Now()) {
				return nil, fmt.Errorf("token not valid yet")
			}
		}
	} else {
		return nil, fmt.Errorf("token not valid yet")
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

	// Check if the token is for the correct audience
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["aud"] != nil {
		if aud, ok := claims["aud"].(string); ok {
			if aud != os.Getenv("JWTAUDIENCE") {
				return nil, fmt.Errorf("token audience mismatch")
			}
		}
	} else {
		return nil, fmt.Errorf("token audience mismatch")
	}

	// Check if the token is for the correct issuer
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["iss"] != nil {
		if iss, ok := claims["iss"].(string); ok {
			if iss != os.Getenv("JWTISSUER") {
				return nil, fmt.Errorf("token issuer mismatch")
			}
		}
	} else {
		return nil, fmt.Errorf("token issuer mismatch")
	}

	return token, nil
}

func GetBearerToken(authorizationHeader string) (string, error) {
	if len(authorizationHeader) <= 0 {
		return "", fmt.Errorf("authorization header is empty")
	}
	headerArray := strings.Split(authorizationHeader, " ")
	if len(headerArray) != 2 {
		return "", fmt.Errorf("invalid authorization header format")
	}
	if headerArray[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
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
