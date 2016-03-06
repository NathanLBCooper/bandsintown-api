package datatypes

import(
	"time"
)

type EventSearchParams struct {
	Artists []string
	Location string
	Radius int
	Date []time.Time
	Page int
	PerPage int
}