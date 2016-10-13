package datatypes

// Artist represents core information about an Artist
// A component of ArtistInfo.go and Event.go
type Artist struct {
	Name string `json:"name"`
	MbID string `json:"mbid"`
	URL  string `json:"url"`
}
