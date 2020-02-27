package ssnclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
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
	defer vsResp.Body.Close()

	// 200 signifies the signature is valid
	if vsResp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

// VerifySignerRequest describes the JSON structure for making a request to the verify signer API
type VerifySignerRequest struct {
	Signer      string `json:"signer"`
	Ssn_account string `json:"ssn_account"`
}

// VerifySignerResponse describes the JSON structure for the response from the verify signer API
type VerifySignerResponse struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// VerifySigner checks whether the provided signer is a signer on an SSN account
func VerifySigner(signer, ssnAccount, api string) (bool, error) {
	// Prepare JSON request
	req := VerifySignerRequest{
		Signer:      signer,
		Ssn_account: ssnAccount,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "VerifySignature: Marshal request body") {
		return false, err
	}

	// Send the request to the API and get the response
	vsReq, err := http.NewRequest("POST", api+"/verify/signer", bytes.NewBuffer(reqBody))
	if ssn.Log(err, "VerifySignature: build HTTP request") {
		return false, err
	}
	vsReq.Header.Add("Content-Type", "application/json")

	vsResp, err := http.DefaultClient.Do(vsReq)
	if ssn.Log(err, "VerifySignature: Send HTTP request") {
		return false, err
	}
	defer vsResp.Body.Close()

	// 200 signifies the signature is valid
	if vsResp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
