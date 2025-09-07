package handler

import (
	"encoding/json"
	"log"
	"login-api/internal/auth"
	"login-api/internal/model"
	"login-api/internal/storage"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// AuthHandler menampung dependensi untuk handler otentikasi.
type AuthHandler struct {
	Store  storage.UserStore
	JwtKey []byte
}

// NewAuthHandler adalah constructor untuk AuthHandler.
func NewAuthHandler(store storage.UserStore, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		Store:  store,
		JwtKey: jwtKey,
	}
}

// RegisterHandler menangani registrasi pengguna baru.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Bad request"}`, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"message":"Server error"}`, http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Username:     creds.Username,
		PasswordHash: string(hashedPassword),
	}

	// Menggunakan Store Interface untuk membuat user
	if err := h.Store.CreateUser(newUser); err != nil {
		log.Printf("Failed to register user: %v\n", err)

		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.Response{Message: "Username already exists", Success: false})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{Message: "Registration successful!", Success: true})
}

// LoginHandler menangani proses login pengguna.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Bad request"}`, http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(creds.Username)
	if !ok || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{Message: "Invalid credentials", Success: false})
		return
	}

	tokenString, err := auth.GenerateJWT(creds.Username, h.JwtKey)
	if err != nil {
		http.Error(w, `{"message":"Could not generate token"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Message: "Login successful!",
		Token:   tokenString,
		Success: true,
	})
}

// StatusHandler adalah handler untuk rute yang dilindungi.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Message: "Welcome to the protected area!", Success: true})
}