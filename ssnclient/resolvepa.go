package ssnclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
)

/*
*
* Resolve Payment Addresses - V2
*
 */

// ResolverRequestDepreciated describes the JSON structure for making a request to the payment address resolver V1 API.
// This method is now depreciated
type ResolverRequestDepreciated struct {
	Asset_issuer    string `json:"asset_issuer,omitempty"`
	Public_key      string `json:"public_key,omitempty"`
	Payment_address string `json:"payment_address"`
}

// ResolverRequest describes the JSON structure for making a request to the payment address resolver V2 API.
type ResolverRequest struct {
	Hash        string `json:"hash"`
	Signature   string `json:"signature"`
	Signer      string `json:"signer"`
	Ssn_account string `json:"ssn_account"`
}

// ResolverResponse describes the JSON structure for the response from the payment address resolver API
type ResolverResponse struct {
	Network_address string                   `json:"network_address,omitempty"`
	Public_key      string                   `json:"public_key,omitempty"` // Depreciated
	Asset_code      string                   `json:"asset_code,omitempty"` // Depreciated
	Payment_type    string                   `json:"payment_type,omitempty"`
	Service_name    string                   `json:"service_name,omitempty"`
	Encrypted       string                   `json:"encrypted,omitempty"` // Depreciated
	Details         *ResolverResponseDetails `json:"details,omitempty"`
	Status          int                      `json:"status,omitempty"`
	Title           string                   `json:"title,omitempty"`
	Signature       string                   `json:"signature,omitempty"`
}

// ResolverResponseDetails describes the JSON structure for the nested details part of the response from the payment address resolver API
type ResolverResponseDetails struct {
	Payment_info                string                               `json:"payment_info,omitempty"`
	Memo                        string                               `json:"memo,omitempty"`
	Recurring_payment_frequency string                               `json:"recurring_payment_frequency,omitempty"` // Depreciated
	Recurring_payment_interval  string                               `json:"recurring_payment_interval,omitempty"`  // Depreciated
	Recurring_payment_start     string                               `json:"recurring_payment_start,omitempty"`     // Depreciated
	Payment                     map[string]ResolverPaymentDetails    `json:"payment,omitempty"`
	Service_fee                 map[string]ResolverServiceFeeDetails `json:"service_fee,omitempty"`
}

// ResolverPaymentDetails describes the JSON structure for the nested payment details part of the response from the payment address resolver API
type ResolverPaymentDetails struct {
	Amount     string `json:"amount,omitempty"`
	Asset_code string `json:"asset_code,omitempty"`
}

// ResolverServiceFeeDetails describes the JSON structure for the nested service fee details part of the response from the payment address resolver API
type ResolverServiceFeeDetails struct {
	Amount     string `json:"amount,omitempty"`
	Asset_code string `json:"asset_code,omitempty"`
}

// ResolvePA sends a payment address to the PA service for resolving. The resolverURL should be a resolver that supports the V2 design
func ResolvePA(paymentAddress, hash, signature, signer, ssnAcc, resolverURL string) (ResolverResponse, error) {
	// Prepare JSON request
	req := ResolverRequest{
		Hash:        hash,
		Signature:   signature,
		Signer:      signer,
		Ssn_account: ssnAcc,
	}

	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "ResolvePA: Marshal request body") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}

	// Send the request to the API and get the reponse
	paReq, err := http.NewRequest("POST", resolverURL+"/resolve/"+paymentAddress, bytes.NewBuffer(reqBody))
	if ssn.Log(err, "ResolvePA: Build HTTP request") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}
	paReq.Header.Add("Content-Type", "application/json")

	paResp, err := http.DefaultClient.Do(paReq)
	if ssn.Log(err, "ResolvePA: Send HTTP request") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}
	defer paResp.Body.Close()

	// Unmarshall the PA response
	body, err := ioutil.ReadAll(paResp.Body)
	if ssn.Log(err, "ResolvePA: Read response body") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}

	resp := ResolverResponse{}
	err = json.Unmarshal(body, &resp)
	if ssn.Log(err, "ResolvePA: Unmarshal response body") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}

	if resp.Network_address == "" {
		return ResolverResponse{Status: 400, Title: "Payment destination missing"}, nil
	}

	return resp, nil
}
