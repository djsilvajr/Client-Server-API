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
	// Rota principal
	r.With(handlers.ContextMiddleware(10*time.Second)).Get("/", handlers.TestHandler)

	//Rota de cotação
	r.With(handlers.ContextMiddleware(200*time.Millisecond)).Post("/cotacao", handlers.CotacaoHandler)

	return r
}
