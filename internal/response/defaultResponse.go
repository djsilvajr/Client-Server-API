package response

import (
	"encoding/json"
	"net/http"
)

type Empty struct{}

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// Exemplo de WriteJSON genérico:
func WriteJSON[T any](w http.ResponseWriter, code int, status bool, msg string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response[T]{Status: status, Message: msg, Data: data})
}
