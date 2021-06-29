//
//
// File generated from our OpenAPI spec
//
//

package stripe

// OrderReturnParams is the set of parameters that can be used when returning orders.
type OrderReturnParams struct {
	Params `form:"*"`
}

// OrderReturnListParams is the set of parameters that can be used when listing order returns.
type OrderReturnListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Order        *string           `form:"order"`
}

// OrderReturn is the resource representing an order return.
// For more details see https://stripe.com/docs/api#order_returns.
type OrderReturn struct {
	APIResource
	Amount   int64        `json:"amount"`
	Created  int64        `json:"created"`
	Currency Currency     `json:"currency"`
	ID       string       `json:"id"`
	Items    []*OrderItem `json:"items"`
	Livemode bool         `json:"livemode"`
	Object   string       `json:"object"`
	Order    *Order       `json:"order"`
	Refund   *Refund      `json:"refund"`
}

// OrderReturnList is a list of order returns as retrieved from a list endpoint.
type OrderReturnList struct {
	APIResource
	ListMeta
	Data []*OrderReturn `json:"data"`
}
