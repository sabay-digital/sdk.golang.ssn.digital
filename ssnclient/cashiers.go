package ssnclient

/*
*
* Cashier/Payment Provider Specific
*
 */

// StatusResponse is the root level JSON object for the cashier endpoint /v1/status
type StatusResponse struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
}

// InfoResponse is the root level JSON object for the cashier endpoint /v1/info
type InfoResponse struct {
	Assets_issued    []Asset `json:"assets_issued"`
	Payment_provider string  `json:"payment_provider"`
	Payment_type     string  `json:"payment_type"`
	Authorization    string  `json:"authorization"`
}

// Asset is the JSON object that describes assets issued by a cashier
type Asset struct {
	Asset_code   string `json:"asset_code"`
	Limit_amount int    `json:"limit_amount,omitempty"`
	Asset_Issuer string `json:"asset_issuer"`
}
