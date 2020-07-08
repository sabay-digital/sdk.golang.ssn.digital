package ssn

// Offer describes the JSON structure related to an individual offer
type Offer struct {
	Id                   string     `json:"id,omitempty"`
	Paging_token         string     `json:"paging_token,omitempty"`
	Seller               string     `json:"seller,omitempty"`
	Selling              OfferAsset `json:"selling,omitempty"`
	Buying               OfferAsset `json:"buying,omitempty"`
	Amount               string     `json:"amount,omitempty"`
	Price_r              OfferPrice `json:"price_r,omitempty"`
	Last_modified_ledger int        `json:"last_modified_ledger,omitempty"`
	Last_modified_time   string     `json:"last_modified_time,omitempty"`
}

// OfferAsset describes the JSON structure related to an offers assets
type OfferAsset struct {
	Asset_type   string `json:"asset_type,omitempty"`
	Asset_code   string `json:"asset_code,omitempty"`
	Asset_issuer string `json:"asset_issuer,omitempty"`
}

// OfferPrice describes the JSON structure related to an offers price
type OfferPrice struct {
	N int `json:"n,omitempty"`
	D int `json:"d,omitempty"`
}
