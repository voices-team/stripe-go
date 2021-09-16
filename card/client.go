//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the TODO APIs
package card

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke TODO APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Update updates a card's properties.
func Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Update(id, params)
}

// Update updates a card's properties.
func (c Client) Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Update updates a card's properties.
func Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Update(id, params)
}

// Update updates a card's properties.
func (c Client) Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s",
		stripe.StringValue(params.Account),
		id,
	)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Del removes a card.
func Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Del(id, params)
}

// Del removes a card.
func (c Client) Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}

// Del removes a card.
func Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Del(id, params)
}

// Del removes a card.
func (c Client) Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s",
		stripe.StringValue(params.Account),
		id,
	)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
