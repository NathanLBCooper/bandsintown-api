package apiclient

import(
	"fmt"
	"net/http"
	"github.com/dghubble/sling"
	"bandsintown-api/datatypes"
)

// VenueService provides methods for Venue related requests
type VenueService struct {
	AppId string
	Sling *sling.Sling
}

// NewVenueService returns a new ArtistService.
func NewVenueService(httpClient *http.Client, baseUrl string, appId string) *VenueService {
	return &VenueService{
		Sling: sling.New().Client(httpClient).Base(baseUrl),
		AppId: appId,
	}
}

// Returns all upcoming events for a single venue.
// https://www.bandsintown.com/api/1.0/requests#venues-events
// http://api.bandsintown.com/venues/id/events.format //todo name of venueId
func (service *VenueService) Events(venueId int) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]deserialisableEvent)
	apiError := new(datatypes.ApiError)
	path := fmt.Sprintf("venues/%v/events.%v?app_id=%v", venueId, format, service.AppId)
	resp, err := service.Sling.New().Get(path).Receive(deserialisableEvents, apiError)

	if err == nil {
		err = apiError
	}

	return newEvents(*deserialisableEvents), resp, err
}