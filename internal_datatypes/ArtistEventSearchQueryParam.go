package internal_datatypes

import "time"

type ArtistEventSearchQueryParam struct {
	Date string `url:"date,omitempty"`
}

// Create ArtistEventSearchParam from []time.Time
func NewArtistEventSearchParam(params *[]time.Time) *ArtistEventSearchQueryParam {
	return &ArtistEventSearchQueryParam{
		Date: formatSearchTimes(*params),
	}
}