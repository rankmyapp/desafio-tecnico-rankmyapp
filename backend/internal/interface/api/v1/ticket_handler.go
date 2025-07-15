package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/backend-challenge/internal/interface/api/response"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/ticket"
)

type TicketHandler struct {
	listCatalogUC ticket.ListCatalogUseCaseInterface
}

// NewTicketHandler cria uma nova instância de TicketHandler
func NewTicketHandler(listCatalogUC ticket.ListCatalogUseCaseInterface) *TicketHandler {
	return &TicketHandler{
		listCatalogUC: listCatalogUC,
	}
}

// RegisterRoutes registra as rotas relacionadas a tickets
func (h *TicketHandler) RegisterRoutes(rg *gin.RouterGroup) {
	tickets := rg.Group("/tickets")
	tickets.GET("/catalog", h.ListCatalog)
}

// ListCatalog godoc
// @Summary Lista o catálogo de tickets disponíveis
// @Description Retorna todos os ingressos cujo estoque seja maior que zero
// @Tags Tickets
// @Produce json
// @Success 200 {array} ticket.TicketOutput
// @Failure 500 {object} response.ErrorResponse "Erro interno ao buscar catálogo de tickets"
// @Router /tickets/catalog [get]
func (h *TicketHandler) ListCatalog(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.listCatalogUC.Execute(ctx)
	if err != nil {
		response.JSONError(c, http.StatusInternalServerError, "error listing catalog", err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
