package auth

import (
	"login-api/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTokens membuat access token dan refresh token.
func GenerateTokens(email string, jwtKey []byte) (string, string, error) {
	accessTokenExpirationTime := time.Now().Add(15 * time.Minute)
	accessClaims := &model.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpirationTime),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenExpirationTime := time.Now().Add(24 * 7 * time.Hour)
	refreshClaims := &model.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpirationTime),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}