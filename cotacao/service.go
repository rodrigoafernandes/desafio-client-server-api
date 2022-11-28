package cotacao

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
)

type QuotationService interface {
	GetUSDQuotation() (ws.Cotacao, error)
}

type quotationServiceImpl struct {
	client ws.EconomiaWSClient
	repo   Repository
}

func NewQuotationService(wsClient ws.EconomiaWSClient, repository Repository) (QuotationService, error) {
	return quotationServiceImpl{
		client: wsClient,
		repo:   repository,
	}, nil
}

func (qs quotationServiceImpl) GetUSDQuotation() (ws.Cotacao, error) {
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
