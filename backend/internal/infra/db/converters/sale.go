package converters

import (
	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
	"github.com/otaviomart1ns/backend-challenge/internal/util"
)

// ToSaleEntity converte o modelo SQLC para uma entidade de domínio
func ToSaleEntity(s models.Sale) *entities.Sale {
	return &entities.Sale{
		ID:        s.ID,
		TicketID:  s.TicketID,
		UserID:    s.UserID,
		Payment:   entities.PaymentType(s.PaymentType.String),
		CreatedAt: s.CreatedAt.Time,
	}
}

// ToSQLCCreateSaleParams converte a entidade do domínio para os parâmetros da query CreateSale
func ToSQLCCreateSaleParams(s *entities.Sale) models.CreateSaleParams {
	return models.CreateSaleParams{
		ID:          s.ID,
		TicketID:    s.TicketID,
		UserID:      s.UserID,
		PaymentType: util.ToNullString(string(s.Payment)),
	}
}
