package stripe

import "encoding/json"

// The reason for a FinancialAccount closure.
type FinancialAccountCloseReason string

// List of values that FinancialAccountCloseReason can take
const (
	FinancialAccountCloseReasonAccountRejected  FinancialAccountCloseReason = "account_rejected"
	FinancialAccountCloseReasonClosedByPlatform FinancialAccountCloseReason = "closed_by_platform"
	FinancialAccountCloseReasonOther            FinancialAccountCloseReason = "other"
)

// The feature of a FinancialAccount.
type FinancialAccountFeatureName string

// List of values that FinancialAccountFeatureName can take
const (
	FinancialAccountFeatureCardIssuing                     FinancialAccountFeatureName = "card_issuing"
	FinancialAccountFeatureOutboundPaymentsAch             FinancialAccountFeatureName = "outbound_payments.ach"
	FinancialAccountFeatureOutboundPaymentsUsDomesticWire  FinancialAccountFeatureName = "outbound_payments.us_domestic_wire"
	FinancialAccountFeatureFinancialAddressesAba           FinancialAccountFeatureName = "financial_addresses.aba"
	FinancialAccountFeatureDepositInsurance                FinancialAccountFeatureName = "deposit_insurance"
	FinancialAccountFeatureIntraStripeFlows                FinancialAccountFeatureName = "intra_stripe_flows"
	FinancialAccountFeatureInboundTransfersCard            FinancialAccountFeatureName = "inbound_transfers.card"
	FinancialAccountFeatureInboundTransfersAch             FinancialAccountFeatureName = "inbound_transfers.ach"
	FinancialAccountFeatureOutboundTransfersAch            FinancialAccountFeatureName = "outbound_transfers.ach"
	FinancialAccountFeatureOutboundTransfersUsDomesticWire FinancialAccountFeatureName = "outbound_transfers.us_domestic_wire"
	FinancialAccountFeatureRemoteDepositCapture            FinancialAccountFeatureName = "remote_deposit_capture"
)

// Currently, FinancialAccounts can only carry balances in USD.
type FinancialAccountBalanceFunds struct {
	USD int64 `json:"usd"`
}

// The single multi-currency balance of the FinancialAccount.
// Positive values represent money that belongs to the user while negative values represent funds the user owes.
// Currently, FinancialAccounts can only carry balances in USD.
type FinancialAccountBalance struct {
	// Funds the user can spend right now.
	Cash FinancialAccountBalanceFunds `json:"cash"`
	// Funds not spendable yet, but will become available at a later time.
	InboundPending FinancialAccountBalanceFunds `json:"inbound_pending"`
	// Funds in the account, but not spendable because they are being held for pending outbound flows.
	OutboundPending FinancialAccountBalanceFunds `json:"outbound_pending"`
}

// The feature status of a FinancialAccount.
type FinancialAccountFeatureStatus string

// List of values that FinancialAccountFeatureStatus can take
const (
	FinancialAccountFeatureStatusActive     FinancialAccountFeatureStatus = "active"
	FinancialAccountFeatureStatusRestricted FinancialAccountFeatureStatus = "restricted"
	FinancialAccountFeatureStatusPending    FinancialAccountFeatureStatus = "pending"
)

// The feature status code of a FinancialAccount.
type FinancialAccountFeatureStatusCode string

// List of values that FinancialAccountFeatureStatusCode can take
const (
	FinancialAccountFeatureStatusCodeCapabilityNotRequested          FinancialAccountFeatureStatusCode = "capability_not_requested"
	FinancialAccountFeatureStatusCodeRequirementsPastDue             FinancialAccountFeatureStatusCode = "requirements_past_due"
	FinancialAccountFeatureStatusCodeRequirementsPendingVerification FinancialAccountFeatureStatusCode = "requirements_pending_verification"
	FinancialAccountFeatureStatusCodeRejectedUnsupportedBusiness     FinancialAccountFeatureStatusCode = "rejected_unsupported_business"
	FinancialAccountFeatureStatusCodeRejectedOther                   FinancialAccountFeatureStatusCode = "rejected_other"
	FinancialAccountFeatureStatusCodeRestrictedByPlatform            FinancialAccountFeatureStatusCode = "restricted_by_platform"
	FinancialAccountFeatureStatusCodeRestrictedOther                 FinancialAccountFeatureStatusCode = "restricted_other"
	FinancialAccountFeatureStatusCodeActivating                      FinancialAccountFeatureStatusCode = "activating"
	FinancialAccountFeatureStatusCodeFinancialAccountClosed          FinancialAccountFeatureStatusCode = "financial_account_closed"
)

// The feature status resolution of a FinancialAccount.
type FinancialAccountFeatureStatusResolution string

