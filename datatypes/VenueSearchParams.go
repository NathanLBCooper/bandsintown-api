package datatypes

type VenueSearchParams struct {
	Query string
	Location string
	Radius int
	Page int
	PerPage int
}