package converters_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/converters"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

func TestToSaleEntity(t *testing.T) {
	now := time.Now()

	model := models.Sale{
		ID:          "sale123",
		TicketID:    "ticket456",
		UserID:      "user789",
		PaymentType: sql.NullString{String: "CREDIT_CARD", Valid: true},
		CreatedAt:   sql.NullTime{Time: now, Valid: true},
	}

	expected := &entities.Sale{
		ID:        "sale123",
		TicketID:  "ticket456",
		UserID:    "user789",
		Payment:   entities.CreditCard,
		CreatedAt: now,
	}

	entity := converters.ToSaleEntity(model)

	assert.Equal(t, expected, entity)
}

func TestToSQLCCreateSaleParams(t *testing.T) {
	entity := &entities.Sale{
		ID:       "sale123",
		TicketID: "ticket456",
		UserID:   "user789",
		Payment:  entities.CreditCard,
	}

	params := converters.ToSQLCCreateSaleParams(entity)

	assert.Equal(t, "sale123", params.ID)
	assert.Equal(t, "ticket456", params.TicketID)
	assert.Equal(t, "user789", params.UserID)
	assert.True(t, params.PaymentType.Valid)
	assert.Equal(t, "CREDIT_CARD", params.PaymentType.String)
}
