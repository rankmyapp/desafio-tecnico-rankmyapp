package repositories

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
)

// TicketRepository define as operações de persistência para a entidade Ticket
type TicketRepository interface {
	Update(ctx context.Context, ticket *entities.Ticket) error
	GetByID(ctx context.Context, id string) (*entities.Ticket, error)
	GetAll(ctx context.Context) ([]*entities.Ticket, error)
}
