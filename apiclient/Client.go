package apiclient

import(
	"net/http"
)

// Client for the Bands in Town API
type Client struct {
	ArtistService *ArtistService
	// other service endpoints...
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		ArtistService: NewArtistService(httpClient),
	}
}