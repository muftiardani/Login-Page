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

// RegisterHandler menangani permintaan pendaftaran pengguna baru.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Menerima permintaan pendaftaran dari alamat IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	// Mengamankan password dengan hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("KRITIS: Gagal melakukan hash password untuk email %s: %v", creds.Email, err)
		http.Error(w, `{"message":"Terjadi kesalahan internal pada server."}`, http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Email:        creds.Email,
		PasswordHash: string(hashedPassword),
	}

	// Mencoba membuat pengguna baru di database
	if err := h.Store.CreateUser(newUser); err != nil {
		log.Printf("PERINGATAN: Gagal mendaftarkan pengguna baru: %v", err)
		w.WriteHeader(http.StatusConflict) // Status 409: Terjadi konflik
		json.NewEncoder(w).Encode(model.Response{Message: "Email ini sudah terdaftar. Silakan gunakan email lain.", Success: false})
		return
	}

	log.Printf("INFO: Pengguna dengan email %s berhasil terdaftar.", creds.Email)
	w.WriteHeader(http.StatusCreated) // Status 201: Berhasil dibuat
	json.NewEncoder(w).Encode(model.Response{Message: "Pendaftaran berhasil! Silakan masuk.", Success: true})
}

// LoginHandler menangani permintaan login dari pengguna.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Menerima permintaan login dari alamat IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	// Memeriksa kredensial pengguna
	user, ok := h.Store.GetUser(creds.Email)
	if !ok || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		log.Printf("PERINGATAN: Upaya login gagal untuk email '%s' dari IP: %s", creds.Email, r.RemoteAddr)
		w.WriteHeader(http.StatusUnauthorized) // Status 401: Tidak terotorisasi
		json.NewEncoder(w).Encode(model.Response{Message: "Email atau kata sandi yang Anda masukkan salah.", Success: false})
		return
	}

	// Jika berhasil, buatkan token JWT
	tokenString, err := auth.GenerateJWT(creds.Email, h.JwtKey)
	if err != nil {
		log.Printf("KRITIS: Gagal membuat token JWT untuk %s: %v", creds.Email, err)
		http.Error(w, `{"message":"Gagal membuat token autentikasi."}`, http.StatusInternalServerError)
		return
	}

	log.Printf("INFO: Pengguna %s berhasil masuk.", creds.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Message: "Login berhasil!",
		Token:   tokenString,
		Success: true,
	})
}

// StatusHandler adalah contoh halaman yang dilindungi.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Halaman terproteksi diakses dari IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Message: "Selamat datang di area terproteksi!", Success: true})
}