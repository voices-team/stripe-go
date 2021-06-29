//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// PaymentIntentCancellationReason is the list of allowed values for the cancelation reason.
type PaymentIntentCancellationReason string

// List of values that PaymentIntentCancellationReason can take.
const (
	PaymentIntentCancellationReasonAbandoned           PaymentIntentCancellationReason = "abandoned"
	PaymentIntentCancellationReasonAutomatic           PaymentIntentCancellationReason = "automatic"
	PaymentIntentCancellationReasonDuplicate           PaymentIntentCancellationReason = "duplicate"
	PaymentIntentCancellationReasonFailedInvoice       PaymentIntentCancellationReason = "failed_invoice"
	PaymentIntentCancellationReasonFraudulent          PaymentIntentCancellationReason = "fraudulent"
	PaymentIntentCancellationReasonRequestedByCustomer PaymentIntentCancellationReason = "requested_by_customer"
	PaymentIntentCancellationReasonVoidInvoice         PaymentIntentCancellationReason = "void_invoice"
)

// PaymentIntentCaptureMethod is the list of allowed values for the capture method.
type PaymentIntentCaptureMethod string

// List of values that PaymentIntentCaptureMethod can take.
const (
	PaymentIntentCaptureMethodAutomatic PaymentIntentCaptureMethod = "automatic"
	PaymentIntentCaptureMethodManual    PaymentIntentCaptureMethod = "manual"
)

// PaymentIntentConfirmationMethod is the list of allowed values for the confirmation method.
type PaymentIntentConfirmationMethod string

// List of values that PaymentIntentConfirmationMethod can take.
const (
	PaymentIntentConfirmationMethodAutomatic PaymentIntentConfirmationMethod = "automatic"
	PaymentIntentConfirmationMethodManual    PaymentIntentConfirmationMethod = "manual"
)

// PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule is the list of allowed values for payment_schedule.
type PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType is the list of allowed values for transaction_type.
type PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod is the list of allowed values for verification_method.
type PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethodInstant       PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

type PaymentIntentPaymentMethodOptionsCardNetwork string

const (
	PaymentIntentPaymentMethodOptionsCardNetworkAmex            PaymentIntentPaymentMethodOptionsCardNetwork = "amex"
	PaymentIntentPaymentMethodOptionsCardNetworkCartesBancaires PaymentIntentPaymentMethodOptionsCardNetwork = "cartes_bancaires"
	PaymentIntentPaymentMethodOptionsCardNetworkDiners          PaymentIntentPaymentMethodOptionsCardNetwork = "diners"
	PaymentIntentPaymentMethodOptionsCardNetworkDiscover        PaymentIntentPaymentMethodOptionsCardNetwork = "discover"
	PaymentIntentPaymentMethodOptionsCardNetworkInterac         PaymentIntentPaymentMethodOptionsCardNetwork = "interac"
	PaymentIntentPaymentMethodOptionsCardNetworkJcb             PaymentIntentPaymentMethodOptionsCardNetwork = "jcb"
	PaymentIntentPaymentMethodOptionsCardNetworkMastercard      PaymentIntentPaymentMethodOptionsCardNetwork = "mastercard"
	PaymentIntentPaymentMethodOptionsCardNetworkUnionpay        PaymentIntentPaymentMethodOptionsCardNetwork = "unionpay"
	PaymentIntentPaymentMethodOptionsCardNetworkUnknown         PaymentIntentPaymentMethodOptionsCardNetwork = "unknown"
	PaymentIntentPaymentMethodOptionsCardNetworkVisa            PaymentIntentPaymentMethodOptionsCardNetwork = "visa"
)

// PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure is the list of allowed values
// //controlling when to request 3D Secure on a PaymentIntent.
type PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure can take.
const (
	PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureAny           PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic     PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure = "challenge_only"
)

type PaymentIntentPaymentMethodOptionsSofortPreferredLanguage string

const (
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageDe PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "de"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageEn PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "en"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageEs PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "es"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageFr PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "fr"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageIt PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "it"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageNl PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "nl"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguagePl PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "pl"
)

