package ssn

// Payment describes the JSON structure related to an individual payment
type Payment struct {
	Id               string `json:"id,omitempty"`
	Paging_token     string `json:"paging_token,omitempty"`
	Source_account   string `json:"source_account,omitempty"`
	Type             string `json:"type,omitempty"`
	Type_i           int    `json:"type_i,omitempty"`
	Created_at       string `json:"created_at,omitempty"`
	Transaction_hash string `json:"transaction_hash,omitempty"`
	Starting_balance string `json:"starting_balance,omitempty"`
	Funder           string `json:"funder,omitempty"`
	Account          string `json:"account,omitempty"`
	Asset_type       string `json:"asset_type,omitempty"`
	Asset_code       string `json:"asset_code,omitempty"`
	Asset_issuer     string `json:"asset_issuer,omitempty"`
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	Amount           string `json:"amount,omitempty"`
}
