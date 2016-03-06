package internal_datatypes

import(
	"bandsintown-api/datatypes"
)

type SerialisableEventOnSaleSoonParams struct {
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventOnSaleSoonParams(params* datatypes.EventOnSaleSoonParams, appId string) *SerialisableEventOnSaleSoonParams {
	return &SerialisableEventOnSaleSoonParams{
		Location: params.Location,
		Radius: params.Radius,
		ApiId: appId,
	}
}