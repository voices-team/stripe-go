//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"
import "github.com/stripe/stripe-go/v72/form"

type SubscriptionScheduleDefaultSettingsBillingCycleAnchor string

const (
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic  SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

type SubscriptionScheduleDefaultSettingsCollectionMethod string

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

type SubscriptionSchedulePhaseCollectionMethod string

const (
	SubscriptionSchedulePhaseCollectionMethodChargeAutomatically SubscriptionSchedulePhaseCollectionMethod = "charge_automatically"
	SubscriptionSchedulePhaseCollectionMethodSendInvoice         SubscriptionSchedulePhaseCollectionMethod = "send_invoice"
)

type SubscriptionSchedulePhaseProrationBehavior string

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

type SubscriptionScheduleDefaultSettingsAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}
type SubscriptionScheduleDefaultSettingsBillingThresholdsParams struct {
	AmountGte               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}
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
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}
type SubscriptionSchedulePhaseBillingThresholdsParams struct {
	AmountGte               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}
type SubscriptionSchedulePhaseInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}
type SubscriptionSchedulePhaseItemBillingThresholdsParams struct {
	UsageGTE *int64 `form:"usage_gte"`
}
type SubscriptionSchedulePhaseItemPriceDataRecurringParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}
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
type SubscriptionScheduleDefaultSettingsBillingThresholds struct {
	AmountGte               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}
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
type SubscriptionSchedulePhaseBillingThresholds struct {
	AmountGte               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}
type SubscriptionSchedulePhaseInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}
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
