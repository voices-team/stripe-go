//
//
// File generated from our OpenAPI spec
//
//

package stripe

// EphemeralKeyParams is the set of parameters that can be used when creating
// an ephemeral key.
type EphemeralKeyParams struct {
	Params      `form:"*"`
	Customer    *string `form:"customer"`
	IssuingCard *string `form:"issuing_card"`
}

// EphemeralKey is the resource representing a Stripe ephemeral key. This is used by Mobile SDKs
// to for example manage a Customer's payment methods.
type EphemeralKey struct {
	APIResource
	Created  int64  `json:"created"`
	Expires  int64  `json:"expires"`
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	Object   string `json:"object"`
	Secret   string `json:"secret"`
}
