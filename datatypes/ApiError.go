package datatypes

import(
	"fmt"
)

type ApiError struct {
	Errors []string `json:"errors"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("response: %v", e.Errors)
}

func (apiError *ApiError) HasErrors() (bool) {
	return !(apiError == nil || len(apiError.Errors) == 0)
}
