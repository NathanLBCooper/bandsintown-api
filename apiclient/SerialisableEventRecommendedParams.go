package apiclient

import(
	"bandsintown-api/datatypes"
)

type serialisableEventRecommendedParams struct{
	serialisableEventSearchParams
	OnlyRecommendations bool `url:"only_recs,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func newSerialisableEventRecommendedParams(params* datatypes.EventRecommendedParams, appId string) *serialisableEventRecommendedParams {
	return &serialisableEventRecommendedParams{
		serialisableEventSearchParams: *newSerialisableEventSearchParams(&params.EventSearchParams, appId),
		OnlyRecommendations: params.OnlyRecommendations,
	}
}