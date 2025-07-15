package ticket_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase/ticket"
)

// Mock para TicketService
type mockTicketService struct {
	mock.Mock
}

// GetByID implements services.TicketService.
func (m *mockTicketService) GetByID(ctx context.Context, id string) (*entities.Ticket, error) {
	panic("unimplemented")
}

func (m *mockTicketService) ListCatalog(ctx context.Context) ([]*entities.Ticket, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Ticket), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestListCatalogUseCase_Execute_Success(t *testing.T) {
	mockService := new(mockTicketService)
	useCase := ticket.NewListCatalogUseCase(mockService)

	expectedTickets := []*entities.Ticket{
		{
			ID:       "1",
			Type:     entities.TicketType("VIP"),
			Price:    150.0,
			Quantity: 5,
		},
		{
			ID:       "2",
			Type:     entities.TicketType("STANDARD"),
			Price:    80.0,
			Quantity: 10,
		},
	}

	mockService.
		On("ListCatalog", mock.Anything).
		Return(expectedTickets, nil)

	result, err := useCase.Execute(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)

	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "VIP", result[0].Type)
	assert.Equal(t, 150.0, result[0].Price)
	assert.Equal(t, 5, result[0].Quantity)

	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, "STANDARD", result[1].Type)
	assert.Equal(t, 80.0, result[1].Price)
	assert.Equal(t, 10, result[1].Quantity)

	mockService.AssertExpectations(t)
}

func TestListCatalogUseCase_Execute_Error(t *testing.T) {
	mockService := new(mockTicketService)
	useCase := ticket.NewListCatalogUseCase(mockService)

	mockService.
		On("ListCatalog", mock.Anything).
		Return(nil, errors.New("service failure"))

	result, err := useCase.Execute(context.Background())

	assert.Nil(t, result)
	assert.EqualError(t, err, "service failure")
	mockService.AssertExpectations(t)
}
