package converters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/converters"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

func TestToTicketEntity(t *testing.T) {
	model := models.Ticket{
		ID:       "ticket123",
		Type:     "VIP",
		Price:    150.0,
		Quantity: 5,
	}

	expected := &entities.Ticket{
		ID:       "ticket123",
		Type:     entities.VIP,
		Price:    150.0,
		Quantity: 5,
	}

	result := converters.ToTicketEntity(model)

	assert.Equal(t, expected, result)
}

func TestToTicketEntityList(t *testing.T) {
	models := []models.Ticket{
		{ID: "1", Type: "GENERAL_AREA", Price: 50.0, Quantity: 10},
		{ID: "2", Type: "VIP", Price: 150.0, Quantity: 5},
	}

	result := converters.ToTicketEntityList(models)

	assert.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, entities.GeneralArea, result[0].Type)
	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, entities.VIP, result[1].Type)
}

func TestToSQLCUpdateTicketQuantityParams(t *testing.T) {
	entity := &entities.Ticket{
		ID:       "ticket456",
		Quantity: 7,
	}

	params := converters.ToSQLCUpdateTicketQuantityParams(entity)

	assert.Equal(t, "ticket456", params.ID)
	assert.Equal(t, int32(7), params.Quantity)
}
