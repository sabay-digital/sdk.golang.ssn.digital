package ssnclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.sabay.com/payment-network/sdk/sdk.golang.ssn.digital/ssn"
)

/*
*
* Create and Stream Payments
*
 */

// CreatePaymentRequest describes the JSON structure for making a request to the create payment API
type CreatePaymentRequest struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	Asset_code   string `json:"asset_code"`
	Asset_issuer string `json:"asset_issuer"`
	Memo         string `json:"memo"`
}

// CreatePaymentResponse describes the JSON structure for the response from the create payment API
type CreatePaymentResponse struct {
	Status       int    `json:"status,omitempty"`
	Envelope_xdr string `json:"envelope_xdr,omitempty"`
	Title        string `json:"title,omitempty"`
}

// CreatePayment sends transaction information to the SSN API to build an XDR envelope
func CreatePayment(from, to, amount, assetCode, assetIssuer, memo, api string) (string, error) {
	// Prepare JSON request
	req := CreatePaymentRequest{
		From:         from,
		To:           to,
		Amount:       amount,
		Asset_code:   assetCode,
		Asset_issuer: assetIssuer,
		Memo:         memo,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "CreatePayment: Marshal request body") {
		return "", err
	}

	// Send the request to the API and get the reponse
	cpReq, err := http.NewRequest("POST", api+"/create/transaction", bytes.NewBuffer(reqBody))
	if ssn.Log(err, "CreatePayment: Build HTTP request") {
		return "", err
	}
	cpReq.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(cpReq)
	if ssn.Log(err, "CreatePayment: Send HTTP request") {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "CreatePayment: Read response body") {
		return "", err
	}

	// Return the transaction envelope for signing
	cpResp := CreatePaymentResponse{}
	err = json.Unmarshal(body, &cpResp)
	if ssn.Log(err, "CreatePayment: Unmarshal response body") {
		return "", err
	}
	return cpResp.Envelope_xdr, nil
}
