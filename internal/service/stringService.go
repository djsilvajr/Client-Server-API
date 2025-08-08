package service

import (
	"log"
	"my-app/internal/requests"
)

func CountString(r requests.StringCountRequest) string {
	log.Println(r)
	return ""
}
