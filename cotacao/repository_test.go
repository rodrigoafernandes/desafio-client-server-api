package cotacao

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"testing"
)

func TestGivenSuccessfullySaveOnDBWhenSaveCotacaoThenShouldNotReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("um erro inesperado ocorreu ao instanciar o banco de dados de teste")
	}
	defer db.Close()
	cfg := config.ServerConfig{
		DbConnectionString:               "",
		DbTransactionTimeoutMilliseconds: 0,
	}
	cotacao := CotacaoDB{
		ID:         0,
		Code:       "A",
		CodeIn:     "A",
		Name:       "A",
		High:       "A",
		Low:        "A",
		VarBid:     "A",
		PctChange:  "A",
		Bid:        "A",
		Ask:        "A",
		Timestamp:  "A",
		CreateDate: "A",
	}
	mock.ExpectExec("INSERT INTO cotacoes").WithArgs(
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
		cotacao.CreateDate).WillReturnResult(
		sqlmock.NewResult(1, 1),
	)
	repository := NewRepository(db, cfg)
	err = repository.Save(cotacao)
	if err != nil {
		t.Errorf("nenhum erro deveria ser retornado quando nenhum erro ocorrer ao salvar no DB. err %s", err.Error())
	}
}

func TestGivenDBErrorWhenSaveCotacaoThenShouldReturnsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("um erro inesperado ocorreu ao instanciar o banco de dados de teste")
	}
	defer db.Close()
	cfg := config.ServerConfig{
		DbConnectionString:               "",
		DbTransactionTimeoutMilliseconds: 0,
	}
	cotacao := CotacaoDB{
		ID:         0,
		Code:       "A",
		CodeIn:     "A",
		Name:       "A",
		High:       "A",
		Low:        "A",
		VarBid:     "A",
		PctChange:  "A",
		Bid:        "A",
		Ask:        "A",
		Timestamp:  "A",
		CreateDate: "A",
	}
	mock.ExpectExec("INSERT INTO cotacoes").WithArgs(
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
		cotacao.CreateDate).WillReturnError(
		errors.New("error access db"),
	)
	repository := NewRepository(db, cfg)
	err = repository.Save(cotacao)
	if err == nil {
		t.Error("deve retornar erro ao falhar quando salvar no banco de dados")
	}
}
