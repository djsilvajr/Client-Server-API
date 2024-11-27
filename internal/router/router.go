package router

import (
	"net/http"
	"time"

	"my-app/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// Setup configura as rotas do servidor
func Setup() http.Handler {
	r := chi.NewRouter()

	// Middleware de contexto
	//r.Use(handlers.ContextMiddleware)

	// Rota principal
	//r.Get("/", handlers.TestHandler)
	r.With(ContextMiddleware(2 * time.Second)).Get("/", handlers.TestHandler)

	return r
}