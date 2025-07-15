# Projeto RankMyApp - Golang

## 📌 Descrição

Este projeto é uma aplicação backend desenvolvida para o desafio técnico da **RankMyApp**, utilizando:

- Backend em **Golang (Gin)**
- Banco de dados **MySQL**
- Mensageria com **RabbitMQ**
- Orquestração com **Docker Compose**
- Geração de código tipado com **SQLC**
- Documentação com **Swagger**
- Scripts e automações com **Makefile**

---

## 🚀 Tecnologias Utilizadas

| Camada       | Tecnologias                                |
|--------------|---------------------------------------------|
| Backend      | Golang, Gin, SQLC, RabbitMQ, Swagger           |
| Banco de Dados | MySQL                               |
| DevOps       | Docker, Docker Compose, Makefile           |
| Testes       | Go Test                      |

---

## 📦 Estrutura do Projeto

```
desafio-tecnico-rankmyapp/
├── backend/
│   ├── cmd/api/
│   ├── docs
│   ├── internal/
│   │   ├── config/
│   │   ├── domain/
│   │   ├── infra/
│   │   ├── interface/
│   │   ├── usecase/
│   │   └── utils/
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── load_test
│   ├── artillery_test.yaml
│   ├── process.js
│   ├── report.html
│   └── report.json
├── .air.toml
├── docker-compose.yml
├── .env
├── .env.example
├── Makefile
├── insomnia_endpoints.yaml
├── populate.sql
├── sqlc.yaml
└── README.md
```

---

## ⚙️ Requisitos

- Docker e Docker Compose
- Make (GNU Make)

---

## 🔧 Configuração `.env`

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# Banco de Dados
MYSQL_HOST=mysql
MYSQL_PORT=3306
MYSQL_DATABASE=tickets_db
MYSQL_USER=tickets_user
MYSQL_PASSWORD=tickets_passwd
MYSQL_ROOT_PASSWORD=root

# RabbitMQ
RABBITMQ_HOST=rabbit
RABBITMQ_DEFAULT_USER=tickets_user
RABBITMQ_DEFAULT_PASS=tickets_passwd
RABBITMQ_MANAGEMENT_PORT=15672
RABBITMQ_PORT=5672

# API
API_PORT=8080
GIN_MODE=release
```

---

## 🛠️ Como Rodar o Projeto

### 1. Subir os containers

```bash
make build
make up
```
### 2. Rodar as migrações

```bash
make migrationup
```
### 3. Populate

```bash
make populate
```

### 4. Acessar os serviços

| Serviço       | URL                            |
|---------------|---------------------------------|
| API Backend   | http://localhost:8080/api/v1/          |
| Swagger (documentação)   | http://localhost:8080/swagger/index.html       |
| RabbitMQ UI   | http://localhost:15672 (user/pasword) |
| MySQL    | via cliente na porta 3306 (user/pasword)     |

---

## 🧪 Executar Testes (Unitários e de Integração)

### Backend - Cobertura 100%

```bash
make test-backend
```
---

## 📚 Endpoints da API

> Base URL: `http://localhost:8080/api/v1`

> Swagger URL: `http://localhost:8080/swagger/index.html`

1. **GET** `/`

Health Check Status.

##### Response
- **200 OK**

```json
{
	"message": "API RankMyApp rodando 🚀"
}
```

2. **POST** `/tickets/buy`

Comprar um ticket.

##### Body (JSON)

```json
{
    "ticketId": "<uuid_ticket>",
    "paymentType": "CREDIT_CARD",
    "userId": "4af7ebc2-288c-473c-9ae3-541c052ef2fa"
}
```

##### Response
- **201 Created**
```json
{
	"saleId": "5259fc07-10fe-416a-8004-c60b682f53fd",
	"ticketId": "50df0a9b-60fd-11f0-82a0-2e6782acc0d9",
	"userId": "1af7ebc2-288c-473c-9ae3-541c052ef2fa",
	"paymentType": "CREDIT_CARD"
}
```


3. **GET** `/tickets/catalog`

