package config

import (
	"log"
	"os"
)

// Config estrutura de configuração
type Config struct {
	ServerAddress string
}

// Load carrega as configurações do sistema
func Load() *Config {

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
		log.Println("Defaulting SERVER_ADDRESS to :8080")
	}

	return &Config{
		ServerAddress: address,
	}
}