// PaymentIntentSetupFutureUsage is the list of allowed values for SetupFutureUsage.
type PaymentIntentSetupFutureUsage string

// List of values that PaymentIntentSetupFutureUsage can take.
const (
	PaymentIntentSetupFutureUsageOffSession PaymentIntentSetupFutureUsage = "off_session"
	PaymentIntentSetupFutureUsageOnSession  PaymentIntentSetupFutureUsage = "on_session"
)

type PaymentIntentSourceType string

const (
	PaymentIntentSourceTypeAlipayAccount   PaymentIntentSourceType = "alipay_account"
	PaymentIntentSourceTypeBankAccount     PaymentIntentSourceType = "bank_account"
	PaymentIntentSourceTypeBitcoinReceiver PaymentIntentSourceType = "bitcoin_receiver"
	PaymentIntentSourceTypeCard            PaymentIntentSourceType = "card"
	PaymentIntentSourceTypeAccount         PaymentIntentSourceType = "account"
	PaymentIntentSourceTypeSource          PaymentIntentSourceType = "source"
)

// PaymentIntentStatus is the list of allowed values for the payment intent's status.
type PaymentIntentStatus string

// List of values that PaymentIntentStatus can take.
const (
	PaymentIntentStatusCanceled              PaymentIntentStatus = "canceled"
	PaymentIntentStatusProcessing            PaymentIntentStatus = "processing"
	PaymentIntentStatusRequiresAction        PaymentIntentStatus = "requires_action"
	PaymentIntentStatusRequiresCapture       PaymentIntentStatus = "requires_capture"
	PaymentIntentStatusRequiresConfirmation  PaymentIntentStatus = "requires_confirmation"
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "requires_payment_method"
	PaymentIntentStatusSucceeded             PaymentIntentStatus = "succeeded"
)

// PaymentIntentCancelParams is the set of parameters that can be used when canceling a payment intent.
type PaymentIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}

