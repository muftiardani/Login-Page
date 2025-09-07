package main

import (
	"log"
	"login-api/internal/config"
	"login-api/internal/handler"
	"login-api/internal/router"
	"login-api/internal/storage/memory"
	"net/http"
)

func main() {
	cfg := config.New()

	jwtKey := []byte(cfg.JWTSecretKey)
	addr := cfg.ServerAddress

	// Inisialisasi
	userStore := memory.NewMemoryUserStore()
	authHandler := handler.NewAuthHandler(userStore, jwtKey) // Perbarui constructor jika perlu
	r := router.NewRouter(authHandler)

	// Jalankan Server
	log.Printf("Go API Server with JWT running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}