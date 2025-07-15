package sale

import (
	"context"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/services"
)

type ProcessSaleUseCaseInterface interface {
	Execute(ctx context.Context, input ProcessSaleInput) (*ProcessSaleOutput, error)
}

type ProcessSaleUseCase struct {
	service services.SaleService
}

func NewProcessSaleUseCase(service services.SaleService) *ProcessSaleUseCase {
	return &ProcessSaleUseCase{
		service: service,
	}
}

func (uc *ProcessSaleUseCase) Execute(ctx context.Context, input ProcessSaleInput) (*ProcessSaleOutput, error) {
	output, err := uc.service.ProcessSale(ctx, services.ProcessSaleInput{
		TicketID:    input.TicketID,
		UserID:      input.UserID,
		PaymentType: input.PaymentType,
	})
	if err != nil {
		return nil, err
	}

	return &ProcessSaleOutput{
		SaleID:      output.SaleID,
		TicketID:    output.TicketID,
		UserID:      output.UserID,
		PaymentType: output.PaymentType,
	}, nil
}
