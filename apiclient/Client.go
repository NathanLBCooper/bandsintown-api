package apiclient

import (
	"net/http"
)

const format = "json"
const bandsInTownBaseURL = "http://api.bandsintown.com"

// Client for the Bands in Town API
type Client struct {
	// Service Implementations
	ArtistService *ArtistService
	EventService  *EventService
	VenueService  *VenueService
}

// NewClientDetailed returns a new Client
func NewClientDetailed(httpClient *http.Client, baseURL string, appID string) *Client {
	return &Client{
		ArtistService: NewArtistService(httpClient, baseURL, appID),
		EventService:  NewEventService(httpClient, baseURL, appID),
		VenueService:  NewVenueService(httpClient, baseURL, appID),
	}
}

// NewClient returns a new Client // todo, this is nil???
func NewClient(appID string) *Client {
	return NewClientDetailed(nil, bandsInTownBaseURL, appID)
}
