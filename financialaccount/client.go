// Package financialaccount provides the /financial_accounts APIs
package financialaccount

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /financial_accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new financial account.
func New(params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	return getC().New(params)
}

// New creates a new financial account.
func (c Client) New(params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	if params == nil {
		params = &stripe.FinancialAccountParams{}
	}

	// force 'Stripe-Version' header to enable financial accounts beta feature
	params.Headers = map[string][]string{
		"Stripe-Version": {fmt.Sprintf("%s;financial_accounts_beta=v3", stripe.APIVersion)},
	}

	account := &stripe.FinancialAccount{}
	err := c.B.Call(http.MethodPost, "/v1/financial_accounts", c.Key, params, account)
	return account, err
}

// GetByID returns the details of a financial account.
func GetByID(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	return getC().GetByID(id, params)
}

// GetByID returns the details of a financial account.
func (c Client) GetByID(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/financial_accounts/%s", id)
	account := &stripe.FinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Update updates a financial account's properties.
func Update(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	return getC().Update(id, params)
}

// Update updates a financial account's properties.
func (c Client) Update(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/financial_accounts/%s", id)
	account := &stripe.FinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Close closes a financial account.
func Close(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	return getC().Close(id, params)
}

// Close closes a financial account.
func (c Client) Close(id string, params *stripe.FinancialAccountParams) (*stripe.FinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/financial_accounts/%s/close", id)
	account := &stripe.FinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// List returns a list of financial accounts.
func List(params *stripe.FinancialAccountListParams) *Iter {
	return getC().List(params)
}

// List returns a list of financial accounts.
func (c Client) List(listParams *stripe.FinancialAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/financial_accounts", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for accounts.
type Iter struct {
	*stripe.Iter
}

// FinancialAccount returns the account which the iterator is currently pointing to.
func (i *Iter) FinancialAccount() *stripe.FinancialAccount {
	return i.Current().(*stripe.FinancialAccount)
}

// FinancialAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialAccountList() *stripe.FinancialAccountList {
	return i.List().(*stripe.FinancialAccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
