//
//
// File generated from our OpenAPI spec
//
//

package stripe

type LineItemListParams struct {
	ListParams `form:"*"`
	Session    *string `form:"-"` // Included in URL
}

// LineItemDiscount represent the details of one discount applied to a line item.
type LineItemDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// LineItemTax represent the details of one tax rate applied to a line item.
type LineItemTax struct {
	Amount int64    `json:"amount"`
	Rate   *TaxRate `json:"rate"`
}

// LineItem is the resource representing a line item.
type LineItem struct {
	AmountSubtotal int64               `json:"amount_subtotal"`
	AmountTotal    int64               `json:"amount_total"`
	Currency       Currency            `json:"currency"`
	Description    string              `json:"description"`
	Discounts      []*LineItemDiscount `json:"discounts"`
	ID             string              `json:"id"`
	Object         string              `json:"object"`
	Price          *Price              `json:"price"`
	Quantity       int64               `json:"quantity"`
	Taxes          []*LineItemTax      `json:"taxes"`
}

// LineItemList is a list of prices as returned from a list endpoint.
type LineItemList struct {
	APIResource
	ListMeta
	Data []*LineItem `json:"data"`
}
