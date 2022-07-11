package eraspacelog

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValue interface{}) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load file .env: ", err)
		panic(err)
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue.(string)
	}

	return value
}
