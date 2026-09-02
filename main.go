package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"

	"github.com/joho/godotenv"
)

var (
	discord_token = ""
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error while loading the .env file!")
	}

	discord_token := os.Getenv("DISCORD_TOKEN")
	if discord_token == "" {
		log.Fatal("Failed to get the DISCORD_TOKEN from the .env file!")
	}

	client, err := disgo.New(discord_token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),

		bot.WithEventListenerFunc(func(e *events.Ready) {
			fmt.Println("Bot is connected as", e.User.Username)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
