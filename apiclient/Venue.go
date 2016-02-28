package apiclient

type Venue struct{
	Id int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	Url string `json:"url"`
	Latitude int `json:"latitude"`
	Longitude int `json:"longitude"`
}