//
//
// File generated from our OpenAPI spec
//
//

package stripe

// BalanceParams is the set of parameters that can be used when retrieving a balance.
// For more details see https://stripe.com/docs/api#balance.
type BalanceParams struct {
	Params `form:"*"`
}
type BalanceAvailableSourceTypes struct {
	BankAccount int64 `json:"bank_account"`
	Card        int64 `json:"card"`
	FPX         int64 `json:"fpx"`
}

// Funds that are available to be transferred or paid out, whether automatically by Stripe or explicitly via the [Transfers API](https://stripe.com/docs/api#transfers) or [Payouts API](https://stripe.com/docs/api#payouts). The available balance for each currency and payment type can be found in the `source_types` property.
type BalanceAvailable struct {
	Amount      int64                        `json:"amount"`
	Currency    Currency                     `json:"currency"`
	SourceTypes *BalanceAvailableSourceTypes `json:"source_types"`
}
type BalanceConnectReservedSourceTypes struct {
	BankAccount int64 `json:"bank_account"`
	Card        int64 `json:"card"`
	FPX         int64 `json:"fpx"`
}

// Funds held due to negative balances on connected Custom accounts. The connect reserve balance for each currency and payment type can be found in the `source_types` property.
type BalanceConnectReserved struct {
	Amount      int64                              `json:"amount"`
	Currency    Currency                           `json:"currency"`
	SourceTypes *BalanceConnectReservedSourceTypes `json:"source_types"`
}
type BalanceInstantAvailableSourceTypes struct {
	BankAccount int64 `json:"bank_account"`
	Card        int64 `json:"card"`
	FPX         int64 `json:"fpx"`
}

// Funds that can be paid out using Instant Payouts.
type BalanceInstantAvailable struct {
	Amount      int64                               `json:"amount"`
	Currency    Currency                            `json:"currency"`
	SourceTypes *BalanceInstantAvailableSourceTypes `json:"source_types"`
}
type BalanceIssuingAvailableSourceTypes struct {
	BankAccount int64 `json:"bank_account"`
	Card        int64 `json:"card"`
	FPX         int64 `json:"fpx"`
}

// Funds that are available for use.
type BalanceIssuingAvailable struct {
	Amount      int64                               `json:"amount"`
	Currency    Currency                            `json:"currency"`
	SourceTypes *BalanceIssuingAvailableSourceTypes `json:"source_types"`
}
type BalanceIssuing struct {
	Available []*BalanceIssuingAvailable `json:"available"`
}
type BalancePendingSourceTypes struct {
	BankAccount int64 `json:"bank_account"`
	Card        int64 `json:"card"`
	FPX         int64 `json:"fpx"`
}

// Funds that are not yet available in the balance, due to the 7-day rolling pay cycle. The pending balance for each currency, and for each payment type, can be found in the `source_types` property.
type BalancePending struct {
	Amount      int64                      `json:"amount"`
	Currency    Currency                   `json:"currency"`
	SourceTypes *BalancePendingSourceTypes `json:"source_types"`
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	APIResource
	Available        []*BalanceAvailable        `json:"available"`
	ConnectReserved  []*BalanceConnectReserved  `json:"connect_reserved"`
	InstantAvailable []*BalanceInstantAvailable `json:"instant_available"`
	Issuing          *BalanceIssuing            `json:"issuing"`
	Livemode         bool                       `json:"livemode"`
	Object           string                     `json:"object"`
	Pending          []*BalancePending          `json:"pending"`
}
