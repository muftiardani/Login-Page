package middleware

import (
	"log"
	"login-api/internal/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// NewJwtMiddleware membuat lapisan pelindung untuk memeriksa token JWT dari cookie.
func NewJwtMiddleware(jwtKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("access_token")
			if err != nil {
				if err == http.ErrNoCookie {
					log.Printf("PERINGATAN: Permintaan ke '%s' dari IP %s ditolak karena tidak ada token.", r.URL.Path, r.RemoteAddr)
					http.Error(w, `{"message":"Token otentikasi tidak ditemukan."}`, http.StatusUnauthorized)
					return
				}
				http.Error(w, `{"message":"Permintaan tidak valid."}`, http.StatusBadRequest)
				return
			}

			tokenString := c.Value
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