package datatypes

import(
	"bandsintown-api/customtimes"
)

type SerialisableEventSearchParams struct {
	Artists []string `url:"artists[],omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Datetime []customtimes.CustomSearchTime `url:"datetime,omitempty"`
	Page int `url:"page,omitempty"`
	Perpage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventSearchParams(params* EventSearchParams, appId string) *SerialisableEventSearchParams {
	return &SerialisableEventSearchParams{
		Artists: params.Artists,
		Location: params.Location,
		Radius: params.Radius,
		Datetime: customtimes.NewCustomSearchTimes(params.Datetime),
		Page: params.Page,
		Perpage: params.Perpage,
		ApiId: appId,
	}
}