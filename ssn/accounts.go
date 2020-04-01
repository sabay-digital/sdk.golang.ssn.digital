package ssn

import "encoding/json"

// Account describes the JSON structure related to an individual account
type Account struct {
	Id                   string                     `json:"id,omitempty"`
	Account_ID           string                     `json:"account_id,omitempty"`
	Sequence             string                     `json:"sequence,omitempty"`
	Subentry_count       int                        `json:"subentry_count"`
	Home_domain          string                     `json:"home_domain,omitempty"`
	Last_modified_ledger int                        `json:"last_modified_ledger"`
	Thresholds           *Thresholds                `json:"thresholds,omitempty"`
	Flags                *Flags                     `json:"flags,omitempty"`
	Balances             []Balances                 `json:"balances,omitempty"`
	Signers              []Signers                  `json:"signers,omitempty"`
	Data                 *Data                      `json:"data,omitempty"`
	Type                 string                     `json:"type,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Status               int                        `json:"status,omitempty"`
	Detail               string                     `json:"detail,omitempty"`
	Extras               map[string]json.RawMessage `json:"extras,omitempty"`
}

// Thresholds describes the JSON structure of an accounts threshold values
type Thresholds struct {
	Low_threshold  int `json:"low_threshold"`
	Med_threshold  int `json:"med_threshold"`
	High_threshold int `json:"high_threshold"`
}

// Flags describes the JSON structure of an accounts flag values
type Flags struct {
	Auth_required  bool `json:"auth_required"`
	Auth_revocable bool `json:"auth_revocable"`
	Auth_immutable bool `json:"auth_immutable"`
}

// Balances describes the JSON structure of an individual balance on an account
type Balances struct {
	Balance                        string `json:"balance,omitempty"`
	Buying_liabilities             string `json:"buying_liabilities,omitempty"`
	Selling_liabilities            string `json:"selling_liabilities,omitempty"`
	Limit                          string `json:"limit,omitempty"`
	Last_modified_ledger           int    `json:"last_modified_ledger"`
	Is_authorized                  bool   `json:"is_authorized"`
	Asset_type                     string `json:"asset_type,omitempty"`
	Asset_code                     string `json:"asset_code,omitempty"`
	Asset_issuer                   string `json:"asset_issuer,omitempty"`
	Asset_issuer_service_name      string `json:"asset_issuer_service_name,omitempty"`
	Asset_issuer_registration_name string `json:"asset_issuer_registration_name,omitempty"`
	Asset_issuer_home_domain       string `json:"asset_issuer_home_domain,omitempty"`
	Asset_issuer_net_payout        string `json:"asset_issuer_net_payout,omitempty"`
}

// Signers describes the JSON structure of an individual signer on an account
type Signers struct {
	Weight int    `json:"weight"`
	Key    string `json:"key,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Data describes the JSON structure of all possible KYC fields on an account
type Data struct {
	// Shared
	Registration_name            string `json:"registration_name,omitempty"`
	Registration_country         string `json:"registration_country,omitempty"`
	Registration_authority       string `json:"registration_authority,omitempty"`
	Registration_no              string `json:"registration_no,omitempty"`
	Contact_name                 string `json:"contact_name,omitempty"`
	Contact_mobile_number        string `json:"contact_mobile_number,omitempty"`
	Document_registration        string `json:"document_registration,omitempty"`
	Document_patent              string `json:"document_patent,omitempty"`
	Document_tax                 string `json:"document_tax,omitempty"`
	Document_registration_url    string `json:"document_registration_url,omitempty"`
	Document_registration_sha256 string `json:"document_registration_sha256,omitempty"`
	Document_patent_url          string `json:"document_patent_url,omitempty"`
	Document_patent_sha256       string `json:"document_patent_sha256,omitempty"`
	Document_tax_url             string `json:"document_tax_url,omitempty"`
	Document_tax_sha256          string `json:"document_tax_sha256,omitempty"`
	// Payment Provider Specific
	Net_payout                  string `json:"net_payout,omitempty"`
	Trust_approval_email        string `json:"trust_approval_email,omitempty"`
	Nbc_license_holder          string `json:"nbc_license_holder,omitempty"`
	Document_nbc_license        string `json:"document_nbc_license,omitempty"`
	Document_nbc_license_url    string `json:"document_nbc_license_url,omitempty"`
	Document_nbc_license_sha256 string `json:"document_nbc_license_sha256,omitempty"`
	Settlement_bank_bic         string `json:"settlement_bank_bic,omitempty"`
	Settlement_bank_code        string `json:"settlement_bank_code,omitempty"`
	Settlement_bank_name        string `json:"settlement_bank_name,omitempty"`
	Settlement_ssn_account      string `json:"settlement_ssn_account,omitempty"`
	// Merchant Specific
	Service_name          string `json:"service_name,omitempty"`
	Payout_spread         string `json:"payout_spread,omitempty"`
	Contact_email_address string `json:"contact_email_address,omitempty"`
	Contact_street        string `json:"contact_street,omitempty"`
	Contact_city          string `json:"contact_city,omitempty"`
	Contact_province      string `json:"contact_province,omitempty"`
	Contact_postal_code   string `json:"contact_postal_code,omitempty"`
	Contact_country       string `json:"contact_country,omitempty"`
}