// List of values that FinancialAccountFeatureStatusResolution can take
const (
	FinancialAccountFeatureStatusResolutionRemoveRestriction  FinancialAccountFeatureStatusResolution = "remove_restriction"
	FinancialAccountFeatureStatusResolutionProvideInformation FinancialAccountFeatureStatusResolution = "provide_information"
	FinancialAccountFeatureStatusResolutionContactStripe      FinancialAccountFeatureStatusResolution = "contact_stripe"
)

// The feature status restriction of a FinancialAccount.
type FinancialAccountFeatureStatusRestriction string

// List of values that FinancialAccountFeatureStatusRestriction can take
const (
	FinancialAccountFeatureStatusRestrictionInboundFlows  FinancialAccountFeatureStatusRestriction = "inbound_flows"
	FinancialAccountFeatureStatusRestrictionOutboundFlows FinancialAccountFeatureStatusRestriction = "outbound_flows"
)

// Encodes whether a FinancialAccount has access to a particular feature, with a status enum and associated status_details.
type FinancialAccountFeature struct {
	// Whether the FinancialAccount should have the Feature.
	Requested bool `json:"requested"`
	// Whether the Feature is operational.
	Status FinancialAccountFeatureStatus `json:"status"`
	// Represents the reason why the status is pending or restricted.
	Code FinancialAccountFeatureStatusCode `json:"code"`
	//	Represents what the user should do, if anything, to activate the Feature.
	Resolution FinancialAccountFeatureStatusResolution `json:"resolution"`
	// The platform_restrictions that are restricting this Feature.
	Restriction FinancialAccountFeatureStatusRestriction `json:"restriction"`
}

// Identifying information for the ABA address.
type FinancialAddressABA struct {
	// The name of the person or business that owns the bank account.
	AccountHolderName string `json:"account_holder_name"`
	// The account number. This field is not included by default. To include it in the response, expand the account_number field.
	AccountNumber string `json:"account_number"`
	// The last four characters of the account number.
	AccountNumberLast4 string `json:"account_number_last4"`
	// Name of the bank.
	BankName string `json:"bank_name"`
	// Routing number for the account.
	RoutingNumber string `json:"routing_number"`
}

