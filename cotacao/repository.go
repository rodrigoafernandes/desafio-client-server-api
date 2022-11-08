package cotacao

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(cotacao CotacaoDB) error {
	_, err := r.db.Exec(`
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
                      create_date,
                      payload
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
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
		cotacao.Payload,
	)
	return err
}
