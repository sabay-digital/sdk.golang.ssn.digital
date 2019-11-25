package ssn

import "encoding/json"

// Account describes the JSON structure related to an individual account
type Account struct {
	Id             string                     `json:"id,omitempty"`
	Account_ID     string                     `json:"account_id,omitempty"`
	Sequence       string                     `json:"sequence,omitempty"`
	Subentry_count int                        `json:"subentry_count,omitempty"`
	Thresholds     *thresholds                `json:"thresholds,omitempty"`
	Flags          *flags                     `json:"flags,omitempty"`
	Balances       []balances                 `json:"balances,omitempty"`
	Signers        []signers                  `json:"signers,omitempty"`
	Type           string                     `json:"type,omitempty"`
	Title          string                     `json:"title,omitempty"`
	Status         int                        `json:"status,omitempty"`
	Detail         string                     `json:"detail,omitempty"`
	Extras         map[string]json.RawMessage `json:"extras,omitempty"`
}

// thresholds describes the JSON structure of an accounts threshold values
type thresholds struct {
	Low_threshold  int `json:"low_threshold"`
	Med_threshold  int `json:"med_threshold"`
	High_threshold int `json:"high_threshold"`
}

// flags describes the JSON structure of an accounts flag values
type flags struct {
	Auth_required  bool `json:"auth_required"`
	Auth_revocable bool `json:"auth_revocable"`
	Auth_immutable bool `json:"auth_immutable"`
}

// balances describes the JSON structure of an individual balance on an account
type balances struct {
	Balance              string `json:"balance,omitempty"`
	Buying_liabilities   string `json:"buying_liabilities,omitempty"`
	Selling_liabilities  string `json:"selling_liabilities,omitempty"`
	Limit                string `json:"limit,omitempty"`
	Last_modified_ledger int    `json:"last_modified_ledger,omitempty"`
	Is_authorized        bool   `json:"is_authorized,omitempty"`
	Asset_type           string `json:"asset_type,omitempty"`
	Asset_code           string `json:"asset_code,omitempty"`
	Asset_issuer         string `json:"asset_issuer,omitempty"`
}

// signers describes the JSON structure of an individual signer on an account
type signers struct {
	Weight int    `json:"weight,omitempty"`
	Key    string `json:"key,omitempty"`
	Type   string `json:"type,omitempty"`
}
