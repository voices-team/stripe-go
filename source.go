//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type SourceType string

const (
	SourceTypeAchCreditTransfer  SourceType = "ach_credit_transfer"
	SourceTypeAchDebit           SourceType = "ach_debit"
	SourceTypeACSSDebit          SourceType = "acss_debit"
	SourceTypeAlipay             SourceType = "alipay"
	SourceTypeAuBECSDebit        SourceType = "au_becs_debit"
	SourceTypeBancontact         SourceType = "bancontact"
	SourceTypeCard               SourceType = "card"
	SourceTypeCardPresent        SourceType = "card_present"
	SourceTypeEPS                SourceType = "eps"
	SourceTypeGiropay            SourceType = "giropay"
	SourceTypeIdeal              SourceType = "ideal"
	SourceTypeKlarna             SourceType = "klarna"
	SourceTypeMultibanco         SourceType = "multibanco"
	SourceTypeP24                SourceType = "p24"
	SourceTypeSepaCreditTransfer SourceType = "sepa_credit_transfer"
	SourceTypeSepaDebit          SourceType = "sepa_debit"
	SourceTypeSofort             SourceType = "sofort"
	SourceTypeThreeDSecure       SourceType = "three_d_secure"
	SourceTypeWechat             SourceType = "wechat"
)

// SourceOwnerParams is the set of parameters allowed for the owner hash on
// source creation or update.
type SourceOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// SourceMandateAcceptanceParams describes the set of parameters allowed for the `acceptance`
// hash on source creation or update.
type SourceMandateAcceptanceParams struct {
	Date      *int64                                `form:"date"`
	IP        *string                               `form:"ip"`
	Offline   *SourceMandateAcceptanceOfflineParams `form:"offline"`
	Online    *SourceMandateAcceptanceOnlineParams  `form:"online"`
	Status    *string                               `form:"status"`
	Type      *string                               `form:"type"`
	UserAgent *string                               `form:"user_agent"`
}

// SourceMandateAcceptanceOnlineParams describes the set of parameters for online accepted mandate
type SourceMandateAcceptanceOnlineParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}
type SourceDetachParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
}
type SourceParams struct {
	Params              `form:"*"`
	Amount              *int64                   `form:"amount"`
	ClientSecret        *string                  `form:"client_secret"`
	Currency            *string                  `form:"currency"`
	Customer            *string                  `form:"customer"`
	Flow                *string                  `form:"flow"`
	Mandate             *SourceMandateParams     `form:"mandate"`
	OriginalSource      *string                  `form:"original_source"`
	Owner               *SourceOwnerParams       `form:"owner"`
	Receiver            *SourceReceiverParams    `form:"receiver"`
	Redirect            *SourceRedirectParams    `form:"redirect"`
	SourceOrder         *SourceSourceOrderParams `form:"source_order"`
	StatementDescriptor *string                  `form:"statement_descriptor"`
	Token               *string                  `form:"token"`
	Type                *string                  `form:"type"`
	Usage               *string                  `form:"usage"`
}

// SourceMandateAcceptanceOfflineParams describes the set of parameters for offline accepted mandate
type SourceMandateAcceptanceOfflineParams struct {
	ContactEmail *string `form:"contact_email"`
}

// SourceMandateParams describes the set of parameters allowed for the `mandate` hash on
// source creation or update.
type SourceMandateParams struct {
	Acceptance         *SourceMandateAcceptanceParams `form:"acceptance"`
	Amount             *int64                         `form:"amount"`
	Currency           *string                        `form:"currency"`
	Interval           *string                        `form:"interval"`
	NotificationMethod *string                        `form:"notification_method"`
}
type SourceSourceOrderItemParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}
type SourceSourceOrderParams struct {
	Items    []*SourceSourceOrderItemParams `form:"items"`
	Shipping *ShippingDetailsParams         `form:"shipping"`
}

