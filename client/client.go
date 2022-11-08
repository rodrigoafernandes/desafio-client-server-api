package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
)

func main() {
	config.Setup()
	economiaWSClient, err := ws.NewEconomiaWSClient(config.ServerCFG)
	if err != nil {
		panic(err)
	}
	response, err := economiaWSClient.GetUSDQuotationFromBRL()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Dolar: %s\n", response.Bid)
	fmt.Println("client app")
}
