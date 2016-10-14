package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/NathanLBCooper/bandsintown-api/datatypes"
)

func TestArtistGetEventCanReceiveResponse(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const actualResponse = `[{"id":11224258,"url":"http://www.bandsintown.com/event/11224258?app_id=myappId",` +
		`"datetime":"2016-04-05T19:00:00","ticket_url":"http://www.bandsintown.com/event/11224258/buy_tickets?` +
		`app_id=myappId\u0026came_from=233","artists":[{"name":"Weezer","url":"http://www.bandsintown.com/Weezer",` +
		`"mbid":"6fe07aa5-fec0-4eca-a456-f29bff451b04"}],"venue":{"id":1015552,"url":"http://www.bandsintown.com/` +
		`venue/1015552","name":"O2 BRIXTON ACADEMY","city":"Brixton","region":"London","country":"United Kingdom",` +
		`"latitude":51.4620184,"longitude":-0.1152248},"ticket_status":"available","on_sale_datetime":null}]`

	expectedResponse := datatypes.Event{
		ID:        11224258,
		URL:       "http://www.bandsintown.com/event/11224258?app_id=myappId",
		Datetime:  time.Date(2016, 04, 05, 19, 0, 0, 0, time.UTC),
		TicketURL: "http://www.bandsintown.com/event/11224258/buy_tickets?app_id=myappId&came_from=233",
		Artists: []datatypes.Artist{
			datatypes.Artist{
				Name: "Weezer",
				MbID: "6fe07aa5-fec0-4eca-a456-f29bff451b04",
				URL:  "http://www.bandsintown.com/Weezer",
			},
		},
		Venue: datatypes.Venue{
			ID:        1015552,
			Name:      "O2 BRIXTON ACADEMY",
			City:      "Brixton",
			Region:    "London",
			Country:   "United Kingdom",
			URL:       "http://www.bandsintown.com/venue/1015552",
			Latitude:  51.4620184,
			Longitude: -0.1152248,
		},
		TicketStatus:   "available",
		OnSaleDatetime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	const artistName = "Foo"
	getPath := fmt.Sprintf("/artists/%v/events.json", artistName)
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	// Doesn't matter
	params := datatypes.ArtistEventSearchParam{
		Name: artistName,
	}

	// Act
	result, _, err := client.ArtistService.GetEvents(params)

	// Assert
	if err != nil {
		test.Errorf("expected err to be nil, got %v", err)
	}

	if len(result) != 1 {
		test.Errorf("expected len(result) to be 1, got %v", len(result))
	}

	gig := result[0]
	gigJSON, _ := json.Marshal(gig)
	expectedGigJSON, _ := json.Marshal(expectedResponse)

	if !bytes.Equal(gigJSON, expectedGigJSON) {
		test.Errorf("Gig Json: expected %v, got %v", string(gigJSON), string(expectedGigJSON))
	}
}

func TestArtistGetEventsByNameProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const artistName = "Maybeshewill"
	getPath := fmt.Sprintf("/artists/%v/events.json", artistName)
	var method, host, path, rawQuery string
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	params := datatypes.ArtistEventSearchParam{
		Name: artistName,
		Date: []time.Time{
			time.Date(2016, time.October, 5, 1, 2, 3, 4, time.UTC),
			time.Date(2017, time.November, 4, 3, 8, 9, 12, time.UTC),
		},
	}

	const expectedRawQuery = "app_id=myappId&date=2016-10-05%2C2017-11-04"

	client.ArtistService.GetEvents(params)

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

func TestArtistGetEventsByMbIDProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const artistName = "Maybeshewill"
	const MbID = "62495254-237e-4e9e-8ffb-31fede562cfd"
	getPath := fmt.Sprintf("/artists/%v/events.json", MbID)
	var method, host, path, rawQuery string
	mux.HandleFunc(getPath, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	mux.HandleFunc(fmt.Sprintf("/artists/%v/events.json", artistName),
		func(w http.ResponseWriter, r *http.Request) {
			test.Errorf("The Api should be called with the mbid, not the artist name")
		},
	)

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	// MbID will take presidence over Name
	params := datatypes.ArtistEventSearchParam{
		Name: artistName,
		MbID: MbID,
		Date: []time.Time{
			time.Date(2016, time.October, 5, 1, 2, 3, 4, time.UTC),
			time.Date(2017, time.November, 4, 3, 8, 9, 12, time.UTC),
		},
	}

	const expectedRawQuery = "app_id=myappId&date=2016-10-05%2C2017-11-04"

	client.ArtistService.GetEvents(params)

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
