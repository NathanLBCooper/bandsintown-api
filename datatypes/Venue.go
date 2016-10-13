package datatypes

// Venue represents a venue, returned by the Venue Search api call
// https://www.bandsintown.com/api/1.0/requests#venues-search
type Venue struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	URL       string  `json:"url"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
