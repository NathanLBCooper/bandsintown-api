package apiclient

import(
	"bandsintown-api/customtimes"
	"bandsintown-api/datatypes"
)

type serialisableEventSearchParams struct {
	Artists []string `url:"artists[],omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Datetime []customtimes.CustomSearchTime `url:"datetime,omitempty"`
	Page int `url:"page,omitempty"`
	Perpage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func newSerialisableEventSearchParams(params* datatypes.EventSearchParams, appId string)*serialisableEventSearchParams {
	return &serialisableEventSearchParams{
		Artists: params.Artists,
		Location: params.Location,
		Radius: params.Radius,
		Datetime: customtimes.NewCustomSearchTimes(params.Datetime),
		Page: params.Page,
		Perpage: params.Perpage,
		ApiId: appId,
	}
}