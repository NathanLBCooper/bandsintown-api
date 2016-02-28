package apiclient

type ApiError struct {
	Errors []string `json:"errors"`
}