package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if tokenString == "" {
			http.Error(w, "Forbidden: Missing token", http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err)
			http.Error(w, "Forbidden: Invalid token", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "user", token)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
