ifneq (,$(wildcard ./.env))
  include .env
  export
endif

COMPOSE_FILE = docker-compose.yaml

.PHONY: help build up down db-up migrationup migrationdown sqlc populate test-backend

help:
	@echo "Targets: build, up, down, db-up, migrationup, migrationdown, sqlc, populate, test-backend"

build: ## Builda todas as imagens Docker
	docker compose -f $(COMPOSE_FILE) build

up: ## Sobe containers (detached)
	docker compose -f $(COMPOSE_FILE) up -d

down: ## Desliga e remove containers
	docker compose -f $(COMPOSE_FILE) down --remove-orphans

db-up: ## Sobe somente o serviço de banco
	docker compose -f $(COMPOSE_FILE) up -d mysql

migrationup: ## Aplica migrações do banco com MySQL
	migrate -path backend/internal/infra/db/migrations \
	  -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" \
	  -verbose up

migrationdown: ## Reverte migrações
	migrate -path backend/internal/infra/db/migrations \
	  -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" \
	  -verbose down

sqlc:
	docker run --rm -v $$(pwd):/src -w /src sqlc/sqlc generate

populate: ## Popula o banco com tickets de teste
	docker exec -i $(shell docker compose ps -q mysql) sh -c 'mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D $(MYSQL_DATABASE) -e "\
	INSERT INTO tickets (id, type, price, quantity) VALUES (UUID(), '\''GENERAL_AREA'\'', 95, 10); \
	INSERT INTO tickets (id, type, price, quantity) VALUES (UUID(), '\''GRANDSTAND'\'', 175, 5); \
	INSERT INTO tickets (id, type, price, quantity) VALUES (UUID(), '\''VIP'\'', 750, 2); \
	INSERT INTO tickets (id, type, price, quantity) VALUES (UUID(), '\''GOLDEN_CIRCLE'\'', 1250, 1);"'


test-backend: ## Roda testes unitários com cobertura
	@MYSQL_URL="mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" \
	RABBITMQ_URL="amqp://$(RABBITMQ_DEFAULT_USER):$(RABBITMQ_DEFAULT_PASS)@localhost:$(RABBITMQ_PORT)/" \
	API_PORT=$(API_PORT) \
	GIN_MODE=$(GIN_MODE) \
	cd backend && go test -v -coverprofile=coverage.out ./...