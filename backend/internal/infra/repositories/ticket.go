package repositories

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/converters"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

type ticketRepository struct {
	queries *models.Queries
}

func NewTicketRepository(q *models.Queries) repositories.TicketRepository {
	return &ticketRepository{queries: q}
}

func (r *ticketRepository) Update(ctx context.Context, ticket *entities.Ticket) error {
	params := converters.ToSQLCUpdateTicketQuantityParams(ticket)
	return r.queries.UpdateTicketQuantity(ctx, params)
}

func (r *ticketRepository) GetAll(ctx context.Context) ([]*entities.Ticket, error) {
	tickets, err := r.queries.GetAllTickets(ctx)
	if err != nil {
		return nil, err
	}
	return converters.ToTicketEntityList(tickets), nil
}

func (r *ticketRepository) GetByID(ctx context.Context, id string) (*entities.Ticket, error) {
	ticket, err := r.queries.GetTicketByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return converters.ToTicketEntity(ticket), nil
}

// Garante em tempo de compilação que a implementação satisfaz a interface
var _ repositories.TicketRepository = (*ticketRepository)(nil)
