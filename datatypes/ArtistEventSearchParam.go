package datatypes

import "time"

type ArtistEventSearchParam struct {
	Name string
	MbId string
	Date []time.Time
}
