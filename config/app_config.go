package config

import "github.com/kelseyhightower/envconfig"

type ClientConfig struct {
	CotacaoServerUrl                       string `envconfig:"SERVER_URL"`
	CotacaoServerPort                      string `envconfig:"SERVER_PORT"`
	CotacaoServerClientTimeoutMilliseconds int    `envconfig:"SERVER_CLIENT_TIMEOUT_MS"`
	ArquivoOutputPath                      string `envconfig:"ARQUIVO_OUTPUT_PATH"`
}

type ServerConfig struct {
	EconomiaWSUrl                    string `envconfig:"ECONOMIA_WS_CLIENT_URL"`
	EconomiaWSTimeoutMilliseconds    int    `envconfig:"ECONOMIA_WS_CLIENT_TIMEOUT_MS"`
	DbConnectionString               string `envconfig:"DB_CONNECTION_STRING"`
	DbTransactionTimeoutMilliseconds int    `envconfig:"DB_TRANSACTION_TIMEOUT_MS"`
}

var ClientCFG ClientConfig
var ServerCFG ServerConfig

func SetupClient() {
	var clientConfig ClientConfig
	if err := envconfig.Process("COTACAO", &clientConfig); err != nil {
		panic(err)
	}
	ClientCFG = clientConfig
}

func SetupServer() {
	var serverConfig ServerConfig
	if err := envconfig.Process("SERVER", &serverConfig); err != nil {
		panic(err)
	}
	ServerCFG = serverConfig
}
