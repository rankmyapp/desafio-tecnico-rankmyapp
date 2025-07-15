package mocks

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type TicketRepository struct {
	mock.Mock
}

func (m *TicketRepository) GetByID(ctx context.Context, id string) (*entities.Ticket, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Ticket), args.Error(1)
}

func (m *TicketRepository) Update(ctx context.Context, ticket *entities.Ticket) error {
	args := m.Called(ctx, ticket)
	return args.Error(0)
}

func (m *TicketRepository) GetAll(ctx context.Context) ([]*entities.Ticket, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entities.Ticket), args.Error(1)
}
