package ssnclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/stellar/go/txnbuild"
)

/*
*
* Sign and Submit Transactions
*
 */

// SignTxn takes a base64 encoded XDR envelope and signs it with the provided secret key
func SignTxn(xdr, signer, networkPassphrase string) (string, error) {
	// Deserialise the provided transaction
	tx, err := txnbuild.TransactionFromXDR(xdr)
	if ssn.Log(err, "SignTxn: Build transaction from XDR") {
		return "", err
	}

	// Load the envelope
	txe, _ := tx.Transaction()

	// Add a signature
	txe, err = txe.SignWithKeyString(networkPassphrase, signer)
	if ssn.Log(err, "SignTxn: Sign transaction with key string") {
		return "", err
	}

	// Serialise the transaction
	b64, err := txe.Base64()
	if ssn.Log(err, "SignTxn: Encode transaction to base 64") {
		return "", err
	}

	return b64, nil
}

type signRequest struct {
	Envelope_xdr string `json:"envelope_xdr"`
	Access_token string `json:"access_token"`
}

// SignTxnService takes a base64 encoded XDR envelope and sends it to the specified sign service API
func SignTxnService(xdr, signer, JWTkey string) (string, error) {
	// Generate JWT
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
		Issuer:    "SSN",
		Subject:   "Sign txn request",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTkey))
	if ssn.Log(err, "SignTxnService: Generate JWT") {
		return "", err
	}

	// Prepare the request body
	sig := signRequest{
		Envelope_xdr: xdr,
		Access_token: signedToken,
	}

	reqBody, err := json.Marshal(sig)
	if ssn.Log(err, "SignTxnService: Marshal request") {
		return "", err
	}

	// Build the request
	req, err := http.NewRequest("POST", signer, bytes.NewReader(reqBody))
	if ssn.Log(err, "SignTxnService: Build HTTP request") {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if ssn.Log(err, "SignTxnService: Send HTTP request") {
		return "", err
	}
	defer res.Body.Close()

	// Get the response body
	respBody, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "SignTxnService: Read response body") {
		return "", err
	}

	// Extract the envelope
	err = json.Unmarshal(respBody, &sig)
	if ssn.Log(err, "SignTxnService: Unmarshal response body") {
		return "", err
	}

	return sig.Envelope_xdr, nil
}

// SubmitTransactionRequest describes the JSON structure for making a request to the submit transaction API
type SubmitTransactionRequest struct {
	Tx string `json:"tx"`
}

// SubmitTransactionResponse describes the JSON structure for the response from the submit transaction API
type SubmitTransactionResponse struct {
	Hash            string                 `json:"hash,omitempty"`
	Ledger          int32                  `json:"ledger,omitempty"`
	Envelope_xdr    string                 `json:"envelope_xdr,omitempty"`
	Result_xdr      string                 `json:"result_xdr,omitempty"`
	Result_meta_xdr string                 `json:"result_meta_xdr,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Title           string                 `json:"title,omitempty"`
	Status          int                    `json:"status,omitempty"`
	Detail          string                 `json:"detail,omitempty"`
	Extras          map[string]interface{} `json:"extras,omitempty"`
}

// SubmitTxn takes a base64 encoded XDR envelope and submits it to the network via provided API
func SubmitTxn(xdr, api string) (string, error) {
	// Prepare JSON request
	req := SubmitTransactionRequest{
		Tx: xdr,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "VerifySignature: Marshal request body") {
		return "", err
	}

	stReq, err := http.NewRequest("POST", api+"/transactions", bytes.NewBuffer(reqBody))
	if ssn.Log(err, "SubmitTxn: Build HTTP request") {
		return "", err
	}
	stReq.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(stReq)
	if ssn.Log(err, "SubmitTxn: Send HTTP request") {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "SubmitTxn: Read response body") {
		return "", err
	}

	apiResp := SubmitTransactionResponse{}
	err = json.Unmarshal(body, &apiResp)
	if ssn.Log(err, "SubmitTxn: Unmarshal response body") {
		return "", err
	}

	if apiResp.Status != 200 {
		mesg := strconv.Itoa(apiResp.Status) + ": " + apiResp.Title + ": " + fmt.Sprintf("%v", apiResp.Extras["result_codes"])
		return "", errors.New(mesg)
	}
	return apiResp.Hash, nil
}
