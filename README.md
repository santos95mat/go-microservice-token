# go-microservice-token <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

## Gerador e validador de Token para um determinado usuário.

#### - Para construção deste microservice foram seguidos os padrões de SOLID
#### - usado GORM para conectar ao banco de dados
#### - banco de dados Postgres
#### - usado Fiber para criação do microservice


## .ENV

#### Para conseguir rodar a aplicação sem erros, você deve criar o arquivo .env com as seguintes variáveis

```
# Porta onde a API vai rodar como mostra o exemplo abaixo
PORT=3000

# URL de conexão com o banco de dados como mostra o exemplo abaixo
DB_URL="host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
```


## Rodando a aplicação

```bash
# baixar todas as dependencias do microservice
$ go mod tidy

# rodar a migrate primeiro para criação da entidade no banco de dados
$ go run cmd/migrate/main.go

# start
$ go run cmd/api/main.go
```

## JSON

#### Para criar um Token você deve mandar no body um JSON com a seguinte informação
##### - Se o id informado não for um uuid um erro será gerado
##### - email para onde o token será enviado
##### - validation é a validade do token em minutos

```JSON
{
  "user_id": "4cb16abd-be70-420d-82af-ecbf21c133a0",
  "email" : "email@email.com",
  "validation": 10
}
```

#### Para Validar o Token gerado você deve mandar no body um JSON com as seguintes informações
```JSON
{
  "user_id": "4cb16abd-be70-420d-82af-ecbf21c133a0",
  "token": "$r5D"
}
```
