package api_client

import (
	"fmt"
	"testing"
	"time"
	"net/http"
	"net/http/httptest"
	"net/url"
	"encoding/json"
	"bandsintown-api/datatypes"
	"bytes"
)

func TestSearchCanReceiveSearchResponse(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	const actualResponse = `[{"id":11224258,"url":"http://www.bandsintown.com/event/11224258?app_id=myappId",` +
	`"datetime":"2016-04-05T19:00:00","ticket_url":"http://www.bandsintown.com/event/11224258/buy_tickets?` +
	`app_id=myappId\u0026came_from=233","artists":[{"name":"Weezer","url":"http://www.bandsintown.com/Weezer",` +
	`"mbid":"6fe07aa5-fec0-4eca-a456-f29bff451b04"}],"venue":{"id":1015552,"url":"http://www.bandsintown.com/` +
	`venue/1015552","name":"O2 BRIXTON ACADEMY","city":"Brixton","region":"London","country":"United Kingdom",` +
	`"latitude":51.4620184,"longitude":-0.1152248},"ticket_status":"available","on_sale_datetime":null}]`

	expectedResponse := datatypes.Event{
		Id: 11224258,
		Url: "http://www.bandsintown.com/event/11224258?app_id=myappId",
		Datetime: time.Date(2016, 04, 05, 19, 0, 0, 0, time.UTC),
		TicketUrl: "http://www.bandsintown.com/event/11224258/buy_tickets?app_id=myappId&came_from=233",
		Artists: []datatypes.Artist{
			datatypes.Artist{
				Name: "Weezer",
				Mbid: "6fe07aa5-fec0-4eca-a456-f29bff451b04",
				Url: "http://www.bandsintown.com/Weezer",
			},
		},
		Venue: datatypes.Venue{
			Id: 1015552,
			Name: "O2 BRIXTON ACADEMY",
			City: "Brixton",
			Region: "London",
			Country: "United Kingdom",
			Url: "http://www.bandsintown.com/venue/1015552",
			Latitude: 51.4620184,
			Longitude: -0.1152248,
		},
		TicketStatus: "available",
		OnSaleDatetime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	mux.HandleFunc("/events/search.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, actualResponse)
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	// Doesn't matter
	params := datatypes.EventSearchParams{
		Artists: []string{"Foo"},
		Location: "London,UK",
		Radius: 10,
	}

	// Act
	result, _, err := client.EventService.Search(params)

	// Assert
	if(err != nil){
		test.Errorf("expected err to be nil, got %v", err)
	}

	if(len(result) != 1){
		test.Errorf("expected len(result) to be 1, got %v", len(result))
	}

	gig := result[0]
	gigJson, _ := json.Marshal(gig)
	expectedGigJson, _ := json.Marshal(expectedResponse)

	if(!bytes.Equal(gigJson, expectedGigJson)){
		test.Errorf("Gig Json: expected %v, got %v", string(gigJson), string(expectedGigJson))
	}
}

func TestSearchProvidesCorrectQuery(test *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	var method, host, path, rawQuery string
	mux.HandleFunc("/events/search.json", func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		host = r.Host
		path = r.URL.Path
		rawQuery = r.URL.RawQuery
	})

	client := NewClientDetailed(httpClient, "http://example.com", "myappId")

	params := datatypes.EventSearchParams{
		Artists: []string{"Foo"},
		Location: "London,UK",
		Date: []time.Time {
			time.Date(2016, time.January, 1, 2, 3, 4, 5, time.UTC),
			time.Date(2017, time.February, 6, 7, 9, 10, 11, time.UTC),
		},
		Radius: 10,
		Page: 1,
		PerPage: 100,
	}

	const expectedRawQuery = "app_id=myappId&artists%5B%5D=Foo&date=2016-01-01%2C2017-02-06" +
	"&location=London%2CUK&page=1&per_page=100&radius=10"

	client.EventService.Search(params)

	if(method != "GET"){test.Errorf("expected method to be GET, got %v", method)}
	if(host != "example.com"){test.Errorf("expected host to be example.com, got %v", host)}
	if(path != "/events/search.json"){test.Errorf("expected path to be /events/search.json, got %v", path)}
	if(rawQuery != expectedRawQuery){test.Errorf("expected rawQuery to be %v, got %v", expectedRawQuery, rawQuery)}
}

func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	client := &http.Client{Transport: transport}
	return client, mux, server
}