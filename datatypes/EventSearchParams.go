package datatypes

import (
	"time"
)

// EventSearchParams is the parameter for the Events Search api call
// https://www.bandsintown.com/api/1.0/requests#events-search
type EventSearchParams struct {
	Artists  []string
	Location string
	Radius   int
	Date     []time.Time
	Page     int
	PerPage  int
}
