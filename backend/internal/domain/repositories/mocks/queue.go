package mocks

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type QueuePublisher struct {
	mock.Mock
}

func (m *QueuePublisher) PublishPurchaseValidation(ctx context.Context, sale *entities.Sale) error {
	args := m.Called(ctx, sale)
	return args.Error(0)
}
