package services

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase"
)

// TicketService define o contrato do serviço de catálogo de ingressos
type TicketService interface {
	ListCatalog(ctx context.Context) ([]*entities.Ticket, error)
	GetByID(ctx context.Context, id string) (*entities.Ticket, error)
}

type ticketService struct {
	ticketRepo repositories.TicketRepository
}

// NewTicketService retorna uma instância concreta de TicketService
func NewTicketService(ticketRepo repositories.TicketRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
	}
}

// ListCatalog retorna todos os tickets disponíveis no catálogo
func (s *ticketService) ListCatalog(ctx context.Context) ([]*entities.Ticket, error) {
	tickets, err := s.ticketRepo.GetAll(ctx)
	if err != nil {
		return nil, usecase.ErrTicketListFailed
	}
	return tickets, nil
}

// GetByID retorna um ticket específico pelo seu ID
func (s *ticketService) GetByID(ctx context.Context, id string) (*entities.Ticket, error) {
	ticket, err := s.ticketRepo.GetByID(ctx, id)
	if err != nil {
		return nil, usecase.ErrTicketNotFound
	}
	return ticket, nil
}
