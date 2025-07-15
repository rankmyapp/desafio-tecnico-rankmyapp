package converters

import (
	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

// ToTicketEntity converte o modelo SQLC para uma entidade de domínio
func ToTicketEntity(t models.Ticket) *entities.Ticket {
	return &entities.Ticket{
		ID:       t.ID,
		Type:     entities.TicketType(t.Type),
		Price:    t.Price,
		Quantity: int(t.Quantity),
	}
}

// ToTicketEntityList converte a lista SQLC de tickets para uma lista de entidades
func ToTicketEntityList(tickets []models.Ticket) []*entities.Ticket {
	result := make([]*entities.Ticket, 0, len(tickets))
	for _, t := range tickets {
		result = append(result, ToTicketEntity(t))
	}
	return result
}

// ToSQLCUpdateTicketQuantityParams converte a entidade do domínio para os parâmetros da query UpdateTicketQuantity
func ToSQLCUpdateTicketQuantityParams(t *entities.Ticket) models.UpdateTicketQuantityParams {
	return models.UpdateTicketQuantityParams{
		ID:       t.ID,
		Quantity: int32(t.Quantity),
	}
}
