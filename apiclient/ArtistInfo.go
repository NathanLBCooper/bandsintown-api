package apiclient

// todo, this

type ArtistInfo struct{
	Name string `json:"name"`
	Mbid  int `json:"mbid"`
	Url string `json:"url"`
	UpcomingGigCount int `json:"upcoming_events_count"`
}

type Artist struct{
	Name string `json:"name"`
	Mbid  int `json:"mbid"`
	Url string `json:"url"`
}