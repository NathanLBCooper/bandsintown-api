package main

import(
	"fmt"
	"time"
	"bandsintown-api/apiclient"
	"bandsintown-api/apiclient/customtimes"
)

func main() {
	client := apiclient.NewClient(nil)
	var result []apiclient.Event
	var err error
	params := apiclient.EventSearchParams{
		Artists: []string{"Weezer", "Kayne West"},
		Datetime: []customtimes.SearchCustomTime{customtimes.SearchCustomTime{ Time: time.Now() }, customtimes.SearchCustomTime{ Time: time.Now() }}, //todo not working
	}
	result, _, err = client.EventService.Search(params)
	fmt.Println(result)
	fmt.Println(err)

	for _, item := range result {
		fmt.Println(item)
	}

	fmt.Println(result[0].Artists)
	fmt.Println(result[0].Datetime)
	fmt.Println(result[0].TicketUrl)
}