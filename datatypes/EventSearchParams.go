package datatypes

import(
	"time"
)

// todo enforce required fields
// todo enforce field validations
type EventSearchParams struct {
	Artists []string
	Location string
	Radius int
	Date []time.Time
	Page int
	Perpage int
}