// PaymentIntentCaptureParams is the set of parameters that can be used when capturing a payment intent.
type PaymentIntentCaptureParams struct {
	Params                    `form:"*"`
	AmountToCapture           *int64                           `form:"amount_to_capture"`
	ApplicationFeeAmount      *int64                           `form:"application_fee_amount"`
	StatementDescriptor       *string                          `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                          `form:"statement_descriptor_suffix"`
	TransferData              *PaymentIntentTransferDataParams `form:"transfer_data"`
}
type PaymentIntentConfirmMandateDataCustomerAcceptanceOfflineParams struct{}
type PaymentIntentConfirmMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}
type PaymentIntentConfirmMandateDataCustomerAcceptanceParams struct {
	AcceptedAt *int64                                                          `form:"accepted_at"`
	Offline    *PaymentIntentConfirmMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *PaymentIntentConfirmMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       *string                                                         `form:"type"`
}
type PaymentIntentConfirmMandateDataParams struct {
	CustomerAcceptance *PaymentIntentConfirmMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}
type PaymentIntentConfirmPaymentMethodDataACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}
type PaymentIntentConfirmPaymentMethodDataAfterpayClearpayParams struct{}
type PaymentIntentConfirmPaymentMethodDataAlipayParams struct{}
type PaymentIntentConfirmPaymentMethodDataAuBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BsbNumber     *string `form:"bsb_number"`
}
type PaymentIntentConfirmPaymentMethodDataBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}
type PaymentIntentConfirmPaymentMethodDataBancontactParams struct{}
type PaymentIntentConfirmPaymentMethodDataBillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}
type PaymentIntentConfirmPaymentMethodDataEPSParams struct {
	Bank *string `form:"bank"`
}
type PaymentIntentConfirmPaymentMethodDataFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}
type PaymentIntentConfirmPaymentMethodDataGiropayParams struct{}
type PaymentIntentConfirmPaymentMethodDataGrabpayParams struct{}
type PaymentIntentConfirmPaymentMethodDataIdealParams struct {
	Bank *string `form:"bank"`
}
type PaymentIntentConfirmPaymentMethodDataInteracPresentParams struct{}
type PaymentIntentConfirmPaymentMethodDataOXXOParams struct{}
type PaymentIntentConfirmPaymentMethodDataP24Params struct {
	Bank *string `form:"bank"`
}
type PaymentIntentConfirmPaymentMethodDataSepaDebitParams struct {
	Iban *string `form:"iban"`
}
type PaymentIntentConfirmPaymentMethodDataSofortParams struct {
	Country *string `form:"country"`
}
type PaymentIntentConfirmPaymentMethodDataParams struct {
	ACSSDebit        *PaymentIntentConfirmPaymentMethodDataACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentConfirmPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentConfirmPaymentMethodDataAlipayParams           `form:"alipay"`
	AuBECSDebit      *PaymentIntentConfirmPaymentMethodDataAuBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentIntentConfirmPaymentMethodDataBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentIntentConfirmPaymentMethodDataBancontactParams       `form:"bancontact"`
	BillingDetails   *PaymentIntentConfirmPaymentMethodDataBillingDetailsParams   `form:"billing_details"`
	EPS              *PaymentIntentConfirmPaymentMethodDataEPSParams              `form:"eps"`
	FPX              *PaymentIntentConfirmPaymentMethodDataFPXParams              `form:"fpx"`
	Giropay          *PaymentIntentConfirmPaymentMethodDataGiropayParams          `form:"giropay"`
	Grabpay          *PaymentIntentConfirmPaymentMethodDataGrabpayParams          `form:"grabpay"`
	Ideal            *PaymentIntentConfirmPaymentMethodDataIdealParams            `form:"ideal"`
	InteracPresent   *PaymentIntentConfirmPaymentMethodDataInteracPresentParams   `form:"interac_present"`
	Metadata         map[string]string                                            `form:"metadata"`
	OXXO             *PaymentIntentConfirmPaymentMethodDataOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentConfirmPaymentMethodDataP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentConfirmPaymentMethodDataSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentConfirmPaymentMethodDataSofortParams           `form:"sofort"`
	Type             *string                                                      `form:"type"`
}
type PaymentIntentConfirmPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string `form:"custom_mandate_url"`
	IntervalDescription *string `form:"interval_description"`
	PaymentSchedule     *string `form:"payment_schedule"`
	TransactionType     *string `form:"transaction_type"`
}
type PaymentIntentConfirmPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *PaymentIntentConfirmPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                `form:"verification_method"`
}
type PaymentIntentConfirmPaymentMethodOptionsAfterpayClearpayParams struct {
	Reference *string `form:"reference"`
}
type PaymentIntentConfirmPaymentMethodOptionsAlipayParams struct{}
type PaymentIntentConfirmPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}
type PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsPlanParams struct {
	Count    *int64  `form:"count"`
	Interval *string `form:"interval"`
	Type     *string `form:"type"`
}
type PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsParams struct {
	Enabled *bool                                                               `form:"enabled"`
	Plan    *PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}
type PaymentIntentConfirmPaymentMethodOptionsCardParams struct {
	CVCToken            *string                                                         `form:"cvc_token"`
	Installments        *PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	Moto                *bool                                                           `form:"moto"`
	Network             *string                                                         `form:"network"`
	RequestThreeDSecure *string                                                         `form:"request_three_d_secure"`
}
type PaymentIntentConfirmPaymentMethodOptionsCardPresentParams struct{}
type PaymentIntentConfirmPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}
type PaymentIntentConfirmPaymentMethodOptionsP24Params struct {
	TosShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}
type PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}
type PaymentIntentConfirmPaymentMethodOptionsSepaDebitParams struct {
	MandateOptions *PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}
type PaymentIntentConfirmPaymentMethodOptionsSofortParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}
type PaymentIntentConfirmPaymentMethodOptionsParams struct {
	ACSSDebit        *PaymentIntentConfirmPaymentMethodOptionsACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentConfirmPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentConfirmPaymentMethodOptionsAlipayParams           `form:"alipay"`
	Bancontact       *PaymentIntentConfirmPaymentMethodOptionsBancontactParams       `form:"bancontact"`
	Card             *PaymentIntentConfirmPaymentMethodOptionsCardParams             `form:"card"`
	CardPresent      *PaymentIntentConfirmPaymentMethodOptionsCardPresentParams      `form:"card_present"`
	OXXO             *PaymentIntentConfirmPaymentMethodOptionsOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentConfirmPaymentMethodOptionsP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentConfirmPaymentMethodOptionsSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentConfirmPaymentMethodOptionsSofortParams           `form:"sofort"`
}

// PaymentIntentConfirmParams is the set of parameters that can be used when confirming a payment intent.
type PaymentIntentConfirmParams struct {
	Params                `form:"*"`
	ErrorOnRequiresAction *bool   `form:"error_on_requires_action"`
	Mandate               *string `form:"mandate"`
	// MandateData *[todo({"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentIntentConfirmMandateDataParams"}} | {"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentIntentConfirmMandateDataParams"}})] `form:"mandate_data"`
	// OffSession *[todo({"shape":"primitive","primitive":"boolean"} | {"shape":"primitive","primitive":"string"})] `form:"off_session"`
	PaymentMethod        *string                                         `form:"payment_method"`
	PaymentMethodData    *PaymentIntentConfirmPaymentMethodDataParams    `form:"payment_method_data"`
	PaymentMethodOptions *PaymentIntentConfirmPaymentMethodOptionsParams `form:"payment_method_options"`
	ReceiptEmail         *string                                         `form:"receipt_email"`
	ReturnURL            *string                                         `form:"return_url"`
	SetupFutureUsage     *string                                         `form:"setup_future_usage"`
	Shipping             *ShippingDetailsParams                          `form:"shipping"`
	UseStripeSDK         *bool                                           `form:"use_stripe_sdk"`
}

// PaymentIntentMandateDataCustomerAcceptanceOfflineParams is the set of parameters for the customer
// acceptance of an offline mandate.
type PaymentIntentMandateDataCustomerAcceptanceOfflineParams struct{}

// PaymentIntentMandateDataCustomerAcceptanceOnlineParams is the set of parameters for the customer
// acceptance of an online mandate.
type PaymentIntentMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// PaymentIntentMandateDataCustomerAcceptanceParams is the set of parameters for the customer
// acceptance of a mandate.
type PaymentIntentMandateDataCustomerAcceptanceParams struct {
	AcceptedAt int64                                                    `form:"accepted_at"`
	Offline    *PaymentIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *PaymentIntentMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       MandateCustomerAcceptanceType                            `form:"type"`
}

// PaymentIntentMandateDataParams is the set of parameters controlling the creation of the mandate
// associated with this PaymentIntent.
type PaymentIntentMandateDataParams struct {
	CustomerAcceptance *PaymentIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}
type PaymentIntentPaymentMethodDataACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}
type PaymentIntentPaymentMethodDataAfterpayClearpayParams struct{}
type PaymentIntentPaymentMethodDataAlipayParams struct{}
type PaymentIntentPaymentMethodDataAuBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BsbNumber     *string `form:"bsb_number"`
}
type PaymentIntentPaymentMethodDataBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}
type PaymentIntentPaymentMethodDataBancontactParams struct{}
type PaymentIntentPaymentMethodDataBillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}
type PaymentIntentPaymentMethodDataEPSParams struct {
	Bank *string `form:"bank"`
}
type PaymentIntentPaymentMethodDataFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}
type PaymentIntentPaymentMethodDataGiropayParams struct{}
type PaymentIntentPaymentMethodDataGrabpayParams struct{}
type PaymentIntentPaymentMethodDataIdealParams struct {
	Bank *string `form:"bank"`
}
type PaymentIntentPaymentMethodDataInteracPresentParams struct{}
type PaymentIntentPaymentMethodDataOXXOParams struct{}
type PaymentIntentPaymentMethodDataP24Params struct {
	Bank *string `form:"bank"`
}
type PaymentIntentPaymentMethodDataSepaDebitParams struct {
	Iban *string `form:"iban"`
}
type PaymentIntentPaymentMethodDataSofortParams struct {
	Country *string `form:"country"`
}

// PaymentIntentPaymentMethodDataParams represents the type-specific parameters associated with a
// payment method on payment intent.
type PaymentIntentPaymentMethodDataParams struct {
	ACSSDebit        *PaymentIntentPaymentMethodDataACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodDataAlipayParams           `form:"alipay"`
	AuBECSDebit      *PaymentIntentPaymentMethodDataAuBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentIntentPaymentMethodDataBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentIntentPaymentMethodDataBancontactParams       `form:"bancontact"`
	BillingDetails   *PaymentIntentPaymentMethodDataBillingDetailsParams   `form:"billing_details"`
	EPS              *PaymentIntentPaymentMethodDataEPSParams              `form:"eps"`
	FPX              *PaymentIntentPaymentMethodDataFPXParams              `form:"fpx"`
	Giropay          *PaymentIntentPaymentMethodDataGiropayParams          `form:"giropay"`
	Grabpay          *PaymentIntentPaymentMethodDataGrabpayParams          `form:"grabpay"`
	Ideal            *PaymentIntentPaymentMethodDataIdealParams            `form:"ideal"`
	InteracPresent   *PaymentIntentPaymentMethodDataInteracPresentParams   `form:"interac_present"`
	Metadata         map[string]string                                     `form:"metadata"`
	OXXO             *PaymentIntentPaymentMethodDataOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentPaymentMethodDataP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentPaymentMethodDataSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentPaymentMethodDataSofortParams           `form:"sofort"`
	Type             *string                                               `form:"type"`
}

// PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsParams represents the mandate options
// for ACSS on the payment intent.
type PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string `form:"custom_mandate_url"`
	IntervalDescription *string `form:"interval_description"`
	PaymentSchedule     *string `form:"payment_schedule"`
	TransactionType     *string `form:"transaction_type"`
}

// PaymentIntentPaymentMethodOptionsACSSDebitParams represents the ACSS debit-specific options
// applieed to a PaymentIntent
type PaymentIntentPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                         `form:"verification_method"`
}

// PaymentIntentPaymentMethodOptionsAfterpayClearpayParams represents the AfterpayClearpay-specific options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsAfterpayClearpayParams struct {
	Reference *string `form:"reference"`
}

// PaymentIntentPaymentMethodOptionsAlipayParams represents the Alipay-specific options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsAlipayParams struct{}

// PaymentIntentPaymentMethodOptionsBancontactParams represents the bancontact-specific options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams represents details about the
// installment plan chosen for this payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams struct {
	Count    *int64  `form:"count"`
	Interval *string `form:"interval"`
	Type     *string `form:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsParams controls whether to enable installment
// plans for this payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallmentsParams struct {
	Enabled *bool                                                        `form:"enabled"`
	Plan    *PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}

