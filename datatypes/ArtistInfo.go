package datatypes

// todo, very similar to Artist

type ArtistInfo struct{
	Artist
	UpcomingGigCount int `json:"upcoming_events_count"`
}