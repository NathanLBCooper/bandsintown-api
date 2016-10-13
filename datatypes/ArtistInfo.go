package datatypes

// ArtistInfo is the basic info for a single artist returned by the Artist get api call
// https://www.bandsintown.com/api/1.0/requests#artists-get
type ArtistInfo struct {
	Artist
	UpcomingGigCount int `json:"upcoming_events_count"`
}
