package datatypes

// EventRecommendedParams is the parameter for the Events Recommmended api call
// https://www.bandsintown.com/api/1.0/requests#events-recommended
type EventRecommendedParams struct {
	EventSearchParams
	OnlyRecommendations bool
}
