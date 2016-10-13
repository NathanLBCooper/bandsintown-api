package internaldatatypes

import (
	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

// SerialisableVenueSearchParams is the parameter for the Venue Search api call
// Internal types are directly compatible with the api
// https://www.bandsintown.com/api/1.0/requests#venues-search
type SerialisableVenueSearchParams struct {
	Query    string `url:"query,omitempty"`
	Location string `url:"location,omitempty"`
	Radius   int    `url:"radius,omitempty"`
	Page     int    `url:"page,omitempty"`
	PerPage  int    `url:"per_page,omitempty"`
	AppID    string `url:"app_id,omitempty"`
}

// NewSerialisableVenueSearchParams creates SerialisableVenueSearchParams from VenueSearchParams
func NewSerialisableVenueSearchParams(params *datatypes.VenueSearchParams, appID string) *SerialisableVenueSearchParams {
	return &SerialisableVenueSearchParams{
		Query:    params.Query,
		Location: params.Location,
		Radius:   params.Radius,
		Page:     params.Page,
		PerPage:  params.PerPage,
		AppID:    appID,
	}
}
