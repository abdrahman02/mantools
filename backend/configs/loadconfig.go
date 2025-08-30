package configs

import (
	"log"
	"os"
	"strconv"
)

type domainConfig struct {
	Host string
	Port int
	FrontendBaseURL string
}

type Config struct {
	DomainConfig domainConfig
}

func LoadConfig() Config {
	port := getEnvAsInt("PORT", 8080)
	return Config{
		DomainConfig: domainConfig{
			Host: os.Getenv("HOST"),
			Port: port,
			FrontendBaseURL: os.Getenv("FRONTEND_BASE_URL"),
		},
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Invalid value for %s: %s, using default %d", key, valStr, defaultValue)
		return defaultValue
	}

	return val
}