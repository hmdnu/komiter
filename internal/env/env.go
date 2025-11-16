package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_URL   = ""
	API_TOKEN = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: Failed loading .env")
	}
	API_URL = os.Getenv("API_URL")
	API_TOKEN = os.Getenv("API_TOKEN")
}
