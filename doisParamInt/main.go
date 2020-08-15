package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("<YOUR SLACK BOT TOKEN>")

	definition := &slacker.CommandDefinition{
		Description: "Repeat a word a number of times!",
		Example:     "repeat hello 10",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.StringParam("word", "Hello!") //define esse parametro em string
			number := request.IntegerParam("number", 1)   //define esse parametro em inteiro
			for i := 0; i < number; i++ {
				response.Reply(word)
			}
		},
	}

	bot.Command("repeat <word> <number>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
