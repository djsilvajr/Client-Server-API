package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"my-app/internal/requests"
	"my-app/internal/response"
	"my-app/internal/service"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// feche o body ao final (boa prática)
	defer r.Body.Close()

	// Decodificador que falha se houver campos desconhecidos (ajuda a evitar payloads inválidos)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var payload requests.LoginUser
	if err := dec.Decode(&payload); err != nil {
		// Mensagem clara ao cliente e log simples para debug
		log.Printf("failed to decode login payload: %v\n", err)
		response.WriteJSON(w, http.StatusUnprocessableEntity, false, "JSON inválido", response.Empty{})
		return
	}

	// Verifica se o contexto já foi cancelado antes de processar
	select {
	case <-r.Context().Done():
		log.Println("request canceled before login processing:", r.Context().Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
		return
	default:
		// continua
	}

	// Chama o serviço de login
	retorno, err := service.LoginUser(payload)
	if err != nil {
		// Use errors.As para suportar erros embrulhados (wrapping)
		var serr *response.ServiceError
		if errors.As(err, &serr) {
			// erro tratado pela camada de serviço (ex.: credenciais inválidas)
			response.WriteJSON(w, serr.StatusCode, false, serr.Message, response.Empty{})
			return
		}

		// erro inesperado
		log.Printf("unexpected error in LoginUser service: %v\n", err)
		response.WriteJSON(w, http.StatusInternalServerError, false, "Erro interno", response.Empty{})
		return
	}

	// Retorno de sucesso (200)
	response.WriteJSON(w, http.StatusOK, true, "OK", retorno)
}

func GetValidationMessage(w http.ResponseWriter, r *http.Request) {
	// Respeita cancelamento do contexto
	select {
	case <-r.Context().Done():
		log.Println("request canceled in GetValidationMessage:", r.Context().Err())
		http.Error(w, "Request canceled or timeout", http.StatusRequestTimeout)
		return
	default:
		// Monta resposta de validação
		// Usa WriteJSON se disponível para resposta consistente
		response.WriteJSON(w, http.StatusOK, true, "Token válido", map[string]any{})
	}
}