// PaymentIntentPaymentMethodOptionsCardParams represents the card-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsCardParams struct {
	CVCToken            *string                                                  `form:"cvc_token"`
	Installments        *PaymentIntentPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	Moto                *bool                                                    `form:"moto"`
	Network             *string                                                  `form:"network"`
	RequestThreeDSecure *string                                                  `form:"request_three_d_secure"`
}
type PaymentIntentPaymentMethodOptionsCardPresentParams struct{}

// PaymentIntentPaymentMethodOptionsOXXOParams represents the OXXO-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}
type PaymentIntentPaymentMethodOptionsP24Params struct {
	TosShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}
type PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}
type PaymentIntentPaymentMethodOptionsSepaDebitParams struct {
	MandateOptions *PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// PaymentIntentPaymentMethodOptionsSofortParams represents the sofort-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsSofortParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// PaymentIntentPaymentMethodOptionsParams represents the type-specific payment method options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsParams struct {
	ACSSDebit        *PaymentIntentPaymentMethodOptionsACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipayParams           `form:"alipay"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontactParams       `form:"bancontact"`
	Card             *PaymentIntentPaymentMethodOptionsCardParams             `form:"card"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresentParams      `form:"card_present"`
	OXXO             *PaymentIntentPaymentMethodOptionsOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentPaymentMethodOptionsP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofortParams           `form:"sofort"`
}

