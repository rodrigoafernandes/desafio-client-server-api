# Client-Server-API

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

Será desenvolvido duas aplicações
- client.go
- server.go

Os requisitos para cumprir este desafio são:

O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.

Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.

O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Pacotes externos utilizados

- sqlite3
- envconfig

### Sqlite3
Driver para conexão com o banco de dados Sqlite

### envconfig
Pacote fornecido por kelseyhightower, o [envconfig](https://github.com/kelseyhightower/envconfig), para leitura das variáveis de ambiente determinadas e transformar em uma struct, facilitando e centralizando a configuração.<br/>
Foi necessário a utilização deste pacote, pois em meus testes, nenhuma das chamadas a api [Economia](https://economia.awesomeapi.com.br/json/last/USD-BRL), retornava em 200ms ou menos.<br/>
Foram utilizadas as variáveis de ambiente:
```
SERVER_ECONOMIA_WS_CLIENT_TIMEOUT_MS
SERVER_ECONOMIA_WS_CLIENT_URL
SERVER_DB_TRANSACTION_TIMEOUT_MS
SERVER_DB_CONNECTION_STRING
COTACAO_SERVER_URL
COTACAO_SERVER_PORT
COTACAO_SERVER_CLIENT_TIMEOUT_MS
COTACAO_ARQUIVO_OUTPUT_PATH
```
Exemplo de configuração das variáveis para utilização das aplicações
```shell
export SERVER_ECONOMIA_WS_CLIENT_TIMEOUT_MS=200
export SERVER_ECONOMIA_WS_CLIENT_URL=https://economia.awesomeapi.com.br
export SERVER_DB_TRANSACTION_TIMEOUT_MS=10
export SERVER_DB_CONNECTION_STRING=cotacoes

export COTACAO_SERVER_URL=http://localhost
export COTACAO_SERVER_PORT=8080
export COTACAO_SERVER_CLIENT_TIMEOUT_MS=300
export COTACAO_ARQUIVO_OUTPUT_PATH=results.txt
```
Caso não seja setado valores para as variáveis de ambiente, há valores default configurado na aplicação para o correto funcionamento