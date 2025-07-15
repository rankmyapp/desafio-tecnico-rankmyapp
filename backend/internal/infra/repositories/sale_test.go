package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/testhelpers"
)

func TestSaleRepository_CreateAndGetByID(t *testing.T) {
	db, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewSaleRepository(queries)

	_, err := db.Exec(`INSERT INTO tickets (id, type, price, quantity) VALUES (?, ?, ?, ?)`,
		"ticket123", "VIP", 200.0, 10)
	assert.NoError(t, err)

	sale := &entities.Sale{
		ID:        uuid.NewString(),
		TicketID:  "ticket123",
		UserID:    "user456",
		Payment:   entities.CreditCard,
		CreatedAt: time.Now(),
	}

	err = repo.Create(context.Background(), sale)
	assert.NoError(t, err)

	found, err := repo.GetByID(context.Background(), sale.ID)
	assert.NoError(t, err)
	assert.Equal(t, sale.ID, found.ID)
	assert.Equal(t, sale.TicketID, found.TicketID)
	assert.Equal(t, sale.UserID, found.UserID)
	assert.Equal(t, sale.Payment, found.Payment)
}

func TestSaleRepository_GetByID_NotFound(t *testing.T) {
	_, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewSaleRepository(queries)

	ctx := context.Background()

	_, err := repo.GetByID(ctx, "non-existent-id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
