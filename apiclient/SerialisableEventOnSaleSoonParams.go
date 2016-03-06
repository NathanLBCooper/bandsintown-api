package apiclient

import(
	"bandsintown-api/datatypes"
)

type serialisableEventOnSaleSoonParams struct {
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func newSerialisableEventOnSaleSoonParams(params* datatypes.EventOnSaleSoonParams, appId string)*serialisableEventOnSaleSoonParams {
	return &serialisableEventOnSaleSoonParams{
		Location: params.Location,
		Radius: params.Radius,
		ApiId: appId,
	}
}