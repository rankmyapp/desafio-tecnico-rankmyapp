package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler representa o handler responsável por tratar requisições HTTP relacionadas ao recurso verificação de status (healthcheck)
type HealthHandler struct{}

// NewHealthHandler instancia o handler responsável pelo healthcheck.
// Como não há dependências, retorna um struct vazio.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// RegisterRoutes registra as rotas relacionadas ao recurso healthcheck
func (h *HealthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", h.HealthCheck)
}

// HealthCheck godoc
// @Summary Verifica o status da API
// @Description Endpoint de health check para confirmar se a API está rodando
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API RankMyApp rodando 🚀",
	})
}
