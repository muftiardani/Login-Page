package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	JWTSecretKey  string
	DatabaseURL   string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("PERINGATAN: Tidak dapat memuat file .env")
	}

	jwtKey := getEnvOrPanic("JWT_SECRET_KEY")
	dbURL := getEnvOrPanic("DATABASE_URL")

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecretKey:  jwtKey,
		DatabaseURL:   dbURL,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("FATAL: Environment variable %s tidak diatur.", key)
	}
	return value
}