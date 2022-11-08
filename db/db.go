package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
)

const InitDb string = `
	CREATE TABLE IF NOT EXISTS cotacoes(
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    code TEXT NOT NULL,
	    codein TEXT NOT NULL,
	    name TEXT NOT NULL,
	    high TEXT NOT NULL,
	    low TEXT NOT NULL,
	    varBid TEXT NOT NULL,
	    pctChange TEXT NOT NULL,
	    bid TEXT NOT NULL,
	    ask TEXT NOT NULL,
	    timestamp TEXT NOT NULL,
	    create_date TEXT NOT NULL,
	    payload TEXT NOT NULL
	)
`

var DB *sql.DB

func SetupDatabase(cfg config.ServerConfig) {
	driverName := "sqlite3"
	connectionString := fmt.Sprintf("%s.db", cfg.DbConnectionString)
	dbTemp, err := sql.Open(driverName, connectionString)
	if err != nil {
		panic(err)
	}

	if _, err = dbTemp.Exec(InitDb); err != nil {
		panic(err)
	}

	DB = dbTemp
}