// PaymentIntentTransferDataParams is the set of parameters allowed for the transfer hash.
type PaymentIntentTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// PaymentIntentParams is the set of parameters that can be used when handling a payment intent.
type PaymentIntentParams struct {
	Params                `form:"*"`
	Amount                *int64                          `form:"amount"`
	ApplicationFeeAmount  *int64                          `form:"application_fee_amount"`
	CaptureMethod         *string                         `form:"capture_method"`
	ClientSecret          *string                         `form:"client_secret"`
	Confirm               *bool                           `form:"confirm"`
	ConfirmationMethod    *string                         `form:"confirmation_method"`
	Currency              *string                         `form:"currency"`
	Customer              *string                         `form:"customer"`
	Description           *string                         `form:"description"`
	ErrorOnRequiresAction *bool                           `form:"error_on_requires_action"`
	Mandate               *string                         `form:"mandate"`
	MandateData           *PaymentIntentMandateDataParams `form:"mandate_data"`
	// OffSession *[todo({"shape":"primitive","primitive":"boolean"} | {"shape":"primitive","primitive":"string"})] `form:"off_session"`
	OnBehalfOf                *string                                  `form:"on_behalf_of"`
	PaymentMethod             *string                                  `form:"payment_method"`
	PaymentMethodData         *PaymentIntentPaymentMethodDataParams    `form:"payment_method_data"`
	PaymentMethodOptions      *PaymentIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes        []*string                                `form:"payment_method_types"`
	ReceiptEmail              *string                                  `form:"receipt_email"`
	ReturnURL                 *string                                  `form:"return_url"`
	SetupFutureUsage          *string                                  `form:"setup_future_usage"`
	Shipping                  *ShippingDetailsParams                   `form:"shipping"`
	StatementDescriptor       *string                                  `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                                  `form:"statement_descriptor_suffix"`
	TransferData              *PaymentIntentTransferDataParams         `form:"transfer_data"`
	TransferGroup             *string                                  `form:"transfer_group"`
	UseStripeSDK              *bool                                    `form:"use_stripe_sdk"`
}

// PaymentIntentListParams is the set of parameters that can be used when listing payment intents.
// For more details see https://stripe.com/docs/api#list_payouts.
type PaymentIntentListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
}

// PaymentIntentNextActionAlipayHandleRedirect represents the resource for the next action of type
// "handle_alipay_redirect".
type PaymentIntentNextActionAlipayHandleRedirect struct {
	NativeData string `json:"native_data"`
	NativeURL  string `json:"native_url"`
	ReturnURL  string `json:"return_url"`
	URL        string `json:"url"`
}

// PaymentIntentNextActionOXXODisplayDetails represents the resource for the next action of type
// "oxxo_display_details".
type PaymentIntentNextActionOXXODisplayDetails struct {
	ExpiresAfter     int64  `json:"expires_after"`
	HostedVoucherURL string `json:"hosted_voucher_url"`
	Number           string `json:"number"`
}

// PaymentIntentNextActionRedirectToURL represents the resource for the next action of type
// "redirect_to_url".
type PaymentIntentNextActionRedirectToURL struct {
	ReturnURL string `json:"return_url"`
	URL       string `json:"url"`
}

// PaymentIntentNextActionUseStripeSDK represents the resource for the next action of typee "use_stripe_sdk".
type PaymentIntentNextActionUseStripeSDK struct{}

// PaymentIntentNextActionVerifyWithMicrodeposits represents the resource for the next action of type
// "verify_with_microdeposits".
type PaymentIntentNextActionVerifyWithMicrodeposits struct {
	ArrivalDate           int64  `json:"arrival_date"`
	HostedVerificationURL string `json:"hosted_verification_url"`
}

// PaymentIntentNextAction represents the type of action to take on a payment intent.
type PaymentIntentNextAction struct {
	AlipayHandleRedirect    *PaymentIntentNextActionAlipayHandleRedirect    `json:"alipay_handle_redirect"`
	OXXODisplayDetails      *PaymentIntentNextActionOXXODisplayDetails      `json:"oxxo_display_details"`
	RedirectToURL           *PaymentIntentNextActionRedirectToURL           `json:"redirect_to_url"`
	Type                    string                                          `json:"type"`
	UseStripeSDK            *PaymentIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *PaymentIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}
type PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlan struct {
	Count    int64   `json:"count"`
	Interval *string `json:"interval"`
	Type     string  `json:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlan describe a specific card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlan struct {
	Count    int64   `json:"count"`
	Interval *string `json:"interval"`
	Type     string  `json:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallments describe the installment options available for
// a card associated with that payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallments struct {
	AvailablePlans []*PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlan `json:"available_plans"`
	Enabled        bool                                                              `json:"enabled"`
	Plan           *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan            `json:"plan"`
}

// PaymentIntentPaymentMethodOptionsACSSDebitMandateOptions describe the mandate options for acss debit
// associated with that payment intent.
type PaymentIntentPaymentMethodOptionsACSSDebitMandateOptions struct {
	CustomMandateURL    string                                                                  `json:"custom_mandate_url"`
	IntervalDescription string                                                                  `json:"interval_description"`
	PaymentSchedule     PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	TransactionType     PaymentIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// PaymentIntentPaymentMethodOptionsACSSDebit describes the ACSS debit-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsACSSDebit struct {
	MandateOptions     *PaymentIntentPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod PaymentIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// PaymentIntentPaymentMethodOptionsAfterpayClearpay describes the AfterpayClearpay-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsAfterpayClearpay struct {
	Reference string `json:"reference"`
}

// PaymentIntentPaymentMethodOptionsAlipay is the set of Alipay-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsAlipay struct{}

// PaymentIntentPaymentMethodOptionsBancontact is the set of bancontact-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsBancontact struct {
	PreferredLanguage string `json:"preferred_language"`
}

// PaymentIntentPaymentMethodOptionsCard is the set of card-specific options associated with that
// payment intent.
type PaymentIntentPaymentMethodOptionsCard struct {
	Installments        *PaymentIntentPaymentMethodOptionsCardInstallments       `json:"installments"`
	Network             PaymentIntentPaymentMethodOptionsCardNetwork             `json:"network"`
	RequestThreeDSecure PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type PaymentIntentPaymentMethodOptionsCardPresent struct{}

// PaymentIntentPaymentMethodOptionsOXXO is the set of OXXO-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsOXXO struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}
type PaymentIntentPaymentMethodOptionsP24 struct{}
type PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions struct{}
type PaymentIntentPaymentMethodOptionsSepaDebit struct {
	MandateOptions *PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions `json:"mandate_options"`
}

// PaymentIntentPaymentMethodOptionsSofort is the set of sofort-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsSofort struct {
	PreferredLanguage PaymentIntentPaymentMethodOptionsSofortPreferredLanguage `json:"preferred_language"`
}

// PaymentIntentPaymentMethodOptions is the set of payment method-specific options associated with
// that payment intent.
type PaymentIntentPaymentMethodOptions struct {
	ACSSDebit        *PaymentIntentPaymentMethodOptionsACSSDebit        `json:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipay           `json:"alipay"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontact       `json:"bancontact"`
	Card             *PaymentIntentPaymentMethodOptionsCard             `json:"card"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresent      `json:"card_present"`
	OXXO             *PaymentIntentPaymentMethodOptionsOXXO             `json:"oxxo"`
	P24              *PaymentIntentPaymentMethodOptionsP24              `json:"p24"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebit        `json:"sepa_debit"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofort           `json:"sofort"`
}

// PaymentIntentTransferData represents the information for the transfer associated with a payment intent.
type PaymentIntentTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}

