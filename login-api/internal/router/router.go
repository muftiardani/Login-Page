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

	r.HandleFunc("/api/register", authHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", authHandler.LoginHandler).Methods("POST")

	protectedRoutes := r.PathPrefix("/api").Subrouter()
	
	jwtAuthMiddleware := middleware.NewJwtMiddleware(authHandler.JwtKey)
	protectedRoutes.Use(jwtAuthMiddleware)

	protectedRoutes.HandleFunc("/status", handler.StatusHandler).Methods("GET")
	protectedRoutes.HandleFunc("/user/password", authHandler.ChangePasswordHandler).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	return handler
}