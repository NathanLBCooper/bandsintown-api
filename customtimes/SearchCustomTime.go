package customtimes

import(
	"time"
)

type CustomSearchTime struct {
	time.Time
}

const customSearchTimeFormat = "2006-01-02"

func (ct *CustomSearchTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == "null" {
		return nil
	}
	ct.Time, err = time.Parse(customSearchTimeFormat, string(b))
	return
}

func (ct *CustomSearchTime) MarshalJSON() ([]byte, error) {
	if(ct == nil){
		return []byte("null"), nil
	}
	return []byte(ct.Time.Format(customSearchTimeFormat)), nil
}

func NewCustomSearchTimes(times []time.Time) []CustomSearchTime{
	customTimes := make([]CustomSearchTime, len(times))
	for i,time := range times{
		customTimes[i] = CustomSearchTime{ Time: time }
	}

	return customTimes
}