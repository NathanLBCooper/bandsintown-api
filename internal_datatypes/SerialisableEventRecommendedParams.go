package internal_datatypes

import(
	"bandsintown-api/datatypes"
)

type SerialisableEventRecommendedParams struct{
	SerialisableEventSearchParams
	OnlyRecommendations bool `url:"only_recs,omitempty"`
}

// Create SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventRecommendedParams(params* datatypes.EventRecommendedParams, appId string) *SerialisableEventRecommendedParams {
	return &SerialisableEventRecommendedParams{
		SerialisableEventSearchParams: *NewSerialisableEventSearchParams(&params.EventSearchParams, appId),
		OnlyRecommendations: params.OnlyRecommendations,
	}
}