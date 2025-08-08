package handlers

import (
	"encoding/json"
	"log"
	"my-app/internal/requests"
	"my-app/internal/response"
	"my-app/internal/service"
	"net/http"
	"time"
)

func StringCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload requests.StringCountRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteJSON[map[string]any](w, http.StatusUnprocessableEntity, false, "Json Inválido", nil)
		return
	}

	select {
	case <-time.After(1 * time.Second):

		retorno, err := service.CountString(payload)
		if err != nil {
			response.WriteJSON[map[string]any](w, http.StatusUnprocessableEntity, false, "Erro ao executar ação de contar string", nil)
			return
		}

		log.Println(retorno)
		return
		//json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Go with Context!"})
	case <-ctx.Done():
		log.Println("Contexto cancelado:", ctx.Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
	}
}
