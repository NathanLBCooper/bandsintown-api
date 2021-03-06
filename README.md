# bandsintown-api-v1

**!!! It looks like the bandsintown api v1 has been discontinued (https://www.bandsintown.com/api/1.0/). Nothing to see here then.**

Packages:
> github.com/NathanLBCooper/bandsintown-api-v1/apiclient
> 
> github.com/NathanLBCooper/bandsintown-api-v1/datatypes

Installation:
>go get github.com/NathanLBCooper/bandsintown-api-v1

### Overview:
A client library for the Bandsintown Concert API.

### Example usage:

Below is an example usage of requesting [Recommended Events](https://www.bandsintown.com/api/1.0/requests#events-recommended), for a fan of Kayne West and 65daysofstatic in the upcoming year:

	package main

	import (
		"fmt"
		"time"

		"github.com/NathanLBCooper/bandsintown-api-v1/apiclient"
		"github.com/NathanLBCooper/bandsintown-api-v1/datatypes"
	)

	func main() {
		client := apiclient.NewClient("some_api_id")

		layout := "2006-01-02T15:04:05.000Z"
		startTime, _ := time.Parse(layout, "2015-11-12T11:45:26.371Z")
		endTime, _ := time.Parse(layout, "2016-11-12T11:45:26.371Z")

		params := datatypes.EventRecommendedParams{
			EventSearchParams: datatypes.EventSearchParams{
				Artists:  []string{"Kayne West", "65daysofstatic"},
				Date:     []time.Time{startTime, endTime},
				Location: "London,UK",
			},
			OnlyRecommendations: true,
		}

		result, _, err := client.EventService.Recommended(params)

		fmt.Println(result)

		fmt.Println("Error:")
		fmt.Println(err)

		fmt.Println("Recommended Events:")
		for _, item := range result {
			fmt.Println(item)
		}
	}

