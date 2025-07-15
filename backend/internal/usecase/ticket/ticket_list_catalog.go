package ticket

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
)

type ListCatalogUseCaseInterface interface {
	Execute(ctx context.Context) ([]*TicketOutput, error)
}

type ListCatalogUseCase struct {
	service services.TicketService
}

func NewListCatalogUseCase(service services.TicketService) *ListCatalogUseCase {
	return &ListCatalogUseCase{
		service: service,
	}
}

func (uc *ListCatalogUseCase) Execute(ctx context.Context) ([]*TicketOutput, error) {
	tickets, err := uc.service.ListCatalog(ctx)
	if err != nil {
		return nil, err
	}

	var result []*TicketOutput
	for _, t := range tickets {
		result = append(result, &TicketOutput{
			ID:       t.ID,
			Type:     string(t.Type),
			Price:    t.Price,
			Quantity: t.Quantity,
		})
	}

	return result, nil
}
