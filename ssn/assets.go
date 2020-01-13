package ssn

// Asset describes the JSON structure related to an individual asset as defined by Horizon
// from the endpoint /assets
type Asset struct {
	Links        tomlLink `json:"_links"`
	Asset_type   string   `json:"asset_type"`
	Asset_code   string   `json:"asset_code"`
	Asset_issuer string   `json:"asset_issuer"`
	Paging_token string   `json:"paging_token"`
	Amount       string   `json:"amount"`
	Num_accounts int      `json:"num_accounts"`
	Flags        flags    `json:"flags"`
}

// tomlLink describes the JSON structure related to an embedded TOML link
type tomlLink struct {
	Toml link `json:"toml"`
}

type link struct {
	Href string `json:"href"`
}

// Issuer describes the JSON structure related to an individual issuer
type Issuer struct {
	Asset_issuer string    `json:"asset_issuer,omitempty"`
	Asset_code   string    `json:"asset_code,omitempty"`
	Asset_type   string    `json:"asset_type,omitempty"`
	Flags        flags     `json:"flags,omitempty"`
	Amount       string    `json:"amount,omitempty"`
	Num_accounts int       `json:"num_accounts,omitempty"`
	Holders      []Holders `json:"holders,omitempty"`
	Title        string    `json:"title,omitempty"`
	Status       int       `json:"status,omitempty"`
	Detail       string    `json:"detail,omitempty"`
}

// Holders describes the JSON structure related to an individual asset holder
type Holders struct {
	Account_id                string   `json:"account_id,omitempty"`
	Balance                   balances `json:"balance,omitempty"`
	Asset_holder_service_name string   `json:"asset_holder_service_name,omitempty"`
	Asset_holder_reg_name     string   `json:"asset_holder_reg_name,omitempty"`
	Asset_holder_home_domain  string   `json:"asset_holder_home_domain,omitempty"`
}
