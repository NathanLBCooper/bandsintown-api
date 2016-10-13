package datatypes

// EventOnSaleSoonParams is the parameter for the Events On Sale Soon api call
// https://www.bandsintown.com/api/1.0/requests#events-on-sale-soon
type EventOnSaleSoonParams struct {
	Location string `url:"location,omitempty"`
	Radius   int    `url:"radius,omitempty"`
}
