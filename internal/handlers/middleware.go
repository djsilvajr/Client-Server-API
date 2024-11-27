package handlers

import (
	"context"
	"net/http"
	"time"
)

// ContextMiddleware adiciona contexto com timeout
// func ContextMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
// 		defer cancel()
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

//ContextMiddleware adiciona contexto com timeout
func ContextMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}