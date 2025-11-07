package handlers

import (
	"context"
	"my-app/internal/response"
	"my-app/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ContextMiddleware adiciona contexto com timeout
// func ContextMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
// 		defer cancel()
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// ContextMiddleware adiciona contexto com timeout
func ContextMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Missing Authorization header", response.Empty{})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Invalid Authorization header format", response.Empty{})
			return
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return service.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Invalid token", response.Empty{})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Invalid token claims", response.Empty{})
			return
		}
		email, ok := claims["email"].(string)
		if !ok || email == "" {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Email not found in token", response.Empty{})
			return
		}

		// Use exported validation function
		if !service.ValidateServerToken(email, tokenString) {
			response.WriteJSON(w, http.StatusUnauthorized, false, "Token does not match stored token", response.Empty{})
			return
		}

		next.ServeHTTP(w, r)
	})
}
