package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string
var EVOURL string
var EVOTOKEN string
var WUZURL string

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	DBURL = os.Getenv("DB_URL")
	EVOURL = os.Getenv("EVO_URL")
	EVOTOKEN = os.Getenv("EVO_TOKEN")
	WUZURL = os.Getenv("WUZ_URL")
}
