package service

import (
	//"log"
	"my-app/internal/requests"
)

type CountStringResponse struct {
	Valor string `json:"valor"`
}

func CountString(r requests.StringCountRequest) ([]CountStringResponse, error) {
	resp := []CountStringResponse{
		{Valor: r.Valor},
	}
	//log.Println(r)
	return resp, nil
}
