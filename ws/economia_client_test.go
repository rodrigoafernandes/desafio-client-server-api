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

type testingScenario func(t *testing.T)

type scenario struct {
	name      string
	cfg       config.ServerConfig
	tScenario testingScenario
}

func TestEconomiaClient(t *testing.T) {
	hClient := httpClientMock{}
	scenarios := []scenario{
		givenNoErrorsWhenGetUSDBRLQuotationThenShouldReturnsCotacaoStructScenario(hClient),
		givenUrlErrorWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient),
		givenTimeoutWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient),
		givenJsonResponseErrorWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient),
	}
	for _, scn := range scenarios {
		t.Run(scn.name, scn.tScenario)
	}
}

func givenNoErrorsWhenGetUSDBRLQuotationThenShouldReturnsCotacaoStructScenario(hClient httpClientMock) scenario {
	cfg := config.ServerConfig{
		EconomiaWSUrl:                 "",
		EconomiaWSTimeoutMilliseconds: 0,
	}
	economiaResponse := economiaWSResponse{
		FromCurrencyToCurrency: Cotacao{
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
		},
	}
	return createScenario(
		"givenNoErrorsWhenGetUSDBRLQuotationsThenShouldReturnsCotacaoStruct",
		cfg,
		func(t *testing.T) {
			b, _ := json.Marshal(economiaResponse)
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(string(b))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			economiaClient, _ := NewEconomiaWSClient(cfg, hClient)
			cotacao, err := economiaClient.GetUSDQuotationFromBRL()
			if err != nil {
				t.Errorf("error %s", err.Error())
			}
			if len(cotacao.Code) < 1 {
				t.Errorf("deveria retornar o codigo da cotação")
			}
		},
	)
}

func givenUrlErrorWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient httpClientMock) scenario {
	cfg := config.ServerConfig{
		EconomiaWSUrl:                 "://test.api.com.br",
		EconomiaWSTimeoutMilliseconds: 0,
	}
	return createScenario(
		"givenUrlErrorWhenGetUSDBRLQuotationsThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			economiaClient, _ := NewEconomiaWSClient(cfg, hClient)
			cotacao, err := economiaClient.GetUSDQuotationFromBRL()
			if err == nil {
				t.Errorf("error deveria ser retornado")
			}
			if len(cotacao.Code) > 0 {
				t.Errorf("Não deveria retornar nenhuma informação sobre a cotação quando um erro ocorrer")
			}
		},
	)
}

func givenTimeoutWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient httpClientMock) scenario {
	cfg := config.ServerConfig{
		EconomiaWSUrl:                 "https://test.api.com.br",
		EconomiaWSTimeoutMilliseconds: 1,
	}
	expectedErr := errors.New("client timeout")
	economiaResponse := economiaWSResponse{
		FromCurrencyToCurrency: Cotacao{
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
		},
	}
	return createScenario(
		"givenTimeoutWhenGetUSDBRLQuotationThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			b, _ := json.Marshal(economiaResponse)
			response := &http.Response{
				StatusCode: http.StatusRequestTimeout,
				Body:       io.NopCloser(bytes.NewBufferString(string(b))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, expectedErr)
			economiaClient, _ := NewEconomiaWSClient(cfg, hClient)
			cotacao, err := economiaClient.GetUSDQuotationFromBRL()
			if err == nil {
				t.Errorf("deveria retornar erro ao ocorrer timeout da solicitação")
			}
			if len(cotacao.Code) > 0 {
				t.Errorf("não deveria retornar nenhuma informação da cotação quando um erro ocorrer")
			}
		},
	)
}

func givenJsonResponseErrorWhenGetUSDBRLQuotationThenShouldReturnsErrorScenario(hClient httpClientMock) scenario {
	cfg := config.ServerConfig{
		EconomiaWSUrl:                 "",
		EconomiaWSTimeoutMilliseconds: 0,
	}
	return createScenario(
		"givenJsonResponseErrorWhenGetUSDBRLQuotationThenShouldReturnsError",
		cfg,
		func(t *testing.T) {
			response := &http.Response{
				StatusCode: http.StatusMultiStatus,
				Body:       io.NopCloser(bytes.NewBufferString(string("abc"))),
				Close:      false,
			}
			hClient.On("Do", mock.Anything).Return(response, nil)
			economiaClient, _ := NewEconomiaWSClient(cfg, hClient)
			cotacao, err := economiaClient.GetUSDQuotationFromBRL()
			if err == nil {
				t.Errorf("deveria retornar erro ao ocorrer erro de marshall do response body")
			}
			if len(cotacao.Code) > 0 {
				t.Errorf("não deveria retornar nenhuma informação da cotação quando um erro ocorrer")
			}
		},
	)
}

func createScenario(name string, cfg config.ServerConfig, tScenario testingScenario) scenario {
	return scenario{
		name:      name,
		cfg:       cfg,
		tScenario: tScenario,
	}
}
