package main

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/api"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/db"
)

func main() {
	config.SetupServer()
	db.SetupDatabase(config.ServerCFG)
	api.Setup()
}
