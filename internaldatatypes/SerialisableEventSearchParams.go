package internaldatatypes

import (
	"github.com/NathanLBCooper/bandsintown-api-v1/datatypes"
)

// SerialisableEventSearchParams is the parameter for the Events Search api call
// Internal types are directly compatible with the api
// https://www.bandsintown.com/api/1.0/requests#events-search
type SerialisableEventSearchParams struct {
	Artists  []string `url:"artists[],omitempty"`
	Location string   `url:"location,omitempty"`
	Radius   int      `url:"radius,omitempty"`
	Date     string   `url:"date,omitempty"`
	Page     int      `url:"page,omitempty"`
	PerPage  int      `url:"per_page,omitempty"`
	AppID    string   `url:"app_id,omitempty"`
}

// NewSerialisableEventSearchParams creates SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventSearchParams(params *datatypes.EventSearchParams, appID string) *SerialisableEventSearchParams {
	return &SerialisableEventSearchParams{
		Artists:  params.Artists,
		Location: params.Location,
		Radius:   params.Radius,
		Date:     formatSearchTimes(params.Date),
		Page:     params.Page,
		PerPage:  params.PerPage,
		AppID:    appID,
	}
}
