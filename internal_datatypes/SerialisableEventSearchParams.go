package internal_datatypes

import(
	"time"
	"strings"
	"bandsintown-api/datatypes"
)

const customSearchTimeFormat = "2006-01-02"

type SerialisableEventSearchParams struct {
	Artists []string `url:"artists[],omitempty"`
	Location string `url:"location,omitempty"`
	Radius int `url:"radius,omitempty"`
	Date string `url:"date,omitempty"`
	Page int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
	ApiId string `url:"app_id,omitempty"`
}

func formatSearchTimes(times []time.Time) string {
	timeStrs := make([]string, len(times))
	for i,time := range times{
		timeStrs[i] = time.Format(customSearchTimeFormat)
	}

	return strings.Join(timeStrs, ",");
}

// Create SerialisableEventSearchParams from EventSearchParams
func NewSerialisableEventSearchParams(params* datatypes.EventSearchParams, appId string)*SerialisableEventSearchParams {
	return &SerialisableEventSearchParams{
		Artists: params.Artists,
		Location: params.Location,
		Radius: params.Radius,
		Date: formatSearchTimes(params.Date),
		Page: params.Page,
		PerPage: params.PerPage,
		ApiId: appId,
	}
}