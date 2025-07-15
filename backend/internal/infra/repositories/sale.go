package repositories

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/converters"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

type saleRepository struct {
	queries *models.Queries
}

func NewSaleRepository(q *models.Queries) repositories.SaleRepository {
	return &saleRepository{queries: q}
}

func (r *saleRepository) Create(ctx context.Context, sale *entities.Sale) error {
	params := converters.ToSQLCCreateSaleParams(sale)
	return r.queries.CreateSale(ctx, params)
}

func (r *saleRepository) GetByID(ctx context.Context, id string) (*entities.Sale, error) {
	sale, err := r.queries.GetSaleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return converters.ToSaleEntity(sale), nil
}

// Garante em tempo de compilação que a implementação satisfaz a interface
var _ repositories.SaleRepository = (*saleRepository)(nil)
