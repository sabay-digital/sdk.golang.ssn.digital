package ssnclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/agl/ed25519/extra25519"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/stellar/go/strkey"
	"golang.org/x/crypto/nacl/box"
)

/*
*
* Resolve Payment Addresses
*
 */

// ResolverRequest describes the JSON structure for making a request to the payment address resolver API
type ResolverRequest struct {
	Asset_issuer    string `json:"asset_issuer,omitempty"`
	Public_key      string `json:"public_key,omitempty"`
	Payment_address string `json:"payment_address"`
}

// ResolverResponse describes the JSON structure for the response from the payment address resolver API
type ResolverResponse struct {
	Network_address string                   `json:"network_address,omitempty"`
	Public_key      string                   `json:"public_key,omitempty"`
	Asset_code      string                   `json:"asset_code,omitempty"`
	Payment_type    string                   `json:"payment_type,omitempty"`
	Service_name    string                   `json:"service_name,omitempty"`
	Encrypted       string                   `json:"encrypted,omitempty"`
	Details         *ResolverResponseDetails `json:"details,omitempty"`
	Status          int                      `json:"status,omitempty"`
	Title           string                   `json:"title,omitempty"`
	Signature       string                   `json:"signature,omitempty"`
}

// ResolverResponseDetails describes the JSON structure for the nested details part of the response from the payment address resolver API
type ResolverResponseDetails struct {
	Payment_info                string                     `json:"payment_info,omitempty"`
	Memo                        string                     `json:"memo,omitempty"`
	Recurring_payment_frequency string                     `json:"recurring_payment_frequency,omitempty"`
	Recurring_payment_interval  string                     `json:"recurring_payment_interval,omitempty"`
	Recurring_payment_start     string                     `json:"recurring_payment_start,omitempty"`
	Payment                     *ResolverPaymentDetails    `json:"payment,omitempty"`
	Service_fee                 *ResolverServiceFeeDetails `json:"service_fee,omitempty"`
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

// ResolvePA sends a payment address to the PA service for resolving.
// If the returned response is encrypted it will attempt to decrypt using the provided keys
func ResolvePA(assetIssuer, publicKey, paymentAddress, resolverURL, decryptionKey string) (ResolverResponse, error) {
	// Prepare JSON request
	req := ResolverRequest{
		Asset_issuer:    assetIssuer,
		Public_key:      publicKey,
		Payment_address: paymentAddress,
	}
	reqBody, err := json.Marshal(req)
	if ssn.Log(err, "ResolvePA: Marshal request body") {
		return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
	}

	// Send the request to the API and get the reponse
	paReq, err := http.NewRequest("POST", resolverURL, bytes.NewBuffer(reqBody))
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

	if resp.Encrypted != "" {
		// Convert keys
		decodePk := strkey.MustDecode(strkey.VersionByteAccountID, resp.Public_key)
		decodeLocalSk := strkey.MustDecode(strkey.VersionByteSeed, decryptionKey)
		var curvePublic, curvePrivate, edPublic [32]byte
		var edPrivate [64]byte
		copy(edPublic[:], decodePk)
		copy(edPrivate[:], decodeLocalSk)
		extra25519.PublicKeyToCurve25519(&curvePublic, &edPublic)
		extra25519.PrivateKeyToCurve25519(&curvePrivate, &edPrivate)

		// Pre - box
		var sharedDecryptKey [32]byte
		box.Precompute(&sharedDecryptKey, &curvePublic, &curvePrivate)

		// Decrypt message
		encrypted, err := base64.StdEncoding.DecodeString(resp.Encrypted)
		if ssn.Log(err, "ResolvePA: Decode encrypted string") {
			return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
		}
		var decryptNonce [24]byte
		copy(decryptNonce[:], encrypted[:24])
		decrypted, ok := box.OpenAfterPrecomputation(nil, encrypted[24:], &decryptNonce, &sharedDecryptKey)
		if !ok {
			return ResolverResponse{Status: 403, Title: "Decryption Failed"}, nil
		}

		// Unmarshal the decrypted response
		err = json.Unmarshal(decrypted, &resp)
		if ssn.Log(err, "ResolvePA: Unmarshal decrypted response") {
			return ResolverResponse{Status: 500, Title: "Internal System Error"}, err
		}
	}

	if resp.Network_address == "" {
		return ResolverResponse{Status: 400, Title: "Payment destination missing"}, nil
	}

	return resp, nil
}
