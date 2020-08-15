package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("xoxb-640873308674-1179766554435-BUB2LMfjEZT73Nx4pH5cznrE") //inicia o bot passando o token dele

	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("e ae")
		},
	}

	bot.Command("ola", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
