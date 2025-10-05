# ListOrders - How to run and test

This document explains how to run and ListOrders feature and test it through the three available interfaces:

- Web (HTTP REST) - using the provided `api/list_eorders.http` file with the VS Code REST Client extension
- GraphQL - using the GraphQL Playground
- gRPC - using `evans` (gRPC CLI)

Prerequisites

- The application and its dependencies (MySQL and RabbitMQ) running via `docker-compose` or locally with compatible services.
- Optional: VS Code with REST Client extension, `evans` for gRPC.

Running the app

Start the full stack with docker-compose:

```bash
docker-compose up --build
```

Ports exposed by the stack (defaults from the repository):

- Web server (REST): localhost:8000
- GraphQL Playground: localhost:8080 (Playground on `/`, GraphQL endpoint `/query`)
- gRPC server: localhost:50051

Web (REST) - quick test using VS Code REST Client

1. Open `api/list_eorders.http` in this repository.
2. Use the REST Client extension (click "Send Request" above each snippet) to run the example.

Example requests in `api/list_eorders.http`:

```http
GET http://localhost:8000/list?page=1&limit=5 HTTP/1.1
Content-Type: application/json

###

GET http://localhost:8000/list?page=2&limit=5 HTTP/1.1
Content-Type: application/json
```

The web server handler returns the following JSON shape (OrderListOutputDTO):

```json
{
  "orders": [
    {"id":"ord-1","price":100,"tax":1,"final_price":101},
    ...
  ],
  "total": 123
}
```

GraphQL - using Playground

1. Open http://localhost:8080/ in your browser.
2. Use the following query in the Playground to list orders:

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

gRPC - using evans

1. Install `evans` if needed.
2. Start an interactive session connected to the server:

```bash
evans -r -p 50051
```

3. Call the `ListOrders` RPC with a payload like:

```
{
  "page": 1,
  "limit": 10
}
```

The gRPC `ListOrders` response contains a repeated `CreateOrderResponse` (used as output DTO) and an integer `total`.

---
Generated on 2025-10-05
