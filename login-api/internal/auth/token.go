package auth

import (
	"login-api/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, jwtKey []byte) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}