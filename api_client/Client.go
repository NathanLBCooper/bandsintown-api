package api_client

import(
	"net/http"
)

const format = "json"
const bandsInTownBaseUrl = "http://api.bandsintown.com"

// Client for the Bands in Town API
type Client struct {
	// Service Implementations
	ArtistService *ArtistService
	EventService *EventService
	VenueService *VenueService
}

// NewClientDetailed returns a new Client
func NewClientDetailed(httpClient *http.Client, baseUrl string, appId string) *Client {
	return &Client{
		ArtistService: NewArtistService(httpClient, baseUrl, appId),
		EventService: NewEventService(httpClient, baseUrl, appId),
		VenueService: NewVenueService(httpClient, baseUrl, appId),
	}
}

// NewClient returns a new Client
func NewClient(appId string) *Client {
	return NewClientDetailed(nil, bandsInTownBaseUrl, appId)
}