package sale_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/sale"
)

// Mock para o SaleService
type mockSaleService struct {
	mock.Mock
}

func (m *mockSaleService) ProcessSale(ctx context.Context, input services.ProcessSaleInput) (*services.ProcessSaleOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Get(0).(*services.ProcessSaleOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestProcessSaleUseCase_Execute_Success(t *testing.T) {
	mockService := new(mockSaleService)
	useCase := sale.NewProcessSaleUseCase(mockService)

	input := sale.ProcessSaleInput{
		TicketID:    "ticket1",
		UserID:      "user1",
		PaymentType: "CREDIT_CARD",
	}

	expectedServiceInput := services.ProcessSaleInput{
		TicketID:    "ticket1",
		UserID:      "user1",
		PaymentType: "CREDIT_CARD",
	}

	expectedServiceOutput := &services.ProcessSaleOutput{
		SaleID:      "sale1",
		TicketID:    "ticket1",
		UserID:      "user1",
		PaymentType: "CREDIT_CARD",
	}

	mockService.
		On("ProcessSale", mock.Anything, expectedServiceInput).
		Return(expectedServiceOutput, nil)

	output, err := useCase.Execute(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, "sale1", output.SaleID)
	assert.Equal(t, "ticket1", output.TicketID)
	assert.Equal(t, "user1", output.UserID)
	assert.Equal(t, "CREDIT_CARD", output.PaymentType)

	mockService.AssertExpectations(t)
}

func TestProcessSaleUseCase_Execute_Error(t *testing.T) {
	mockService := new(mockSaleService)
	useCase := sale.NewProcessSaleUseCase(mockService)

	input := sale.ProcessSaleInput{
		TicketID:    "invalid",
		UserID:      "user1",
		PaymentType: "CREDIT_CARD",
	}

	mockService.
		On("ProcessSale", mock.Anything, mock.Anything).
		Return(nil, errors.New("failed to process"))

	output, err := useCase.Execute(context.Background(), input)

	assert.Nil(t, output)
	assert.EqualError(t, err, "failed to process")
	mockService.AssertExpectations(t)
}
