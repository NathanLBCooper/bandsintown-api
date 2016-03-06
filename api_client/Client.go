package api_client

import(
	"net/http"
)

const format = "json"

// Client for the Bands in Town API
type Client struct {
	// Service Implementations
	ArtistService *ArtistService
	EventService *EventService
	VenueService *VenueService
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, baseUrl string, appId string) *Client {
	return &Client{
		ArtistService: NewArtistService(httpClient, baseUrl, appId),
		EventService: NewEventService(httpClient, baseUrl, appId),
		VenueService: NewVenueService(httpClient, baseUrl, appId),
	}
}