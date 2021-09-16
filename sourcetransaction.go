//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of source this transaction is attached to.
type SourceTransactionType string

// List of values that SourceTransactionType can take
const (
	SourceTransactionTypeAchCreditTransfer SourceTransactionType = "ach_credit_transfer"
	SourceTransactionTypeAchDebit          SourceTransactionType = "ach_debit"
	SourceTransactionTypeAlipay            SourceTransactionType = "alipay"
	SourceTransactionTypeBancontact        SourceTransactionType = "bancontact"
	SourceTransactionTypeCard              SourceTransactionType = "card"
	SourceTransactionTypeCardPresent       SourceTransactionType = "card_present"
	SourceTransactionTypeEPS               SourceTransactionType = "eps"
	SourceTransactionTypeGiropay           SourceTransactionType = "giropay"
	SourceTransactionTypeIdeal             SourceTransactionType = "ideal"
	SourceTransactionTypeKlarna            SourceTransactionType = "klarna"
	SourceTransactionTypeMultibanco        SourceTransactionType = "multibanco"
	SourceTransactionTypeP24               SourceTransactionType = "p24"
	SourceTransactionTypeSepaDebit         SourceTransactionType = "sepa_debit"
	SourceTransactionTypeSofort            SourceTransactionType = "sofort"
	SourceTransactionTypeThreeDSecure      SourceTransactionType = "three_d_secure"
	SourceTransactionTypeWechat            SourceTransactionType = "wechat"
)

type SourceTransactionAchCreditTransfer struct {
	CustomerData  string `json:"customer_data"`
	Fingerprint   string `json:"fingerprint"`
	Last4         string `json:"last4"`
	RoutingNumber string `json:"routing_number"`
}
type SourceTransactionChfCreditTransfer struct {
	Reference            string `json:"reference"`
	SenderAddressCountry string `json:"sender_address_country"`
	SenderAddressLine1   string `json:"sender_address_line1"`
	SenderIban           string `json:"sender_iban"`
	SenderName           string `json:"sender_name"`
}
type SourceTransactionGbpCreditTransfer struct {
	Fingerprint         string `json:"fingerprint"`
	FundingMethod       string `json:"funding_method"`
	Last4               string `json:"last4"`
	Reference           string `json:"reference"`
	SenderAccountNumber string `json:"sender_account_number"`
	SenderName          string `json:"sender_name"`
	SenderSortCode      string `json:"sender_sort_code"`
}
type SourceTransactionPaperCheck struct {
	AvailableAt string `json:"available_at"`
	Invoices    string `json:"invoices"`
}
type SourceTransactionSepaCreditTransfer struct {
	Reference  string `json:"reference"`
	SenderIban string `json:"sender_iban"`
	SenderName string `json:"sender_name"`
}

// SourceTransaction is the resource representing a Stripe source transaction.
type SourceTransaction struct {
	AchCreditTransfer  *SourceTransactionAchCreditTransfer  `json:"ach_credit_transfer"`
	Amount             int64                                `json:"amount"`
	ChfCreditTransfer  *SourceTransactionChfCreditTransfer  `json:"chf_credit_transfer"`
	Created            int64                                `json:"created"`
	Currency           Currency                             `json:"currency"`
	GbpCreditTransfer  *SourceTransactionGbpCreditTransfer  `json:"gbp_credit_transfer"`
	ID                 string                               `json:"id"`
	Livemode           bool                                 `json:"livemode"`
	Object             string                               `json:"object"`
	PaperCheck         *SourceTransactionPaperCheck         `json:"paper_check"`
	SepaCreditTransfer *SourceTransactionSepaCreditTransfer `json:"sepa_credit_transfer"`
	Source             string                               `json:"source"`
	Status             string                               `json:"status"`
	Type               SourceTransactionType                `json:"type"`
}

// SourceTransactionList is a list object for SourceTransactions.
type SourceTransactionList struct {
	APIResource
	ListMeta
	Data []*SourceTransaction `json:"data"`
}
