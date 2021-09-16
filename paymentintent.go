//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"
import "github.com/stripe/stripe-go/v72/form"

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

// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.
// One of `month`.
type PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanInterval string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanInterval can take
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanIntervalMonth PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanInterval = "month"
)

// Type of installment plan, one of `fixed_count`.
type PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanType string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanType can take
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanTypeFixedCount PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanType = "fixed_count"
)

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval is the interval of a card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval can take.
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsPlanIntervalMonth PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval = "month"
)

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType is the type of a card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType can take.
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsPlanTypeFixedCount PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType = "fixed_count"
)

// Selected network to process this payment intent on. Depends on the available networks of the card attached to the payment intent. Can be only set confirm-time.
type PaymentIntentPaymentMethodOptionsCardNetwork string

// List of values that PaymentIntentPaymentMethodOptionsCardNetwork can take
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

// Preferred language of the SOFORT authorization page that the customer is redirected to.
type PaymentIntentPaymentMethodOptionsSofortPreferredLanguage string

// List of values that PaymentIntentPaymentMethodOptionsSofortPreferredLanguage can take
const (
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageDe PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "de"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageEn PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "en"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageEs PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "es"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageFr PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "fr"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageIt PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "it"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguageNl PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "nl"
	PaymentIntentPaymentMethodOptionsSofortPreferredLanguagePl PaymentIntentPaymentMethodOptionsSofortPreferredLanguage = "pl"
)

type PaymentIntentPaymentMethodOptionsWechatPayClient string

// List of values that PaymentIntentPaymentMethodOptionsWechatPayClient can take
const (
	PaymentIntentPaymentMethodOptionsWechatPayClientAndroid PaymentIntentPaymentMethodOptionsWechatPayClient = "android"
	PaymentIntentPaymentMethodOptionsWechatPayClientIos     PaymentIntentPaymentMethodOptionsWechatPayClient = "ios"
	PaymentIntentPaymentMethodOptionsWechatPayClientWeb     PaymentIntentPaymentMethodOptionsWechatPayClient = "web"
)

// PaymentIntentSetupFutureUsage is the list of allowed values for SetupFutureUsage.
type PaymentIntentSetupFutureUsage string

// List of values that PaymentIntentSetupFutureUsage can take.
const (
	PaymentIntentSetupFutureUsageOffSession PaymentIntentSetupFutureUsage = "off_session"
	PaymentIntentSetupFutureUsageOnSession  PaymentIntentSetupFutureUsage = "on_session"
)

type PaymentIntentSourceType string

// List of values that PaymentIntentSourceType can take
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

// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
type PaymentIntentConfirmMandateDataCustomerAcceptanceOfflineParams struct{}

// If this is a Mandate accepted online, this hash contains details about the online acceptance.
type PaymentIntentConfirmMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type PaymentIntentConfirmMandateDataCustomerAcceptanceParams struct {
	AcceptedAt *int64                                                          `form:"accepted_at"`
	Offline    *PaymentIntentConfirmMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *PaymentIntentConfirmMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       *string                                                         `form:"type"`
}

