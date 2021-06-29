//
//
// File generated from our OpenAPI spec
//
//

package stripe

type PaymentSourceListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	Object     *string `form:"object"`
}
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
type PaymentSourceList struct {
	APIResource
	ListMeta
	Data []*PaymentSource `json:"data"`
}
