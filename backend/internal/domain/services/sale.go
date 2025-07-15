package services

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase"
)

// SaleService define o contrato do serviço de vendas
type SaleService interface {
	ProcessSale(ctx context.Context, input ProcessSaleInput) (*ProcessSaleOutput, error)
}

// ProcessSaleInput representa a entrada para processar uma venda
type ProcessSaleInput struct {
	TicketID    string
	UserID      string
	PaymentType string
}

// ProcessSaleOutput representa a saída de uma venda processada
type ProcessSaleOutput struct {
	SaleID      string
	TicketID    string
	UserID      string
	PaymentType string
}

type saleService struct {
	ticketRepo repositories.TicketRepository
	saleRepo   repositories.SaleRepository
	queue      repositories.QueuePublisher
}

// NewSaleService retorna uma instância do serviço de vendas
func NewSaleService(
	ticketRepo repositories.TicketRepository,
	saleRepo repositories.SaleRepository,
	queue repositories.QueuePublisher,
) *saleService {
	return &saleService{
		ticketRepo: ticketRepo,
		saleRepo:   saleRepo,
		queue:      queue,
	}
}

func (s *saleService) ProcessSale(ctx context.Context, input ProcessSaleInput) (*ProcessSaleOutput, error) {
	// 1. Validar método de pagamento
	if strings.ToUpper(input.PaymentType) != string(entities.CreditCard) {
		return nil, usecase.ErrInvalidPayment
	}

	// 2. Buscar ticket
	ticket, err := s.ticketRepo.GetByID(ctx, input.TicketID)
	if err != nil {
		return nil, usecase.ErrTicketNotFound
	}

	// 3. Verificar disponibilidade
	if ticket.Quantity <= 0 {
		return nil, usecase.ErrOutOfStock
	}

	// 4. Criar venda
	sale := &entities.Sale{
		ID:        uuid.NewString(),
		TicketID:  ticket.ID,
		UserID:    input.UserID,
		Payment:   entities.CreditCard,
		CreatedAt: time.Now(),
	}

	err = s.saleRepo.Create(ctx, sale)
	if err != nil {
		return nil, usecase.ErrSaleCreation
	}

	// 5. Atualizar estoque (reduzir quantidade)
	ticket.Quantity--
	err = s.ticketRepo.Update(ctx, ticket)
	if err != nil {
		return nil, usecase.ErrStockUpdate
	}

	// 6. Publicar na fila
	err = s.queue.PublishPurchaseValidation(ctx, sale)
	if err != nil {
		return nil, usecase.ErrQueuePublish
	}

	// 7. Retornar resultado
	return &ProcessSaleOutput{
		SaleID:      sale.ID,
		TicketID:    sale.TicketID,
		UserID:      sale.UserID,
		PaymentType: string(sale.Payment),
	}, nil
}
