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
		//Location: "Boston,MA",
		Date: []time.Time{ time.Now().AddDate(0,0,0), time.Now().AddDate(0,1,0) },
	}
	result, _, err = client.EventService.Search(params)
	fmt.Println("Results")
	fmt.Println(result)
	fmt.Println("err")
	fmt.Println(err)

	fmt.Println("items")
	for _, item := range result {
		fmt.Println(item)
	}
}