package internal_datatypes

import(
	"time"
	"strings"
)

const customSearchTimeFormat = "2006-01-02"

func formatSearchTimes(times []time.Time) string {
	timeStrs := make([]string, len(times))
	for i,time := range times{
		timeStrs[i] = time.Format(customSearchTimeFormat)
	}

	return strings.Join(timeStrs, ",");
}