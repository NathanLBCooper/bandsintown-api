package apiclient

import(
	"bandsintown-api/datatypes"
)

type serialisableVenueSearchParams struct {
	Query string `url:"query,omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Page int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create serialisableVenueSearchParams from VenueSearchParams
func newSerialisableVenueSearchParams(params* datatypes.VenueSearchParams, appId string) *serialisableVenueSearchParams {
	return &serialisableVenueSearchParams{
		Query: params.Query,
		Location: params.Location,
		Radius: params.Radius,
		Page: params.Page,
		PerPage: params.PerPage,
		ApiId: appId,
	}
}