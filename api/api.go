package api

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/cotacao"
	"log"
	"net/http"
)

func Setup(quotationService cotacao.QuotationService) {
	controller, err := cotacao.NewController(quotationService)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/cotacao", controller.GetCotacaoUSD)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
