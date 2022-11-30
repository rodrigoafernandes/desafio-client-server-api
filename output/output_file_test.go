package output

import (
	"github.com/rodrigoafernandes/desafio-client-server-api/config"
	"os"
	"testing"
)

func TestWriteBid(t *testing.T) {
	cfg := config.ClientConfig{
		ArquivoOutputPath: "",
	}
	bid := 1.1
	resultOutput, _ := NewResultOutput(cfg)
	err := resultOutput.WriteQuotationResult(bid)
	if err != nil {
		t.Errorf("nenhum erro deve ser retornado quando não ocorrer um erro ao gravar o resultado no arquivo. err: %s", err.Error())
	}
	err = os.Remove("results.txt")
	if err != nil {
		t.Fatalf("erro ao excluir o arquivo de teste. err: %s", err.Error())
	}
}

func TestErrorOpenFileToWrite(t *testing.T) {
	filename := "https://server-file.test.com/file"
	cfg := config.ClientConfig{
		ArquivoOutputPath: filename,
	}
	bid := 1.1
	resultOutput, _ := NewResultOutput(cfg)
	err := resultOutput.WriteQuotationResult(bid)
	if err == nil {
		t.Error("deve retornar erro ao não conseguir abrir o arquivo para gravar o bid.")
	}
}
