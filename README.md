# bandsintown-api

Packages:
> github.com/NathanLBCooper/bandsintown-api/api_client
> 
> github.com/NathanLBCooper/bandsintown-api/datatypes

Installation:
>go get github.com/NathanLBCooper/bandsintown-api

### Overview:
A client library for the Bandsintown Concert API.

### Example usage:

Below is an example usage of the Recommended 

	package main
	
	import(
		"fmt"
		"time"
		"bandsintown-api/api_client"
		"bandsintown-api/datatypes"
	)
	
	func main() {
		client := api_client.NewClient("some_api_id" )
		var result []datatypes.Event
		var err error
	
		params := datatypes.EventRecommendedParams{
			EventSearchParams : datatypes.EventSearchParams{
				Artists: []string{"Weezer", "Kayne West"},
				Date: []time.Time{time.Now().AddDate(0, 0, 0), time.Now().AddDate(10, 1, 0) },
				Location: "London,UK",
			},
			OnlyRecommendations: true,
		}
	
		result, _, err = client.EventService.Recommended(params)
		fmt.Println("Results")
		fmt.Println(result)
		fmt.Println("err")
		fmt.Println(err)
	
		fmt.Println("items")
		for _, item := range result {
			fmt.Println(item)
		}
	}
