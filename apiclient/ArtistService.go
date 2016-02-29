package apiclient

import(
	"fmt"
	"net/http"
	"github.com/dghubble/sling"
)

const baseURL = "http://api.bandsintown.com"
const format = "json"
const appId = "some_api_id"


// ArtistService provides methods for Artist related requests
type ArtistService struct {
	sling *sling.Sling
}

// NewIssueService returns a new IssueService.
func NewArtistService(httpClient *http.Client) *ArtistService {
	return &ArtistService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

// Returns basic information for a single artist, including the number of upcoming events.
// Useful in determining if an artist is on tour without requesting the event data.
// https://www.bandsintown.com/api/1.0/requests#artists-get
func (s *ArtistService) GetInfo(name string) (ArtistInfo, *http.Response, error) {
	artistInfo := new(ArtistInfo)
	apiError := new(ApiError)
	path := fmt.Sprintf("artists/%v.%v?app_id=%v", name, format, appId)
	resp, err := s.sling.New().Get(path).Receive(artistInfo, apiError)
	if err == nil {
		err = apiError
	}
	return *artistInfo, resp, err
}

// Returns events for a single artists.
func (s *ArtistService) GetEvents(name string) ([]Event, *http.Response, error) {
	events := new([]Event)
	apiError := new(ApiError)
	path := fmt.Sprintf("artists/%v/events.%v?app_id=%v", name, format, appId)
	resp, err := s.sling.New().Get(path).Receive(events, apiError)
	if err == nil {
		err = apiError
	}
	return *events, resp, err
}