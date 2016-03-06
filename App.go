package main

import(
	"fmt"
//	"time"
	"bandsintown-api/apiclient"
	"bandsintown-api/datatypes"
)

func main() {
	client := apiclient.NewClient(nil, "http://api.bandsintown.com", "some_api_id" )
	var result []datatypes.Event
	var err error


	/*params := datatypes.EventRecommendedParams{
		EventSearchParams : datatypes.EventSearchParams{
			Artists: []string{"Weezer", "Kayne West"},
			Date: []time.Time{time.Now().AddDate(0, 0, 0), time.Now().AddDate(10, 1, 0) },
			Location: "London,UK",
		},
		OnlyRecommendations: true,
	}*/

	/*params := datatypes.EventOnSaleSoonParams{
		Location: "London,UK",
		Radius: 150,
	}*/


	result, _, err = client.EventService.Daily()
	fmt.Println("Results")
	fmt.Println(result)
	fmt.Println("err")
	fmt.Println(err)

	fmt.Println("items")
	for _, item := range result {
		fmt.Println(item)
	}
}