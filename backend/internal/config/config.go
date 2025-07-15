package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLURL    string
	RabbitMQURL string
	Port        string
	GinMode     string
	Environment string
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

// Load carrega as variáveis de ambiente
func Load() *Config {
	env := getEnv("ENVIRONMENT", "local")

	if env != "production" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println("Aviso: .env.local não encontrado. Variáveis de ambiente devem estar definidas manualmente.")
		} else {
			log.Println(".env.local carregado com sucesso.")
		}
	}

	mysqlURL, ok := os.LookupEnv("MYSQL_URL")
	if !ok || mysqlURL == "" {
		log.Fatal("MYSQL_URL não definida")
	}

	rabbitMQURL, ok := os.LookupEnv("RABBITMQ_URL")
	if !ok || rabbitMQURL == "" {
		log.Fatal("RABBITMQ_URL não definida")
	}

	return &Config{
		MySQLURL:    mysqlURL,
		RabbitMQURL: rabbitMQURL,
		Port:        getEnv("API_PORT", "8080"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		Environment: env,
	}
}
