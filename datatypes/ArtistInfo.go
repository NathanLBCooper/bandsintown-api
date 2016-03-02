package datatypes

// todo, very similar to Artist

type ArtistInfo struct{
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	Url string `json:"url"`
	UpcomingGigCount int `json:"upcoming_events_count"`
}