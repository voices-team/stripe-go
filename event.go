//
//
// File generated from our OpenAPI spec
//
//

package stripe

// EventParams is the set of parameters that can be used when retrieving events.
// For more details see https://stripe.com/docs/api#retrieve_events.
type EventParams struct {
	Params `form:"*"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams      `form:"*"`
	Created         *int64            `form:"created"`
	CreatedRange    *RangeQueryParams `form:"created"`
	DeliverySuccess *bool             `form:"delivery_success"`
	Type            *string           `form:"type"`
	Types           []*string         `form:"types"`
}

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	APIResource
	Account         string        `json:"account"`
	APIVersion      string        `json:"api_version"`
	Created         int64         `json:"created"`
	Data            *EventData    `json:"data"`
	ID              string        `json:"id"`
	Livemode        bool          `json:"livemode"`
	Object          string        `json:"object"`
	PendingWebhooks int64         `json:"pending_webhooks"`
	Request         *EventRequest `json:"request"`
	Type            string        `json:"type"`
}

// EventRequest contains information on a request that created an event.
type EventRequest struct {
	ID             string `json:"id"`
	IdempotencyKey string `json:"idempotency_key"`
}
type EventDataObject struct{}
type EventDataPreviousAttributes struct{}

// EventData is the unmarshalled object as a map.
type EventData struct {
	Object             *EventDataObject             `json:"object"`
	PreviousAttributes *EventDataPreviousAttributes `json:"previous_attributes"`
}

// EventList is a list of events as retrieved from a list endpoint.
type EventList struct {
	APIResource
	ListMeta
	Data []*Event `json:"data"`
}
