package repositories

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
)

// SaleRepository define as operações de persistência para a entidade Sale
type SaleRepository interface {
	Create(ctx context.Context, sale *entities.Sale) error
	GetByID(ctx context.Context, id string) (*entities.Sale, error)
}
