package ssnclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"git.sabay.com/payment-network/sdk/sdk.golang.ssn.digital/ssn"
)

/*
*
* Add, Approve, Remove and Verify Trustlines
*
 */

// VerifyTrustRequest describes the JSON structure for making a request to verify trust API
type VerifyTrustRequest struct {
	Account      string `json:"account"`
	Asset_code   string `json:"asset_code"`
	Asset_issuer string `json:"asset_issuer"`
}

// VerifyTrust checks whether the provided asset and assetIssuer is trusted by destination
func VerifyTrust(destination, asset, assetIssuer, api string) (bool, error) {
	// Prepare JSON request
	req := VerifyTrustRequest{
		Account:      destination,
		Asset_code:   asset,
		Asset_issuer: assetIssuer,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "VerifyTrust: Marshal request body") {
		return false, err
	}

	// Send the request to the API and get the reponse
	vtReq, err := http.NewRequest("POST", api+"/verify/trust", bytes.NewBuffer(reqBody))
	if ssn.Log(err, "VerifyTrust: Build HTTP request") {
		return false, err
	}
	vtResp, err := http.DefaultClient.Do(vtReq)
	if ssn.Log(err, "VerifyTrust: Send HTTP request") {
		return false, err
	}

	// 200 signifies the cashier is trusted
	if vtResp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
