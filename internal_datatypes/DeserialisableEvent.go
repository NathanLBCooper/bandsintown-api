package internal_datatypes

import(
	"bandsintown-api/datatypes"
)

// Event with CustomDate serialisation
type DeserialisableEvent struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime CustomResponseTime `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []datatypes.Artist `json:"artists"`
	Venue datatypes.Venue `json:"venue"`
	Status string `json:"status"`
	TicketStatus string `json:"ticket_status"`
	OnSaleDatetime CustomResponseTime `json:"on_sale_datetime"`
}

// Create Event from DeserialisableEvent
func NewEvent(event* DeserialisableEvent) *datatypes.Event {
	return &datatypes.Event{
		Id: event.Id,
		Url: event.Url,
		Datetime: event.Datetime.Time,
		TicketUrl: event.TicketUrl,
		Artists: event.Artists,
		Venue: event.Venue,
		Status: event.Status,
		TicketStatus: event.TicketStatus,
		OnSaleDatetime: event.OnSaleDatetime.Time,
	}
}

func NewEvents(deserialisableEvents []DeserialisableEvent) []datatypes.Event{
	events := make([]datatypes.Event, len(deserialisableEvents))
	for i,dEvent := range deserialisableEvents{
		events[i] = *NewEvent(&dEvent)
	}

	return events
}