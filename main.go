package main

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/api"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/cotacao"
	"github.com/rodrigoafernandes/desafio-client-server-api/db"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
)

func main() {
	config.SetupServer()
	db.SetupDatabase(config.ServerCFG)
	wsClient, err := ws.NewEconomiaWSClient(config.ServerCFG)
	if err != nil {
		panic(err)
	}
	repository := cotacao.NewRepository(db.DB, config.ServerCFG)
	quotationService, err := cotacao.NewQuotationService(wsClient, repository)
	if err != nil {
		panic(err)
	}
	api.Setup(quotationService)
}
