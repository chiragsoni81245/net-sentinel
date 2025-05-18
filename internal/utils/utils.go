package utils

import (
	"fmt"
	"time"

	"github.com/chiragsoni81245/net-sentinel/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId int, config *config.Config) (string, error) {
    expirationTime := time.Duration(config.Server.TokenExpiration) * time.Hour

    claims := jwt.MapClaims{
        "userId": userId,
		"exp":      time.Now().Add(expirationTime).Unix(),
		"iat":      time.Now().Unix(),
	}

    _token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return _token.SignedString(config.Server.Secret)
}

func ValidateJWT(tokenString string, config *config.Config) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Ensure token uses expected signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return config.Server.Secret, nil
	})

    if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("userId not found in token")
	}

    return userId, nil
}
