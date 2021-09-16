//
//
// File generated from our OpenAPI spec
//
//

// Package bankaccount provides the TODO APIs
package bankaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke TODO APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Update updates a bank account's properties.
func Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Update(id, params)
}

// Update updates a bank account's properties.
func (c Client) Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Update updates a bank account's properties.
func Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Update(id, params)
}

// Update updates a bank account's properties.
func (c Client) Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s",
		stripe.StringValue(params.Account),
		id,
	)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Del removes a bank account.
func Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Del(id, params)
}

// Del removes a bank account.
func (c Client) Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Del removes a bank account.
func Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Del(id, params)
}

// Del removes a bank account.
func (c Client) Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s",
		stripe.StringValue(params.Account),
		id,
	)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Verify is the method for the `POST /v1/customers/{customer}/sources/{id}/verify` API.
func Verify(id string, params *stripe.BankAccountVerifyParams) (*stripe.BankAccount, error) {
	return getC().Verify(id, params)
}

// Verify is the method for the `POST /v1/customers/{customer}/sources/{id}/verify` API.
func (c Client) Verify(id string, params *stripe.BankAccountVerifyParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s/verify",
		stripe.StringValue(params.Customer),
		id,
	)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
