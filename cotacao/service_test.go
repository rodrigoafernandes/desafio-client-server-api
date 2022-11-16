package cotacao

import (
	"errors"
	"github.com/rodrigoafernandes/desafio-client-server-api/ws"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

type EconomiaWSClientMock struct {
	mock.Mock
}

func (m EconomiaWSClientMock) GetUSDQuotationFromBRL() (ws.Cotacao, error) {
	args := m.Called()
	return args.Get(0).(ws.Cotacao), args.Error(1)
}

type RepositoryMock struct {
	mock.Mock
}

func (r RepositoryMock) Save(cotacao CotacaoDB) error {
	args := r.Called()
	return args.Error(0)
}

type TestingScenario func(t *testing.T)

type Scenario struct {
	name            string
	expectedCotacao ws.Cotacao
	expectedError   error
	tScenario       TestingScenario
}

func TestQuotationService(t *testing.T) {
	wsClient := EconomiaWSClientMock{}
	repo := RepositoryMock{}
	scenarios := []Scenario{
		GivenNoErrorsWhenGetUSDQuotationThenShouldReturnsCotacaoStructScenario(wsClient, repo),
		GivenHttpErrorWhenGetUSDQuotationThenShouldReturnsErrorScenario(wsClient, repo),
		GivenDBErrorWhenGetUSDQuotationThenShouldReturnsErrorScenario(wsClient, repo),
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, scenario.tScenario)
	}
}

func createScenario(name string, expectedCotacao ws.Cotacao, expectedError error, tScenario TestingScenario) Scenario {
	return Scenario{
		name:            name,
		expectedCotacao: expectedCotacao,
		expectedError:   expectedError,
		tScenario:       tScenario,
	}
}

func GivenNoErrorsWhenGetUSDQuotationThenShouldReturnsCotacaoStructScenario(wsClient EconomiaWSClientMock, repo RepositoryMock) Scenario {
	cotacao := ws.Cotacao{
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
	return createScenario("GivenNoErrorsWhenGetUSDQuotationThenShouldReturnsCotacaoStruct", cotacao, nil, func(t *testing.T) {
		cotacao := ws.Cotacao{
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
		wsClient.On("GetUSDQuotationFromBRL").Return(cotacao, nil)
		repo.On("Save", mock.Anything).Return(nil)
		svc, err := NewQuotationService(wsClient, repo)
		actualCotacao, err := svc.GetUSDQuotation()
		if err != nil {
			t.Errorf("erro ao buscar cotação. err: %s", err.Error())
		}
		if len(strings.TrimSpace(actualCotacao.Code)) < 1 {
			t.Errorf("erro ao buscar cotação. err: %s", err.Error())
		}
	})
}

func GivenHttpErrorWhenGetUSDQuotationThenShouldReturnsErrorScenario(wsClient EconomiaWSClientMock, repo RepositoryMock) Scenario {
	cotacao := ws.Cotacao{}
	err := errors.New("falha ao buscar cotacao na api")
	return createScenario("GivenHttpErrorWhenGetUSDQuotationThenShouldReturnsError", cotacao, err, func(t *testing.T) {
		wsClient.On("GetUSDQuotationFromBRL").Return(cotacao, err)
		svc, err := NewQuotationService(wsClient, repo)
		actualCotacao, err := svc.GetUSDQuotation()
		if err == nil {
			t.Errorf("ao buscar uma cotação e algum erro ocorrer na chamada da api, é necessário retornar um erro. err: %s", err.Error())
		}
		if len(strings.TrimSpace(actualCotacao.Code)) > 0 {
			t.Error("ao buscar uma cotação e algum erro ocorrer na chamada da api, não deveria retornar nenhum dado da cotação.")
		}
	})
}

func GivenDBErrorWhenGetUSDQuotationThenShouldReturnsErrorScenario(wsClient EconomiaWSClientMock, repo RepositoryMock) Scenario {
	cotacao := ws.Cotacao{
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
	err := errors.New("falha ao persistir a cotação no DB")
	return createScenario("GivenHttpErrorWhenGetUSDQuotationThenShouldReturnsError", cotacao, err, func(t *testing.T) {
		wsClient.On("GetUSDQuotationFromBRL").Return(cotacao, nil)
		repo.On("Save", mock.Anything).Return(err)
		svc, err := NewQuotationService(wsClient, repo)
		actualCotacao, err := svc.GetUSDQuotation()
		if err == nil {
			t.Errorf("ao buscar uma cotação e algum erro ocorrer na persistencia no DB, é necessário retornar um erro. err: %s", err.Error())
		}
		if len(strings.TrimSpace(actualCotacao.Code)) > 0 {
			t.Error("ao buscar uma cotação e algum erro ocorrer na persistencia no DB, não deveria retornar nenhum dado da cotação.")
		}
	})
}
