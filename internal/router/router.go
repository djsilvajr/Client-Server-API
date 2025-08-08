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

	//Rotas de pratica com string
	r.With(handlers.ContextMiddleware(10*time.Second)).Get("/strings/contar", handlers.StringCount)

	return r
}
