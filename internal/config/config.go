package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config estrutura de configuração
type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
}

// Load carrega as configurações do sistema
func Load() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
		log.Println("Defaulting SERVER_ADDRESS to :8080")
	}

	return &Config{
		ServerAddress: address,
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
	}
}
