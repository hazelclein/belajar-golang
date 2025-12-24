package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DBPath          string
	PokeAPIBaseURL  string
	DefaultPageSize int
	MaxPageSize     int
}

var AppConfig *Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file tidak ditemukan, menggunakan default values")
	}

	port := getEnv("PORT", "8080")
	dbPath := getEnv("DB_PATH", "./pokemon.db")
	pokeAPIBaseURL := getEnv("POKEAPI_BASE_URL", "https://pokeapi.co/api/v2")
	defaultPageSize := getEnvAsInt("DEFAULT_PAGE_SIZE", 10)
	maxPageSize := getEnvAsInt("MAX_PAGE_SIZE", 100)

	AppConfig = &Config{
		Port:            port,
		DBPath:          dbPath,
		PokeAPIBaseURL:  pokeAPIBaseURL,
		DefaultPageSize: defaultPageSize,
		MaxPageSize:     maxPageSize,
	}

	log.Println("Environment variables loaded")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}