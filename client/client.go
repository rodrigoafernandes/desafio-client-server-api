package main

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/output"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
)

func main() {
	config.SetupClient()
	cotacaoClient, err := ws.NewCotacaoWSClient(config.ClientCFG)
	if err != nil {
		panic(err)
	}
	ro, err := output.NewResultOutput(config.ClientCFG)
	if err != nil {
		panic(err)
	}
	bid, err := cotacaoClient.GetUSDBRLQuotation()
	if err != nil {
		panic(err)
	}
	err = ro.WriteQuotationResult(bid)
	if err != nil {
		panic(err)
	}
}
