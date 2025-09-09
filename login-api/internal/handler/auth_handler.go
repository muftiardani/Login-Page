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

type AuthHandler struct {
	Store  storage.UserStore
	JwtKey []byte
}

func NewAuthHandler(store storage.UserStore, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		Store:  store,
		JwtKey: jwtKey,
	}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Permintaan tidak valid."}`, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"message":"Terjadi kesalahan internal."}`, http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Email:        creds.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := h.Store.CreateUser(newUser); err != nil {
		log.Printf("Gagal mendaftarkan pengguna: %v\n", err)
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.Response{Message: "Email ini sudah terdaftar, silakan gunakan email lain.", Success: false})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{Message: "Pendaftaran berhasil!", Success: true})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Permintaan tidak valid."}`, http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(creds.Email)
	if !ok || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{Message: "Email atau kata sandi salah.", Success: false})
		return
	}

	tokenString, err := auth.GenerateJWT(creds.Email, h.JwtKey)
	if err != nil {
		http.Error(w, `{"message":"Gagal membuat token."}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Message: "Login berhasil!",
		Token:   tokenString,
		Success: true,
	})
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Message: "Selamat datang di area terproteksi!", Success: true})
}