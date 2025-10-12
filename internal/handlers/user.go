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

// HelloHandler responde com uma mensagem JSON
func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//var payload requests.StringCountRequest
	var payload requests.LoginUser

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteJSON[map[string]any](w, http.StatusUnprocessableEntity, false, "Json Inválido", nil)
		return
	}

	select {
	case <-time.After(1 * time.Second):

		retorno, err := service.LoginUser(payload)
		if err != nil {
			if serr, ok := err.(*response.ServiceError); ok {
				response.WriteJSON(w, serr.StatusCode, false, serr.Message, response.Empty{})
			} else {
				// Erro genérico
				response.WriteJSON(w, http.StatusInternalServerError, false, "Erro interno", response.Empty{})
			}
			return
		}

		response.WriteJSON(w, http.StatusOK, true, "OK", retorno)
		return
	case <-ctx.Done():
		log.Println("Contexto cancelado:", ctx.Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
	}
}

func GetValidationMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(1 * time.Second):

		var validTokenResponse response.ValidTokenResponse
		validTokenResponse.Status = "OK"
		validTokenResponse.Message = "Token valido"
		json.NewEncoder(w).Encode(validTokenResponse)
		return
	case <-ctx.Done():
		log.Println("Contexto cancelado:", ctx.Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
	}

}
