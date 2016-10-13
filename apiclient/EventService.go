package apiclient

import (
	"fmt"
	"net/http"

	"github.com/NathanLBCooper/bandsintown-api/datatypes"
	"github.com/NathanLBCooper/bandsintown-api/internaldatatypes"
	"github.com/dghubble/sling"
)

// EventService provides methods for Event related requests
type EventService struct {
	AppID string
	Sling *sling.Sling
}

// NewEventService returns a new EventService.
func NewEventService(httpClient *http.Client, baseURL string, appID string) *EventService {
	return &EventService{
		Sling: sling.New().Client(httpClient).Base(baseURL),
		AppID: appID,
	}
}

// Search returns events matching search criteria (see below for available params).
// Useful in searching for local events or events within a specific time frame.
// If you are just looking for upcoming events for a single artist use Artists - Events.
// https://www.bandsintown.com/api/1.0/requests#events-search
// http://api.bandsintown.com/events/search.format
func (service *EventService) Search(params datatypes.EventSearchParams) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("events/search.%v", format)
	serialisableParams := internaldatatypes.NewSerialisableEventSearchParams(&params, service.AppID)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}

// Recommended returns recommended events given an artist or list of artists and a location (see below for all available params).
// https://www.bandsintown.com/api/1.0/requests#events-recommended
// http://api.bandsintown.com/events/recommended.format
func (service *EventService) Recommended(params datatypes.EventRecommendedParams) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("events/recommended.%v", format)
	serialisableParams := internaldatatypes.NewSerialisableEventRecommendedParams(&params, service.AppID)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}

// OnSaleSoon returns events going on sale in the next week (including today). Supports location filtering.
// https://www.bandsintown.com/api/1.0/requests#events-on-sale-soon
// http://api.bandsintown.com/events/on_sale_soon.format
func (service *EventService) OnSaleSoon(params datatypes.EventOnSaleSoonParams) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("events/on_sale_soon.%v", format)
	serialisableParams := internaldatatypes.NewSerialisableEventOnSaleSoonParams(&params, service.AppID)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}

// Daily returns events that have been created, updated or deleted in the last day. Useful in syncing data with Bandsintown.
// The daily feed is updated each day at 12:00 PM EST. (17:00 UTC)
// https://www.bandsintown.com/api/1.0/requests#events-daily
// http://api.bandsintown.com/events/daily.format
func (service *EventService) Daily() ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("events/daily.%v", format)
	resp, err := service.Sling.New().Get(path).QueryStruct(
		internaldatatypes.AppIDParam{AppID: service.AppID}).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}
