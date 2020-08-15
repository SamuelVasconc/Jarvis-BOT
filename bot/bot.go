package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shomali11/slacker"
)

type data struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func consulta(id string) string {

	response, _ := http.Get("http://localhost:3001/profile/" + id) //get

	// defer response.Body.Close() //fechar a requisicao assim que acabar

	body, _ := ioutil.ReadAll(response.Body) //ler a resposta http e converter

	var dado data
	json.Unmarshal(body, &dado) //converter json

	return dado.Status //retorna string
}

func main() {
	bot := slacker.NewClient("xoxb-640873308674-1179766554435-BUB2LMfjEZT73Nx4pH5cznrE")

	definition := &slacker.CommandDefinition{
		Description: "Consulta cadastro",
		Example:     "Consultar 266465",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			di := request.Param("id")

			response.Reply("O ID "+di+" "+consulta(di), slacker.WithThreadReply(true)) //Reply só aceita string!!!!!
		},
	}

	another := &slacker.CommandDefinition{
		Description: "Verifica credencial",
		Example:     "Verificar 266465",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Um momento..", slacker.WithThreadReply(true))
			time.Sleep(time.Second)

			di := request.Param("id")

			response.Reply(di+" Verificado!", slacker.WithThreadReply(true)) //Reply só aceita string!!!!!
		},
	}

	bot.Command("consultar <id>", definition) //comando
	bot.Command("verificar <id>", another)

	ctx, cancel := context.WithCancel(context.Background()) //erro e cancel
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil { //tratar erro
		log.Fatal(err)
	}
}
