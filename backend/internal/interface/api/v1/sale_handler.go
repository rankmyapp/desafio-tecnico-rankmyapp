package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/backend-challenge/internal/interface/api/response"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/sale"
)

type SaleHandler struct {
	processSaleUC sale.ProcessSaleUseCaseInterface
}

// NewSaleHandler cria uma nova instância do SaleHandler
func NewSaleHandler(processSaleUC sale.ProcessSaleUseCaseInterface) *SaleHandler {
	return &SaleHandler{
		processSaleUC: processSaleUC,
	}
}

// RegisterRoutes registra as rotas de venda no grupo fornecido
func (h *SaleHandler) RegisterRoutes(rg *gin.RouterGroup) {
	sales := rg.Group("/tickets")
	sales.POST("/buy", h.BuyTicket)
}

// BuyTicket godoc
// @Summary Realiza a compra de um ticket
// @Description Registra uma venda e publica na fila de validação via RabbitMQ
// @Tags Sales
// @Accept  json
// @Produce  json
// @Param input body sale.ProcessSaleInput true "Informações da venda (somente pagamento com cartão de crédito)"
// @Success 201 {object} sale.ProcessSaleOutput
// @Failure 400 {object} response.ErrorResponse "Payload inválido (JSON malformado)"
// @Failure 422 {object} response.ErrorResponse "Erro ao processar venda (falta de estoque, ticket inexistente, etc)"
// @Router /tickets/buy [post]
func (h *SaleHandler) BuyTicket(c *gin.Context) {
	var input sale.ProcessSaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.JSONError(c, http.StatusBadRequest, "invalid JSON payload", err.Error())
		return
	}

	output, err := h.processSaleUC.Execute(c.Request.Context(), input)
	if err != nil {
		response.JSONError(c, http.StatusUnprocessableEntity, "could not process sale", err.Error())
		return
	}

	c.JSON(http.StatusCreated, output)
}
