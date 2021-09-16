//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Shipping address for the order. Required if any of the SKUs are for products that have `shippable` set to true.
type OrderShippingParams struct {
	Address        *AddressParams `form:"address"`
	Carrier        *string        `form:"carrier"`
	Name           *string        `form:"name"`
	Phone          *string        `form:"phone"`
	TrackingNumber *string        `form:"tracking_number"`
}

// OrderParams is the set of parameters that can be used when creating an order.
type OrderParams struct {
	Params                 `form:"*"`
	Coupon                 *string              `form:"coupon"`
	Currency               *string              `form:"currency"`
	Customer               *string              `form:"customer"`
	Email                  *string              `form:"email"`
	Items                  []*OrderItemParams   `form:"items"`
	SelectedShippingMethod *string              `form:"selected_shipping_method"`
	Shipping               *OrderShippingParams `form:"shipping"`
	Status                 *string              `form:"status"`
}

// Filter orders based on when they were paid, fulfilled, canceled, or returned.
type OrderListStatusTransitionsParams struct {
	Canceled       *int64            `form:"canceled"`
	CanceledRange  *RangeQueryParams `form:"canceled"`
	Fulfilled      *int64            `form:"fulfilled"`
	FulfilledRange *RangeQueryParams `form:"fulfilled"`
	Paid           *int64            `form:"paid"`
	PaidRange      *RangeQueryParams `form:"paid"`
	Returned       *int64            `form:"returned"`
	ReturnedRange  *RangeQueryParams `form:"returned"`
}

// OrderListParams is the set of parameters that can be used when listing orders.
type OrderListParams struct {
	ListParams        `form:"*"`
	Created           *int64                            `form:"created"`
	CreatedRange      *RangeQueryParams                 `form:"created"`
	Customer          *string                           `form:"customer"`
	IDs               []*string                         `form:"ids"`
	Status            *string                           `form:"status"`
	StatusTransitions *OrderListStatusTransitionsParams `form:"status_transitions"`
	UpstreamIDs       []*string                         `form:"upstream_ids"`
}

// OrderPayParams is the set of parameters that can be used when paying orders.
type OrderPayParams struct {
	Params         `form:"*"`
	ApplicationFee *int64  `form:"application_fee"`
	Customer       *string `form:"customer"`
	Email          *string `form:"email"`
	Source         *string `form:"source"`
}

// OrderItemParams is the set of parameters describing an order item on order creation or update.
type OrderItemParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}

// List of items to return.
type OrderReturnOrderItemParams struct {
	Amount      *int64  `form:"amount"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}

// Return all or part of an order. The order must have a status of paid or fulfilled before it can be returned. Once all items have been returned, the order will become canceled or returned depending on which status the order started in.
type OrderReturnOrderParams struct {
	Params `form:"*"`
	Items  []*OrderReturnOrderItemParams `form:"items"`
}

// The estimated delivery date for the given shipping method. Can be either a specific date or a range.
type OrderShippingMethodDeliveryEstimate struct {
	Date     string `json:"date"`
	Earliest string `json:"earliest"`
	Latest   string `json:"latest"`
	Type     string `json:"type"`
}

// A list of supported shipping methods for this order. The desired shipping method can be specified either by updating the order, or when paying it.
type OrderShippingMethod struct {
	Amount           int64                                `json:"amount"`
	Currency         Currency                             `json:"currency"`
	DeliveryEstimate *OrderShippingMethodDeliveryEstimate `json:"delivery_estimate"`
	Description      string                               `json:"description"`
	ID               string                               `json:"id"`
}

// The timestamps at which the order status was updated.
type OrderStatusTransitions struct {
	Canceled int64 `json:"canceled"`
	Fulfiled int64 `json:"fulfiled"`
	Paid     int64 `json:"paid"`
	Returned int64 `json:"returned"`
}

// Order is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#orders_legacy.
type Order struct {
	APIResource
	Amount                 int64                   `json:"amount"`
	AmountReturned         int64                   `json:"amount_returned"`
	Application            string                  `json:"application"`
	ApplicationFee         int64                   `json:"application_fee"`
	Charge                 *Charge                 `json:"charge"`
	Created                int64                   `json:"created"`
	Currency               Currency                `json:"currency"`
	Customer               *Customer               `json:"customer"`
	Email                  string                  `json:"email"`
	ExternalCouponCode     string                  `json:"external_coupon_code"`
	ID                     string                  `json:"id"`
	Items                  []*OrderItem            `json:"items"`
	Livemode               bool                    `json:"livemode"`
	Metadata               map[string]string       `json:"metadata"`
	Object                 string                  `json:"object"`
	Returns                *OrderReturnList        `json:"returns"`
	SelectedShippingMethod string                  `json:"selected_shipping_method"`
	Shipping               *ShippingDetails        `json:"shipping"`
	ShippingMethods        []*OrderShippingMethod  `json:"shipping_methods"`
	Status                 string                  `json:"status"`
	StatusTransitions      *OrderStatusTransitions `json:"status_transitions"`
	Updated                int64                   `json:"updated"`
	UpstreamID             string                  `json:"upstream_id"`
}

// OrderList is a list of orders as retrieved from a list endpoint.
type OrderList struct {
	APIResource
	ListMeta
	Data []*Order `json:"data"`
}

// UnmarshalJSON handles deserialization of an Order.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *Order) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		o.ID = id
		return nil
	}

	type order Order
	var v order
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = Order(v)
	return nil
}
