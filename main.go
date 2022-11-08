package main

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/cotacao"
	"github.com/rodrigoafernandes/desafio-client-server-api/db"
)

func main() {
	config.SetupServer()
	db.SetupDatabase(config.ServerCFG)

	var ct cotacao.CotacaoDB

	db.DB.QueryRow(`SELECT
    						id, code, codein, name, high, low, varBid, ask, timestamp, create_date
					FROM cotacoes`).Scan(
		&ct.ID,
		&ct.Code,
		&ct.CodeIn,
		&ct.Name,
		&ct.High,
		&ct.Low,
		&ct.VarBid,
		&ct.PctChange,
		&ct.Bid,
		&ct.Ask,
		&ct.Timestamp,
		&ct.CreateDate)
	fmt.Println(ct)
	fmt.Println("server app")
}
