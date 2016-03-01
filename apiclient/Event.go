package apiclient

import(
	"time"
)

// Event with CustomDate serialisation
type deserialisableEvent struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime CustomTime `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []Artist `json:"artists"`
	Status string `json:"status"`// todo enum
	TicketStatus string `json:"ticket_status"`// todo enum
	OnSaleDatetime CustomTime `json:"on_sale_datetime"`
}

// Event with normal time.Time (RFC 3339) serialisation
type Event struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime time.Time `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []Artist `json:"artists"`
	Status string `json:"status"`// todo enum
	TicketStatus string `json:"ticket_status"`// todo enum
	OnSaleDatetime time.Time `json:"on_sale_datetime"`
}

// Create Event from deserialisableEvent
func newEvent(event* deserialisableEvent) *Event {
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

func newEvents(deserialisableEvents []deserialisableEvent) []Event{
	events := make([]Event, len(deserialisableEvents))
	for i,dEvent := range deserialisableEvents{
		events[i] = *newEvent(&dEvent)
	}

	return events
}