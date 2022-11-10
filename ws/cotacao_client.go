package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"net/http"
	"strconv"
	"time"
)

type CotacaoWSClient struct {
	client  *http.Client
	url     string
	port    int
	timeout int
}

func NewCotacaoWSClient(cfg config.ClientConfig) (CotacaoWSClient, error) {
	cotacaoClient := CotacaoWSClient{
		client:  http.DefaultClient,
		url:     cfg.CotacaoServerUrl,
		port:    cfg.CotacaoServerPort,
		timeout: cfg.CotacaoServerClientTimeoutMilliseconds,
	}
	if cotacaoClient.timeout < 1 {
		cotacaoClient.timeout = 300
	}
	if cotacaoClient.url == "" {
		cotacaoClient.url = "http://localhost"
	}
	if cotacaoClient.port < 1 {
		cotacaoClient.port = 8080
	}
	return cotacaoClient, nil
}

func (cotacaoClient CotacaoWSClient) GetUSDBRLQuotation() (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cotacaoClient.timeout)*time.Millisecond)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s:%d/cotacao", cotacaoClient.url, cotacaoClient.port), nil)
	if err != nil {
		return 0, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	var responseBody Cotacao
	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return 0, nil
	}
	bid, err := strconv.ParseFloat(responseBody.Bid, 64)
	if err != nil {
		return 0, err
	}
	return bid, nil
}
