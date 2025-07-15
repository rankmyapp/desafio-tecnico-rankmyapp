package v1_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/otaviomart1ns/backend-challenge/internal/interface/api/v1"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/sale"
)

// Mock do usecase
type mockProcessSaleUC struct {
	mock.Mock
}

func (m *mockProcessSaleUC) Execute(ctx context.Context, input sale.ProcessSaleInput) (*sale.ProcessSaleOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Get(0).(*sale.ProcessSaleOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

func setupRouterWithSaleHandler(mockUC *mockProcessSaleUC) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := v1.NewSaleHandler(mockUC)
	api := router.Group("/")
	handler.RegisterRoutes(api)

	return router
}

func TestSaleHandler_BuyTicket_Success(t *testing.T) {
	mockUC := new(mockProcessSaleUC)
	router := setupRouterWithSaleHandler(mockUC)

	inputJSON := `{"ticketId":"abc","userId":"xyz","paymentType":"CREDIT_CARD"}`

	expectedOutput := &sale.ProcessSaleOutput{
		SaleID:      "s1",
		TicketID:    "abc",
		UserID:      "xyz",
		PaymentType: "CREDIT_CARD",
	}

	mockUC.
		On("Execute", mock.Anything, sale.ProcessSaleInput{
			TicketID:    "abc",
			UserID:      "xyz",
			PaymentType: "CREDIT_CARD",
		}).
		Return(expectedOutput, nil)

	req, _ := http.NewRequest(http.MethodPost, "/tickets/buy", bytes.NewBufferString(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"saleId":"s1","ticketId":"abc","userId":"xyz","paymentType":"CREDIT_CARD"}`, rec.Body.String())
}

func TestSaleHandler_BuyTicket_InvalidJSON(t *testing.T) {
	mockUC := new(mockProcessSaleUC)
	router := setupRouterWithSaleHandler(mockUC)

	req, _ := http.NewRequest(http.MethodPost, "/tickets/buy", bytes.NewBufferString(`{invalid-json}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid JSON payload")
}

func TestSaleHandler_BuyTicket_ServiceError(t *testing.T) {
	mockUC := new(mockProcessSaleUC)
	router := setupRouterWithSaleHandler(mockUC)

	inputJSON := `{"ticketId":"abc","userId":"xyz","paymentType":"CREDIT_CARD"}`

	mockUC.
		On("Execute", mock.Anything, sale.ProcessSaleInput{
			TicketID:    "abc",
			UserID:      "xyz",
			PaymentType: "CREDIT_CARD",
		}).
		Return(nil, errors.New("out of stock"))

	req, _ := http.NewRequest(http.MethodPost, "/tickets/buy", bytes.NewBufferString(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	assert.Contains(t, rec.Body.String(), "could not process sale")
}
