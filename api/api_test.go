package api

import (
	"encoding/json"
	"errors"
	"github.com/rodrigoafernandes/desafio-client-server-api/cotacao"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type quotationServiceMock struct {
	mock.Mock
}

func (qs quotationServiceMock) GetUSDQuotation() (ws.Cotacao, error) {
	args := qs.Called()
	return args.Get(0).(ws.Cotacao), args.Error(1)
}

type scenario struct {
	name               string
	expectedCotacao    ws.Cotacao
	expectedHttpStatus int
	tScenario          testingScenario
}

type testingScenario func(t *testing.T)

func TestApi(t *testing.T) {
	service := quotationServiceMock{}
	scenarios := []scenario{
		givenQuotationFoundWhenSearchUSDBRLQuotationThenShouldReturnsHttpStatusOkAndCotacaoJsonScenario(service),
		givenErrorGettingQuotationWhenSearchUSDToBRLQuotationThenShouldReturnsHttpStatusServiceUnavailableScenario(service),
	}
	for _, scn := range scenarios {
		t.Run(scn.name, scn.tScenario)
	}
}

func createScenario(name string, expectedCotacao ws.Cotacao, expectedHttpStatus int, tScenario testingScenario) scenario {
	return scenario{
		name:               name,
		expectedCotacao:    expectedCotacao,
		expectedHttpStatus: expectedHttpStatus,
		tScenario:          tScenario,
	}
}

func givenQuotationFoundWhenSearchUSDBRLQuotationThenShouldReturnsHttpStatusOkAndCotacaoJsonScenario(service quotationServiceMock) scenario {
	cotacaoResponse := ws.Cotacao{
		Code:       "USD",
		CodeIn:     "BRL",
		Name:       "Dólar Americano/Real Brasileiro",
		High:       "5.3519",
		Low:        "5.2603",
		VarBid:     "-0.021",
		PctChange:  "-0.4",
		Bid:        "5.3027",
		Ask:        "5.3052",
		Timestamp:  "1668458599",
		CreateDate: "2022-11-14 17:43:19",
	}
	return createScenario(
		"GivenQuotationFoundWhenSearchUSDToBRLQuotationThenShouldReturnsHttpStatusOkAndCotacaoJson",
		cotacaoResponse,
		http.StatusOK,
		func(t *testing.T) {
			service.On("GetUSDQuotation").Return(cotacaoResponse, nil)
			cotacaoApi, err := cotacao.NewController(service)
			if err != nil {
				t.Fatalf("erro ao instanciar service. err: %s", err.Error())
			}
			request, err := http.NewRequest(http.MethodGet, "/cotacao", nil)
			if err != nil {
				t.Fatalf("erro ao criar request. err: %s", err.Error())
			}
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(cotacaoApi.GetCotacaoUSD)
			handler.ServeHTTP(res, request)
			if http.StatusOK != res.Code {
				t.Errorf("erro ao chamar endpoint. Http Status Code: %d", res.Code)
			}
			var actual ws.Cotacao
			if err = json.NewDecoder(res.Body).Decode(&actual); err != nil {
				t.Fatalf("erro ao recuperar response body. err: %s", err.Error())
			}
			if len(strings.TrimSpace(actual.Code)) < 1 {
				t.Error("propriedade \"Code\" deve ser retornada.")
			}
		})
}

func givenErrorGettingQuotationWhenSearchUSDToBRLQuotationThenShouldReturnsHttpStatusServiceUnavailableScenario(service quotationServiceMock) scenario {
	cotacaoResponse := ws.Cotacao{}
	err := errors.New("erro ao buscar cotação usd para brl")
	return createScenario(
		"GivenErrorGettingQuotationWhenSearchUSDToBRLQuotationThenShouldReturnsHttpStatusInternalServerError",
		cotacaoResponse,
		http.StatusServiceUnavailable,
		func(t *testing.T) {
			service.On("GetUSDQuotation").Return(cotacaoResponse, err)
			cotacaoApi, err := cotacao.NewController(service)
			if err != nil {
				t.Fatalf("erro ao instanciar service. err: %s", err.Error())
			}
			request, err := http.NewRequest(http.MethodGet, "/cotacao", nil)
			if err != nil {
				t.Fatalf("erro ao criar request. err: %s", err.Error())
			}
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(cotacaoApi.GetCotacaoUSD)
			handler.ServeHTTP(res, request)
			if http.StatusServiceUnavailable != res.Code {
				t.Errorf("ao ocorrer um erro para buscar a cotação, deverá retornar o statu http 503. Http Status Code: %d", res.Code)
			}
			var actual ws.Cotacao
			if err = json.NewDecoder(res.Body).Decode(&actual); err == nil {
				t.Errorf("ao ocorrer um erro para buscar a cotação, não deverá retornar a estrutura json. err: %s", err.Error())
			}
		},
	)
}
