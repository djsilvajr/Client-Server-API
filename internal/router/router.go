package router

import (
	"my-app/internal/handlers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Setup configura as rotas do servidor
func Setup() http.Handler {
	r := chi.NewRouter()
	// Rota principal
	r.With(handlers.ContextMiddleware(10*time.Second)).Get("/", handlers.TestHandler)

	//Users
	r.With(handlers.ContextMiddleware(10*time.Second)).Post("/user/login", handlers.LoginUser)
	r.With(
		handlers.ContextMiddleware(2*time.Second),
		handlers.TokenMiddleware,
	).Get("/user/token/validation", handlers.GetValidationMessage)

	return r
}
