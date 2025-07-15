package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	v1 "github.com/otaviomart1ns/backend-challenge/internal/interface/api/v1"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	h := v1.NewHealthHandler()
	h.RegisterRoutes(router.Group("/"))

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"API RankMyApp rodando 🚀"}`, rec.Body.String())
}
