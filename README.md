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

Below is an example usage of requesting [Recommended Events](https://www.bandsintown.com/api/1.0/requests#events-recommended), for a fan of Kayne West and 65daysofstatic in the upcoming year:

	package main
	
	import(
		"fmt"
		"time"
		"bandsintown-api/api_client"
		"bandsintown-api/datatypes"
	)
	
	func main() {
		client := api_client.NewClient("some_api_id" )
	
		params := datatypes.EventRecommendedParams{
			EventSearchParams : datatypes.EventSearchParams{
				Artists: []string{"Kayne West", "65daysofstatic"},
				Date: []time.Time{time.Now(), time.Now().AddDate(1, 0, 0) },
				Location: "London,UK",
			},
			OnlyRecommendations: true,
		}
	
		result, _, err := client.EventService.Recommended(params)
	
		fmt.Println("Error:")
		fmt.Println(err)
	
		fmt.Println("Recommended Events:")
		for _, item := range result {
			fmt.Println(item)
		}
	}
