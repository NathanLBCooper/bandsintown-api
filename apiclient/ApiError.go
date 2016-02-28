package apiclient

import(
	"fmt"
)

type ApiError struct {
	Errors []string `json:"errors"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("response: %v", e.Errors)
}
