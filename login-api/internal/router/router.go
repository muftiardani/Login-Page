package router

import (
	"login-api/internal/handler"
	"login-api/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter(authHandler *handler.AuthHandler) http.Handler {
	r := mux.NewRouter()

	// Terapkan rate limiter hanya pada endpoint login
	loginHandler := middleware.RateLimiterMiddleware(http.HandlerFunc(authHandler.LoginHandler))
	r.Handle("/api/login", loginHandler).Methods("POST")

	r.HandleFunc("/api/register", authHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/refresh", authHandler.RefreshTokenHandler).Methods("POST")
	r.HandleFunc("/api/logout", authHandler.LogoutHandler).Methods("POST")

	protectedRoutes := r.PathPrefix("/api").Subrouter()
	
	jwtAuthMiddleware := middleware.NewJwtMiddleware(authHandler.JwtKey)
	protectedRoutes.Use(jwtAuthMiddleware)

	protectedRoutes.HandleFunc("/status", handler.StatusHandler).Methods("GET")
	protectedRoutes.HandleFunc("/user/password", authHandler.ChangePasswordHandler).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	return handler
}