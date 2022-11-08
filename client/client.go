package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
)

func main() {
	config.Setup()
	cfg := config.ClientCFG
	fmt.Println(cfg.CotacaoServerUrl)
	fmt.Println("client app")
}
