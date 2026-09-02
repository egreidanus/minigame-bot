package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"

	"github.com/egreidanus/minigamebot/minigamebot"
)

func MessageHandler(b *minigamebot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {
		// TODO: handle message
	})
}
