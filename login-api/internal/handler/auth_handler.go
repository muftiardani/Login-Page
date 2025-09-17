package handler

import (
	"encoding/json"
	"errors"
	"log"
	"login-api/internal/model"
	"login-api/internal/service"
	"login-api/internal/validator"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler menangani permintaan HTTP terkait autentikasi.
type AuthHandler struct {
	AuthSvc *service.AuthService
	JwtKey  []byte
}

// NewAuthHandler membuat instance AuthHandler baru dengan AuthSvc yang disuntikkan.
func NewAuthHandler(authSvc *service.AuthService, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		AuthSvc: authSvc,
		JwtKey:  jwtKey,
	}
}

// RegisterHandler menangani permintaan pendaftaran pengguna baru.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	err := h.AuthSvc.RegisterUser(creds)
	if err != nil {
		// Menentukan kode status berdasarkan jenis error dari service
		if errors.Is(err, service.ErrEmailExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(model.Response{Message: err.Error(), Success: false})
			return
		}
		// Untuk error validasi lainnya
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{Message: err.Error(), Success: false})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{Message: "Pendaftaran berhasil! Silakan masuk.", Success: true})
}

// LoginHandler menangani permintaan login dan mengirimkan token melalui HttpOnly cookie.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"message":"Format permintaan tidak sesuai."}`, http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.AuthSvc.LoginUser(creds)
	if err != nil {
		if errors.Is(err, validator.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.Response{Message: err.Error(), Success: false})
			return
		}
		log.Printf("KRITIS: Gagal membuat token JWT: %v", err)
		http.Error(w, `{"message":"Terjadi kesalahan internal pada server."}`, http.StatusInternalServerError)
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

	user, ok := h.AuthSvc.UserStore.GetUser(req.Email) // Akses UserStore melalui AuthSvc
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

	if err := h.AuthSvc.UserStore.UpdateUser(req.Email, user); err != nil {
		log.Printf("ERROR: Gagal memperbarui password untuk email %s: %v", req.Email, err)
		http.Error(w, `{"message":"Gagal memperbarui kata sandi."}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Message: "Kata sandi berhasil diubah.", Success: true})
}

// StatusHandler adalah contoh halaman yang dilindungi.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Message: "Selamat datang di area terproteksi!", Success: true})
}