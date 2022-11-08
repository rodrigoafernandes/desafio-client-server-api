package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
)

func main() {
	config.SetupClient()
	fmt.Println("client app")
}
