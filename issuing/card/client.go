//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /issuing/cards APIs
package card

import (
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /issuing/cards APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing card.
func New(params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().New(params)
}

// New creates a new issuing card.
func (c Client) New(params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	// force 'Stripe-Version' header to enable financial accounts beta feature
	if params == nil {
		params = &stripe.IssuingCardParams{}
	}
	params.Beta = &stripe.TreasuryBetaFeatures

	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, "/v1/issuing/cards", c.Key, params, card)
	return card, err
}

// Get returns the details of an issuing card.
func Get(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing card.
func (c Client) Get(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	// force 'Stripe-Version' header to enable financial accounts beta feature
	if params == nil {
		params = &stripe.IssuingCardParams{}
	}
	params.Beta = &stripe.TreasuryBetaFeatures

	path := stripe.FormatURLPath("/v1/issuing/cards/%s", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Update updates an issuing card's properties.
func Update(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().Update(id, params)
}

// Update updates an issuing card's properties.
func (c Client) Update(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	// force 'Stripe-Version' header to enable financial accounts beta feature
	if params == nil {
		params = &stripe.IssuingCardParams{}
	}
	params.Beta = &stripe.TreasuryBetaFeatures

	path := stripe.FormatURLPath("/v1/issuing/cards/%s", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// List returns a list of issuing cards.
func List(params *stripe.IssuingCardListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing cards.
func (c Client) List(listParams *stripe.IssuingCardListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			// force 'Stripe-Version' header to enable financial accounts beta feature
			if p == nil {
				p = &stripe.Params{}
			}
			p.Beta = &stripe.TreasuryBetaFeatures

			list := &stripe.IssuingCardList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cards", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing cards.
type Iter struct {
	*stripe.Iter
}

// IssuingCard returns the issuing card which the iterator is currently pointing to.
func (i *Iter) IssuingCard() *stripe.IssuingCard {
	return i.Current().(*stripe.IssuingCard)
}

// IssuingCardList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardList() *stripe.IssuingCardList {
	return i.List().(*stripe.IssuingCardList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