// PaymentIntent is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payment_intents.
type PaymentIntent struct {
	APIResource
	Amount                    int64                              `json:"amount"`
	AmountCapturable          int64                              `json:"amount_capturable"`
	AmountReceived            int64                              `json:"amount_received"`
	Application               *Application                       `json:"application"`
	ApplicationFeeAmount      int64                              `json:"application_fee_amount"`
	CanceledAt                int64                              `json:"canceled_at"`
	CancellationReason        PaymentIntentCancellationReason    `json:"cancellation_reason"`
	CaptureMethod             PaymentIntentCaptureMethod         `json:"capture_method"`
	Charges                   *ChargeList                        `json:"charges"`
	ClientSecret              string                             `json:"client_secret"`
	ConfirmationMethod        PaymentIntentConfirmationMethod    `json:"confirmation_method"`
	Created                   int64                              `json:"created"`
	Currency                  Currency                           `json:"currency"`
	Customer                  *Customer                          `json:"customer"`
	Description               string                             `json:"description"`
	ID                        string                             `json:"id"`
	Invoice                   *Invoice                           `json:"invoice"`
	LastPaymentError          *Error                             `json:"last_payment_error"`
	Livemode                  bool                               `json:"livemode"`
	Metadata                  map[string]string                  `json:"metadata"`
	NextAction                *PaymentIntentNextAction           `json:"next_action"`
	Object                    string                             `json:"object"`
	OnBehalfOf                *Account                           `json:"on_behalf_of"`
	PaymentMethod             *PaymentMethod                     `json:"payment_method"`
	PaymentMethodOptions      *PaymentIntentPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes        []string                           `json:"payment_method_types"`
	ReceiptEmail              string                             `json:"receipt_email"`
	Review                    *Review                            `json:"review"`
	SetupFutureUsage          PaymentIntentSetupFutureUsage      `json:"setup_future_usage"`
	Shipping                  *ShippingDetails                   `json:"shipping"`
	Source                    *PaymentIntentSource               `json:"source"`
	StatementDescriptor       string                             `json:"statement_descriptor"`
	StatementDescriptorSuffix string                             `json:"statement_descriptor_suffix"`
	Status                    PaymentIntentStatus                `json:"status"`
	TransferData              *PaymentIntentTransferData         `json:"transfer_data"`
	TransferGroup             string                             `json:"transfer_group"`
}
type PaymentIntentSource struct {
	ID   string                  `json:"id"`
	Type PaymentIntentSourceType `json:"object"`

	Account         *Account         `json:"-"`
	AlipayAccount   *AlipayAccount   `json:"-"`
	BankAccount     *BankAccount     `json:"-"`
	BitcoinReceiver *BitcoinReceiver `json:"-"`
	Card            *Card            `json:"-"`
	PaymentSource   *PaymentSource   `json:"-"`
}

