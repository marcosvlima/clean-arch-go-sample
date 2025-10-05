# clean-arch-go-sample

Este repositório é um exemplo de arquitetura limpa em Go com exposição das features via Web (REST), GraphQL e gRPC.

## Como executar (Docker Compose)

As instruções abaixo iniciam o banco (MySQL), RabbitMQ e a aplicação usando o `docker-compose.yml` presente na raiz do projeto.

1. Construa e inicie os serviços:

```bash
docker-compose up --build
```

2. Ports expostas (padrão):

- Web (REST): http://localhost:8000
- GraphQL Playground: http://localhost:8080 (Playground em `/`, endpoint `/query`)
- gRPC: localhost:50051

3. Para parar e remover os containers:

```bash
docker-compose down
```

## Testando a feature ListOrders

Você pode testar a feature `ListOrders` por três interfaces. Abaixo estão os passos e exemplos.

### 1) REST (via Web)

- Abra `api/list_eorders.http` no VS Code e use a extensão REST Client (clique em "Send Request").
- Exemplos nel arquivo:

```http
GET http://localhost:8000/list?page=1&limit=5 HTTP/1.1
Content-Type: application/json

###

GET http://localhost:8000/list?page=2&limit=5 HTTP/1.1
Content-Type: application/json
```

Ou use curl:

```bash
curl "http://localhost:8000/list?page=1&limit=5"
```

Resposta esperada (exemplo):

```json
{
	"orders": [
		{"id":"ord-1","price":100,"tax":1,"final_price":101}
	],
	"total": 1
}
```

### 2) GraphQL (Playground)

1. Abra http://localhost:8080/ no seu navegador.
2. Use a query abaixo:

```graphql
query ListOrders($page: Int, $limit: Int) {
	ListOrders(page: $page, limit: $limit) {
		Id
		Price
		Tax
		FinalPrice
	}
}

# Variables
{
	"page": 1,
	"limit": 10
}
```

### 3) gRPC (evans)

1. Instale `evans` (https://github.com/ktr0731/evans).
2. Conecte e chame o RPC `ListOrders`:

```bash
evans -r -p 50051
# dentro do evans: escolha o serviço pb.OrderService e execute ListOrders
```

Payload de exemplo (no evans):

```
{
	"page": 1,
	"limit": 10
}
```

## Troubleshooting rápido

- Se o container `app` reiniciar continuamente, verifique os logs:

```bash
docker-compose logs app --follow
```

- Erros comuns:
	- Conexão com RabbitMQ (certifique-se que o serviço rabbitmq está healthy)
	- Erros de DB (verifique credenciais em `cmd/ordersystem/.env`)
	- Panic em GraphQL se os resolvers não estiverem implementados (procure por mensagens `not implemented` nos logs)

---

Para documentação detalhada de cada interface veja `docs/CREATE_ORDER.md` e `docs/LIST_ORDERS.md`.

