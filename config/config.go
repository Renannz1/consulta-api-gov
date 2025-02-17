package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representa as configurações do sistema
type Config struct {
	PortalAPIKey string
	PortalAPIURL string
}

// LoadConfig carrega as configurações do .env
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando variáveis de ambiente")
	}

	return Config{
		PortalAPIKey: os.Getenv("PORTAL_TRANSPARENCIA_API_KEY"),
		PortalAPIURL: os.Getenv("PORTAL_TRANSPARENCIA_URL"),
	}
}
