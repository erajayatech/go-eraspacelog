package eraspacelog

import (
	"encoding/json"
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

func Dump(i interface{}) string {
	return string(ToByte(i))
}

func ToByte(i interface{}) []byte {
	byte_, _ := JSONMarshal(i)
	return byte_
}

func JSONMarshal(data interface{}) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
