package apiclient

import(
	"time"
)

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