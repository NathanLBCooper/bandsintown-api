package internaldatatypes

import (
	"github.com/NathanLBCooper/bandsintown-api-v1/datatypes"
)

// DeserialisableEvent in an datatypes.Event.go with CustomDate serialisation
type DeserialisableEvent struct {
	ID             int                `json:"id"`
	URL            string             `json:"url"`
	Datetime       CustomResponseTime `json:"datetime"`
	TicketURL      string             `json:"ticket_url"`
	Artists        []datatypes.Artist `json:"artists"`
	Venue          datatypes.Venue    `json:"venue"`
	Status         string             `json:"status"`
	TicketStatus   string             `json:"ticket_status"`
	OnSaleDatetime CustomResponseTime `json:"on_sale_datetime"`
}

// NewEvent creates Event from DeserialisableEvent
func NewEvent(event *DeserialisableEvent) *datatypes.Event {
	return &datatypes.Event{
		ID:             event.ID,
		URL:            event.URL,
		Datetime:       event.Datetime.Time,
		TicketURL:      event.TicketURL,
		Artists:        event.Artists,
		Venue:          event.Venue,
		Status:         event.Status,
		TicketStatus:   event.TicketStatus,
		OnSaleDatetime: event.OnSaleDatetime.Time,
	}
}

// NewEvents creates multiple Events from multiple DeserialisableEvents
func NewEvents(deserialisableEvents []DeserialisableEvent) []datatypes.Event {
	events := make([]datatypes.Event, len(deserialisableEvents))
	for i, dEvent := range deserialisableEvents {
		events[i] = *NewEvent(&dEvent)
	}

	return events
}
