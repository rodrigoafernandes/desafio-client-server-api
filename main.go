package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/db"
)

func main() {
	config.SetupServer()
	db.SetupDatabase(config.ServerCFG)
	fmt.Println("server app")
}
