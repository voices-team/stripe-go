//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// CardAvailablePayoutMethod is a set of available payout methods for the card.
type CardAvailablePayoutMethod string

// List of values that CardAvailablePayoutMethod can take.
const (
	CardAvailablePayoutMethodInstant  CardAvailablePayoutMethod = "instant"
	CardAvailablePayoutMethodStandard CardAvailablePayoutMethod = "standard"
)

type CardOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// CardParams is the set of parameters that can be used when creating or updating a card.
// For more details see https://stripe.com/docs/api#create_card and https://stripe.com/docs/api#update_card.
// //
// Note that while form annotations are used for tokenization and updates,
// cards have some unusual logic on creates that necessitates manual handling
// of all parameters. See AppendToAsCardSourceOrExternalAccount.
type CardParams struct {
	Params             `form:"*"`
	Account            *string          `form:"-"` // Included in URL
	Customer           *string          `form:"-"` // Included in URL
	AccountHolderName  *string          `form:"account_holder_name"`
	AccountHolderType  *string          `form:"account_holder_type"`
	AddressCity        *string          `form:"address_city"`
	AddressCountry     *string          `form:"address_country"`
	AddressLine1       *string          `form:"address_line1"`
	AddressLine2       *string          `form:"address_line2"`
	AddressState       *string          `form:"address_state"`
	AddressZip         *string          `form:"address_zip"`
	DefaultForCurrency *bool            `form:"default_for_currency"`
	ExpMonth           *string          `form:"exp_month"`
	ExpYear            *string          `form:"exp_year"`
	Name               *string          `form:"name"`
	Owner              *CardOwnerParams `form:"owner"`
}

// Card is the resource representing a Stripe credit/debit card.
// For more details see https://stripe.com/docs/api#cards.
type Card struct {
	APIResource
	Account                *Account                    `json:"account"`
	AddressCity            string                      `json:"address_city"`
	AddressCountry         string                      `json:"address_country"`
	AddressLine1           string                      `json:"address_line1"`
	AddressLine1Check      string                      `json:"address_line1_check"`
	AddressLine2           string                      `json:"address_line2"`
	AddressState           string                      `json:"address_state"`
	AddressZip             string                      `json:"address_zip"`
	AddressZipCheck        string                      `json:"address_zip_check"`
	AvailablePayoutMethods []CardAvailablePayoutMethod `json:"available_payout_methods"`
	Brand                  string                      `json:"brand"`
	Country                string                      `json:"country"`
	Currency               Currency                    `json:"currency"`
	Customer               *Customer                   `json:"customer"`
	CVCCheck               string                      `json:"cvc_check"`
	DefaultForCurrency     bool                        `json:"default_for_currency"`
	Deleted                bool                        `json:"deleted"`
	Description            string                      `json:"description"`
	DynamicLast4           string                      `json:"dynamic_last4"`
	ExpMonth               int64                       `json:"exp_month"`
	ExpYear                int64                       `json:"exp_year"`
	Fingerprint            string                      `json:"fingerprint"`
	Funding                string                      `json:"funding"`
	ID                     string                      `json:"id"`
	Iin                    string                      `json:"iin"`
	Issuer                 string                      `json:"issuer"`
	Last4                  string                      `json:"last4"`
	Metadata               map[string]string           `json:"metadata"`
	Name                   string                      `json:"name"`
	Object                 string                      `json:"object"`
	Recipient              *Recipient                  `json:"recipient"`
	TokenizationMethod     string                      `json:"tokenization_method"`
}

// UnmarshalJSON handles deserialization of a Card.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Card) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type card Card
	var v card
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Card(v)
	return nil
}
