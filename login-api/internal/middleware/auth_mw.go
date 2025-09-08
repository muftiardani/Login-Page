package middleware

import (
	"log"
	"login-api/internal/model"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwtMiddleware(jwtKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"message":"Token otentikasi tidak ditemukan. Harap sertakan di header Authorization."}`, http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, `{"message":"Format token tidak valid. Pastikan menggunakan format 'Bearer <token>'."}`, http.StatusUnauthorized)
				return
			}

			claims := &model.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, `{"message":"Token tidak valid atau telah kedaluwarsa. Silakan login kembali."}`, http.StatusUnauthorized)
				return
			}

			log.Printf("Akses terotentikasi ke %s dari pengguna: %s", r.URL.Path, claims.Username)

			next.ServeHTTP(w, r)
		})
	}
}