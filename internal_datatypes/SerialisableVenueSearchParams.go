package internal_datatypes

import(
	"bandsintown-api/datatypes"
)

type SerialisableVenueSearchParams struct {
	Query string `url:"query,omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Page int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableVenueSearchParams from VenueSearchParams
func NewSerialisableVenueSearchParams(params* datatypes.VenueSearchParams, appId string) *SerialisableVenueSearchParams {
	return &SerialisableVenueSearchParams{
		Query: params.Query,
		Location: params.Location,
		Radius: params.Radius,
		Page: params.Page,
		PerPage: params.PerPage,
		ApiId: appId,
	}
}