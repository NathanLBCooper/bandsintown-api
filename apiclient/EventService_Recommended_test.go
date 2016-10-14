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

func TestRecommendedCanReceiveRecommendResponse(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const actualResponse = `[{"id":21224258,"url":"https://soundcloud.com/realironchef?app_id=myappId",` +
		`"datetime":"2017-03-04T20:01:02","ticket_url":"https://ironchef.bandcamp.com?` +
		`app_id=myappId\u0026came_from=233","artists":[{"name":"Iron Chef","url":"https://ironchef.bandcamp.com",` +
		`"mbid":"7fe07aa5-fec0-4eca-a456-f29bff451b04"}],"venue":{"id":2015552,"url":"http://www.bandsintown.com/` +
		`venue/9000","name":"Purple Turtle","city":"Reading","region":"Berkshire","country":"Wessex",` +
		`"latitude":57.6000,"longitude":13.6833},"ticket_status":"unavailable","on_sale_datetime":null}]`

	expectedResponse := datatypes.Event{
		ID:        21224258,
		URL:       "https://soundcloud.com/realironchef?app_id=myappId",
		Datetime:  time.Date(2017, 03, 04, 20, 1, 2, 0, time.UTC),
		TicketURL: "https://ironchef.bandcamp.com?app_id=myappId&came_from=233",
		Artists: []datatypes.Artist{
			datatypes.Artist{
				Name: "Iron Chef",
				MbID: "7fe07aa5-fec0-4eca-a456-f29bff451b04",
				URL:  "https://ironchef.bandcamp.com",
			},
		},
		Venue: datatypes.Venue{
			ID:        2015552,
			Name:      "Purple Turtle",
			City:      "Reading",
			Region:    "Berkshire",
			Country:   "Wessex",
			URL:       "http://www.bandsintown.com/venue/9000",
			Latitude:  57.6000,
			Longitude: 13.6833,
		},
		TicketStatus:   "unavailable",
		OnSaleDatetime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	mux.HandleFunc("/events/recommended.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	// Doesn't matter
	params := datatypes.EventRecommendedParams{
		EventSearchParams: datatypes.EventSearchParams{
			Artists:  []string{"Foo"},
			Location: "London,UK",
			Radius:   10,
		},
		OnlyRecommendations: false,
	}
	// Act
	result, _, err := client.EventService.Recommended(params)

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

func TestRecommendedProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	var method, host, path, rawQuery string
	mux.HandleFunc("/events/recommended.json", func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example2.com", "myappId")

	params := datatypes.EventRecommendedParams{
		EventSearchParams: datatypes.EventSearchParams{
			Artists:  []string{"Foo", "Bar"},
			Location: "London,UK",
			Date: []time.Time{
				time.Date(2016, time.March, 1, 2, 3, 4, 5, time.UTC),
				time.Date(2017, time.April, 6, 7, 9, 10, 11, time.UTC),
			},
			Radius:  12,
			Page:    2,
			PerPage: 1,
		},
		OnlyRecommendations: true,
	}

	const expectedRawQuery = "app_id=myappId&artists%5B%5D=Foo&artists%5B%5D=Bar&date=2016-03-01%2C2017-04-06" +
		"&location=London%2CUK&only_recs=true&page=2&per_page=1&radius=12"

	client.EventService.Recommended(params)

	if method != "GET" {
		test.Errorf("expected method to be GET, got %v", method)
	}
	if host != "example2.com" {
		test.Errorf("expected host to be example.com, got %v", host)
	}
	if path != "/events/recommended.json" {
		test.Errorf("expected path to be /events/search.json, got %v", path)
	}
	if rawQuery != expectedRawQuery {
		test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)
	}
}
