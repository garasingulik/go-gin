package config

import (
	lib "gsk/app/lib"

	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// App - stores configuration
var App *Config
var Server *ServerConfig

// Init - read .env file & populate app's configuration
func Init() error {
	// Create logger
	log := lib.CreateLogger("config")

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Errorf("%v", err)
	}

	helper := lib.StringHelper{}

	// Load environment variables & populate App
	App = &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	Server = &ServerConfig{
		Port:       helper.ToInt(getEnv("HTTP_SERVER_PORT", "3000")),
		EnableCORS: getEnv("HTTP_CORS_ENABLED", "false") == "true",
	}

	return nil
}
