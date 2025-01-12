# Client-Server API in Go

Este projeto implementa um sistema client-server em Go que consulta a cotação do dólar, armazena os dados em um banco de dados SQLite e grava o valor em um arquivo.

## Descrição do Projeto

O sistema consiste em dois componentes principais:

1. **Servidor (`server.go`)**:

   - Exibe um endpoint HTTP na rota `/cotacao` (porta `8080`).
   - Consome a API de cotação de dólar: [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL).
   - Armazena as cotações no banco de dados SQLite.
   - Implementa timeouts usando o pacote `context`:
     - 200ms para chamar a API de cotação.
     - 10ms para salvar no banco de dados.

2. **Cliente (`client.go`)**:
   - Faz uma requisição HTTP ao servidor para obter a cotação.
   - Recebe apenas o valor do campo `bid` da resposta JSON.
   - Grava o valor recebido em um arquivo `cotacao.txt` no formato:
     ```
     Dólar: {valor}
     ```
   - Implementa um timeout de 300ms para comunicação com o servidor.

## Requisitos

- **Dependências do Go**:
  - [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) para manipulação do banco SQLite.

## Estrutura do Projeto

```bash
client-server-api-go/
├── client/
│   └── client.go # Código do cliente
├── server/
│   └── server.go # Código do servidor
├── go.mod # Gerenciamento de dependências
├── go.sum # Lockfile das dependências
├── cotacao.txt # Arquivo de saída com a cotação do dólar
├── quotes.db # Banco de dados SQLite
└── README.md # Documentação do projeto

```

## Como Executar

#### 1. Clone este repositório:

```bash
git clone git@github.com:AugustoSandim/client-server-api-go.git
cd client-server-api-go
```

#### 2. Instale as dependências:

```bash
go mod tidy
```

#### 3. Execute o servidor:

```bash
go run server/server.go
```

#### 4. Em outro terminal, execute o cliente:

```bash
go run client/client.go
```

#### 5. Verifique o arquivo `cotacao.txt` gerado:

> cotacao.txt: Contém o valor atual da cotação do dólar.
> cotacoes.db: Banco de dados SQLite com os registros das cotações.

```bash
cat cotacao.txt
```

## Configuração do Banco de Dados

O banco de dados SQLite é criado automaticamente no diretório raiz do projeto como `cotacoes.db`. A tabela utilizada para armazenar as cotações é definida no código do servidor.

#### Estrutura da Tabela

```sql
CREATE TABLE IF NOT EXISTS quotes (
    id INTEGER PRIMARY KEY,
    bid TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```
