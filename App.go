package main

import(
	"fmt"
	"time"
	"bandsintown-api/apiclient"
	"bandsintown-api/datatypes"
)

func main() {
	client := apiclient.NewClient(nil, "http://api.bandsintown.com", "some_api_id" )
	var result []datatypes.Event
	var err error
	params := datatypes.EventSearchParams{
		Artists: []string{"Weezer", "Kayne West"},
		Datetime: []time.Time{ time.Now(), time.Now() }, // todo, not working
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