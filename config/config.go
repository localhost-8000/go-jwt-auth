package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)


func GetDotEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}