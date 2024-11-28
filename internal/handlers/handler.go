package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
)

// HelloHandler responde com uma mensagem JSON
func TestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):

		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Go with Context!"})
	case <-ctx.Done():
		log.Println("Contexto cancelado:", ctx.Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
	}
}

