package infrastructure

import (
	"time"

	"github.com/OGARCIIA/examen-backend/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(), // 5 min de vida
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}
