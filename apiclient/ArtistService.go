package apiclient

import(
	"fmt"
	"net/http"
	"github.com/dghubble/sling"
	"bandsintown-api/datatypes"
	"bandsintown-api/internal_datatypes"
)

// ArtistService provides methods for Artist related requests
type ArtistService struct {
	AppId string
	Sling *sling.Sling
}

// NewArtistService returns a new ArtistService.
func NewArtistService(httpClient *http.Client, baseUrl string, appId string) *ArtistService {
	return &ArtistService{
		Sling: sling.New().Client(httpClient).Base(baseUrl),
		AppId: appId,
	}
}

// Returns basic information for a single artist, including the number of upcoming events.
// Useful in determining if an artist is on tour without requesting the event data.
// https://www.bandsintown.com/api/1.0/requests#artists-get
// http://api.bandsintown.com/artists/name.format
func (service *ArtistService) GetInfo(name string) (datatypes.ArtistInfo, *http.Response, error) {
	artistInfo := new(datatypes.ArtistInfo)
	apiError := new(datatypes.ApiError)
	path := fmt.Sprintf("artists/%v.%v?app_id=%v", name, format, service.AppId)
	resp, err := service.Sling.New().Get(path).Receive(artistInfo, apiError)
	if err == nil {
		err = apiError
	}
	return *artistInfo, resp, err
}

// Returns events for a single artist.
// https://www.bandsintown.com/api/1.0/requests#artists-events
// http://api.bandsintown.com/artists/name/events.format
func (service *ArtistService) GetEvents(name string) ([]datatypes.Event, *http.Response, error) {
	deserialisableEvents := new([]internal_datatypes.DeserialisableEvent)
	apiError := new(datatypes.ApiError)
	path := fmt.Sprintf("artists/%v/events.%v?app_id=%v", name, format, service.AppId)
	resp, err := service.Sling.New().Get(path).Receive(deserialisableEvents, apiError)

	if err == nil {
		err = apiError
	}

	return internal_datatypes.NewEvents(*deserialisableEvents), resp, err
}