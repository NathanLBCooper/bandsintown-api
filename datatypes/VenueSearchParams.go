package datatypes

// VenueSearchParams is the parameter for the Venue Search api call
// https://www.bandsintown.com/api/1.0/requests#venues-search
type VenueSearchParams struct {
	Query    string
	Location string
	Radius   int
	Page     int
	PerPage  int
}
