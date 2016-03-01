package apiclient

import(
	"time"
)

type SearchCustomTime struct {
	time.Time
}

const searchCustomTimeFormat = "2006-01-02"

func (ct *SearchCustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == "null" {
		return nil
	}
	ct.Time, err = time.Parse(searchCustomTimeFormat, string(b))
	return
}

func (ct *SearchCustomTime) MarshalJSON() ([]byte, error) {
	if(ct == nil){
		return []byte("null"), nil
	}
	return []byte(ct.Time.Format(searchCustomTimeFormat)), nil
}