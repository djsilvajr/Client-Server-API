package main

import (
	"fmt"
	"log"
	"net/http"

	"my-app/internal/config"
	"my-app/internal/db"
	"my-app/internal/router"
)

func main() {
	fmt.Println("INICIANDO SERVER")
	cfg := config.Load()
	log.Println("CONFIG LOADED ON PORT", cfg.ServerAddress)

	dbConn, err := db.Connect(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("CONNECTED TO DB SUCCESSFULLY")
	defer dbConn.Close()

	r := router.Setup()
	//fmt.Println(r)

	// Inicializa servidor HTTP
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	log.Printf("SERVER RUNNING ON %s", cfg.ServerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
