//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"
import "github.com/stripe/stripe-go/v72/form"

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionScheduleDefaultSettingsBillingCycleAnchor string

// List of values that SubscriptionScheduleDefaultSettingsBillingCycleAnchor can take
const (
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic  SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.
type SubscriptionScheduleDefaultSettingsCollectionMethod string

// List of values that SubscriptionScheduleDefaultSettingsCollectionMethod can take
const (
	SubscriptionScheduleDefaultSettingsCollectionMethodChargeAutomatically SubscriptionScheduleDefaultSettingsCollectionMethod = "charge_automatically"
	SubscriptionScheduleDefaultSettingsCollectionMethodSendInvoice         SubscriptionScheduleDefaultSettingsCollectionMethod = "send_invoice"
)

// SubscriptionScheduleEndBehavior describe what happens to a schedule when it ends.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take.
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.
type SubscriptionSchedulePhaseCollectionMethod string

// List of values that SubscriptionSchedulePhaseCollectionMethod can take
const (
	SubscriptionSchedulePhaseCollectionMethodChargeAutomatically SubscriptionSchedulePhaseCollectionMethod = "charge_automatically"
	SubscriptionSchedulePhaseCollectionMethodSendInvoice         SubscriptionSchedulePhaseCollectionMethod = "send_invoice"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type SubscriptionSchedulePhaseProrationBehavior string

// List of values that SubscriptionSchedulePhaseProrationBehavior can take
const (
	SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    SubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	SubscriptionSchedulePhaseProrationBehaviorCreateProrations SubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	SubscriptionSchedulePhaseProrationBehaviorNone             SubscriptionSchedulePhaseProrationBehavior = "none"
)

// SubscriptionScheduleStatus is the list of allowed values for the schedule's status.
type SubscriptionScheduleStatus string

// List of values that SubscriptionScheduleStatus can take.
const (
	SubscriptionScheduleStatusActive     SubscriptionScheduleStatus = "active"
	SubscriptionScheduleStatusCanceled   SubscriptionScheduleStatus = "canceled"
	SubscriptionScheduleStatusCompleted  SubscriptionScheduleStatus = "completed"
	SubscriptionScheduleStatusNotStarted SubscriptionScheduleStatus = "not_started"
	SubscriptionScheduleStatusReleased   SubscriptionScheduleStatus = "released"
)

// SubscriptionSchedulePhaseBillingCycleAnchor is the list of allowed values for the
// schedule's billing_cycle_anchor.
type SubscriptionSchedulePhaseBillingCycleAnchor string

// List of values for SubscriptionSchedulePhaseBillingCycleAnchor
const (
	SubscriptionSchedulePhaseBillingCycleAnchorAutomatic  SubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart SubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Default settings for automatic tax computation.
type SubscriptionScheduleDefaultSettingsAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleDefaultSettingsBillingThresholdsParams struct {
	AmountGte               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}

// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
type SubscriptionScheduleDefaultSettingsTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// SubscriptionScheduleDefaultSettingsParams is the set of parameters
// representing the subscription schedule’s default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	ApplicationFeePercent *float64                                                    `form:"application_fee_percent"`
	AutomaticTax          *SubscriptionScheduleDefaultSettingsAutomaticTaxParams      `form:"automatic_tax"`
	BillingCycleAnchor    *string                                                     `form:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionScheduleDefaultSettingsBillingThresholdsParams `form:"billing_thresholds"`
	CollectionMethod      *string                                                     `form:"collection_method"`
	DefaultPaymentMethod  *string                                                     `form:"default_payment_method"`
	InvoiceSettings       *SubscriptionScheduleDefaultSettingsInvoiceSettingsParams   `form:"invoice_settings"`
	TransferData          *SubscriptionScheduleDefaultSettingsTransferDataParams      `form:"transfer_data"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams is a structure representing the parameters
// to create an inline price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams struct {
	Currency          *string  `form:"currency"`
	Product           *string  `form:"product"`
	TaxBehavior       *string  `form:"tax_behavior"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// SubscriptionSchedulePhaseAddInvoiceItemParams is a structure representing the parameters allowed to control
// the invoice items to add at the start of a phase.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	Price     *string                                                 `form:"price"`
	PriceData *SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams `form:"price_data"`
	Quantity  *int64                                                  `form:"quantity"`
	TaxRates  []*string                                               `form:"tax_rates"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseBillingThresholdsParams struct {
	AmountGte               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}

// All invoices will be billed using the specified settings.
type SubscriptionSchedulePhaseInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseItemBillingThresholdsParams struct {
	UsageGTE *int64 `form:"usage_gte"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type SubscriptionSchedulePhaseItemPriceDataRecurringParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type SubscriptionSchedulePhaseItemPriceDataParams struct {
	Currency          *string                                                `form:"currency"`
	Product           *string                                                `form:"product"`
	Recurring         *SubscriptionSchedulePhaseItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                                `form:"tax_behavior"`
	UnitAmount        *int64                                                 `form:"unit_amount"`
	UnitAmountDecimal *float64                                               `form:"unit_amount_decimal,high_precision"`
}

// SubscriptionSchedulePhaseItemParams is a structure representing the parameters allowed to control
// a specic plan on a phase on a subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholdsParams `form:"billing_thresholds"`
	Plan              *string                                               `form:"plan"`
	Price             *string                                               `form:"price"`
	PriceData         *SubscriptionSchedulePhaseItemPriceDataParams         `form:"price_data"`
	Quantity          *int64                                                `form:"quantity"`
	TaxRates          []*string                                             `form:"tax_rates"`
}

// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
type SubscriptionSchedulePhaseTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// SubscriptionSchedulePhaseParams is a structure representing the parameters allowed to control
// a phase on a subscription schedule.
type SubscriptionSchedulePhaseParams struct {
	AddInvoiceItems                     []*SubscriptionSchedulePhaseAddInvoiceItemParams  `form:"add_invoice_items"`
	ApplicationFeePercent               *float64                                          `form:"application_fee_percent"`
	AutomaticTax                        *SubscriptionSchedulePhaseAutomaticTaxParams      `form:"automatic_tax"`
	BillingCycleAnchor                  *string                                           `form:"billing_cycle_anchor"`
	BillingThresholds                   *SubscriptionSchedulePhaseBillingThresholdsParams `form:"billing_thresholds"`
	CollectionMethod                    *string                                           `form:"collection_method"`
	Coupon                              *string                                           `form:"coupon"`
	DefaultPaymentMethod                *string                                           `form:"default_payment_method"`
	DefaultTaxRates                     []*string                                         `form:"default_tax_rates"`
	EndDateTODOINCOMPATIBLECOMBINATION  *int64                                            `form:"end_dateTODO_INCOMPATIBLE_COMBINATION"`
	InvoiceSettings                     *SubscriptionSchedulePhaseInvoiceSettingsParams   `form:"invoice_settings"`
	Items                               []*SubscriptionSchedulePhaseItemParams            `form:"items"`
	Iterations                          *int64                                            `form:"iterations"`
	ProrationBehavior                   *string                                           `form:"proration_behavior"`
	StartDate                           *int64                                            `form:"start_date"`
	StartDateNow                        *bool                                             `form:"-"` // See custom AppendTo
	TransferData                        *SubscriptionSchedulePhaseTransferDataParams      `form:"transfer_data"`
	Trial                               *bool                                             `form:"trial"`
	TrialEndTODOINCOMPATIBLECOMBINATION *int64                                            `form:"trial_endTODO_INCOMPATIBLE_COMBINATION"`
}

// AppendTo implements custom encoding logic for SubscriptionSchedulePhaseParams.
func (s *SubscriptionSchedulePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// SubscriptionScheduleParams is the set of parameters that can be used when creating or updating a
// subscription schedule.
type SubscriptionScheduleParams struct {
	Params            `form:"*"`
	Customer          *string                                    `form:"customer"`
	DefaultSettings   *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	EndBehavior       *string                                    `form:"end_behavior"`
	FromSubscription  *string                                    `form:"from_subscription"`
	Phases            []*SubscriptionSchedulePhaseParams         `form:"phases"`
	ProrationBehavior *string                                    `form:"proration_behavior"`
	StartDate         *int64                                     `form:"start_date"`
	StartDateNow      *bool                                      `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams.
func (s *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// SubscriptionScheduleCancelParams is the set of parameters that can be used when canceling a
// subscription schedule.
type SubscriptionScheduleCancelParams struct {
	Params     `form:"*"`
	InvoiceNow *bool `form:"invoice_now"`
	Prorate    *bool `form:"prorate"`
}

// SubscriptionScheduleReleaseParams is the set of parameters that can be used when releasing a
// subscription schedule.
type SubscriptionScheduleReleaseParams struct {
	Params             `form:"*"`
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// SubscriptionScheduleListParams is the set of parameters that can be used when listing
// subscription schedules.
type SubscriptionScheduleListParams struct {
	ListParams       `form:"*"`
	CanceledAt       *int64            `form:"canceled_at"`
	CanceledAtRange  *RangeQueryParams `form:"canceled_at"`
	CompletedAt      *int64            `form:"completed_at"`
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Customer         *string           `form:"customer"`
	ReleasedAt       *int64            `form:"released_at"`
	ReleasedAtRange  *RangeQueryParams `form:"released_at"`
	Scheduled        *bool             `form:"scheduled"`
}

// SubscriptionScheduleCurrentPhase is a structure the current phase's period.
type SubscriptionScheduleCurrentPhase struct {
	EndDate   int64 `json:"end_date"`
	StartDate int64 `json:"start_date"`
}
type SubscriptionScheduleDefaultSettingsAutomaticTax struct {
	Enabled bool `json:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionScheduleDefaultSettingsBillingThresholds struct {
	AmountGte               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}

// The subscription schedule's default invoice settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type SubscriptionScheduleDefaultSettingsTransferData struct {
	AmountPercent float64  `json:"amount_percent"`
	Destination   *Account `json:"destination"`
}

// SubscriptionScheduleDefaultSettings is a structure representing the
// subscription schedule’s default settings.
type SubscriptionScheduleDefaultSettings struct {
	ApplicationFeePercent float64                                               `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionScheduleDefaultSettingsAutomaticTax      `json:"automatic_tax"`
	BillingCycleAnchor    SubscriptionScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionScheduleDefaultSettingsBillingThresholds `json:"billing_thresholds"`
	CollectionMethod      SubscriptionScheduleDefaultSettingsCollectionMethod   `json:"collection_method"`
	DefaultPaymentMethod  *PaymentMethod                                        `json:"default_payment_method"`
	InvoiceSettings       *SubscriptionScheduleDefaultSettingsInvoiceSettings   `json:"invoice_settings"`
	TransferData          *SubscriptionScheduleDefaultSettingsTransferData      `json:"transfer_data"`
}

// SubscriptionSchedulePhaseAddInvoiceItem represents the invoice items to add when the phase starts.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	Price    *Price     `json:"price"`
	Quantity int64      `json:"quantity"`
	TaxRates []*TaxRate `json:"tax_rates"`
}
type SubscriptionSchedulePhaseAutomaticTax struct {
	Enabled bool `json:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionSchedulePhaseBillingThresholds struct {
	AmountGte               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}

// The invoice settings applicable during this phase.
type SubscriptionSchedulePhaseInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type SubscriptionSchedulePhaseItemBillingThresholds struct {
	UsageGTE int64 `json:"usage_gte"`
}

// SubscriptionSchedulePhaseItem represents plan details for a given phase
type SubscriptionSchedulePhaseItem struct {
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholds `json:"billing_thresholds"`
	Plan              *Plan                                           `json:"plan"`
	Price             *Price                                          `json:"price"`
	Quantity          int64                                           `json:"quantity"`
	TaxRates          []*TaxRate                                      `json:"tax_rates"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type SubscriptionSchedulePhaseTransferData struct {
	AmountPercent float64  `json:"amount_percent"`
	Destination   *Account `json:"destination"`
}

// SubscriptionSchedulePhase is a structure a phase of a subscription schedule.
type SubscriptionSchedulePhase struct {
	AddInvoiceItems       []*SubscriptionSchedulePhaseAddInvoiceItem  `json:"add_invoice_items"`
	ApplicationFeePercent float64                                     `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionSchedulePhaseAutomaticTax      `json:"automatic_tax"`
	BillingCycleAnchor    SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionSchedulePhaseBillingThresholds `json:"billing_thresholds"`
	CollectionMethod      SubscriptionSchedulePhaseCollectionMethod   `json:"collection_method"`
	Coupon                *Coupon                                     `json:"coupon"`
	DefaultPaymentMethod  *PaymentMethod                              `json:"default_payment_method"`
	DefaultTaxRates       []*TaxRate                                  `json:"default_tax_rates"`
	EndDate               int64                                       `json:"end_date"`
	InvoiceSettings       *SubscriptionSchedulePhaseInvoiceSettings   `json:"invoice_settings"`
	Items                 []*SubscriptionSchedulePhaseItem            `json:"items"`
	ProrationBehavior     SubscriptionSchedulePhaseProrationBehavior  `json:"proration_behavior"`
	StartDate             int64                                       `json:"start_date"`
	TransferData          *SubscriptionSchedulePhaseTransferData      `json:"transfer_data"`
	TrialEnd              int64                                       `json:"trial_end"`
}

// SubscriptionSchedule is the resource representing a Stripe subscription schedule.
type SubscriptionSchedule struct {
	APIResource
	CanceledAt           int64                                `json:"canceled_at"`
	CompletedAt          int64                                `json:"completed_at"`
	Created              int64                                `json:"created"`
	CurrentPhase         *SubscriptionScheduleCurrentPhase    `json:"current_phase"`
	Customer             *Customer                            `json:"customer"`
	DefaultSettings      *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	EndBehavior          SubscriptionScheduleEndBehavior      `json:"end_behavior"`
	ID                   string                               `json:"id"`
	Livemode             bool                                 `json:"livemode"`
	Metadata             map[string]string                    `json:"metadata"`
	Object               string                               `json:"object"`
	Phases               []*SubscriptionSchedulePhase         `json:"phases"`
	ReleasedAt           int64                                `json:"released_at"`
	ReleasedSubscription string                               `json:"released_subscription"`
	Status               SubscriptionScheduleStatus           `json:"status"`
	Subscription         *Subscription                        `json:"subscription"`
}

// SubscriptionScheduleList is a list object for subscription schedules.
type SubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*SubscriptionSchedule `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionSchedule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionSchedule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscriptionSchedule SubscriptionSchedule
	var v subscriptionSchedule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionSchedule(v)
	return nil
}
