package apiclient

import(
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02T15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == "null" {
		return nil
	}
	ct.Time, err = time.Parse(customTimeFormat, string(b))
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if(ct == nil){
		return []byte("null"), nil
	}
	return []byte(ct.Time.Format(customTimeFormat)), nil
}
