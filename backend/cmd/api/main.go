package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/backend-challenge/internal/config"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/queue"
	dbRepo "github.com/otaviomart1ns/backend-challenge/internal/infra/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/interface/api"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/sale"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/ticket"

	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
)

// @title 			API - Desafio RankMyApp
// @version 		1.0
// @description 	Documentação da API REST para o desafio da RankMyApp
// @contact.name    Otavio Martins
// @contact.email   taviomartins01@gmail.com
// @license.name    MIT
// @host 			localhost:8080
// @BasePath 		/api/v1
// @schemes 		http
func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)

	// Conecta ao banco de dados
	db, err := sql.Open("mysql", cfg.MySQLURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	// Instancia os métodos SQL gerados pelo sqlc
	queries := models.New(db)

	// Conecta ao RabbitMQ
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Inicializa fila
	queuePublisher, err := queue.NewRabbitMQPublisher(conn, "validate-purchase")
	if err != nil {
		log.Fatalf("Erro ao inicializar publisher RabbitMQ: %v", err)
	}

	// Injeta os repositórios implementando a camada de persistência
	ticketRepo := dbRepo.NewTicketRepository(queries)
	saleRepo := dbRepo.NewSaleRepository(queries)

	// Instancia os serviços que encapsulam as regras de negócio
	ticketService := services.NewTicketService(ticketRepo)
	saleService := services.NewSaleService(ticketRepo, saleRepo, queuePublisher)

	// Injeta os casos de uso da aplicação, com base nas regras de negócio
	listCatalogUC := ticket.NewListCatalogUseCase(ticketService)
	processSaleUC := sale.NewProcessSaleUseCase(saleService)

	// Cria o roteador da API e injeta dependências nos handlers e registra as rotas
	router := api.NewRouter(api.RouteDependencies{
		ListCatalogUC: listCatalogUC,
		ProcessSaleUC: processSaleUC,
	})

	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
