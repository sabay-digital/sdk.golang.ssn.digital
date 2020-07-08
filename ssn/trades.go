package ssn

// Trade describes the JSON structure related to an individual trade
type Trade struct {
	Id                   string     `json:"id,omitempty"`
	Paging_token         string     `json:"paging_token,omitempty"`
	Ledger_close_time    string     `json:"ledger_close_time,omitempty"`
	Offer_id             string     `json:"offer_id,omitempty"`
	Base_offer_id        string     `json:"base_offer_id,omitempty"`
	Base_account         string     `json:"base_account,omitempty"`
	Base_amount          string     `json:"base_amount,omitempty"`
	Base_asset_type      string     `json:"base_asset_type,omitempty"`
	Base_asset_code      string     `json:"base_asset_code,omitempty"`
	Base_asset_issuer    string     `json:"base_asset_issuer,omitempty"`
	Counter_offer_id     string     `json:"counter_offer_id,omitempty"`
	Counter_account      string     `json:"counter_account,omitempty"`
	Counter_amount       string     `json:"counter_amount,omitempty"`
	Counter_asset_type   string     `json:"counter_asset_type,omitempty"`
	Counter_asset_code   string     `json:"counter_asset_code,omitempty"`
	Counter_asset_issuer string     `json:"counter_asset_issuer,omitempty"`
	Base_is_seller       bool       `json:"base_is_seller"`
	Price                OfferPrice `json:"price,omitempty"`
}
