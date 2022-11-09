package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/output"
)

func main() {
	config.SetupClient()
	ro, err := output.NewResultOutput(config.ClientCFG)
	if err != nil {
		panic(err)
	}
	err = ro.WriteQuotationResult(5.1592)
	if err != nil {
		panic(err)
	}
	fmt.Println("client app")
}
