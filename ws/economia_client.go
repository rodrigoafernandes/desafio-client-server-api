package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"net/http"
	"time"
)

type EconomiaWSClient struct {
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

func NewEconomiaWSClient(cfg config.ServerConfig) (EconomiaWSClient, error) {
	economiaWSClient := EconomiaWSClient{
		client:  http.DefaultClient,
		timeout: cfg.EconomiaWSTimeoutMilliseconds,
		url:     cfg.EconomiaWSUrl,
	}
	if economiaWSClient.timeout < 1 {
		return EconomiaWSClient{}, errors.New("milliseconds timeout must be greater then zero")
	}
	if economiaWSClient.url == "" {
		return EconomiaWSClient{}, errors.New("economia ws client url must be informed")
	}
	return economiaWSClient, nil
}

func (eWSClient EconomiaWSClient) GetUSDQuotationFromBRL() (Cotacao, error) {
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
