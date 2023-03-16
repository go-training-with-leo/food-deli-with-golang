package dotenv

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot load .env file: ", err)
	}

	fmt.Println("Loaded .env file!")
}
