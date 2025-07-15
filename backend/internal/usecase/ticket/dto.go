package ticket

type TicketOutput struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
