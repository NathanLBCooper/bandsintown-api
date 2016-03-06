package internal_datatypes

import(
	"bandsintown-api/datatypes"
)

type SerialisableEventSearchParams struct {
	Artists []string `url:"artists[],omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Date string `url:"date,omitempty"`
	Page int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventSearchParams(params* datatypes.EventSearchParams, appId string)*SerialisableEventSearchParams {
	return &SerialisableEventSearchParams{
		Artists: params.Artists,
		Location: params.Location,
		Radius: params.Radius,
		Date: formatSearchTimes(params.Date),
		Page: params.Page,
		PerPage: params.PerPage,
		ApiId: appId,
	}
}