//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List sources for a specified customer.
type PaymentSourceListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	Object     *string `form:"object"`
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://stripe.com/docs/api#update_customer) to have a new default_source.
type PaymentSourceParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	Source   *string `form:"source"`
}

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BankAccount or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	APIResource
	Deleted bool `json:"deleted"`
}

// PaymentSourceList is a list of PaymentSources as retrieved from a list endpoint.
type PaymentSourceList struct {
	APIResource
	ListMeta
	Data []*PaymentSource `json:"data"`
}
