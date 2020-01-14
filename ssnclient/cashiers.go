package ssnclient

/*
*
* Cashier/Payment Provider Specific
*
 */

// CashierStatusResponse is the root level JSON object for the cashier endpoint /v1/status
type CashierStatusResponse struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
}

// CashierInfoResponse is the root level JSON object for the cashier endpoint /v1/info
type CashierInfoResponse struct {
	Assets_issued    []CashierIssuedAsset `json:"assets_issued"`
	Payment_provider string               `json:"payment_provider"`
	Payment_type     string               `json:"payment_type"`
	Authorization    string               `json:"authorization"`
}

// CashierIssuedAsset is the JSON object that describes an asset issued by a cashier
type CashierIssuedAsset struct {
	Asset_code   string `json:"asset_code"`
	Limit_amount int    `json:"limit_amount,omitempty"`
	Asset_Issuer string `json:"asset_issuer"`
}
