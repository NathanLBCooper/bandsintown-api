package datatypes

type ArtistInfo struct{
	Artist
	UpcomingGigCount int `json:"upcoming_events_count"`
}