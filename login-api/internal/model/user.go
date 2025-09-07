package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Response struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Success bool   `json:"success"`
}