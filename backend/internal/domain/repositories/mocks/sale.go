package mocks

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type SaleRepository struct {
	mock.Mock
}

func (m *SaleRepository) Create(ctx context.Context, sale *entities.Sale) error {
	args := m.Called(ctx, sale)
	return args.Error(0)
}

func (m *SaleRepository) GetByID(ctx context.Context, id string) (*entities.Sale, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Sale), args.Error(1)
}
