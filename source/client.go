//
//
// File generated from our OpenAPI spec
//
//

// Package source provides the /sources APIs
package source

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new source.
func New(params *stripe.SourceParams) (*stripe.Source, error) {
	return getC().New(params)
}

// New creates a new source.
func (c Client) New(params *stripe.SourceParams) (*stripe.Source, error) {
	source := &stripe.Source{}
	err := c.B.Call(http.MethodPost, "/v1/sources", c.Key, params, source)
	return source, err
}

// Get returns the details of a source.
func Get(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	return getC().Get(id, params)
}

// Get returns the details of a source.
func (c Client) Get(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/v1/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, source)
	return source, err
}

// Update updates a source's properties.
func Update(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	return getC().Update(id, params)
}

// Update updates a source's properties.
func (c Client) Update(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/v1/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, source)
	return source, err
}

// Detach is the method for the `DELETE /v1/customers/{customer}/sources/{id}` API.
func Detach(id string, params *stripe.SourceDetachParams) (*stripe.Source, error) {
	return getC().Detach(id, params)
}

// Detach is the method for the `DELETE /v1/customers/{customer}/sources/{id}` API.
func (c Client) Detach(id string, params *stripe.SourceDetachParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, source)
	return source, err
}

// Verify is the method for the `POST /v1/sources/{source}/verify` API.
func Verify(id string, params *stripe.SourceVerifyParams) (*stripe.Source, error) {
	return getC().Verify(id, params)
}

// Verify is the method for the `POST /v1/sources/{source}/verify` API.
func (c Client) Verify(id string, params *stripe.SourceVerifyParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/v1/sources/%s/verify", id)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, source)
	return source, err
}

// SourceTransactions is the method for the `GET /v1/sources/{source}/source_transactions` API.
func SourceTransactions(params *stripe.SourceSourceTransactionsParams) *TransactionIter {
	return getC().SourceTransactions(params)
}

// SourceTransactions is the method for the `GET /v1/sources/{source}/source_transactions` API.
func (c Client) SourceTransactions(listParams *stripe.SourceSourceTransactionsParams) *TransactionIter {
	path := stripe.FormatURLPath(
		"/v1/sources/%s/source_transactions",
		stripe.StringValue(listParams.Source),
	)
	return &TransactionIter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.SourceTransactionList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// TransactionIter is an iterator for source transactions.
type TransactionIter struct {
	*stripe.Iter
}

// SourceTransaction returns the source transaction which the iterator is currently pointing to.
func (i *TransactionIter) SourceTransaction() *stripe.SourceTransaction {
	return i.Current().(*stripe.SourceTransaction)
}

// SourceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *TransactionIter) SourceTransactionList() *stripe.SourceTransactionList {
	return i.List().(*stripe.SourceTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
