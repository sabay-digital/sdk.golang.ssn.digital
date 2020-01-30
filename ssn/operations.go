package ssn

import "encoding/json"

// Operation describes the JSON structure related to an individual operation
type Operation struct {
	Id                     string                     `json:"id,omitempty"`
	Paging_token           string                     `json:"paging_token,omitempty"`
	Transaction_successful bool                       `json:"transaction_successful,omitempty"`
	Type                   string                     `json:"type,omitempty"`
	Type_i                 int                        `json:"type_i,omitempty"`
	Account                string                     `json:"account,omitempty"`
	Funder                 string                     `json:"funder,omitempty"`
	Starting_balance       string                     `json:"starting_balance,omitempty"`
	From                   string                     `json:"from,omitempty"`
	To                     string                     `json:"to,omitempty"`
	Asset_type             string                     `json:"asset_type,omitempty"`
	Asset_code             string                     `json:"asset_code,omitempty"`
	Asset_issuer           string                     `json:"asset_issuer,omitempty"`
	Amount                 string                     `json:"amount,omitempty"`
	Source_asset_code      string                     `json:"source_asset_code,omitempty"`
	Source_asset_issuer    string                     `json:"source_asset_issuer,omitempty"`
	Source_asset_type      string                     `json:"source_asset_type,omitempty"`
	Source_max             string                     `json:"source_max,omitempty"`
	Source_amount          string                     `json:"source_amount,omitempty"`
	Offer_id               int                        `json:"offer_id,omitempty"`
	Buying_asset_code      string                     `json:"buying_asset_code,omitempty"`
	Buying_asset_issuer    string                     `json:"buying_asset_issuer,omitempty"`
	Buying_asset_type      string                     `json:"buying_asset_type,omitempty"`
	Price                  string                     `json:"price,omitempty"`
	Price_r                *offerPrice                `json:"price_r,omitempty"`
	Selling_asset_code     string                     `json:"selling_asset_code,omitempty"`
	Selling_asset_issuer   string                     `json:"selling_asset_issuer,omitempty"`
	Selling_asset_type     string                     `json:"selling_asset_type,omitempty"`
	Buy_amount             string                     `json:"buy_amount,omitempty"`
	Signer_key             string                     `json:"signer_key,omitempty"`
	Signer_weight          int                        `json:"signer_weight,omitempty"`
	Master_key_weight      int                        `json:"master_key_weight,omitempty"`
	Low_threshold          int                        `json:"low_threshold,omitempty"`
	Med_threshold          int                        `json:"med_threshold,omitempty"`
	High_threshold         int                        `json:"high_threshold,omitempty"`
	Home_domain            string                     `json:"home_domain,omitempty"`
	Set_flags              []int                      `json:"set_flags,omitempty"`
	Set_flags_s            []string                   `json:"set_flags_s,omitempty"`
	Clear_flags            []int                      `json:"clear_flags,omitempty"`
	Clear_flags_s          []string                   `json:"clear_flags_s,omitempty"`
	Trustee                string                     `json:"trustee,omitempty"`
	Trustor                string                     `json:"trustor,omitempty"`
	Limit                  string                     `json:"limit,omitempty"`
	Authorize              bool                       `json:"authorize,omitempty"`
	Into                   string                     `json:"into,omitempty"`
	Name                   string                     `json:"name,omitempty"`
	Value                  string                     `json:"value,omitempty"`
	BumpTo                 string                     `json:"bumpTo,omitempty"`
	Title                  string                     `json:"title,omitempty"`
	Status                 int                        `json:"status,omitempty"`
	Detail                 string                     `json:"detail,omitempty"`
	Extras                 map[string]json.RawMessage `json:"extras,omitempty"`
}

type offerPrice struct {
	N int `json:"n,omitempty"`
	D int `json:"d,omitempty"`
}