Listar catálogo de tickets.

##### Response
- **200 OK**

```json
[
	{
		"id": "50df0a9b-60fd-11f0-82a0-2e6782acc0d9",
		"type": "GENERAL_AREA",
		"price": 95,
		"quantity": 10
	},
	{
		"id": "531d0180-60fd-11f0-82a0-2e6782acc0d9",
		"type": "VIP",
		"price": 750,
		"quantity": 2
	},
	{
		"id": "5449cb96-60fd-11f0-82a0-2e6782acc0d9",
		"type": "GOLDEN_CIRCLE",
		"price": 1250,
		"quantity": 1
	},
	{
		"id": "5efe8975-60fd-11f0-82a0-2e6782acc0d9",
		"type": "GRANDSTAND",
		"price": 175,
		"quantity": 5
	}
]
```

---

## 📥 Makefile - Comandos Úteis

```bash
make up              # Sobe os containers
make down            # Derruba os containers
make build           # Builda todas as imagens
make db-up           # Sobe apenas o banco
make migrationup     # Aplica as migrações
make migrationdown   # Reverte a última migração
make sqlc            # Gera código Go com SQLC
make populate        # Popula o banco com dados padrões
make test-backend    # Executa os testes do backend
```

---
## 📊 Teste de Carga com Artillery

### Configuração do Teste

- **Duração total:** 2 minutos (120 segundos)
- **Fases simuladas:**
  - `30s` de rampa progressiva: começa com **20 usuários/segundo**, aumentando até **100 usuários/segundo**
  - `60s` com carga sustentada de **100 usuários/segundo**
  - `30s` de pico intenso com **150 usuários/segundo**

- **Objetivo:** Avaliar estabilidade sob escalada de carga realista e testar resistência em pico agressivo

- **Usuários virtuais simulados (estimado):** +10.000
- **Total de requisições esperadas:** ~24.600

- **Endpoints testados:**
  - `GET /api/v1/tickets/catalog`
  - `POST /api/v1/tickets/buy` com `ticketId` válido (dinamicamente selecionado)

- **Ambiente:** Localhost (Docker)

### Resultados Obtidos

- **Total de requisições simuladas:** 24.600 
- **Usuários virtuais:** +15.000 
- **Duração total do teste:** 2 minutos

#### 1. Distribuição de Respostas HTTP:
| Código | Descrição                         | Ocorrências |
|--------|------------------------------------|-------------|
| 200    | OK (`/api/v1/tickets/catalog`)     | 12.300       |
| 201    | Created (`/api/v1/tickets/buy`)    | 12.300       |
| ❌     | Falhas de execução                 | 0           |

#### ⏱️ Latência (Tempo de Resposta):
| Métrica  | Valor         |
|----------|---------------|
| Média    | **10 ms**    |
| P95      | **23 ms**      |
| P99      | **31 ms**      |
| Máxima   | **2,9 s**     |

> Observação: os picos de até 2.9s são outliers raros. 99.9% das requisições responderam em menos de 116 ms.

#### ⚡ Performance Geral:
- **Throughput médio:** ~195 requisições/segundo
- **Pico de RPS:** ~300
- **Apdex Score:** 0.99 (99,96% dos usuários satisfeitos)

#### 📌 Observações Técnicas:
- A aplicação demonstrou **excelente estabilidade** e **baixa latência mesmo sob carga alta**.
- Nenhuma requisição falhou — todos os fluxos de compra e consulta foram executados com sucesso.
- O backend se mostrou capaz de sustentar **picos de tráfego em cenário realista de estresse** com mais de 300 RPS sem degradação.
- Picos isolados de latência (≤ 0.05% das requisições) não comprometeram a experiência geral.

### [Clique aqui para ver o relatório interativo completo](https://app.artillery.io/share/sh_f0a0f4f9377f4f284eaafe30ed93dded91cb21b32a0e1684f0eb48aa10050ed0)
---

## 🐞 Problemas Comuns

### 1. Porta em uso

Altere as portas no `.env`.

---