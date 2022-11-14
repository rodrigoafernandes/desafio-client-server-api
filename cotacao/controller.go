package cotacao

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	svc QuotationService
}

func NewController(quotationService QuotationService) (Controller, error) {
	controller := Controller{
		svc: quotationService,
	}
	return controller, nil
}

func (c Controller) GetCotacaoUSD(w http.ResponseWriter, r *http.Request) {
	cotacao, err := c.svc.GetUSDQuotation()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(cotacao); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
