# ----------------------------------------------------------------------
# ESTÁGIO DE BUILD (CORRIGIDO E REINSERIDO)
# ----------------------------------------------------------------------
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copia dependências (para cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário 'app'
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o app ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

# ----------------------------------------------------------------------
# ESTÁGIO FINAL - IMAGEM MINIMALISTA
# ----------------------------------------------------------------------
FROM alpine:3.18

WORKDIR /app

# Instala dependências do Alpine
RUN apk add --no-cache libc6-compat

# 1. Copia o binário **(Agora o binário existe!)**
COPY --from=builder /app/app .

# 2. Garante que o binário é executável (como root)
RUN chmod +x /app/app

# 3. Cria e troca para o usuário não-root
RUN adduser -D appuser
USER appuser

EXPOSE 8000 8080 50051

CMD ["./app"]