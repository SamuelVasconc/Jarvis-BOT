package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("xoxb-640873308674-1179766554435-BUB2LMfjEZT73Nx4pH5cznrE")

	definition := &slacker.CommandDefinition{
		Description: "faz eco",
		Example:     "echo ola",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			response.Reply(word, slacker.WithThreadReply(true))
		},
	}

	bot.Command("echo <word>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
