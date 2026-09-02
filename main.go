package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error while loading the .env file!")
	}
}
