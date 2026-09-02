package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Respond with Pong!",
	},
}

func PingCommandHandler(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	return e.CreateMessage(discord.MessageCreate{
		Content: "Pong!",
	})
}
