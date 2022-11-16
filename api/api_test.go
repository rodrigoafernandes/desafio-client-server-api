package api

import (
	"encoding/json"
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
	service.On("GetUSDQuotation").Return(cotacaoResponse, nil)
	cotacaoApi, err := cotacao.NewController(service)
	if err != nil {
		t.Fatalf("erro ao instanciar service. err: %s", err.Error())
	}
	request, err := http.NewRequest(http.MethodGet, "/cotacao", nil)
	if err != nil {
		t.Fatalf("erro ao instanciar service. err: %s", err.Error())
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
}

func createScenario(name string, expectedCotacao ws.Cotacao, expectedHttpStatus int, tScenario testingScenario) scenario {
	return scenario{
		name:               name,
		expectedCotacao:    expectedCotacao,
		expectedHttpStatus: expectedHttpStatus,
		tScenario:          tScenario,
	}
}
