package apiclient

import (
	"fmt"
	"net/http"

	"github.com/NathanLBCooper/bandsintown-api/datatypes"
	"github.com/NathanLBCooper/bandsintown-api/internaldatatypes"
	"github.com/dghubble/sling"
)

// VenueService provides methods for Venue related requests
type VenueService struct {
	AppID string
	Sling *sling.Sling
}

// NewVenueService returns a new ArtistService.
func NewVenueService(httpClient *http.Client, baseURL string, appID string) *VenueService {
	return &VenueService{
		Sling: sling.New().Client(httpClient).Base(baseURL),
		AppID: appID,
	}
}

// Events returns all upcoming events for a single venue.
// https://www.bandsintown.com/api/1.0/requests#venues-events
// http://api.bandsintown.com/venues/id/events.format
func (service *VenueService) Events(venueID int) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("venues/%v/events.%v?app_id=%v", venueID, format, service.AppID)
	resp, err := service.Sling.New().Get(path).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}

// Search returns venues matching a search query (supports location filtering).
// https://www.bandsintown.com/api/1.0/requests#venues-search
// http://api.bandsintown.com/venues/search.format
// todo: Seems consistant with the Api. But I can't even get data from their API manually. Perhaps it's broken on their end?
func (service *VenueService) Search(params datatypes.VenueSearchParams) ([]datatypes.Venue, *http.Response, error) {
	venues := new([]datatypes.Venue)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("venues/search.%v", format)
	serialisableParams := internaldatatypes.NewSerialisableVenueSearchParams(&params, service.AppID)
	resp, err := service.Sling.New().Get(path).QueryStruct(serialisableParams).Receive(venues, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return *venues, resp, err
}
