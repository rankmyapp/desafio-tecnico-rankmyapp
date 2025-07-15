package sale

type ProcessSaleInput struct {
	TicketID    string `json:"ticketId" binding:"required"`
	UserID      string `json:"userId" binding:"required"`
	PaymentType string `json:"paymentType" binding:"required"`
}

type ProcessSaleOutput struct {
	SaleID      string `json:"saleId"`
	TicketID    string `json:"ticketId"`
	UserID      string `json:"userId"`
	PaymentType string `json:"paymentType"`
}
