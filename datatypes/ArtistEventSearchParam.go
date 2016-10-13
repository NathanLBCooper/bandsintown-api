package datatypes

import "time"

// ArtistEventSearchParam is the parameter for the Artist Events api call
// https://www.bandsintown.com/api/1.0/requests#artists-events
type ArtistEventSearchParam struct {
	Name string
	MbID string
	Date []time.Time
}
