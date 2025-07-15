package v1_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v1 "github.com/otaviomart1ns/backend-challenge/internal/interface/api/v1"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/ticket"
)

// Mock do usecase
type mockListCatalogUC struct {
	mock.Mock
}

func (m *mockListCatalogUC) Execute(ctx context.Context) ([]*ticket.TicketOutput, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*ticket.TicketOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

func setupRouterWithTicketHandler(mockUC *mockListCatalogUC) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := v1.NewTicketHandler(mockUC)
	api := router.Group("/")
	handler.RegisterRoutes(api)

	return router
}

func TestTicketHandler_ListCatalog_Success(t *testing.T) {
	mockUC := new(mockListCatalogUC)
	router := setupRouterWithTicketHandler(mockUC)

	expected := []*ticket.TicketOutput{
		{ID: "t1", Type: "VIP", Price: 100.0, Quantity: 10},
		{ID: "t2", Type: "Standard", Price: 50.0, Quantity: 20},
	}

	mockUC.On("Execute", mock.Anything).Return(expected, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tickets/catalog", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `
	[
		{"id":"t1","type":"VIP","price":100.0,"quantity":10},
		{"id":"t2","type":"Standard","price":50.0,"quantity":20}
	]`, rec.Body.String())
}

func TestTicketHandler_ListCatalog_Error(t *testing.T) {
	mockUC := new(mockListCatalogUC)
	router := setupRouterWithTicketHandler(mockUC)

	mockUC.On("Execute", mock.Anything).Return(nil, errors.New("database error"))

	req, _ := http.NewRequest(http.MethodGet, "/tickets/catalog", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "error listing catalog")
}
