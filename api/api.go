package api

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/cotacao"
	"log"
	"net/http"
)

func Setup() {
	controller, err := cotacao.NewController(config.ServerCFG)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/cotacao", controller.GetCotacaoUSD)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
