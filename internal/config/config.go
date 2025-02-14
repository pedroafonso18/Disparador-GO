package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	DBURL = os.Getenv("DB_URL")
}
