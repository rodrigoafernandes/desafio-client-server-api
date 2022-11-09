package output

import (
	"fmt"
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"os"
)

type ResultOutput struct {
	Filename string
}

func NewResultOutput(cfg config.ClientConfig) (ResultOutput, error) {
	filename := cfg.ArquivoOutputPath
	ro := ResultOutput{Filename: filename}
	return ro, nil
}

func (ro ResultOutput) WriteQuotationResult(bid float64) error {
	file, err := os.OpenFile(ro.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("Dolar: %f\n", bid))
	return err
}
