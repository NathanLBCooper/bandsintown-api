package main

import(
	"fmt"
	"bandsintown-api/apiclient"
)

func main() {
	client := apiclient.NewClient(nil)
	result := apiclient.ArtistInfo{}
	result, _, _ = client.ArtistService.GetInfo("Weezer")
	fmt.Println(result)
}