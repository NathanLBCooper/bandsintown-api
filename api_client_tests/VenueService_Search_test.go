package api_client_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/NathanLBCooper/bandsintown-api/api_client"
	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

func TestVenueSearchCanReceiveSearchResponse(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const actualResponse = `[{"id":1015552,"url":"http://www.bandsintown.com/` +
		`venue/1015552","name":"O2 BRIXTON ACADEMY","city":"Brixton","region":"London","country":"United Kingdom",` +
		`"latitude":51.4620184,"longitude":-0.1152248}]`

	expectedResponse := datatypes.Venue{
		Id:        1015552,
		Name:      "O2 BRIXTON ACADEMY",
		City:      "Brixton",
		Region:    "London",
		Country:   "United Kingdom",
		Url:       "http://www.bandsintown.com/venue/1015552",
		Latitude:  51.4620184,
		Longitude: -0.1152248,
	}

	mux.HandleFunc("/venues/search.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := api_client.NewClientDetailed(httpClient, "http://example.com", "myappId")

	// Doesn't matter
	params := datatypes.VenueSearchParams{
		Query: "Foo",
	}

	// Act
	result, _, err := client.VenueService.Search(params)

	// Assert
	if err != nil {
		test.Errorf("expected err to be nil, got %v", err)
	}

	if len(result) != 1 {
		test.Errorf("expected len(result) to be 1, got %v", len(result))
	}

	gig := result[0]
	gigJson, _ := json.Marshal(gig)
	expectedGigJson, _ := json.Marshal(expectedResponse)

	if !bytes.Equal(gigJson, expectedGigJson) {
		test.Errorf("Gig Json: expected %v, got %v", string(gigJson), string(expectedGigJson))
	}
}

func TestVenueSearchProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	var method, host, path, rawQuery string
	mux.HandleFunc("/venues/search.json", func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := api_client.NewClientDetailed(httpClient, "http://example.com", "myappId")

	params := datatypes.VenueSearchParams{
		Query:    "somequerystr",
		Location: "Bath,UK",
		Radius:   4,
		Page:     3,
		PerPage:  9,
	}

	const expectedRawQuery = "app_id=myappId&location=Bath%2CUK&page=3&per_page=9&query=somequerystr&radius=4"

	client.VenueService.Search(params)

	if method != "GET" {
		test.Errorf("expected method to be GET, got %v", method)
	}
	if host != "example.com" {
		test.Errorf("expected host to be example.com, got %v", host)
	}
	if path != "/venues/search.json" {
		test.Errorf("expected path to be /events/search.json, got %v", path)
	}
	if rawQuery != expectedRawQuery {
		test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)
	}
}
