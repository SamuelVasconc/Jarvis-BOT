package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type data struct {
	ID     int    `json:"id"`
	Imagem string `json:"imagem"`
	Nome   string `json:"nome"`
	Adress string `json:"adress"`
}

func main() {
	response, err := http.Get("http://localhost:3001/posts/1") //get

	if err != nil {
		log.Fatal(err) //tratar erro
	}
	defer response.Body.Close() //fechar a requisicao assim que acabar

	body, err := ioutil.ReadAll(response.Body) //ler a resposta http e converter

	if err != nil {
		log.Fatal(err) //tratar erro
	}

	var dado data
	json.Unmarshal(body, &dado)
	fmt.Print("o nome é: ", dado.Nome)
}