// PaymentIntentList is a list of payment intents as retrieved from a list endpoint.
type PaymentIntentList struct {
	APIResource
	ListMeta
	Data []*PaymentIntent `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentIntent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentIntent PaymentIntent
	var v paymentIntent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentIntent(v)
	return nil
}

// UnmarshalJSON handles deserialization of a PaymentIntentSource.
// This custom unmarshaling is needed because the specific type of
// PaymentIntentSource it refers to is specified in the JSON
func (p *PaymentIntentSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentIntentSource PaymentIntentSource
	var v paymentIntentSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentIntentSource(v)
	var err error

	switch p.Type {
	case PaymentIntentSourceTypeAlipayAccount:
		err = json.Unmarshal(data, &p.AlipayAccount)
	case PaymentIntentSourceTypeBankAccount:
		err = json.Unmarshal(data, &p.BankAccount)
	case PaymentIntentSourceTypeBitcoinReceiver:
		err = json.Unmarshal(data, &p.BitcoinReceiver)
	case PaymentIntentSourceTypeCard:
		err = json.Unmarshal(data, &p.Card)
	case PaymentIntentSourceTypeAccount:
		err = json.Unmarshal(data, &p.Account)
	case PaymentIntentSourceTypePaymentSource:
		err = json.Unmarshal(data, &p.PaymentSource)
	}
	return err
}
