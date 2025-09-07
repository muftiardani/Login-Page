package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	JWTSecretKey  string
	DatabaseURL   string
}

func New() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecretKey:  getEnv("JWT_SECRET_KEY", "kunci_rahasia_super_aman_saya"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://postgres:arda123@localhost:5432/LoginDB"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}