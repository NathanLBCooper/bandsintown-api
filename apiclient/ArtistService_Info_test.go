package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

func TestArtistGetInfoCanReceiveResponse(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const artistName = "65daysofstatic"
	const actualResponse = `{"name":"65daysofstatic","url":"http://www.bandsintown.com/65daysofstatic",` +
		`"mbid":"0cd12ab3-9628-45ef-a97b-ff18624f14a0","upcoming_events_count":5}`

	expectedResponse := datatypes.ArtistInfo{
		Artist: datatypes.Artist{
			Name: artistName,
			URL:  "http://www.bandsintown.com/65daysofstatic",
			MbID: "0cd12ab3-9628-45ef-a97b-ff18624f14a0",
		},
		UpcomingGigCount: 5,
	}

	getPath := fmt.Sprintf("/artists/%v.json", artistName)
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := NewClientDetailed(httpClient, "http://example.com", "appid")

	// Act
	result, _, err := client.ArtistService.GetInfoByName(artistName)

	// Assert
	if err != nil {
		test.Errorf("expected err to be nil, got %v", err)
	}

	artistJSON, _ := json.Marshal(result)
	expectedArtistJSON, _ := json.Marshal(expectedResponse)

	if !bytes.Equal(artistJSON, expectedArtistJSON) {
		test.Errorf("Gig Json: expected %v, got %v", string(expectedArtistJSON), string(artistJSON))
	}
}

func TestArtistGetInfoByNameProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const artistName = "65daysofstatic"
	const appID = "appid"
	var method, host, path, rawQuery string
	getPath := fmt.Sprintf("/artists/%v.json", artistName)
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example.com", appID)

	expectedRawQuery := fmt.Sprintf("app_id=%v", appID)

	client.ArtistService.GetInfoByName(artistName)

	if method != "GET" {
		test.Errorf("expected method to be GET, got %v", method)
	}
	if host != "example.com" {
		test.Errorf("expected host to be example.com, got %v", host)
	}
	if path != getPath {
		test.Errorf("expected path to be %v, got %v", getPath, path)
	}
	if rawQuery != expectedRawQuery {
		test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)
	}
}

func TestArtistGetInfoByMbIDProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const MbID = "0cd12ab3-9628-45ef-a97b-ff18624f14a0"
	const appID = "appid"
	var method, host, path, rawQuery string
	getPath := fmt.Sprintf("/artists/mbid_%v.json", MbID)
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example.com", appID)

	expectedRawQuery := fmt.Sprintf("app_id=%v", appID)

	client.ArtistService.GetInfoByMbID(MbID)

	if method != "GET" {
		test.Errorf("expected method to be GET, got %v", method)
	}
	if host != "example.com" {
		test.Errorf("expected host to be example.com, got %v", host)
	}
	if path != getPath {
		test.Errorf("expected path to be %v, got %v", getPath, path)
	}
	if rawQuery != expectedRawQuery {
		test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)
	}
}
