package apiclient

import(
	"bandsintown-api/apiclient/customtimes"
)

// todo enforce required fields
// todo enforce field validations
type EventSearchParams struct {
	Artists []string `url:"artists[],omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Datetime []customtimes.SearchCustomTime `url:"datetime,omitempty"`
	Page int `url:"page,omitempty"`
	Perpage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

// todo make private and provide struct without ApiId