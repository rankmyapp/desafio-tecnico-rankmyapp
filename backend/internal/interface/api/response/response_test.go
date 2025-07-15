package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/backend-challenge/internal/interface/api/response"
	"github.com/stretchr/testify/assert"
)

func TestJSONError(t *testing.T) {
	// Define modo de teste do gin para não poluir stdout
	gin.SetMode(gin.TestMode)

	// Cria um contexto gin com um ResponseRecorder (buffer)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Chama a função que será testada
	response.JSONError(c, http.StatusBadRequest, "invalid input", "missing field: name")

	// Verifica o status HTTP
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Decodifica a resposta JSON
	var resp response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Verifica o conteúdo retornado
	assert.Equal(t, "invalid input", resp.Message)
	assert.Equal(t, "missing field: name", resp.Details)
}
