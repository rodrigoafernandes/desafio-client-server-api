package ws

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

type testingScenarioClient func(t *testing.T)

type scenarioClient struct {
	name      string
	cfg       config.ClientConfig
	tScenario testingScenarioClient
}

func createScenarioClient(name string, cfg config.ClientConfig, tScenario testingScenarioClient) scenarioClient {
	return scenarioClient{
		name:      name,
		cfg:       cfg,
		tScenario: tScenario,
	}
}

func TestCotacaoClient(t *testing.T) {
	hClient := httpClientMock{}
	scenarios := []scenarioClient{
		givenNoErrorsWhenSearchCotacaoThenShouldReturnsCotacaoStructScenario(hClient),
		givenMalformadURLWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient),
		givenAnyErrorOnHttpCallWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient),
		givenHttpStatusIsNotSucessWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient),
		givenResponseBodyFormatErrorWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient),
		givenBidValueFormatErrorWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient),
	}
	for _, scn := range scenarios {
		t.Run(scn.name, scn.tScenario)
	}
}

func givenNoErrorsWhenSearchCotacaoThenShouldReturnsCotacaoStructScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "https://cotacao-api.test.com.br",
		CotacaoServerPort:                      8080,
		CotacaoServerClientTimeoutMilliseconds: 300,
	}
	cotacaoResponse := Cotacao{
		Code:       "USD",
		CodeIn:     "BRL",
		Name:       "Dólar Americano/Real Brasileiro",
		High:       "5.4293",
		Low:        "5.3495",
		VarBid:     "-0.0457",
		PctChange:  "-0.84",
		Bid:        "5.3633",
		Ask:        "5.3649",
		Timestamp:  "1669661481",
		CreateDate: "2022-11-28 15:51:21",
	}
	return createScenarioClient(
		"givenNoErrorsWhenSearchCotacaoThenShouldReturnsCotacaoStruct",
		cfg,
		func(t *testing.T) {
			b, _ := json.Marshal(cotacaoResponse)
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(string(b))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err != nil {
				t.Errorf("nenhum erro deveria ser retornado quando a chamada for bem sucedida, err %s", err.Error())
			}
			if bid == 0.0 {
				t.Errorf("o valor de bid deve ser retornado quando nenhum erro ocorrer")
			}
		},
	)
}

func givenMalformadURLWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "://cotacao-api.test.com.br",
		CotacaoServerPort:                      0,
		CotacaoServerClientTimeoutMilliseconds: 0,
	}
	return createScenarioClient(
		"givenMalformadURLWhenSearchCotacaoThenShouldReturnsError",
		config.ClientConfig{},
		func(t *testing.T) {
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err == nil {
				t.Error("deve retornar erro quando a url não estiver formatada corretamente")
			}
			if bid > 0.0 {
				t.Error("não deve retornar nenhum valor quando ocorrer um erro")
			}
		},
	)
}

func givenAnyErrorOnHttpCallWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "",
		CotacaoServerPort:                      8080,
		CotacaoServerClientTimeoutMilliseconds: 300,
	}
	cotacaoResponse := Cotacao{
		Code:       "USD",
		CodeIn:     "BRL",
		Name:       "Dólar Americano/Real Brasileiro",
		High:       "5.4293",
		Low:        "5.3495",
		VarBid:     "-0.0457",
		PctChange:  "-0.84",
		Bid:        "5.3633",
		Ask:        "5.3649",
		Timestamp:  "1669661481",
		CreateDate: "2022-11-28 15:51:21",
	}
	errExpected := errors.New("client timeout")
	return createScenarioClient(
		"givenAnyErrorOnHttpCallWhenSearchCotacaoThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			b, _ := json.Marshal(cotacaoResponse)
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(string(b))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, errExpected)
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err == nil {
				t.Error("deve retornar erro quando ocorrer timeout na chamada a api")
			}
			if bid > 0.0 {
				t.Error("não deve retornar nenhum valor quando ocorrer um erro")
			}
		},
	)
}

func givenHttpStatusIsNotSucessWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "",
		CotacaoServerPort:                      8080,
		CotacaoServerClientTimeoutMilliseconds: 300,
	}
	return createScenarioClient(
		"givenHttpStatusIsNotSucessWhenSearchCotacaoThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			response := &http.Response{
				StatusCode: http.StatusServiceUnavailable,
				Body:       io.NopCloser(bytes.NewBufferString("SERVICE UNAVAILABLE")),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err == nil {
				t.Error("deve retornar erro quando status http for diferente de sucesso na chamada a api")
			}
			if bid > 0.0 {
				t.Error("não deve retornar nenhum valor quando ocorrer um erro")
			}
		},
	)
}

func givenResponseBodyFormatErrorWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "",
		CotacaoServerPort:                      8080,
		CotacaoServerClientTimeoutMilliseconds: 300,
	}
	return createScenarioClient(
		"givenResponseBodyFormatErrorWhenSearchCotacaoThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("OK")),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err == nil {
				t.Error("deve retornar erro quando o response body for diferente do experado na chamada a api")
			}
			if bid > 0.0 {
				t.Error("não deve retornar nenhum valor quando ocorrer um erro")
			}
		},
	)
}

func givenBidValueFormatErrorWhenSearchCotacaoThenShouldReturnsErrorScenario(hClient httpClientMock) scenarioClient {
	cfg := config.ClientConfig{
		CotacaoServerUrl:                       "https://cotacao-api.test.com.br",
		CotacaoServerPort:                      8080,
		CotacaoServerClientTimeoutMilliseconds: 300,
	}
	cotacaoResponse := Cotacao{
		Code:       "USD",
		CodeIn:     "BRL",
		Name:       "Dólar Americano/Real Brasileiro",
		High:       "5.4293",
		Low:        "5.3495",
		VarBid:     "-0.0457",
		PctChange:  "-0.84",
		Bid:        "ABC",
		Ask:        "5.3649",
		Timestamp:  "1669661481",
		CreateDate: "2022-11-28 15:51:21",
	}
	return createScenarioClient(
		"givenBidValueFormatErrorWhenSearchCotacaoThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			b, _ := json.Marshal(cotacaoResponse)
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(string(b))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			cotacaoClient, _ := NewCotacaoWSClient(cfg, hClient)
			bid, err := cotacaoClient.GetUSDBRLQuotation()
			if err == nil {
				t.Error("deve retornar erro quando o formato do atributo \"bid\" não for decimal na resposta da chamada a api")
			}
			if bid > 0.0 {
				t.Error("não deve retornar nenhum valor quando ocorrer um erro")
			}
		},
	)
}