// SourceReceiverParams is the set of parameters allowed for the `receiver` hash on
// source creation or update.
type SourceReceiverParams struct {
	RefundAttributesMethod *string `form:"refund_attributes_method"`
}
type SourceRedirectParams struct {
	ReturnURL *string `form:"return_url"`
}
type SourceVerifyParams struct {
	Params `form:"*"`
	Values []*string `form:"values"`
}
type SourceSourceTransactionsParams struct {
	Params        `form:"*"`
	EndingBefore  *string `form:"ending_before"`
	Limit         *int64  `form:"limit"`
	StartingAfter *string `form:"starting_after"`
}
type SourceAchCreditTransfer struct {
	AccountNumber           string `json:"account_number"`
	BankName                string `json:"bank_name"`
	Fingerprint             string `json:"fingerprint"`
	RefundAccountHolderName string `json:"refund_account_holder_name"`
	RefundAccountHolderType string `json:"refund_account_holder_type"`
	RefundRoutingNumber     string `json:"refund_routing_number"`
	RoutingNumber           string `json:"routing_number"`
	SwiftCode               string `json:"swift_code"`
}
type SourceAchDebit struct {
	BankName      string `json:"bank_name"`
	Country       string `json:"country"`
	Fingerprint   string `json:"fingerprint"`
	Last4         string `json:"last4"`
	RoutingNumber string `json:"routing_number"`
	Type          string `json:"type"`
}
type SourceACSSDebit struct {
	BankAddressCity       string `json:"bank_address_city"`
	BankAddressLine1      string `json:"bank_address_line_1"`
	BankAddressLine2      string `json:"bank_address_line_2"`
	BankAddressPostalCode string `json:"bank_address_postal_code"`
	BankName              string `json:"bank_name"`
	Category              string `json:"category"`
	Country               string `json:"country"`
	Fingerprint           string `json:"fingerprint"`
	Last4                 string `json:"last4"`
	RoutingNumber         string `json:"routing_number"`
}
type SourceAlipay struct {
	DataString          string `json:"data_string"`
	NativeURL           string `json:"native_url"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceAuBECSDebit struct {
	BsbNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
}
type SourceBancontact struct {
	BankCode            string `json:"bank_code"`
	BankName            string `json:"bank_name"`
	Bic                 string `json:"bic"`
	IbanLast4           string `json:"iban_last4"`
	PreferredLanguage   string `json:"preferred_language"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceCard struct {
	AddressLine1Check  string `json:"address_line1_check"`
	AddressZipCheck    string `json:"address_zip_check"`
	Brand              string `json:"brand"`
	Country            string `json:"country"`
	CVCCheck           string `json:"cvc_check"`
	Description        string `json:"description"`
	DynamicLast4       string `json:"dynamic_last4"`
	ExpMonth           int64  `json:"exp_month"`
	ExpYear            int64  `json:"exp_year"`
	Fingerprint        string `json:"fingerprint"`
	Funding            string `json:"funding"`
	Iin                string `json:"iin"`
	Issuer             string `json:"issuer"`
	Last4              string `json:"last4"`
	Name               string `json:"name"`
	ThreeDSecure       string `json:"three_d_secure"`
	TokenizationMethod string `json:"tokenization_method"`
}
type SourceCardPresent struct {
	ApplicationCryptogram          string `json:"application_cryptogram"`
	ApplicationPreferredName       string `json:"application_preferred_name"`
	AuthorizationCode              string `json:"authorization_code"`
	AuthorizationResponseCode      string `json:"authorization_response_code"`
	Brand                          string `json:"brand"`
	Country                        string `json:"country"`
	CvmType                        string `json:"cvm_type"`
	DataType                       string `json:"data_type"`
	DedicatedFileName              string `json:"dedicated_file_name"`
	Description                    string `json:"description"`
	EmvAuthData                    string `json:"emv_auth_data"`
	EvidenceCustomerSignature      string `json:"evidence_customer_signature"`
	EvidenceTransactionCertificate string `json:"evidence_transaction_certificate"`
	ExpMonth                       int64  `json:"exp_month"`
	ExpYear                        int64  `json:"exp_year"`
	Fingerprint                    string `json:"fingerprint"`
	Funding                        string `json:"funding"`
	Iin                            string `json:"iin"`
	Issuer                         string `json:"issuer"`
	Last4                          string `json:"last4"`
	PosDeviceID                    string `json:"pos_device_id"`
	PosEntryMode                   string `json:"pos_entry_mode"`
	Reader                         string `json:"reader"`
	ReadMethod                     string `json:"read_method"`
	TerminalVerificationResults    string `json:"terminal_verification_results"`
	TransactionStatusInformation   string `json:"transaction_status_information"`
}
type SourceCodeVerification struct {
	AttemptsRemaining int64  `json:"attempts_remaining"`
	Status            string `json:"status"`
}
type SourceEPS struct {
	Reference           string `json:"reference"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceGiropay struct {
	BankCode            string `json:"bank_code"`
	BankName            string `json:"bank_name"`
	Bic                 string `json:"bic"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceIdeal struct {
	Bank                string `json:"bank"`
	Bic                 string `json:"bic"`
	IbanLast4           string `json:"iban_last4"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceKlarna struct {
	BackgroundImageURL              string `json:"background_image_url"`
	ClientToken                     string `json:"client_token"`
	FirstName                       string `json:"first_name"`
	LastName                        string `json:"last_name"`
	Locale                          string `json:"locale"`
	LogoURL                         string `json:"logo_url"`
	PageTitle                       string `json:"page_title"`
	PayLaterAssetUrlsDescriptive    string `json:"pay_later_asset_urls_descriptive"`
	PayLaterAssetUrlsStandard       string `json:"pay_later_asset_urls_standard"`
	PayLaterName                    string `json:"pay_later_name"`
	PayLaterRedirectURL             string `json:"pay_later_redirect_url"`
	PaymentMethodCategories         string `json:"payment_method_categories"`
	PayNowAssetUrlsDescriptive      string `json:"pay_now_asset_urls_descriptive"`
	PayNowAssetUrlsStandard         string `json:"pay_now_asset_urls_standard"`
	PayNowName                      string `json:"pay_now_name"`
	PayNowRedirectURL               string `json:"pay_now_redirect_url"`
	PayOverTimeAssetUrlsDescriptive string `json:"pay_over_time_asset_urls_descriptive"`
	PayOverTimeAssetUrlsStandard    string `json:"pay_over_time_asset_urls_standard"`
	PayOverTimeName                 string `json:"pay_over_time_name"`
	PayOverTimeRedirectURL          string `json:"pay_over_time_redirect_url"`
	PurchaseCountry                 string `json:"purchase_country"`
	PurchaseType                    string `json:"purchase_type"`
	RedirectURL                     string `json:"redirect_url"`
	ShippingDelay                   int64  `json:"shipping_delay"`
	ShippingFirstName               string `json:"shipping_first_name"`
	ShippingLastName                string `json:"shipping_last_name"`
}
type SourceMultibanco struct {
	Entity                               string `json:"entity"`
	Reference                            string `json:"reference"`
	RefundAccountHolderAddressCity       string `json:"refund_account_holder_address_city"`
	RefundAccountHolderAddressCountry    string `json:"refund_account_holder_address_country"`
	RefundAccountHolderAddressLine1      string `json:"refund_account_holder_address_line1"`
	RefundAccountHolderAddressLine2      string `json:"refund_account_holder_address_line2"`
	RefundAccountHolderAddressPostalCode string `json:"refund_account_holder_address_postal_code"`
	RefundAccountHolderAddressState      string `json:"refund_account_holder_address_state"`
	RefundAccountHolderName              string `json:"refund_account_holder_name"`
	RefundIban                           string `json:"refund_iban"`
}

// SourceOwner describes the owner hash on a source.
type SourceOwner struct {
	Address         *Address `json:"address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone"`
	VerifiedAddress *Address `json:"verified_address"`
	VerifiedEmail   string   `json:"verified_email"`
	VerifiedName    string   `json:"verified_name"`
	VerifiedPhone   string   `json:"verified_phone"`
}
type SourceP24 struct {
	Reference string `json:"reference"`
}
type SourceReceiver struct {
	Address                string `json:"address"`
	AmountCharged          int64  `json:"amount_charged"`
	AmountReceived         int64  `json:"amount_received"`
	AmountReturned         int64  `json:"amount_returned"`
	RefundAttributesMethod string `json:"refund_attributes_method"`
	RefundAttributesStatus string `json:"refund_attributes_status"`
}
type SourceRedirect struct {
	FailureReason string `json:"failure_reason"`
	ReturnURL     string `json:"return_url"`
	Status        string `json:"status"`
	URL           string `json:"url"`
}
type SourceSepaCreditTransfer struct {
	BankName                             string `json:"bank_name"`
	Bic                                  string `json:"bic"`
	Iban                                 string `json:"iban"`
	RefundAccountHolderAddressCity       string `json:"refund_account_holder_address_city"`
	RefundAccountHolderAddressCountry    string `json:"refund_account_holder_address_country"`
	RefundAccountHolderAddressLine1      string `json:"refund_account_holder_address_line1"`
	RefundAccountHolderAddressLine2      string `json:"refund_account_holder_address_line2"`
	RefundAccountHolderAddressPostalCode string `json:"refund_account_holder_address_postal_code"`
	RefundAccountHolderAddressState      string `json:"refund_account_holder_address_state"`
	RefundAccountHolderName              string `json:"refund_account_holder_name"`
	RefundIban                           string `json:"refund_iban"`
}
type SourceSepaDebit struct {
	BankCode         string `json:"bank_code"`
	BranchCode       string `json:"branch_code"`
	Country          string `json:"country"`
	Fingerprint      string `json:"fingerprint"`
	Last4            string `json:"last4"`
	MandateReference string `json:"mandate_reference"`
	MandateURL       string `json:"mandate_url"`
}
type SourceSofort struct {
	BankCode            string `json:"bank_code"`
	BankName            string `json:"bank_name"`
	Bic                 string `json:"bic"`
	Country             string `json:"country"`
	IbanLast4           string `json:"iban_last4"`
	PreferredLanguage   string `json:"preferred_language"`
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceSourceOrderItem struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Description string   `json:"description"`
	Parent      string   `json:"parent"`
	Quantity    int64    `json:"quantity"`
	Type        string   `json:"type"`
}

// SourceSourceOrder describes a source order for a source.
type SourceSourceOrder struct {
	Amount   int64                    `json:"amount"`
	Currency Currency                 `json:"currency"`
	Email    string                   `json:"email"`
	Items    []*SourceSourceOrderItem `json:"items"`
	Shipping ShippingDetails          `json:"shipping"`
}
type SourceThreeDSecure struct {
	AddressLine1Check  string `json:"address_line1_check"`
	AddressZipCheck    string `json:"address_zip_check"`
	Authenticated      bool   `json:"authenticated"`
	Brand              string `json:"brand"`
	Card               string `json:"card"`
	Country            string `json:"country"`
	Customer           string `json:"customer"`
	CVCCheck           string `json:"cvc_check"`
	Description        string `json:"description"`
	DynamicLast4       string `json:"dynamic_last4"`
	ExpMonth           int64  `json:"exp_month"`
	ExpYear            int64  `json:"exp_year"`
	Fingerprint        string `json:"fingerprint"`
	Funding            string `json:"funding"`
	Iin                string `json:"iin"`
	Issuer             string `json:"issuer"`
	Last4              string `json:"last4"`
	Name               string `json:"name"`
	ThreeDSecure       string `json:"three_d_secure"`
	TokenizationMethod string `json:"tokenization_method"`
}
type SourceWechat struct {
	PrepayID            string `json:"prepay_id"`
	QrCodeURL           string `json:"qr_code_url"`
	StatementDescriptor string `json:"statement_descriptor"`
}

// Source is the resource representing a Source.
// For more details see https://stripe.com/docs/api#sources.
type Source struct {
	APIResource
	AchCreditTransfer   *SourceAchCreditTransfer  `json:"ach_credit_transfer"`
	AchDebit            *SourceAchDebit           `json:"ach_debit"`
	ACSSDebit           *SourceACSSDebit          `json:"acss_debit"`
	Alipay              *SourceAlipay             `json:"alipay"`
	Amount              int64                     `json:"amount"`
	AuBECSDebit         *SourceAuBECSDebit        `json:"au_becs_debit"`
	Bancontact          *SourceBancontact         `json:"bancontact"`
	Card                *SourceCard               `json:"card"`
	CardPresent         *SourceCardPresent        `json:"card_present"`
	ClientSecret        string                    `json:"client_secret"`
	CodeVerification    *SourceCodeVerification   `json:"code_verification"`
	Created             int64                     `json:"created"`
	Currency            Currency                  `json:"currency"`
	Customer            string                    `json:"customer"`
	EPS                 *SourceEPS                `json:"eps"`
	Flow                string                    `json:"flow"`
	Giropay             *SourceGiropay            `json:"giropay"`
	ID                  string                    `json:"id"`
	Ideal               *SourceIdeal              `json:"ideal"`
	Klarna              *SourceKlarna             `json:"klarna"`
	Livemode            bool                      `json:"livemode"`
	Metadata            map[string]string         `json:"metadata"`
	Multibanco          *SourceMultibanco         `json:"multibanco"`
	Object              string                    `json:"object"`
	Owner               *SourceOwner              `json:"owner"`
	P24                 *SourceP24                `json:"p24"`
	Receiver            *SourceReceiver           `json:"receiver"`
	Redirect            *SourceRedirect           `json:"redirect"`
	SepaCreditTransfer  *SourceSepaCreditTransfer `json:"sepa_credit_transfer"`
	SepaDebit           *SourceSepaDebit          `json:"sepa_debit"`
	Sofort              *SourceSofort             `json:"sofort"`
	SourceOrder         *SourceSourceOrder        `json:"source_order"`
	StatementDescriptor string                    `json:"statement_descriptor"`
	Status              string                    `json:"status"`
	ThreeDSecure        *SourceThreeDSecure       `json:"three_d_secure"`
	Type                SourceType                `json:"type"`
	Usage               string                    `json:"usage"`
	Wechat              *SourceWechat             `json:"wechat"`
}

// UnmarshalJSON handles deserialization of a Source.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Source) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type source Source
	var v source
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Source(v)
	return nil
}
