package middleware

import (
	"log"
	"login-api/internal/model"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// NewJwtMiddleware membuat lapisan pelindung (middleware) untuk memeriksa token JWT.
func NewJwtMiddleware(jwtKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Printf("PERINGATAN: Permintaan ke '%s' dari IP %s ditolak karena tidak ada token.", r.URL.Path, r.RemoteAddr)
				http.Error(w, `{"message":"Token otentikasi tidak ditemukan. Harap sertakan di header Authorization."}`, http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				log.Printf("PERINGATAN: Format token salah untuk permintaan ke '%s' dari IP %s.", r.URL.Path, r.RemoteAddr)
				http.Error(w, `{"message":"Format token tidak valid. Pastikan menggunakan format 'Bearer <token>'."}`, http.StatusUnauthorized)
				return
			}

			claims := &model.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				log.Printf("PERINGATAN: Token tidak valid atau kedaluwarsa digunakan untuk akses ke '%s' dari IP %s. Error: %v", r.URL.Path, r.RemoteAddr, err)
				http.Error(w, `{"message":"Token tidak valid atau telah kedaluwarsa. Silakan login kembali."}`, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}