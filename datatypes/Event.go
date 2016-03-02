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

// Create Event from DeserialisableEvent
func NewEvent(event* DeserialisableEvent) *Event {
	return &Event{
		Id: event.Id,
		Url: event.Url,
		Datetime: event.Datetime.Time,
		TicketUrl: event.TicketUrl,
		Artists: event.Artists,
		Status: event.Status,
		TicketStatus: event.TicketStatus,
		OnSaleDatetime: event.OnSaleDatetime.Time,
	}
}

func NewEvents(deserialisableEvents []DeserialisableEvent) []Event{
	events := make([]Event, len(deserialisableEvents))
	for i,dEvent := range deserialisableEvents{
		events[i] = *NewEvent(&dEvent)
	}

	return events
}