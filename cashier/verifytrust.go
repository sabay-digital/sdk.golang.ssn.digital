package cashier

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// VerifyTrust checks whether the provided asset and assetIssuer is trusted by destination
func VerifyTrust(destination, asset, assetIssuer, api string) bool {
	// Prepare URL encoded values
	vtValues := url.Values{}
	vtValues.Set("account", destination)
	vtValues.Set("asset_code", asset)
	vtValues.Set("asset_issuer", assetIssuer)
	vtBody := strings.NewReader(vtValues.Encode())

	// Send the request to the API and get the reponse
	vtReq, err := http.NewRequest("POST", api+"/verify/trust", vtBody)
	if err != nil {
		fmt.Println(error.Error(err))
	}
	vtResp, err := http.DefaultClient.Do(vtReq)
	if err != nil {
		fmt.Println(error.Error(err))
	}

	// 200 signifies the cashier is trusted
	if vtResp.StatusCode == 200 {
		return true
	}
	return false
}
