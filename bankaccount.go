//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// BankAccountAvailablePayoutMethod is a set of available payout methods for the card.
type BankAccountAvailablePayoutMethod string

// List of values that CardAvailablePayoutMethod can take.
const (
	BankAccountAvailablePayoutMethodInstant  BankAccountAvailablePayoutMethod = "instant"
	BankAccountAvailablePayoutMethodStandard BankAccountAvailablePayoutMethod = "standard"
)

type BankAccountOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// BankAccountParams is the set of parameters that can be used when updating a
// bank account.
// //
// Note that while form annotations are used for updates, bank accounts have
// some unusual logic on creates that necessitates manual handling of all
// parameters. See AppendToAsSourceOrExternalAccount.
type BankAccountParams struct {
	Params             `form:"*"`
	Account            *string                 `form:"-"` // Included in URL
	Customer           *string                 `form:"-"` // Included in URL
	AccountHolderName  *string                 `form:"account_holder_name"`
	AccountHolderType  *string                 `form:"account_holder_type"`
	AccountType        *string                 `form:"account_type"`
	AddressCity        *string                 `form:"address_city"`
	AddressCountry     *string                 `form:"address_country"`
	AddressLine1       *string                 `form:"address_line1"`
	AddressLine2       *string                 `form:"address_line2"`
	AddressState       *string                 `form:"address_state"`
	AddressZip         *string                 `form:"address_zip"`
	DefaultForCurrency *bool                   `form:"default_for_currency"`
	ExpMonth           *string                 `form:"exp_month"`
	ExpYear            *string                 `form:"exp_year"`
	Name               *string                 `form:"name"`
	Owner              *BankAccountOwnerParams `form:"owner"`
}

// Verify a specified bank account for a given customer.
type BankAccountVerifyParams struct {
	Params   `form:"*"`
	Customer *string  `form:"-"` // Included in URL
	Amounts  []*int64 `form:"amounts"`
}

// BankAccount represents a Stripe bank account.
type BankAccount struct {
	APIResource
	Account                *Account                           `json:"account"`
	AccountHolderName      string                             `json:"account_holder_name"`
	AccountHolderType      string                             `json:"account_holder_type"`
	AccountType            string                             `json:"account_type"`
	AvailablePayoutMethods []BankAccountAvailablePayoutMethod `json:"available_payout_methods"`
	BankName               string                             `json:"bank_name"`
	Country                string                             `json:"country"`
	Currency               Currency                           `json:"currency"`
	Customer               *Customer                          `json:"customer"`
	DefaultForCurrency     bool                               `json:"default_for_currency"`
	Deleted                bool                               `json:"deleted"`
	Fingerprint            string                             `json:"fingerprint"`
	ID                     string                             `json:"id"`
	Last4                  string                             `json:"last4"`
	Metadata               map[string]string                  `json:"metadata"`
	Object                 string                             `json:"object"`
	RoutingNumber          string                             `json:"routing_number"`
	Status                 string                             `json:"status"`
}

// UnmarshalJSON handles deserialization of a BankAccount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BankAccount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type bankAccount BankAccount
	var v bankAccount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BankAccount(v)
	return nil
}
