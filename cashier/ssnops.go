package cashier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/stellar/go/txnbuild"
)

type buildResponse struct {
	Envelope_xdr string `json:"envelope_xdr"`
}

type signRequest struct {
	Xdr_string string `json:"xdr_string"`
}

type submitResponse struct {
	Hash string `json:"hash"`
}

// BuildTxn sends transaction information to the SSN API to build an XDR envelope
func BuildTxn(from, to, amount, assetCode, assetIssuer, memo, api string) string {
	tx := url.Values{}
	tx.Set("from", from)
	tx.Set("to", to)
	tx.Set("amount", amount)
	tx.Set("asset_code", assetCode)
	tx.Set("asset_issuer", assetIssuer)
	tx.Set("memo", memo)

	req, err := http.NewRequest("POST", api+"/create/transaction", strings.NewReader(tx.Encode()))
	if err != nil {
		fmt.Println(error.Error(err))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	apiResp := buildResponse{}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		fmt.Println(error.Error(err))
	}
	return apiResp.Envelope_xdr
}

// SignLocal takes a base64 encoded XDR envelope and signs it with the provided secret key
func SignLocal(xdr, signer, networkPassphrase string) string {
	// Deserialise the provided transaction
	tx, err := txnbuild.TransactionFromXDR(xdr)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// Explicitly set the network where this transaction is to be valid
	tx.Network = networkPassphrase

	// Add a signature
	err = tx.SignWithKeyString(signer)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// Serialise the transaction
	b64, err := tx.Base64()
	if err != nil {
		fmt.Println(error.Error(err))
	}

	return b64
}

// SignService takes a base64 encoded XDR envelope and sends it to the specified sign service
func SignService(xdr, signer string) string {
	// Prepare the request body
	sig := signRequest{
		Xdr_string: xdr,
	}

	reqBody, err := json.Marshal(sig)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// Build the request
	req, err := http.NewRequest("POST", signer, bytes.NewReader(reqBody))
	if err != nil {
		fmt.Println(error.Error(err))
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// Get the response body
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// Extract the envelope
	err = json.Unmarshal(respBody, &sig)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	return sig.Xdr_string
}

// SubmitTxn takes a base64 encoded XDR envelope and submits it to the network via provided API
func SubmitTxn(xdr, api string) string {
	tx := url.Values{}
	tx.Set("tx", xdr)

	req, err := http.NewRequest("POST", api+"/transactions", strings.NewReader(tx.Encode()))
	if err != nil {
		fmt.Println(error.Error(err))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	apiResp := submitResponse{}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		fmt.Println(error.Error(err))
	}
	return apiResp.Hash
}
