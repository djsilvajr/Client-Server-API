package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"my-app/internal/requests"
	"my-app/internal/service"
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


func CotacaoHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON
	var payload requests.CotacaoRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Valida se brl_amount é maior que 0
	if payload.BRLAmount <= 0 {
		http.Error(w, "O valor de 'brl_amount' deve ser maior que 0", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	select {
	case <-time.After(10 * time.Millisecond):

		retorno := service.BrlToUsd(payload) 
		log.Println(retorno)
		return
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Go with Context!"})
	case <-ctx.Done():
		log.Println("Contexto cancelado:", ctx.Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
	}
}
