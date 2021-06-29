//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// SKUParams is the set of parameters allowed on SKU creation or update.
type SKUParams struct {
	Params            `form:"*"`
	Active            *bool                       `form:"active"`
	Attributes        map[string]string           `form:"attributes"`
	Currency          *string                     `form:"currency"`
	ID                *string                     `form:"id"`
	Image             *string                     `form:"image"`
	Inventory         *SKUInventoryParams         `form:"inventory"`
	PackageDimensions *SKUPackageDimensionsParams `form:"package_dimensions"`
	Price             *int64                      `form:"price"`
	Product           *string                     `form:"product"`
}
type SKUInventoryParams struct {
	Quantity *int64  `form:"quantity"`
	Type     *string `form:"type"`
	Value    *string `form:"value"`
}
type SKUPackageDimensionsParams struct {
	Height *float64 `form:"height"`
	Length *float64 `form:"length"`
	Weight *float64 `form:"weight"`
	Width  *float64 `form:"width"`
}

// SKUListParams is the set of parameters that can be used when listing SKUs.
type SKUListParams struct {
	ListParams `form:"*"`
	Active     *bool             `form:"active"`
	Attributes map[string]string `form:"attributes"`
	IDs        []*string         `form:"ids"`
	InStock    *bool             `form:"in_stock"`
	Product    *string           `form:"product"`
}
type SKUInventory struct {
	Quantity int64  `json:"quantity"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}
type SKUPackageDimensions struct {
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Weight float64 `json:"weight"`
	Width  float64 `json:"width"`
}

// SKU is the resource representing a SKU.
// For more details see https://stripe.com/docs/api#skus.
type SKU struct {
	APIResource
	Active            bool                  `json:"active"`
	Attributes        map[string]string     `json:"attributes"`
	Created           int64                 `json:"created"`
	Currency          Currency              `json:"currency"`
	Deleted           bool                  `json:"deleted"`
	ID                string                `json:"id"`
	Image             string                `json:"image"`
	Inventory         *SKUInventory         `json:"inventory"`
	Livemode          bool                  `json:"livemode"`
	Metadata          map[string]string     `json:"metadata"`
	Object            string                `json:"object"`
	PackageDimensions *SKUPackageDimensions `json:"package_dimensions"`
	Price             int64                 `json:"price"`
	Product           *Product              `json:"product"`
	Updated           int64                 `json:"updated"`
}

// SKUList is a list of SKUs as returned from a list endpoint.
type SKUList struct {
	APIResource
	ListMeta
	Data []*SKU `json:"data"`
}

// UnmarshalJSON handles deserialization of a SKU.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SKU) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type sKU SKU
	var v sKU
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SKU(v)
	return nil
}
