package cotacao

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/rodrigoafernandes/desafio-client-server-api/db"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
)

type QuotationService struct {
	client ws.EconomiaWSClient
	repo   Repository
}

func NewQuotationService(cfg config.ServerConfig) (QuotationService, error) {
	wsClient, err := ws.NewEconomiaWSClient(cfg)
	if err != nil {
		return QuotationService{}, err
	}
	return QuotationService{
		client: wsClient,
		repo:   NewRepository(db.DB, cfg),
	}, nil
}

func (qs QuotationService) GetUSDQuotation() (ws.Cotacao, error) {
	cotacao, err := qs.client.GetUSDQuotationFromBRL()
	if err != nil {
		return ws.Cotacao{}, err
	}
	cotacaoDB := CotacaoDB{
		Code:       cotacao.Code,
		CodeIn:     cotacao.CodeIn,
		Name:       cotacao.Name,
		High:       cotacao.High,
		Low:        cotacao.Low,
		VarBid:     cotacao.VarBid,
		PctChange:  cotacao.PctChange,
		Bid:        cotacao.Bid,
		Ask:        cotacao.Ask,
		Timestamp:  cotacao.Timestamp,
		CreateDate: cotacao.CreateDate,
	}
	err = qs.repo.Save(cotacaoDB)
	if err != nil {
		return ws.Cotacao{}, err
	}
	return cotacao, nil
}
