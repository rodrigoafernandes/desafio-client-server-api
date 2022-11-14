package cotacao

import (
	"context"
	"database/sql"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"time"
)

type Repository interface {
	Save(cotacao CotacaoDB) error
}

type RepositoryImpl struct {
	db      *sql.DB
	timeout int
}

func NewRepository(db *sql.DB, cfg config.ServerConfig) RepositoryImpl {
	repository := RepositoryImpl{
		db:      db,
		timeout: cfg.DbTransactionTimeoutMilliseconds,
	}
	if repository.timeout < 1 {
		repository.timeout = 10
	}
	return repository
}

func (r RepositoryImpl) Save(cotacao CotacaoDB) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Millisecond)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO cotacoes (
                      code,
                      codein,
                      name,
                      high,
                      low,
                      varBid,
                      pctChange,
                      bid,
                      ask,
                      timestamp,
                      create_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		cotacao.Code,
		cotacao.CodeIn,
		cotacao.Name,
		cotacao.High,
		cotacao.Low,
		cotacao.VarBid,
		cotacao.PctChange,
		cotacao.Bid,
		cotacao.Ask,
		cotacao.Timestamp,
		cotacao.CreateDate,
	)
	return err
}
