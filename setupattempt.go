//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Preferred language of the Bancontact authorization page that the customer is redirected to.
// Can be one of `en`, `de`, `fr`, or `nl`
type SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage string

// List of values that SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage can take
const (
	SetupAttemptPaymentMethodDetailsBancontactPreferredLanguageDe SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage = "de"
	SetupAttemptPaymentMethodDetailsBancontactPreferredLanguageEn SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage = "en"
	SetupAttemptPaymentMethodDetailsBancontactPreferredLanguageFr SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage = "fr"
	SetupAttemptPaymentMethodDetailsBancontactPreferredLanguageNl SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage = "nl"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow indicates the type of 3D Secure authentication performed.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureResult indicates the outcome of 3D Secure authentication.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResult string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResult can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAuthenticated       SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultFailed              SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "failed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultNotSupported        SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultProcessingError     SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason represents dditional information about why 3D Secure succeeded or failed
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonRejected            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The version of 3D Secure that was used.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion can take
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion102 SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion = "1.0.2"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion210 SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion = "2.1.0"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion220 SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion = "2.2.0"
)

// The customer's bank. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
type SetupAttemptPaymentMethodDetailsIdealBank string

// List of values that SetupAttemptPaymentMethodDetailsIdealBank can take
const (
	SetupAttemptPaymentMethodDetailsIdealBankAbnAmro       SetupAttemptPaymentMethodDetailsIdealBank = "abn_amro"
	SetupAttemptPaymentMethodDetailsIdealBankAsnBank       SetupAttemptPaymentMethodDetailsIdealBank = "asn_bank"
	SetupAttemptPaymentMethodDetailsIdealBankBunq          SetupAttemptPaymentMethodDetailsIdealBank = "bunq"
	SetupAttemptPaymentMethodDetailsIdealBankHandelsbanken SetupAttemptPaymentMethodDetailsIdealBank = "handelsbanken"
	SetupAttemptPaymentMethodDetailsIdealBankIng           SetupAttemptPaymentMethodDetailsIdealBank = "ing"
	SetupAttemptPaymentMethodDetailsIdealBankKnab          SetupAttemptPaymentMethodDetailsIdealBank = "knab"
	SetupAttemptPaymentMethodDetailsIdealBankMoneyou       SetupAttemptPaymentMethodDetailsIdealBank = "moneyou"
	SetupAttemptPaymentMethodDetailsIdealBankRabobank      SetupAttemptPaymentMethodDetailsIdealBank = "rabobank"
	SetupAttemptPaymentMethodDetailsIdealBankRegiobank     SetupAttemptPaymentMethodDetailsIdealBank = "regiobank"
	SetupAttemptPaymentMethodDetailsIdealBankRevolut       SetupAttemptPaymentMethodDetailsIdealBank = "revolut"
	SetupAttemptPaymentMethodDetailsIdealBankSnsBank       SetupAttemptPaymentMethodDetailsIdealBank = "sns_bank"
	SetupAttemptPaymentMethodDetailsIdealBankTriodosBank   SetupAttemptPaymentMethodDetailsIdealBank = "triodos_bank"
	SetupAttemptPaymentMethodDetailsIdealBankVanLanschot   SetupAttemptPaymentMethodDetailsIdealBank = "van_lanschot"
)

// The Bank Identifier Code of the customer's bank.
type SetupAttemptPaymentMethodDetailsIdealBic string

// List of values that SetupAttemptPaymentMethodDetailsIdealBic can take
const (
	SetupAttemptPaymentMethodDetailsIdealBicABNANL2A SetupAttemptPaymentMethodDetailsIdealBic = "ABNANL2A"
	SetupAttemptPaymentMethodDetailsIdealBicASNBNL21 SetupAttemptPaymentMethodDetailsIdealBic = "ASNBNL21"
	SetupAttemptPaymentMethodDetailsIdealBicBUNQNL2A SetupAttemptPaymentMethodDetailsIdealBic = "BUNQNL2A"
	SetupAttemptPaymentMethodDetailsIdealBicFVLBNL22 SetupAttemptPaymentMethodDetailsIdealBic = "FVLBNL22"
	SetupAttemptPaymentMethodDetailsIdealBicHANDNL2A SetupAttemptPaymentMethodDetailsIdealBic = "HANDNL2A"
	SetupAttemptPaymentMethodDetailsIdealBicINGBNL2A SetupAttemptPaymentMethodDetailsIdealBic = "INGBNL2A"
	SetupAttemptPaymentMethodDetailsIdealBicKNABNL2H SetupAttemptPaymentMethodDetailsIdealBic = "KNABNL2H"
	SetupAttemptPaymentMethodDetailsIdealBicMOYONL21 SetupAttemptPaymentMethodDetailsIdealBic = "MOYONL21"
	SetupAttemptPaymentMethodDetailsIdealBicRABONL2U SetupAttemptPaymentMethodDetailsIdealBic = "RABONL2U"
	SetupAttemptPaymentMethodDetailsIdealBicRBRBNL21 SetupAttemptPaymentMethodDetailsIdealBic = "RBRBNL21"
	SetupAttemptPaymentMethodDetailsIdealBicREVOLT21 SetupAttemptPaymentMethodDetailsIdealBic = "REVOLT21"
	SetupAttemptPaymentMethodDetailsIdealBicSNSBNL2A SetupAttemptPaymentMethodDetailsIdealBic = "SNSBNL2A"
	SetupAttemptPaymentMethodDetailsIdealBicTRIONL2U SetupAttemptPaymentMethodDetailsIdealBic = "TRIONL2U"
)

