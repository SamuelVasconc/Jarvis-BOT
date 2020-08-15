package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("xoxb-640873308674-1179766554435-BUB2LMfjEZT73Nx4pH5cznrE")

	definition := &slacker.CommandDefinition{
		Description: "Diz ola",
		Example:     "ola Jarvis",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("ola senhor", slacker.WithThreadReply(true))
		},
	}

	bot.Command("ola Jarvis", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
