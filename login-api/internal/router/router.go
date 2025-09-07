package router

import (
	"login-api/internal/handler"
	"login-api/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// NewRouter membuat dan mengkonfigurasi router aplikasi.
// Ia menerima AuthHandler sebagai dependensi untuk menghubungkan rute ke logika bisnisnya.
func NewRouter(authHandler *handler.AuthHandler) http.Handler {
	// Membuat instance mux router baru
	r := mux.NewRouter()

	// Menetapkan rute publik untuk registrasi dan login
	r.HandleFunc("/api/register", authHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", authHandler.LoginHandler).Methods("POST")

	// Membuat subrouter untuk rute-rute yang memerlukan otentikasi
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	
	// Menerapkan middleware JWT ke subrouter ini
	// Middleware ini dibuat menggunakan NewJwtMiddleware untuk menyuntikkan jwtKey
	jwtAuthMiddleware := middleware.NewJwtMiddleware(authHandler.JwtKey)
	protectedRoutes.Use(jwtAuthMiddleware)

	// Menetapkan rute yang dilindungi
	protectedRoutes.HandleFunc("/status", handler.StatusHandler).Methods("GET")

	// Konfigurasi CORS (Cross-Origin Resource Sharing)
	// Ini memungkinkan front-end di localhost:5173 untuk berkomunikasi dengan API ini
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Membungkus router utama dengan handler CORS
	handler := c.Handler(r)

	return handler
}