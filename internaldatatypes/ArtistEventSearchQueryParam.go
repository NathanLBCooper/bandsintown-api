package internaldatatypes

import "time"

// ArtistEventSearchQueryParam is the parameter for the Artist Events api call
// Internal types are directly compatible with the api
// https://www.bandsintown.com/api/1.0/requests#artists-events
type ArtistEventSearchQueryParam struct {
	Date string `url:"date,omitempty"`
}

// NewArtistEventSearchParam creates ArtistEventSearchParam from []time.Time
func NewArtistEventSearchParam(params *[]time.Time) *ArtistEventSearchQueryParam {
	return &ArtistEventSearchQueryParam{
		Date: formatSearchTimes(*params),
	}
}
