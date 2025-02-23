package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitalizeEnv() {
	// Load the .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func GetAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}
