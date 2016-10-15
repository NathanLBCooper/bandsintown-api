package apiclient

import (
	"fmt"
	"net/http"

	"github.com/NathanLBCooper/bandsintown-api-v1/datatypes"
	"github.com/NathanLBCooper/bandsintown-api-v1/internaldatatypes"

	"github.com/dghubble/sling"
)

// An ArtistService provides methods for Artist related requests
type ArtistService struct {
	AppID string
	Sling *sling.Sling
}

// NewArtistService returns a new ArtistService.
func NewArtistService(httpClient *http.Client, baseURL string, appID string) *ArtistService {
	return &ArtistService{
		Sling: sling.New().Client(httpClient).Base(baseURL),
		AppID: appID,
	}
}

// Returns basic information for a single artist, including the number of upcoming events.
// Useful in determining if an artist is on tour without requesting the event data.
// https://www.bandsintown.com/api/1.0/requests#artists-get
// http://api.bandsintown.com/artists/name.format
func (service *ArtistService) getInfo(param string) (datatypes.ArtistInfo, *http.Response, error) {
	artistInfo := new(datatypes.ArtistInfo)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("artists/%v.%v?app_id=%v", param, format, service.AppID)
	resp, err := service.Sling.New().Get(path).Receive(artistInfo, apiError)
	if err == nil && apiError.HasErrors() {
		err = apiError
	}
	return *artistInfo, resp, err
}

// GetInfoByName returns basic information for a single artist by name, including the number of upcoming events.
// Useful in determining if an artist is on tour without requesting the event data.
// https://www.bandsintown.com/api/1.0/requests#artists-get
// http://api.bandsintown.com/artists/name.format
func (service *ArtistService) GetInfoByName(name string) (datatypes.ArtistInfo, *http.Response, error) {
	return service.getInfo(name)
}

// GetInfoByMbID returns basic information for a single artist by MbID, including the number of upcoming events.
// Useful in determining if an artist is on tour without requesting the event data.
// https://www.bandsintown.com/api/1.0/requests#artists-get
// http://api.bandsintown.com/artists/name.format
func (service *ArtistService) GetInfoByMbID(MbID string) (datatypes.ArtistInfo, *http.Response, error) {
	return service.getInfo(fmt.Sprintf("mbid_%v", MbID))
}

// GetEvents Returns events for a single artist.
// https://www.bandsintown.com/api/1.0/requests#artists-events
// http://api.bandsintown.com/artists/name/events.format
func (service *ArtistService) GetEvents(param datatypes.ArtistEventSearchParam) ([]datatypes.Event, *http.Response, error) {
	var artist string
	if param.MbID != "" {
		artist = param.MbID
	} else {
		artist = param.Name
	}

	deserialisableEvents := new([]internaldatatypes.DeserialisableEvent)
	apiError := new(datatypes.APIError)
	path := fmt.Sprintf("artists/%v/events.%v?app_id=%v", artist, format, service.AppID)

	args := new(internaldatatypes.ArtistEventSearchQueryParam)
	if param.Date != nil {
		args = internaldatatypes.NewArtistEventSearchParam(&param.Date)
	}
	resp, err := service.Sling.New().Get(path).QueryStruct(args).Receive(deserialisableEvents, apiError)

	if err == nil && apiError.HasErrors() {
		err = apiError
	}

	return internaldatatypes.NewEvents(*deserialisableEvents), resp, err
}
