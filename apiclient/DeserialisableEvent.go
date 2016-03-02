package apiclient

import(
	"bandsintown-api/customtimes"
	"bandsintown-api/datatypes"
)

// Event with CustomDate serialisation
type deserialisableEvent struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime customtimes.CustomResponseTime `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []datatypes.Artist `json:"artists"`
	Status string `json:"status"`
	TicketStatus string `json:"ticket_status"`
	OnSaleDatetime customtimes.CustomResponseTime `json:"on_sale_datetime"`
}

// Create Event from DeserialisableEvent
func newEvent(event* deserialisableEvent) *datatypes.Event {
	return &datatypes.Event{
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

func newEvents(deserialisableEvents []deserialisableEvent) []datatypes.Event{
	events := make([]datatypes.Event, len(deserialisableEvents))
	for i,dEvent := range deserialisableEvents{
		events[i] = *newEvent(&dEvent)
	}

	return events
}