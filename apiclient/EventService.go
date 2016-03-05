package apiclient

import(
	"fmt"
	"net/http"
	"github.com/dghubble/sling"
	"bandsintown-api/datatypes"
)

// EventService provides methods for Event related requests
type EventService struct {
	AppId string
	Sling *sling.Sling
}

// NewEventService returns a new EventService.
func NewEventService(httpClient *http.Client, baseUrl string, appId string) *EventService {
	return &EventService{
		Sling: sling.New().Client(httpClient).Base(baseUrl),
		AppId: appId,
	}
}

// Returns events matching search criteria (see below for available params).
// Useful in searching for local events or events within a specific time frame.
// If you are just looking for upcoming events for a single artist use Artists - Events.
// https://www.bandsintown.com/api/1.0/requests#events-search
// http://api.bandsintown.com/events/search.format
func (service *EventService) Search(params datatypes.EventSearchParams) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]deserialisableEvent)
	apiError := new(datatypes.ApiError)
	path := fmt.Sprintf("events/search.%v", format)
	serialisableParams := newSerialisableEventSearchParams(&params, service.AppId)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(deserialisableEvents, apiError)

	if err == nil {
		err = apiError
	}

	return newEvents(*deserialisableEvents), resp, err
}

// Returns recommended events given an artist or list of artists and a location (see below for all available params).
// https://www.bandsintown.com/api/1.0/requests#events-recommended
// http://api.bandsintown.com/events/recommended.format
func (service *EventService) Recommended(params datatypes.EventRecommendedParams) ([]datatypes.Event, *http.Response, error){
	deserialisableEvents := new([]deserialisableEvent)
	apiError := new(datatypes.ApiError)
	path := fmt.Sprintf("events/recommended.%v", format)
	serialisableParams := newSerialisableEventRecommendedParams(&params, service.AppId)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(deserialisableEvents, apiError)

	if err == nil {
		err = apiError
	}

	return newEvents(*deserialisableEvents), resp, err
}