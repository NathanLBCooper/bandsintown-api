package internaldatatypes

import (
	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

// SerialisableEventRecommendedParams is the parameter for the Events Recommmended api call
// Internal types are directly compatible with the api
// https://www.bandsintown.com/api/1.0/requests#events-recommended
type SerialisableEventRecommendedParams struct {
	SerialisableEventSearchParams
	OnlyRecommendations bool `url:"only_recs,omitempty"`
}

// NewSerialisableEventRecommendedParams creates SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventRecommendedParams(params *datatypes.EventRecommendedParams, appID string) *SerialisableEventRecommendedParams {
	return &SerialisableEventRecommendedParams{
		SerialisableEventSearchParams: *NewSerialisableEventSearchParams(&params.EventSearchParams, appID),
		OnlyRecommendations:           params.OnlyRecommendations,
	}
}
