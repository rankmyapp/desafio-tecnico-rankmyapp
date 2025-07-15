package entities

type TicketType string

const (
	GeneralArea  TicketType = "GENERAL_AREA"
	Grandstand   TicketType = "GRANDSTAND"
	VIP          TicketType = "VIP"
	GoldenCircle TicketType = "GOLDEN_CIRCLE"
)

type Ticket struct {
	ID       string     `json:"id"`
	Type     TicketType `json:"type"`
	Price    float64    `json:"price"`
	Quantity int        `json:"quantity"`
}
