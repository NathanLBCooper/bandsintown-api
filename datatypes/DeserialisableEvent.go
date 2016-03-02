package datatypes

import(
	"bandsintown-api/customtimes"
)

// Event with CustomDate serialisation
type DeserialisableEvent struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Datetime customtimes.CustomResponseTime `json:"datetime"`
	TicketUrl string `json:"ticket_url"`
	Artists []Artist `json:"artists"`
	Status string `json:"status"`
	TicketStatus string `json:"ticket_status"`
	OnSaleDatetime customtimes.CustomResponseTime `json:"on_sale_datetime"`
}