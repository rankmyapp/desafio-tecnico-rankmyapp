package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories/mocks"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase"
)

func TestProcessSale_Success(t *testing.T) {
	ctx := context.TODO()

	ticketRepo := new(mocks.TicketRepository)
	saleRepo := new(mocks.SaleRepository)
	queue := new(mocks.QueuePublisher)

	ticket := &entities.Ticket{
		ID:       "ticket123",
		Type:     entities.VIP,
		Price:    150.0,
		Quantity: 5,
	}

	ticketRepo.On("GetByID", ctx, "ticket123").Return(ticket, nil)
	saleRepo.On("Create", ctx, mock.AnythingOfType("*entities.Sale")).Return(nil)
	ticketRepo.On("Update", ctx, mock.AnythingOfType("*entities.Ticket")).Return(nil)
	queue.On("PublishPurchaseValidation", ctx, mock.AnythingOfType("*entities.Sale")).Return(nil)

	service := services.NewSaleService(ticketRepo, saleRepo, queue)
	input := services.ProcessSaleInput{
		TicketID:    "ticket123",
		UserID:      "user456",
		PaymentType: "CREDIT_CARD",
	}

	result, err := service.ProcessSale(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, input.TicketID, result.TicketID)
	assert.Equal(t, input.UserID, result.UserID)
	assert.Equal(t, "CREDIT_CARD", result.PaymentType)

	ticketRepo.AssertExpectations(t)
	saleRepo.AssertExpectations(t)
	queue.AssertExpectations(t)
}

func TestProcessSale_InvalidPayment(t *testing.T) {
	service := services.NewSaleService(nil, nil, nil)

	input := services.ProcessSaleInput{
		TicketID:    "any",
		UserID:      "any",
		PaymentType: "PIX", // inválido
	}

	_, err := service.ProcessSale(context.TODO(), input)
	assert.ErrorIs(t, err, usecase.ErrInvalidPayment)
}

func TestProcessSale_TicketNotFound(t *testing.T) {
	ctx := context.TODO()
	ticketRepo := new(mocks.TicketRepository)

	ticketRepo.On("GetByID", ctx, "not-found").Return((*entities.Ticket)(nil), errors.New("not found"))

	service := services.NewSaleService(ticketRepo, nil, nil)

	input := services.ProcessSaleInput{
		TicketID:    "not-found",
		UserID:      "user",
		PaymentType: "CREDIT_CARD",
	}

	_, err := service.ProcessSale(ctx, input)
	assert.ErrorIs(t, err, usecase.ErrTicketNotFound)
}

func TestProcessSale_OutOfStock(t *testing.T) {
	ctx := context.TODO()
	ticketRepo := new(mocks.TicketRepository)

	ticket := &entities.Ticket{
		ID:       "ticket123",
		Quantity: 0,
	}

	ticketRepo.On("GetByID", ctx, ticket.ID).Return(ticket, nil)

	service := services.NewSaleService(ticketRepo, nil, nil)

	input := services.ProcessSaleInput{
		TicketID:    ticket.ID,
		UserID:      "user",
		PaymentType: "CREDIT_CARD",
	}

	_, err := service.ProcessSale(ctx, input)
	assert.ErrorIs(t, err, usecase.ErrOutOfStock)
}

func TestProcessSale_FailCreateSale(t *testing.T) {
	ctx := context.TODO()
	ticketRepo := new(mocks.TicketRepository)
	saleRepo := new(mocks.SaleRepository)

	ticket := &entities.Ticket{
		ID:       "ticket123",
		Quantity: 1,
	}

	ticketRepo.On("GetByID", ctx, ticket.ID).Return(ticket, nil)
	saleRepo.On("Create", ctx, mock.Anything).Return(errors.New("db error"))

	service := services.NewSaleService(ticketRepo, saleRepo, nil)

	input := services.ProcessSaleInput{
		TicketID:    ticket.ID,
		UserID:      "user",
		PaymentType: "CREDIT_CARD",
	}

	_, err := service.ProcessSale(ctx, input)
	assert.ErrorIs(t, err, usecase.ErrSaleCreation)
}

func TestProcessSale_FailUpdateStock(t *testing.T) {
	ctx := context.TODO()
	ticketRepo := new(mocks.TicketRepository)
	saleRepo := new(mocks.SaleRepository)

	ticket := &entities.Ticket{
		ID:       "ticket123",
		Quantity: 2,
	}

	ticketRepo.On("GetByID", ctx, ticket.ID).Return(ticket, nil)
	saleRepo.On("Create", ctx, mock.Anything).Return(nil)
	ticketRepo.On("Update", ctx, mock.Anything).Return(errors.New("update failed"))

	service := services.NewSaleService(ticketRepo, saleRepo, nil)

	input := services.ProcessSaleInput{
		TicketID:    ticket.ID,
		UserID:      "user",
		PaymentType: "CREDIT_CARD",
	}

	_, err := service.ProcessSale(ctx, input)
	assert.ErrorIs(t, err, usecase.ErrStockUpdate)
}

func TestProcessSale_FailQueuePublish(t *testing.T) {
	ctx := context.TODO()
	ticketRepo := new(mocks.TicketRepository)
	saleRepo := new(mocks.SaleRepository)
	queue := new(mocks.QueuePublisher)

	ticket := &entities.Ticket{
		ID:       "ticket123",
		Quantity: 1,
	}

	ticketRepo.On("GetByID", ctx, ticket.ID).Return(ticket, nil)
	saleRepo.On("Create", ctx, mock.Anything).Return(nil)
	ticketRepo.On("Update", ctx, mock.Anything).Return(nil)
	queue.On("PublishPurchaseValidation", ctx, mock.Anything).Return(errors.New("queue error"))

	service := services.NewSaleService(ticketRepo, saleRepo, queue)

	input := services.ProcessSaleInput{
		TicketID:    ticket.ID,
		UserID:      "user",
		PaymentType: "CREDIT_CARD",
	}

	_, err := service.ProcessSale(ctx, input)
	assert.ErrorIs(t, err, usecase.ErrQueuePublish)
}
