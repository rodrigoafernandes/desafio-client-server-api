package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"net/http"
	"strconv"
	"time"
)

type CotacaoWSClient interface {
	GetUSDBRLQuotation() (float64, error)
}

type CotacaoWSClientImpl struct {
	client  *http.Client
	url     string
	port    int
	timeout int
}

func NewCotacaoWSClient(cfg config.ClientConfig) (CotacaoWSClientImpl, error) {
	cotacaoClient := CotacaoWSClientImpl{
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

func (cotacaoClient CotacaoWSClientImpl) GetUSDBRLQuotation() (float64, error) {
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
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return 0, errors.New("serviço de cotação indisponível")
	}
	var responseBody Cotacao
	if err = json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return 0, err
	}
	bid, err := strconv.ParseFloat(responseBody.Bid, 64)
	if err != nil {
		return 0, err
	}
	return bid, nil
}