// Preferred language of the Sofort authorization page that the customer is redirected to.
// Can be one of `en`, `de`, `fr`, or `nl`
type SetupAttemptPaymentMethodDetailsSofortPreferredLanguage string

// List of values that SetupAttemptPaymentMethodDetailsSofortPreferredLanguage can take
const (
	SetupAttemptPaymentMethodDetailsSofortPreferredLanguageDe SetupAttemptPaymentMethodDetailsSofortPreferredLanguage = "de"
	SetupAttemptPaymentMethodDetailsSofortPreferredLanguageEn SetupAttemptPaymentMethodDetailsSofortPreferredLanguage = "en"
	SetupAttemptPaymentMethodDetailsSofortPreferredLanguageFr SetupAttemptPaymentMethodDetailsSofortPreferredLanguage = "fr"
	SetupAttemptPaymentMethodDetailsSofortPreferredLanguageNl SetupAttemptPaymentMethodDetailsSofortPreferredLanguage = "nl"
)

// SetupAttemptListParams is the set of parameters that can be used when listing setup attempts.
type SetupAttemptListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	SetupIntent  *string           `form:"setup_intent"`
}

// SetupAttemptPaymentMethodDetailsCardThreeDSecure represents details about 3DS associated with the setup attempt's payment method.
type SetupAttemptPaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             SetupAttemptPaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            SetupAttemptPaymentMethodDetailsCardThreeDSecureVersion            `json:"version"`
}
type SetupAttemptPaymentMethodDetailsACSSDebit struct{}
type SetupAttemptPaymentMethodDetailsAUBECSDebit struct{}
type SetupAttemptPaymentMethodDetailsBACSDebit struct{}

// SetupAttemptPaymentMethodDetailsBancontact represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsBancontact struct {
	BankCode                  string                                                      `json:"bank_code"`
	BankName                  string                                                      `json:"bank_name"`
	Bic                       string                                                      `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod                                              `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                                                    `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                                                      `json:"iban_last4"`
	PreferredLanguage         SetupAttemptPaymentMethodDetailsBancontactPreferredLanguage `json:"preferred_language"`
	VerifiedName              string                                                      `json:"verified_name"`
}

// SetupAttemptPaymentMethodDetailsCard represents details about the Card PaymentMethod.
type SetupAttemptPaymentMethodDetailsCard struct {
	ThreeDSecure *SetupAttemptPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
}
type SetupAttemptPaymentMethodDetailsCardPresent struct {
	GeneratedCard *PaymentMethod `json:"generated_card"`
}

// SetupAttemptPaymentMethodDetailsIdeal represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsIdeal struct {
	Bank                      SetupAttemptPaymentMethodDetailsIdealBank `json:"bank"`
	Bic                       SetupAttemptPaymentMethodDetailsIdealBic  `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod                            `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                                  `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                                    `json:"iban_last4"`
	VerifiedName              string                                    `json:"verified_name"`
}
type SetupAttemptPaymentMethodDetailsSepaDebit struct{}

// SetupAttemptPaymentMethodDetailsSofort represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsSofort struct {
	BankCode                  string                                                  `json:"bank_code"`
	BankName                  string                                                  `json:"bank_name"`
	Bic                       string                                                  `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod                                          `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                                                `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                                                  `json:"iban_last4"`
	PreferredLanguage         SetupAttemptPaymentMethodDetailsSofortPreferredLanguage `json:"preferred_language"`
	VerifiedName              string                                                  `json:"verified_name"`
}

// SetupAttemptPaymentMethodDetails represents the details about the PaymentMethod associated with the setup attempt.
type SetupAttemptPaymentMethodDetails struct {
	ACSSDebit   *SetupAttemptPaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *SetupAttemptPaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *SetupAttemptPaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	Bancontact  *SetupAttemptPaymentMethodDetailsBancontact  `json:"bancontact"`
	Card        *SetupAttemptPaymentMethodDetailsCard        `json:"card"`
	CardPresent *SetupAttemptPaymentMethodDetailsCardPresent `json:"card_present"`
	Ideal       *SetupAttemptPaymentMethodDetailsIdeal       `json:"ideal"`
	SepaDebit   *SetupAttemptPaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	Sofort      *SetupAttemptPaymentMethodDetailsSofort      `json:"sofort"`
	Type        string                                       `json:"type"`
}

// SetupAttempt is the resource representing a Stripe setup attempt.
type SetupAttempt struct {
	Application          *Application                      `json:"application"`
	Created              int64                             `json:"created"`
	Customer             *Customer                         `json:"customer"`
	ID                   string                            `json:"id"`
	Livemode             bool                              `json:"livemode"`
	Object               string                            `json:"object"`
	OnBehalfOf           *Account                          `json:"on_behalf_of"`
	PaymentMethod        *PaymentMethod                    `json:"payment_method"`
	PaymentMethodDetails *SetupAttemptPaymentMethodDetails `json:"payment_method_details"`
	SetupError           *Error                            `json:"setup_error"`
	SetupIntent          *SetupIntent                      `json:"setup_intent"`
	Status               string                            `json:"status"`
	Usage                string                            `json:"usage"`
}

// SetupAttemptList is a list of setup attempts as retrieved from a list endpoint.
type SetupAttemptList struct {
	APIResource
	ListMeta
	Data []*SetupAttempt `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupAttempt.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SetupAttempt) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type setupAttempt SetupAttempt
	var v setupAttempt
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SetupAttempt(v)
	return nil
}
