package handler

import (
	"encoding/json"
	"log"
	"login-api/internal/auth"
	"login-api/internal/model"
	"login-api/internal/storage"
	"login-api/internal/validator"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	creds.Email = strings.TrimSpace(creds.Email)

	if err := validator.ValidatePassword(creds.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{Message: err.Error(), Success: false})
		return
	}

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

	if err := h.Store.CreateUser(newUser); err != nil {
		log.Printf("PERINGATAN: Gagal mendaftarkan pengguna baru: %v", err)
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.Response{Message: "Email ini sudah terdaftar. Silakan gunakan email lain.", Success: false})
		return
	}

	log.Printf("INFO: Pengguna dengan email %s berhasil terdaftar.", creds.Email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{Message: "Pendaftaran berhasil! Silakan masuk.", Success: true})
}

// LoginHandler menangani permintaan login dan mengirimkan token melalui HttpOnly cookie.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Menerima permintaan login dari alamat IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(creds.Email)
	if !ok || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		log.Printf("PERINGATAN: Upaya login gagal untuk email '%s' dari IP: %s", creds.Email, r.RemoteAddr)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{Message: "Email atau kata sandi yang Anda masukkan salah.", Success: false})
		return
	}

	accessToken, refreshToken, err := auth.GenerateTokens(creds.Email, h.JwtKey)
	if err != nil {
		log.Printf("KRITIS: Gagal membuat token JWT untuk %s: %v", creds.Email, err)
		http.Error(w, `{"message":"Gagal membuat token autentikasi."}`, http.StatusInternalServerError)
		return
	}

	// Set HttpOnly cookie untuk access token
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})

	// Set HttpOnly cookie untuk refresh token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/api/refresh",
	})

	log.Printf("INFO: Pengguna %s berhasil masuk.", creds.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Message: "Login berhasil!",
		Success: true,
	})
}

// RefreshTokenHandler memvalidasi refresh token dan memberikan access token baru.
func (h *AuthHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, `{"message":"Refresh token tidak ditemukan."}`, http.StatusUnauthorized)
			return
		}
		http.Error(w, `{"message":"Permintaan tidak valid."}`, http.StatusBadRequest)
		return
	}

	tokenStr := c.Value
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return h.JwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, `{"message":"Refresh token tidak valid atau kedaluwarsa."}`, http.StatusUnauthorized)
		return
	}

	// Buat access token baru
	expirationTime := time.Now().Add(15 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(h.JwtKey)
	if err != nil {
		http.Error(w, `{"message":"Gagal membuat token baru."}`, http.StatusInternalServerError)
		return
	}

	// Set cookie access token yang baru
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessTokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Message: "Token berhasil diperbarui.", Success: true})
}

// LogoutHandler menghapus cookie otentikasi.
func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/api/refresh",
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Message: "Logout berhasil.", Success: true})
}

// ChangePasswordHandler menangani permintaan perubahan password.
func (h *AuthHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Menerima permintaan ubah password dari alamat IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email       string `json:"email"`
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(req.Email)
	if !ok {
		http.Error(w, `{"message":"Pengguna tidak ditemukan."}`, http.StatusNotFound)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{Message: "Kata sandi lama salah.", Success: false})
		return
	}
	
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.NewPassword)) == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{Message: "Kata sandi baru tidak boleh sama dengan kata sandi lama.", Success: false})
		return
	}

	if err := validator.ValidatePassword(req.NewPassword); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{Message: err.Error(), Success: false})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("KRITIS: Gagal melakukan hash password baru untuk email %s: %v", req.Email, err)
		http.Error(w, `{"message":"Terjadi kesalahan internal pada server."}`, http.StatusInternalServerError)
		return
	}
	
	user.PasswordHash = string(hashedPassword)

	if err := h.Store.UpdateUser(req.Email, user); err != nil {
		log.Printf("ERROR: Gagal memperbarui password untuk email %s: %v", req.Email, err)
		http.Error(w, `{"message":"Gagal memperbarui kata sandi."}`, http.StatusInternalServerError)
		return
	}

	log.Printf("INFO: Pengguna %s berhasil mengubah kata sandi.", req.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Message: "Kata sandi berhasil diubah.", Success: true})
}

// StatusHandler adalah contoh halaman yang dilindungi.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: Halaman terproteksi diakses dari IP: %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Message: "Selamat datang di area terproteksi!", Success: true})
}