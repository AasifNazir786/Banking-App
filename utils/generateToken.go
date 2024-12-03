package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(userName string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userName,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
