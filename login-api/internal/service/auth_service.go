package service

import (
	"errors"
	"login-api/internal/auth"
	"login-api/internal/model"
	"login-api/internal/storage"
	"login-api/internal/validator"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ErrEmailExists adalah error kustom saat email sudah terdaftar.
var ErrEmailExists = errors.New("email ini sudah terdaftar")

// AuthService menyediakan logika bisnis terkait autentikasi.
type AuthService struct {
	UserStore storage.UserStore
	JwtKey    []byte
}

// NewAuthService membuat instance AuthService baru.
func NewAuthService(store storage.UserStore, jwtKey []byte) *AuthService {
	return &AuthService{
		UserStore: store,
		JwtKey:    jwtKey,
	}
}

// RegisterUser memvalidasi dan mendaftarkan pengguna baru.
func (s *AuthService) RegisterUser(creds model.Credentials) error {
	creds.Email = strings.TrimSpace(creds.Email)
	if _, err := mail.ParseAddress(creds.Email); err != nil {
		return errors.New("format email tidak valid")
	}
	if err := validator.ValidatePassword(creds.Password); err != nil {
		return err
	}

	// Cek apakah pengguna sudah ada
	if _, exists := s.UserStore.GetUser(creds.Email); exists {
		return ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := model.User{
		Email:        creds.Email,
		PasswordHash: string(hashedPassword),
	}

	return s.UserStore.CreateUser(newUser)
}

// LoginUser memverifikasi kredensial dan menghasilkan token.
func (s *AuthService) LoginUser(creds model.Credentials) (string, string, error) {
	user, ok := s.UserStore.GetUser(creds.Email)
	if !ok || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		return "", "", validator.ErrInvalidCredentials
	}

	return auth.GenerateTokens(creds.Email, s.JwtKey)
}