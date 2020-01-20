package ssnclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.sabay.com/payment-network/sdk/sdk.golang.ssn.digital/ssn"
)

/*
*
* Service/Merchant Specific
*
 */

// GetServiceName returns the service_name data key for a specified account public key
func GetServiceName(serviceKey, api string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, api+"/accounts/"+serviceKey, nil)
	if ssn.Log(err, "GetServiceName: Build HTTP request") {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if ssn.Log(err, "GetServiceName: Send HTTP request") {
		return "", err
	}
	defer res.Body.Close()

	// Read the request response
	body, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "GetServiceName: Read HTTP response body") {
		return "", err
	}

	// Take the JSON apart
	account := ssn.Account{}
	err = json.Unmarshal(body, &account)
	if ssn.Log(err, "GetServiceName: Unmarshal response body") {
		return "", err
	}
	return account.Data.Service_name, nil
}
