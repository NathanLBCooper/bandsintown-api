package apiclient

import(
	"net/http"
)

const format = "json"

// todo remove
const baseUrl = "http://api.bandsintown.com" //baseUrl
const appId = "some_api_id" // appId

// Client for the Bands in Town API
type Client struct {
	// Service Implementations
	ArtistService *ArtistService
	// todo, more services
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		ArtistService: NewArtistService(httpClient, baseUrl, appId),
	}
}