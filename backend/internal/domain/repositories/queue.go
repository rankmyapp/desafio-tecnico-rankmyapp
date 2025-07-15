package repositories

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
)

type QueuePublisher interface {
	PublishPurchaseValidation(ctx context.Context, sale *entities.Sale) error
}
