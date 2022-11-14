package output

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"os"
	"strings"
)

type ResultOutput interface {
	WriteQuotationResult(bid float64) error
}

type ResultOutputImpl struct {
	Filename string
}

func NewResultOutput(cfg config.ClientConfig) (ResultOutputImpl, error) {
	filename := cfg.ArquivoOutputPath
	if len(strings.TrimSpace(filename)) == 0 {
		filename = "results.txt"
	}
	ro := ResultOutputImpl{
		Filename: filename,
	}
	return ro, nil
}

func (ro ResultOutputImpl) WriteQuotationResult(bid float64) error {
	file, err := os.OpenFile(ro.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Dolar: %f\n", bid))
	return err
}
