package datatypes

import (
	"time"
)

// Event respresent events, returned by the event service
// Has normal time.Time (RFC 3339) json serialisation
type Event struct {
	ID             int       `json:"id"`
	URL            string    `json:"url"`
	Datetime       time.Time `json:"datetime"`
	TicketURL      string    `json:"ticket_url"`
	Artists        []Artist  `json:"artists"`
	Venue          Venue     `json:"venue"`
	Status         string    `json:"status"`
	TicketStatus   string    `json:"ticket_status"`
	OnSaleDatetime time.Time `json:"on_sale_datetime"`
}
