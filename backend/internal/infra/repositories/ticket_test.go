package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/repositories"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/testhelpers"
)

func TestTicketRepository_GetByID(t *testing.T) {
	db, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewTicketRepository(queries)

	_, err := db.Exec(`INSERT INTO tickets (id, type, price, quantity) VALUES (?, ?, ?, ?)`,
		"ticket1", "VIP", 150.0, 20)
	assert.NoError(t, err)

	ticket, err := repo.GetByID(context.Background(), "ticket1")
	assert.NoError(t, err)
	assert.Equal(t, entities.TicketType("VIP"), ticket.Type)
	assert.Equal(t, float64(150.0), ticket.Price)
	assert.Equal(t, 20, ticket.Quantity)
}

func TestTicketRepository_GetByID_NotFound(t *testing.T) {
	_, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewTicketRepository(queries)

	_, err := repo.GetByID(context.Background(), "does-not-exist")
	assert.Error(t, err)
}

func TestTicketRepository_GetAll(t *testing.T) {
	db, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewTicketRepository(queries)

	_, err := db.Exec(`
		INSERT INTO tickets (id, type, price, quantity) VALUES
		("ticket1", "VIP", 100.0, 10),
		("ticket2", "STANDARD", 50.0, 30)
	`)
	assert.NoError(t, err)

	tickets, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, tickets, 2)
}

func TestTicketRepository_Update(t *testing.T) {
	db, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewTicketRepository(queries)

	// Inserir ticket inicial
	_, err := db.Exec(`INSERT INTO tickets (id, type, price, quantity) VALUES (?, ?, ?, ?)`,
		"ticket123", "STANDARD", 50.0, 30)
	assert.NoError(t, err)

	// Atualizar quantidade
	ticket := &entities.Ticket{
		ID:       "ticket123",
		Quantity: 25,
	}

	err = repo.Update(context.Background(), ticket)
	assert.NoError(t, err)

	// Verificar no banco
	row := db.QueryRow(`SELECT quantity FROM tickets WHERE id = ?`, "ticket123")
	var qty int
	err = row.Scan(&qty)
	assert.NoError(t, err)
	assert.Equal(t, 25, qty)
}

func TestTicketRepository_GetAll_DBError(t *testing.T) {
	db, queries := testhelpers.NewTestDB(t)
	repo := repositories.NewTicketRepository(queries)

	_ = db.Close() // força erro ao tentar acessar

	_, err := repo.GetAll(context.Background())
	assert.Error(t, err)
}
