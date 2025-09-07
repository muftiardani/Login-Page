package main

import (
	"context"
	"fmt"
	"log"
	"login-api/internal/config"
	"login-api/internal/handler"
	"login-api/internal/router"
	"login-api/internal/storage/postgres"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.New()
	
	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Database connected successfully!")

	userStore := postgres.NewPostgresUserStore(dbpool)
	
	jwtKey := []byte(cfg.JWTSecretKey)
	addr := cfg.ServerAddress
	authHandler := handler.NewAuthHandler(userStore, jwtKey)
	r := router.NewRouter(authHandler)

	log.Printf("Go API Server with JWT running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}