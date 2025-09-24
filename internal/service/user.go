package service

import (
	//"log"
	"my-app/internal/requests"
)

func LoginUser(r requests.LoginUser) (map[string]any, error) {
	retorno := map[string]any{
		"token": "fake-jwt-token-123",
		"user": map[string]string{
			"email": r.Email,
			"name":  "Usuário de Teste",
		},
	}
	return retorno, nil
}
