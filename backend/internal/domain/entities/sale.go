package entities

import "time"

type PaymentType string

const (
	CreditCard PaymentType = "CREDIT_CARD"
)

func (p PaymentType) IsValid() bool {
	return p == CreditCard
}

type Sale struct {
	ID        string      `json:"id"`
	TicketID  string      `json:"ticketId"`
	UserID    string      `json:"userId"`
	Payment   PaymentType `json:"paymentType"`
	CreatedAt time.Time   `json:"createdAt"`
}
