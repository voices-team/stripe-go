//
//
// File generated from our OpenAPI spec
//
//

// Package lineitem provides the /checkout/sessions/{session}/line_items APIs
package lineitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /checkout/sessions/{session}/line_items APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of line items.
func List(params *stripe.LineItemListParams) *Iter {
	return getC().List(params)
}

// List returns a list of line items.
func (c Client) List(listParams *stripe.LineItemListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/checkout/sessions/%s/line_items",
		stripe.StringValue(listParams.Session),
	)
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.LineItemList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for line items.
type Iter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *Iter) LineItem() *stripe.LineItem {
	return i.Current().(*stripe.LineItem)
}

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) LineItemList() *stripe.LineItemList {
	return i.List().(*stripe.LineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
