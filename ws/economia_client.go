package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"net/http"
	"time"
)

type EconomiaWSClient interface {
	GetUSDQuotationFromBRL() (Cotacao, error)
}

type EconomiaWSClientImpl struct {
	client  *http.Client
	url     string
	timeout int
}

type economiaWSResponse struct {
	FromCurrencyToCurrency Cotacao `json:"USDBRL"`
}

type Cotacao struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func NewEconomiaWSClient(cfg config.ServerConfig) (EconomiaWSClientImpl, error) {
	economiaWSClient := EconomiaWSClientImpl{
		client:  http.DefaultClient,
		timeout: cfg.EconomiaWSTimeoutMilliseconds,
		url:     cfg.EconomiaWSUrl,
	}
	if economiaWSClient.timeout < 1 {
		economiaWSClient.timeout = 200
	}
	if economiaWSClient.url == "" {
		economiaWSClient.url = "https://economia.awesomeapi.com.br"
	}
	return economiaWSClient, nil
}

func (eWSClient EconomiaWSClientImpl) GetUSDQuotationFromBRL() (Cotacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(eWSClient.timeout)*time.Millisecond)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/json/last/USD-BRL", eWSClient.url), nil)
	if err != nil {
		return Cotacao{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Cotacao{}, err
	}
	defer response.Body.Close()
	var responseBody economiaWSResponse
	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return Cotacao{}, nil
	}
	return responseBody.FromCurrencyToCurrency, nil
}
