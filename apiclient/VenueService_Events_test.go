package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/NathanLBCooper/bandsintown-api-v1/datatypes"
)

func TestVenueEventsCanReceiveEventsResponse(test *testing.T) {
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

	venueID := 1043

	mux.HandleFunc(fmt.Sprintf("/venues/%v/events.json", venueID), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	// Act
	result, _, err := client.VenueService.Events(venueID)

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

func TestVenueEventsProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	venueID := 32
	eventPath := fmt.Sprintf("/venues/%v/events.json", venueID)
	var method, host, path, rawQuery string
	mux.HandleFunc(eventPath, func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	const expectedRawQuery = "app_id=myappId"

	client.VenueService.Events(venueID)

	if method != "GET" {
		test.Errorf("expected method to be GET, got %v", method)
	}
	if host != "example.com" {
		test.Errorf("expected host to be example.com, got %v", host)
	}
	if path != eventPath {
		test.Errorf("expected path to be %v, got %v", eventPath, path)
	}
	if rawQuery != expectedRawQuery {
		test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)
	}
}
