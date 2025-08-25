package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type domainConfig struct {
	Host string
	Port int
}

type Config struct {
	DomainConfig domainConfig
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env is not found!, using system environment")
	}

	portStr := os.Getenv("PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid PORT value: %s, defaulting to 8080\n", portStr)
		port = 8080
	}

	return Config{
		DomainConfig: domainConfig{
			Host: os.Getenv("HOST"),
			Port: port,
		},
	}
}