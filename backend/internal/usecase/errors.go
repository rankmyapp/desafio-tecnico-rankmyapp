package usecase

import "errors"

var (
	// Tickets
	ErrTicketNotFound   = errors.New("ticket not found")
	ErrOutOfStock       = errors.New("ticket is sold out")
	ErrTicketListFailed = errors.New("failed to list tickets")
	ErrStockUpdate      = errors.New("failed to update ticket quantity")

	// Pagamento
	ErrInvalidPayment = errors.New("only CREDIT_CARD payments are accepted")

	// Fila
	ErrQueuePublish = errors.New("failed to publish to queue")

	// Sale
	ErrSaleCreation = errors.New("failed to create sale")
)