// This hash contains details about the Mandate to create
type PaymentIntentConfirmMandateDataParams struct {
	CustomerAcceptance *PaymentIntentConfirmMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type PaymentIntentConfirmPaymentMethodDataACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type PaymentIntentConfirmPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type PaymentIntentConfirmPaymentMethodDataAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type PaymentIntentConfirmPaymentMethodDataAUBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BSBNumber     *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type PaymentIntentConfirmPaymentMethodDataBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type PaymentIntentConfirmPaymentMethodDataBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type PaymentIntentConfirmPaymentMethodDataBillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type PaymentIntentConfirmPaymentMethodDataBoletoParams struct {
	TaxID *string `form:"tax_id"`
}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type PaymentIntentConfirmPaymentMethodDataEPSParams struct {
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type PaymentIntentConfirmPaymentMethodDataFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type PaymentIntentConfirmPaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type PaymentIntentConfirmPaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type PaymentIntentConfirmPaymentMethodDataIdealParams struct {
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type PaymentIntentConfirmPaymentMethodDataInteracPresentParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type PaymentIntentConfirmPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type PaymentIntentConfirmPaymentMethodDataP24Params struct {
	Bank *string `form:"bank"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type PaymentIntentConfirmPaymentMethodDataSepaDebitParams struct {
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type PaymentIntentConfirmPaymentMethodDataSofortParams struct {
	Country *string `form:"country"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type PaymentIntentConfirmPaymentMethodDataWechatPayParams struct{}

// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear
// in the [payment_method](https://stripe.com/docs/api/payment_intents/object#payment_intent_object-payment_method)
// property on the PaymentIntent.
type PaymentIntentConfirmPaymentMethodDataParams struct {
	ACSSDebit        *PaymentIntentConfirmPaymentMethodDataACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentConfirmPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentConfirmPaymentMethodDataAlipayParams           `form:"alipay"`
	AUBECSDebit      *PaymentIntentConfirmPaymentMethodDataAUBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentIntentConfirmPaymentMethodDataBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentIntentConfirmPaymentMethodDataBancontactParams       `form:"bancontact"`
	BillingDetails   *PaymentIntentConfirmPaymentMethodDataBillingDetailsParams   `form:"billing_details"`
	Boleto           *PaymentIntentConfirmPaymentMethodDataBoletoParams           `form:"boleto"`
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
	WechatPay        *PaymentIntentConfirmPaymentMethodDataWechatPayParams        `form:"wechat_pay"`
}

// Additional fields for Mandate creation
type PaymentIntentConfirmPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string `form:"custom_mandate_url"`
	IntervalDescription *string `form:"interval_description"`
	PaymentSchedule     *string `form:"payment_schedule"`
	TransactionType     *string `form:"transaction_type"`
}

// If this is a `acss_debit` PaymentMethod, this sub-hash contains details about the ACSS Debit payment method options.
type PaymentIntentConfirmPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *PaymentIntentConfirmPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                `form:"verification_method"`
}

// If this is a `afterpay_clearpay` PaymentMethod, this sub-hash contains details about the Afterpay Clearpay payment method options.
type PaymentIntentConfirmPaymentMethodOptionsAfterpayClearpayParams struct {
	Reference *string `form:"reference"`
}

// If this is a `alipay` PaymentMethod, this sub-hash contains details about the Alipay payment method options.
type PaymentIntentConfirmPaymentMethodOptionsAlipayParams struct{}

// If this is a `bancontact` PaymentMethod, this sub-hash contains details about the Bancontact payment method options.
type PaymentIntentConfirmPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// If this is a `boleto` PaymentMethod, this sub-hash contains details about the Boleto payment method options.
type PaymentIntentConfirmPaymentMethodOptionsBoletoParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// The selected installment plan to use for this payment attempt.
// This parameter can only be provided during confirmation.
type PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsPlanParams struct {
	Count    *int64  `form:"count"`
	Interval *string `form:"interval"`
	Type     *string `form:"type"`
}

// Installment configuration for payments attempted on this PaymentIntent (Mexico Only).
//
// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
type PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsParams struct {
	Enabled *bool                                                               `form:"enabled"`
	Plan    *PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}

// Configuration for any card payments attempted on this PaymentIntent.
type PaymentIntentConfirmPaymentMethodOptionsCardParams struct {
	CVCToken            *string                                                         `form:"cvc_token"`
	Installments        *PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	Moto                *bool                                                           `form:"moto"`
	Network             *string                                                         `form:"network"`
	RequestThreeDSecure *string                                                         `form:"request_three_d_secure"`
}

// If this is a `card_present` PaymentMethod, this sub-hash contains details about the Card Present payment method options.
type PaymentIntentConfirmPaymentMethodOptionsCardPresentParams struct{}

// If this is a `ideal` PaymentMethod, this sub-hash contains details about the Ideal payment method options.
type PaymentIntentConfirmPaymentMethodOptionsIdealParams struct{}

// If this is a `oxxo` PaymentMethod, this sub-hash contains details about the OXXO payment method options.
type PaymentIntentConfirmPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// If this is a `p24` PaymentMethod, this sub-hash contains details about the Przelewy24 payment method options.
type PaymentIntentConfirmPaymentMethodOptionsP24Params struct {
	TosShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// Additional fields for Mandate creation
type PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}

// If this is a `sepa_debit` PaymentIntent, this sub-hash contains details about the SEPA Debit payment method options.
type PaymentIntentConfirmPaymentMethodOptionsSepaDebitParams struct {
	MandateOptions *PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// If this is a `sofort` PaymentMethod, this sub-hash contains details about the SOFORT payment method options.
type PaymentIntentConfirmPaymentMethodOptionsSofortParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// If this is a `wechat_pay` PaymentMethod, this sub-hash contains details about the WeChat Pay payment method options.
type PaymentIntentConfirmPaymentMethodOptionsWechatPayParams struct {
	AppID  *string `form:"app_id"`
	Client *string `form:"client"`
}

// Payment-method-specific configuration for this PaymentIntent.
type PaymentIntentConfirmPaymentMethodOptionsParams struct {
	ACSSDebit        *PaymentIntentConfirmPaymentMethodOptionsACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentConfirmPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentConfirmPaymentMethodOptionsAlipayParams           `form:"alipay"`
	Bancontact       *PaymentIntentConfirmPaymentMethodOptionsBancontactParams       `form:"bancontact"`
	Boleto           *PaymentIntentConfirmPaymentMethodOptionsBoletoParams           `form:"boleto"`
	Card             *PaymentIntentConfirmPaymentMethodOptionsCardParams             `form:"card"`
	CardPresent      *PaymentIntentConfirmPaymentMethodOptionsCardPresentParams      `form:"card_present"`
	Ideal            *PaymentIntentConfirmPaymentMethodOptionsIdealParams            `form:"ideal"`
	OXXO             *PaymentIntentConfirmPaymentMethodOptionsOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentConfirmPaymentMethodOptionsP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentConfirmPaymentMethodOptionsSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentConfirmPaymentMethodOptionsSofortParams           `form:"sofort"`
	WechatPay        *PaymentIntentConfirmPaymentMethodOptionsWechatPayParams        `form:"wechat_pay"`
}

// PaymentIntentConfirmParams is the set of parameters that can be used when confirming a payment intent.
type PaymentIntentConfirmParams struct {
	Params                `form:"*"`
	ErrorOnRequiresAction *bool   `form:"error_on_requires_action"`
	Mandate               *string `form:"mandate"`
	// MandateData *[todo({"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentIntentConfirmMandateDataParams"}} | {"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentIntentConfirmMandateDataParams"}})] `form:"mandate_data"`
	OffSession           *bool                                           `form:"off_session"`
	OffSessionOneOff     *bool                                           `form:"-"` // See custom AppendTo
	OffSessionRecurring  *bool                                           `form:"-"` // See custom AppendTo
	PaymentMethod        *string                                         `form:"payment_method"`
	PaymentMethodData    *PaymentIntentConfirmPaymentMethodDataParams    `form:"payment_method_data"`
	PaymentMethodOptions *PaymentIntentConfirmPaymentMethodOptionsParams `form:"payment_method_options"`
	ReceiptEmail         *string                                         `form:"receipt_email"`
	ReturnURL            *string                                         `form:"return_url"`
	SetupFutureUsage     *string                                         `form:"setup_future_usage"`
	Shipping             *ShippingDetailsParams                          `form:"shipping"`
	UseStripeSDK         *bool                                           `form:"use_stripe_sdk"`
}

// AppendTo implements custom encoding logic for PaymentIntentConfirmParams.
func (p *PaymentIntentConfirmParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.OffSessionOneOff) {
		body.Add(form.FormatKey(append(keyParts, "off_session")), "one_off")
	}
	if BoolValue(p.OffSessionRecurring) {
		body.Add(form.FormatKey(append(keyParts, "off_session")), "recurring")
	}
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

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type PaymentIntentPaymentMethodDataACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type PaymentIntentPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type PaymentIntentPaymentMethodDataAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type PaymentIntentPaymentMethodDataAUBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BSBNumber     *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type PaymentIntentPaymentMethodDataBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type PaymentIntentPaymentMethodDataBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type PaymentIntentPaymentMethodDataBillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type PaymentIntentPaymentMethodDataBoletoParams struct {
	TaxID *string `form:"tax_id"`
}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type PaymentIntentPaymentMethodDataEPSParams struct {
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type PaymentIntentPaymentMethodDataFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type PaymentIntentPaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type PaymentIntentPaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type PaymentIntentPaymentMethodDataIdealParams struct {
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type PaymentIntentPaymentMethodDataInteracPresentParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type PaymentIntentPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type PaymentIntentPaymentMethodDataP24Params struct {
	Bank *string `form:"bank"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type PaymentIntentPaymentMethodDataSepaDebitParams struct {
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type PaymentIntentPaymentMethodDataSofortParams struct {
	Country *string `form:"country"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type PaymentIntentPaymentMethodDataWechatPayParams struct{}

// PaymentIntentPaymentMethodDataParams represents the type-specific parameters associated with a
// payment method on payment intent.
type PaymentIntentPaymentMethodDataParams struct {
	ACSSDebit        *PaymentIntentPaymentMethodDataACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodDataAlipayParams           `form:"alipay"`
	AUBECSDebit      *PaymentIntentPaymentMethodDataAUBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentIntentPaymentMethodDataBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentIntentPaymentMethodDataBancontactParams       `form:"bancontact"`
	BillingDetails   *PaymentIntentPaymentMethodDataBillingDetailsParams   `form:"billing_details"`
	Boleto           *PaymentIntentPaymentMethodDataBoletoParams           `form:"boleto"`
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
	WechatPay        *PaymentIntentPaymentMethodDataWechatPayParams        `form:"wechat_pay"`
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

// PaymentIntentPaymentMethodOptionsBoletoParams represents the boleto-specific options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsBoletoParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
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

// If this is a `card_present` PaymentMethod, this sub-hash contains details about the Card Present payment method options.
type PaymentIntentPaymentMethodOptionsCardPresentParams struct{}

// If this is a `ideal` PaymentMethod, this sub-hash contains details about the Ideal payment method options.
type PaymentIntentPaymentMethodOptionsIdealParams struct{}

// PaymentIntentPaymentMethodOptionsOXXOParams represents the OXXO-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// If this is a `p24` PaymentMethod, this sub-hash contains details about the Przelewy24 payment method options.
type PaymentIntentPaymentMethodOptionsP24Params struct {
	TosShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// Additional fields for Mandate creation
type PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}

// If this is a `sepa_debit` PaymentIntent, this sub-hash contains details about the SEPA Debit payment method options.
type PaymentIntentPaymentMethodOptionsSepaDebitParams struct {
	MandateOptions *PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// PaymentIntentPaymentMethodOptionsSofortParams represents the sofort-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsSofortParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// PaymentIntentPaymentMethodOptionsWechatPayParams represents the wechat_pay-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsWechatPayParams struct {
	AppID  *string `form:"app_id"`
	Client *string `form:"client"`
}

// PaymentIntentPaymentMethodOptionsParams represents the type-specific payment method options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsParams struct {
	ACSSDebit        *PaymentIntentPaymentMethodOptionsACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipayParams           `form:"alipay"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontactParams       `form:"bancontact"`
	Boleto           *PaymentIntentPaymentMethodOptionsBoletoParams           `form:"boleto"`
	Card             *PaymentIntentPaymentMethodOptionsCardParams             `form:"card"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresentParams      `form:"card_present"`
	Ideal            *PaymentIntentPaymentMethodOptionsIdealParams            `form:"ideal"`
	OXXO             *PaymentIntentPaymentMethodOptionsOXXOParams             `form:"oxxo"`
	P24              *PaymentIntentPaymentMethodOptionsP24Params              `form:"p24"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofortParams           `form:"sofort"`
	WechatPay        *PaymentIntentPaymentMethodOptionsWechatPayParams        `form:"wechat_pay"`
}

// PaymentIntentTransferDataParams is the set of parameters allowed for the transfer hash.
type PaymentIntentTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// PaymentIntentParams is the set of parameters that can be used when handling a payment intent.
type PaymentIntentParams struct {
	Params                    `form:"*"`
	Amount                    *int64                                   `form:"amount"`
	ApplicationFeeAmount      *int64                                   `form:"application_fee_amount"`
	CaptureMethod             *string                                  `form:"capture_method"`
	ClientSecret              *string                                  `form:"client_secret"`
	Confirm                   *bool                                    `form:"confirm"`
	ConfirmationMethod        *string                                  `form:"confirmation_method"`
	Currency                  *string                                  `form:"currency"`
	Customer                  *string                                  `form:"customer"`
	Description               *string                                  `form:"description"`
	ErrorOnRequiresAction     *bool                                    `form:"error_on_requires_action"`
	Mandate                   *string                                  `form:"mandate"`
	MandateData               *PaymentIntentMandateDataParams          `form:"mandate_data"`
	OffSession                *bool                                    `form:"off_session"`
	OffSessionOneOff          *bool                                    `form:"-"` // See custom AppendTo
	OffSessionRecurring       *bool                                    `form:"-"` // See custom AppendTo
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

// AppendTo implements custom encoding logic for PaymentIntentParams.
func (p *PaymentIntentParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.OffSessionOneOff) {
		body.Add(form.FormatKey(append(keyParts, "off_session")), "one_off")
	}
	if BoolValue(p.OffSessionRecurring) {
		body.Add(form.FormatKey(append(keyParts, "off_session")), "recurring")
	}
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

// PaymentIntentNextActionBoletoDisplayDetails represents the resource for the next action of type
// "boleto_display_details".
type PaymentIntentNextActionBoletoDisplayDetails struct {
	ExpiresAt        int64  `json:"expires_at"`
	HostedVoucherURL string `json:"hosted_voucher_url"`
	Number           string `json:"number"`
	PDF              string `json:"pdf"`
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
type PaymentIntentNextActionWechatPayDisplayQrCode struct {
	Data         string `json:"data"`
	ImageDataURL string `json:"image_data_url"`
}

// PaymentIntentNextActionWechatPayRedirectToAndroidApp represents the resource for the next action of type
// "wechat_pay_redirect_to_android_app"
type PaymentIntentNextActionWechatPayRedirectToAndroidApp struct {
	AppID     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	PartnerID string `json:"partner_id"`
	PrepayID  string `json:"prepay_id"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}
type PaymentIntentNextActionWechatPayRedirectToIosApp struct {
	NativeURL string `json:"native_url"`
}

// PaymentIntentNextAction represents the type of action to take on a payment intent.
type PaymentIntentNextAction struct {
	AlipayHandleRedirect          *PaymentIntentNextActionAlipayHandleRedirect          `json:"alipay_handle_redirect"`
	BoletoDisplayDetails          *PaymentIntentNextActionBoletoDisplayDetails          `json:"boleto_display_details"`
	OXXODisplayDetails            *PaymentIntentNextActionOXXODisplayDetails            `json:"oxxo_display_details"`
	RedirectToURL                 *PaymentIntentNextActionRedirectToURL                 `json:"redirect_to_url"`
	Type                          string                                                `json:"type"`
	UseStripeSDK                  *PaymentIntentNextActionUseStripeSDK                  `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits       *PaymentIntentNextActionVerifyWithMicrodeposits       `json:"verify_with_microdeposits"`
	WechatPayDisplayQrCode        *PaymentIntentNextActionWechatPayDisplayQrCode        `json:"wechat_pay_display_qr_code"`
	WechatPayRedirectToAndroidApp *PaymentIntentNextActionWechatPayRedirectToAndroidApp `json:"wechat_pay_redirect_to_android_app"`
	WechatPayRedirectToIosApp     *PaymentIntentNextActionWechatPayRedirectToIosApp     `json:"wechat_pay_redirect_to_ios_app"`
}
type PaymentIntentPaymentMethodOptionsBoleto struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}

// Installment plans that may be selected for this PaymentIntent.
type PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlan struct {
	Count    int64                                                                  `json:"count"`
	Interval PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanInterval `json:"interval"`
	Type     PaymentIntentPaymentMethodOptionsCardInstallmentsAvailablePlanType     `json:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlan describe a specific card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlan struct {
	Count    int64                                                         `json:"count"`
	Interval PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval `json:"interval"`
	Type     PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType     `json:"type"`
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
type PaymentIntentPaymentMethodOptionsIdeal struct{}

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

// PaymentIntentPaymentMethodOptionsWechatPay is the set of wechat_pay-specific options associated
// with that payment intent.
type PaymentIntentPaymentMethodOptionsWechatPay struct {
	AppID  string                                           `json:"app_id"`
	Client PaymentIntentPaymentMethodOptionsWechatPayClient `json:"client"`
}

// PaymentIntentPaymentMethodOptions is the set of payment method-specific options associated with
// that payment intent.
type PaymentIntentPaymentMethodOptions struct {
	ACSSDebit        *PaymentIntentPaymentMethodOptionsACSSDebit        `json:"acss_debit"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipay           `json:"alipay"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontact       `json:"bancontact"`
	Boleto           *PaymentIntentPaymentMethodOptionsBoleto           `json:"boleto"`
	Card             *PaymentIntentPaymentMethodOptionsCard             `json:"card"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresent      `json:"card_present"`
	Ideal            *PaymentIntentPaymentMethodOptionsIdeal            `json:"ideal"`
	OXXO             *PaymentIntentPaymentMethodOptionsOXXO             `json:"oxxo"`
	P24              *PaymentIntentPaymentMethodOptionsP24              `json:"p24"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebit        `json:"sepa_debit"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofort           `json:"sofort"`
	WechatPay        *PaymentIntentPaymentMethodOptionsWechatPay        `json:"wechat_pay"`
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
