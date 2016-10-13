package datatypes

import (
	"fmt"
)

// APIError represents the container for errors from the bandsintown API
type APIError struct {
	Errors []string `json:"errors"`
}

func (apiError APIError) Error() string {
	return fmt.Sprintf("response: %v", apiError.Errors)
}

// HasErrors returns whether there are errors in the conatiner
func (apiError *APIError) HasErrors() bool {
	return !(apiError == nil || len(apiError.Errors) == 0)
}