// Identifying information for the iban address.
type FinancialAddressIban struct {
	// The name of the person or business that owns the bank account.
	AccountHolderName string `json:"account_holder_name"`
	// The BIC/SWIFT code of the account.
	BIC string `json:"bic"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// The IBAN of the account.
	Iban string `json:"iban"`
}

// Identifying information for the sort code address.
type FinancialAddressSortCode struct {
	// The name of the person or business that owns the bank account.
	AccountHolderName string `json:"account_holder_name"`
	// The account number.
	AccountNumber string `json:"account_number"`
	// The six-digit sort code.
	SortCode string `json:"sort_code"`
}

// Identifying information for the sort code address.
type FinancialAddressSPEI struct {
	// The three-digit bank code.
	BankCode string `json:"bank_code"`
	// The short banking institution name.
	BankName string `json:"bank_name"`
	// The CLABE number.
	CLABE string `json:"clabe"`
}

// The list of networks that the address supports.
type FinancialAddressSupportedNetwork string

// List of values that FinancialAddressSupportedNetwork can take
const (
	FinancialAddressSupportedNetworkAch            FinancialAddressSupportedNetwork = "ach"
	FinancialAddressSupportedNetworkUsDomesticWire FinancialAddressSupportedNetwork = "us_domestic_wire"
)

// Identifying information for the SWIFT address.
type FinancialAddressSWIFT struct {
	// The account holder name.
	AccountHolderName string `json:"account_holder_name"`
	// The account number.
	AccountNumber string `json:"account_number"`
	// The bank name.
	BankName string `json:"bank_name"`
	// The SWIFT code.
	SWIFTCode string `json:"swift_code"`
}

// The type of financial address.
type FinancialAddressType string

// List of values that FinancialAddressType can take
const (
	FinancialAddressTypeABA FinancialAddressType = "aba"
)

// The set of credentials that resolve to a FinancialAccount.
type FinancialAddresses struct {
	// Identifying information for the ABA address
	ABA FinancialAddressABA `json:"aba"`
	// Identifying information for the iban address
	Iban FinancialAddressIban `json:"iban"`
	// Identifying information for the id bban address
	IDBban map[string]interface{} `json:"id_bban"`
	// Identifying information for the sort code address
	SortCode FinancialAddressSortCode `json:"sort_code"`
	// Identifying information for the sort code address
	SPEI FinancialAddressSPEI `json:"spei"`
	// The list of networks that the address supports
	SupportedNetworks FinancialAddressSupportedNetwork `json:"supported_networks"`
	// Identifying information for the SWIFT address
	SWIFT FinancialAddressSWIFT `json:"swift"`
	// The type of financial address
	Type FinancialAddressType `json:"type"`
	// Identifying information for the zengin address
	Zengin map[string]interface{} `json:"zengin"`
}

// The platrofm restriction status.
type PlatformRestrictionStatus string

// List of values that PlatformRestrictionStatus can take
const (
	PlatformRestrictionRestricted   PlatformRestrictionStatus = "restricted"
	PlatformRestrictionUnrestricted PlatformRestrictionStatus = "unrestricted"
)

func PlatformRestrictionStatusP(s PlatformRestrictionStatus) *PlatformRestrictionStatus {
	return &s
}

// The financial account status.
type FinancialAccountStatus string

// List of values that FinancialAccountStatus can take
const (
	PlatformRestrictionOpen   FinancialAccountStatus = "open"
	PlatformRestrictionClosed FinancialAccountStatus = "closed"
)

// Encodes whether a FinancialAccount has access to a particular feature, with a status enum and associated status_details.
type FinancialAccountFeatures struct {
	// Contains a Feature encoding the FinancialAccount’s ability to be used with the Issuing product, including attaching cards to and drawing funds from.
	CardIssuing FinancialAccountFeature `json:"card_issuing"`
	// Represents whether this FinancialAccount is eligible for deposit insurance. Various factors determine the insurance amount.
	DepositInsurance FinancialAccountFeature `json:"deposit_insurance"`
	// Contains Features that add FinancialAddresses to the FinancialAccount.
	FinancialAddresses FinancialAccountFeature `json:"financial_addresses"`
	// Contains settings related to adding funds to a FinancialAccount from another Account with the same owner.
	InboundTransfers FinancialAccountFeature `json:"inbound_transfers"`
	// Represents the ability for this FinancialAccount to send money to, or receive money from other FinancialAccounts (for example, via OutboundPayment).
	IntraStripeFlows FinancialAccountFeature `json:"intra_stripe_flows"`
	// Contains Features related to initiating money movement out of the FinancialAccount to someone else’s bucket of money.
	OutboundPayments FinancialAccountFeature `json:"outbound_payments"`
	// Contains a Feature and settings related to moving money out of the FinancialAccount into another Account with the same owner.
	OutboundTransfers FinancialAccountFeature `json:"outbound_transfers"`
}

// The set of functionalities that the platform can restrict on the FinancialAccount.
type FinancialAccountPlatformRestrictions struct {
	// Restricts all inbound money movement.
	InboundFlows PlatformRestrictionStatus `json:"inbound_flows"`
	// Restricts all outbound money movement.
	OutboundFlows PlatformRestrictionStatus `json:"outbound_flows"`
}

// This is an object representing a Stripe financial account.
type FinancialAccount struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The array that contains reasons for a FinancialAccount closure. Non-empty only if the status of the FinancialAccount is "closed".
	AccountClosedReasons []FinancialAccountCloseReason `json:"account_closed_reasons"`
	// The array of paths to active Features in the Features hash.
	ActiveFeatures []FinancialAccountFeatureName `json:"active_features"`
	// The single multi-currency balance of the FinancialAccount.
	Balance FinancialAccountBalance `json:"balance"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Time at which the financial account was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Encodes whether a FinancialAccount has access to a particular feature, with a status enum and associated status_details.
	// Stripe or the platform may control features via the requested field.
	// This field is not included by default. To include it in the response, expand the features field.
	Features FinancialAccountFeatures `json:"features"`
	// The set of credentials that resolve to a FinancialAccount.
	FinancialAddresses FinancialAddresses `json:"financial_addresses"`
	// Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The array of paths to pending Features in the Features hash.
	PendingFeatures []FinancialAccountFeatureName `json:"pending_features"`
	// The set of functionalities that the platform can restrict on the FinancialAccount.
	PlatformRestrictions FinancialAccountPlatformRestrictions `json:"platform_restrictions"`
	// The array of paths to restricted Features in the Features hash.
	RestrictedFeatures []FinancialAccountFeatureName `json:"restricted_features"`
	// Specifying what state the account is in.
	Status FinancialAccountStatus `json:"status"`
	// The currencies the FinancialAccount can hold a balance in. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.
	SupportedCurrencies []string `json:"supported_currencies"`
}

// The card_issuing feature.
type FinancialAccountFeaturesCardIssuingParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The deposit_insurance feature.
type FinancialAccountFeaturesDepositInsuranceParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The financial_addresses aba feature.
type FinancialAccountFeaturesFinancialAddressesABAParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The financial_addresses feature.
type FinancialAccountFeaturesFinancialAddressesParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	ABA *FinancialAccountFeaturesFinancialAddressesABAParams `form:"aba"`
}

