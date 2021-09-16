//
//
// File generated from our OpenAPI spec
//
//

// Package paymentsource provides the /customers/{customer}/sources APIs
package paymentsource

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /customers/{customer}/sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payment source.
func New(params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	return getC().New(params)
}

// New creates a new payment source.
func (c Client) New(params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources",
		stripe.StringValue(params.Customer),
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Get returns the details of a payment source.
func Get(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payment source.
func (c Client) Get(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// List returns a list of payment sources.
func List(params *stripe.PaymentSourceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payment sources.
func (c Client) List(listParams *stripe.PaymentSourceListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources",
		stripe.StringValue(listParams.Customer),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentSource{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment sources.
type Iter struct {
	*stripe.Iter
}

// PaymentSource returns the payment source which the iterator is currently pointing to.
func (i *Iter) PaymentSource() *stripe.PaymentSource {
	return i.Current().(*stripe.PaymentSource)
}

// PaymentSourceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentSourceList() *stripe.PaymentSourceList {
	return i.List().(*stripe.PaymentSourceList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
