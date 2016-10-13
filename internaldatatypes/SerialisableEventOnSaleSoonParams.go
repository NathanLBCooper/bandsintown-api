package internaldatatypes

import (
	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

// SerialisableEventOnSaleSoonParams is the parameter for the Events On Sale Soon api call
// Internal types are directly compatible with the api
// https://www.bandsintown.com/api/1.0/requests#events-on-sale-soon
type SerialisableEventOnSaleSoonParams struct {
	Location string `url:"location,omitempty"`
	Radius   int    `url:"radius,omitempty"`
	AppID    string `url:"app_id,omitempty"`
}

// NewSerialisableEventOnSaleSoonParams creates SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventOnSaleSoonParams(params *datatypes.EventOnSaleSoonParams, appID string) *SerialisableEventOnSaleSoonParams {
	return &SerialisableEventOnSaleSoonParams{
		Location: params.Location,
		Radius:   params.Radius,
		AppID:    appID,
	}
}
