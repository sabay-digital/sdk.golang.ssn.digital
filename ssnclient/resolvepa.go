package ssnclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/stellar/go/strkey"

	"github.com/agl/ed25519/extra25519"
	"golang.org/x/crypto/nacl/box"
)

/*
*
* Resolve Payment Addresses
*
*/


type paResponse struct {
	Network_address string    `json:"network_address"`
	Public_key      string    `json:"public_key"`
	Asset_code      string    `json:"asset_code"`
	Payment_type    string    `json:"payment_type"`
	Service_name    string    `json:"service_name"`
	Encrypted       string    `json:"encrypted"`
	Details         paDetails `json:"details"`
}

type paDetails struct {
	Payment_info string    `json:"payment_info"`
	Memo         string    `json:"memo"`
	Payment      paPayment `json:"payment"`
	Service_fee  paPayment `json:"service_fee"`
}

type paPayment struct {
	Amount     string `json:"amount"`
	Asset_code string `json:"asset_code"`
}

// ResolvePA sends a payment address to the PA service for resolving.
// It also sends the cashiers asset issuing public key and local signing public key to decrypt any response
func ResolvePA(assetIssuer, publicKey, paymentAddress, resolverURL, decryptionKey string) (paResponse, bool) {
	// Prepare URL encoded values
	paValues := url.Values{}
	paValues.Set("asset_issuer", assetIssuer)
	paValues.Set("public_key", publicKey)
	paValues.Set("payment_address", paymentAddress)
	paBody := strings.NewReader(paValues.Encode())

	// Send the request to the PA service and get the reponse
	paReq, err := http.NewRequest("POST", resolverURL, paBody)
	if err != nil {
		fmt.Println(error.Error(err))
		return paResponse{}, true
	}
	paResp, err := http.DefaultClient.Do(paReq)
	if err != nil {
		fmt.Println(error.Error(err))
		return paResponse{}, true
	}

	// Unmarshall the PA response
	body, err := ioutil.ReadAll(paResp.Body)
	if err != nil {
		fmt.Println(error.Error(err))
		return paResponse{}, true
	}
	resp := paResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(error.Error(err))
		return paResponse{}, true
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
		if err != nil {
			fmt.Println(error.Error(err))
			return paResponse{}, true
		}
		var decryptNonce [24]byte
		copy(decryptNonce[:], encrypted[:24])
		decrypted, ok := box.OpenAfterPrecomputation(nil, encrypted[24:], &decryptNonce, &sharedDecryptKey)
		if !ok {
			fmt.Println("Decryption failed")
			return paResponse{}, true
		}

		// Unmarshal the decrypted response
		err = json.Unmarshal(decrypted, &resp)
		if err != nil {
			fmt.Println("Decryption failed")
			return paResponse{}, true
		}
	}

	if resp.Network_address == "" {
		fmt.Println("Decryption failed")
		return paResponse{}, true
	}

	return resp, false
}