// The inbound_transfers ach feature.
type FinancialAccountFeaturesInboundTransfersACHParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The inbound_transfers card feature.
type FinancialAccountFeaturesInboundTransfersCardParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The inbound_transfers us domestic wire feature.
type FinancialAccountFeaturesInboundTransfersUsDomesticWireParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The inbound_transfers feature.
type FinancialAccountFeaturesInboundTransfersParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	ACH  *FinancialAccountFeaturesInboundTransfersACHParams  `form:"ach"`
	Card *FinancialAccountFeaturesInboundTransfersCardParams `form:"card"`
}

// The intra_stripe_flows feature.
type FinancialAccountFeaturesIntraStripeFlowsParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	Requested *bool `form:"requested"`
}

// The outbound_payments feature.
type FinancialAccountFeaturesOutboundPaymentsParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	ACH            *FinancialAccountFeaturesInboundTransfersACHParams            `form:"ach"`
	UsDomesticWire *FinancialAccountFeaturesInboundTransfersUsDomesticWireParams `form:"us_domestic_wire"`
}

// The outbound_transfers feature.
type FinancialAccountFeaturesOutboundTransfersParams struct {
	// Passing true requests the feature for the financial account, if it is not already requested.
	ACH            *FinancialAccountFeaturesInboundTransfersACHParams            `form:"ach"`
	UsDomesticWire *FinancialAccountFeaturesInboundTransfersUsDomesticWireParams `form:"us_domestic_wire"`
}

// Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.
type FinancialAccountFeaturesParams struct {
	// Encodes the FinancialAccount’s ability to be used with the Issuing product, including attaching cards to and drawing funds from the FinancialAccount.
	CardIssuing *FinancialAccountFeaturesCardIssuingParams `form:"card_issuing"`
	// Represents whether this FinancialAccount is eligible for deposit insurance. Various factors determine the insurance amount.
	DepositInsurance *FinancialAccountFeaturesDepositInsuranceParams `form:"deposit_insurance"`
	// Contains Features that add FinancialAddresses to the FinancialAccount.
	FinancialAddresses *FinancialAccountFeaturesFinancialAddressesParams `form:"financial_addresses"`
	// Contains settings related to adding funds to a FinancialAccount from another Account with the same owner.
	InboundTransfers *FinancialAccountFeaturesInboundTransfersParams `form:"inbound_transfers"`
	// Represents the ability for the FinancialAccount to send money to, or receive money from other FinancialAccounts (for example, via OutboundPayment).
	IntraStripeFlows *FinancialAccountFeaturesIntraStripeFlowsParams `form:"intra_stripe_flows"`
	// Includes Features related to initiating money movement out of the FinancialAccount to someone else’s bucket of money.
	OutboundPayments *FinancialAccountFeaturesOutboundPaymentsParams `form:"outbound_payments"`
	// Contains a Feature and settings related to moving money out of the FinancialAccount into another Account with the same owner.
	OutboundTransfers *FinancialAccountFeaturesOutboundTransfersParams `form:"outbound_transfers"`
}

// The set of functionalities that the platform can restrict on the FinancialAccount.
type FinancialAccountPlatformRestrictionsParams struct {
	// Restricts all inbound money movement.
	InboundFlows *PlatformRestrictionStatus `form:"inbound_flows"`
	// Restricts all outbound money movement.
	OutboundFlows *PlatformRestrictionStatus `form:"outbound_flows"`
}

// Retrieves the details of a financial account.
type FinancialAccountParams struct {
	Params `form:"*"`
	// The currencies the FinancialAccount can hold a balance in.
	SupportedCurrencies []*string `form:"supported_currencies"`
	// Encodes whether a FinancialAccount has access to a particular feature. Stripe or the platform can control features via the requested field.
	Features *FinancialAccountFeaturesParams `form:"features"`
	// The set of functionalities that the platform can restrict on the FinancialAccount.
	PlatformRestrictions *FinancialAccountPlatformRestrictionsParams `form:"platform_restrictions"`
}

// Returns a list of financial accounts.
type FinancialAccountListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// FinancialAccountList is a list of FinancialAccounts as retrieved from a list endpoint.
type FinancialAccountList struct {
	APIResource
	ListMeta
	Data []*Account `json:"data"`
}

// UnmarshalJSON handles deserialization of an FinancialAccount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (a *FinancialAccount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		a.ID = id
		return nil
	}

	type account FinancialAccount
	var v account
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*a = FinancialAccount(v)
	return nil
}
