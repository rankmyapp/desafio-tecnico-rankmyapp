package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories/mocks"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
	"github.com/otaviomart1ns/backend-challenge/internal/usecase"
)

func TestListCatalog_Success(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.TicketRepository)

	expected := []*entities.Ticket{
		{ID: "1", Type: entities.GoldenCircle, Price: 300.0, Quantity: 10},
	}

	mockRepo.On("GetAll", ctx).Return(expected, nil)

	service := services.NewTicketService(mockRepo)

	tickets, err := service.ListCatalog(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, tickets)
	mockRepo.AssertExpectations(t)
}

func TestListCatalog_Failure(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.TicketRepository)

	mockRepo.On("GetAll", ctx).Return(([]*entities.Ticket)(nil), errors.New("db error"))

	service := services.NewTicketService(mockRepo)

	tickets, err := service.ListCatalog(ctx)

	assert.Nil(t, tickets)
	assert.ErrorIs(t, err, usecase.ErrTicketListFailed)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.TicketRepository)

	ticket := &entities.Ticket{ID: "abc", Type: entities.VIP, Price: 150.0, Quantity: 5}

	mockRepo.On("GetByID", ctx, "abc").Return(ticket, nil)

	service := services.NewTicketService(mockRepo)

	result, err := service.GetByID(ctx, "abc")

	assert.NoError(t, err)
	assert.Equal(t, ticket, result)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Failure(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.TicketRepository)

	mockRepo.On("GetByID", ctx, "xyz").Return((*entities.Ticket)(nil), errors.New("not found"))

	service := services.NewTicketService(mockRepo)

	result, err := service.GetByID(ctx, "xyz")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, usecase.ErrTicketNotFound)
	mockRepo.AssertExpectations(t)
}
