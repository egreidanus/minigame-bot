package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error while loading the .env file!")
	}

	discord_token := os.Getenv("DISCORD_TOKEN")
	if discord_token == "" {
		log.Fatal("Failed to get the DISCORD_TOKEN from the .env file!")
	}
}
