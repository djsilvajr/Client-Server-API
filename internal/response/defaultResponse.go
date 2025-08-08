package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func WriteJSON[T any](w http.ResponseWriter, code int, status bool, message string, data *T) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Response[T]{Status: status, Message: message, Data: data})
}
