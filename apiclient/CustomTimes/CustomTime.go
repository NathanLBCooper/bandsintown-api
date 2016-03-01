package customtimes

import(
	"time"
)

type ResponseCustomTime struct {
	time.Time
}

const responseCustomTimeFormat = "2006-01-02T15:04:05"

func (ct *ResponseCustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == "null" {
		return nil
	}
	ct.Time, err = time.Parse(responseCustomTimeFormat, string(b))
	return
}

func (ct *ResponseCustomTime) MarshalJSON() ([]byte, error) {
	if(ct == nil){
		return []byte("null"), nil
	}
	return []byte(ct.Time.Format(responseCustomTimeFormat)), nil
}
