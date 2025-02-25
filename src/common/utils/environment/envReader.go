package environment

import (
	"log"
	"os"
	"strconv"
	"strings"

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
	// returns as int
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

func GetAsSlice(key string, defaultValue []string) []string {
	// returns as slice with comma separated values
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}

func GetAsBool(key string, defaultValue bool) bool {
	// returns as bool
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}
