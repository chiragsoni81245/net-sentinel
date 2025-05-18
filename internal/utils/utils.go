package utils

import (
	"fmt"
	"net/http"
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
    return _token.SignedString([]byte(config.Server.Secret))
}

func ValidateJWT(tokenString string, config *config.Config) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Ensure token uses expected signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.Server.Secret), nil
	})

    if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}

	userId, ok := claims["userId"]
	if !ok {
		return 0, fmt.Errorf("userId not found in token")
	}

    return int(userId.(float64)), nil
}

func SendJSON(w http.ResponseWriter, jsonData string, status int) {
    w.Header().Add("Content-Type", "application/json")
    http.Error(w, jsonData, status)
}

