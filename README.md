# DESAFIO RANKMYAPP

API REST em Go (Gin) com:
- Autenticação JWT (signup/login)
- CRUD de usuários (listar, obter, atualizar, deletar), paginação e filtro por nome/email
- Publicação de mensagens no RabbitMQ no signup
- Worker consumidor que envia e-mails via SMTP (MailHog em dev)
- GORM + Postgres, Docker Compose, testes unitários

## Requisitos
- Go 1.22+
- Docker e Docker Compose

## Como rodar com Docker
1. Copie `.env.example` para `.env` (opcional)
2. Suba os serviços:
    # Go Gin JWT API + RabbitMQ + Worker (SMTP)

API REST em Go (Gin) com:
- Autenticação JWT (signup/login)
- CRUD de usuários (listar, obter, atualizar, deletar), paginação e filtro por nome/email
- Publicação de mensagens no RabbitMQ no signup
- Worker consumidor que envia e-mails via SMTP (MailHog em dev)
- GORM + Postgres, Docker Compose, testes unitários

## Requisitos
- Go 1.22+
- Docker e Docker Compose

## Como rodar com Docker
1. Copie `.env.example` para `.env` (opcional)
2. Suba os serviços:
docker compose up -d --build

3. API: http://localhost:8080
4. RabbitMQ UI: http://localhost:15672 (user: guest / pass: guest)
5. MailHog UI: http://localhost:8025

## Rodar local sem Docker
1. Suba Postgres, RabbitMQ e MailHog localmente
2. Exporte variáveis de ambiente (ou use `.env`)
3. Instale dependências: `go mod tidy`
4. API: `go run ./cmd/api`
5. Worker: `go run ./cmd/worker`

## Endpoints
- POST /api/v1/auth/signup
- POST /api/v1/auth/login
- GET /api/v1/users (auth)
- GET /api/v1/users/:id (auth)
- PUT /api/v1/users/:id (auth)
- DELETE /api/v1/users/:id (auth)

Authorization: Bearer <token>

Exemplos cURL:
curl -s -X GET
localhost:8080/health
 curl -s -X POST 
localhost:8080/api/v1/signup
-H "Content-Type: application/json"
-d '{"name":"Kaue","email":"kaue@example.com","password":"StrongPass123!"}'

## Testes
go test ./... -v


## Notas
- Em dev, MailHog captura e-mails (UI em http://localhost:8025).
- Ajuste JWT_EXPIRES_IN (ex.: "1h", "30m").