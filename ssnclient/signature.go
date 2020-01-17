package ssnclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"git.sabay.com/payment-network/sdk/sdk.golang.ssn.digital/ssn"
)

/*
*
* Create and Verify Signatures
*
 */

// VerifySignatureRequest describes the JSON structure for making a request to the verify signature API
type VerifySignatureRequest struct {
	Public_key string `json:"public_key"`
	Signature  string `json:"signature"`
	Message    string `json:"message"`
}

// VerifySignatureResponse describes the JSON structure for the response from the verify signature API
type VerifySignatureResponse struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// VerifySignature checks whether the provided message, signature and public key are valid
func VerifySignature(message, signature, publicKey, api string) (bool, error) {
	// Prepare JSON request
	req := VerifySignatureRequest{
		Public_key: publicKey,
		Signature:  signature,
		Message:    message,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "VerifySignature: Marshal request body") {
		return false, err
	}

	// Send the request to the API and get the response
	vsReq, err := http.NewRequest("POST", api+"/verify/signature", bytes.NewBuffer(reqBody))
	if ssn.Log(err, "VerifySignature: build HTTP request") {
		return false, err
	}
	vsReq.Header.Add("Content-Type", "application/json")

	vsResp, err := http.DefaultClient.Do(vsReq)
	if ssn.Log(err, "VerifySignature: Send HTTP request") {
		return false, err
	}

	// 200 signifies the signature is valid
	if vsResp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
