package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/otaviomart1ns/backend-challenge/docs"
	v1 "github.com/otaviomart1ns/backend-challenge/internal/interface/api/v1"
	"github.com/otaviomart1ns/backend-challenge/internal/interface/middleware"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/sale"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/ticket"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouteDependencies struct {
	ProcessSaleUC *sale.ProcessSaleUseCase
	ListCatalogUC *ticket.ListCatalogUseCase
}

// NewRouter configura e retorna o roteador principal da API
func NewRouter(deps RouteDependencies) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")

	// Healthcheck
	healthHandler := v1.NewHealthHandler()
	healthHandler.RegisterRoutes(api)

	// Catálogo de tickets
	ticketHandler := v1.NewTicketHandler(deps.ListCatalogUC)
	ticketHandler.RegisterRoutes(api)

	// Compra de tickets (venda)
	saleHandler := v1.NewSaleHandler(deps.ProcessSaleUC)
	saleHandler.RegisterRoutes(api)

	return router
}
