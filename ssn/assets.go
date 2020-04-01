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
	Flags        Flags    `json:"flags"`
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
	Asset_issuer  string   `json:"asset_issuer,omitempty"`
	Assets_issued []Assets `json:"assets_issued,omitempty"`
	Title         string   `json:"title,omitempty"`
	Status        int      `json:"status,omitempty"`
	Detail        string   `json:"detail,omitempty"`
}

// Assets describeds the JSON structure related to an individual asset
type Assets struct {
	Asset_code   string       `json:"asset_code,omitempty"`
	Asset_type   string       `json:"asset_type,omitempty"`
	Flags        Flags        `json:"flags,omitempty"`
	Amount       string       `json:"amount,omitempty"`
	Num_accounts int          `json:"num_accounts,omitempty"`
	Trustlines   []Trustlines `json:"trustlines,omitempty"`
}

// Trustlines describes the JSON structure related to an individual trustline
type Trustlines struct {
	Account_id        string   `json:"account_id,omitempty"`
	Balance           Balances `json:"balance,omitempty"`
	Service_name      string   `json:"service_name,omitempty"`
	Registration_name string   `json:"registration_name,omitempty"`
	Home_domain       string   `json:"home_domain,omitempty"`
}
