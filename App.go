package main

import(
	"fmt"
	"bandsintown-api/apiclient"
)

func main() {
	client := apiclient.NewClient(nil)
	var result apiclient.Event
	var err error
	result, _, err = client.ArtistService.GetEvents("Weezer")
	fmt.Println(result)
	fmt.Println(err)
}