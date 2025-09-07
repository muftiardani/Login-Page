package middleware

import (
	"login-api/internal/model"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// NewJwtMiddleware membuat middleware baru yang menggunakan jwtKey yang diberikan untuk validasi.
func NewJwtMiddleware(jwtKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"message":"Missing token"}`, http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, `{"message":"Invalid token format"}`, http.StatusUnauthorized)
				return
			}

			claims := &model.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, `{"message":"Invalid token"}`, http.StatusUnauthorized)
				return
			}

			// Jika token valid, lanjutkan ke handler berikutnya
			next.ServeHTTP(w, r)
		})
	}
}