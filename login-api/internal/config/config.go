package config

import (
	"os"
)

// Config menampung semua konfigurasi untuk aplikasi.
// Nilai-nilai ini dibaca dari environment variables.
type Config struct {
	ServerAddress string
	JWTSecretKey  string
}

// New memuat konfigurasi dari environment variables.
func New() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecretKey:  getEnv("JWT_SECRET_KEY", "kunci_rahasia_super_aman_saya"),
	}
}

// getEnv adalah fungsi helper untuk membaca environment variable
// atau mengembalikan nilai default jika tidak ditemukan.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}