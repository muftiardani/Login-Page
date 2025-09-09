package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Response struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Success bool   `json:"success"`
}