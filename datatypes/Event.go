package datatypes

import(
	"time"
)

// Event with normal time.Time (RFC 3339) serialisation
type Event struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime time.Time `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []Artist `json:"artists"`
	Status string `json:"status"`
	TicketStatus string `json:"ticket_status"`
	OnSaleDatetime time.Time `json:"on_sale_datetime"`
}