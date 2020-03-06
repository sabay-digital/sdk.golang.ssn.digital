package ssnclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
)

/*
*
* Service/Merchant Specific
*
 */

// IsPreauthorizedAccount returns a public key that may potentially hold preauthorization for a service
func IsPreauthorizedAccount(serviceKey, api string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, api+"/accounts/"+serviceKey, nil)
	if ssn.Log(err, "IsPreauthorizedAccount: Build HTTP request") {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if ssn.Log(err, "IsPreauthorizedAccount: Send HTTP request") {
		return "", err
	}
	defer res.Body.Close()

	// Read the request response
	body, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "IsPreauthorizedAccount: Read HTTP response body") {
		return "", err
	}

	// Take the JSON apart
	account := ssn.Account{}
	err = json.Unmarshal(body, &account)
	if ssn.Log(err, "IsPreauthorizedAccount: Unmarshal response body") {
		return "", err
	}

	for i := range account.Balances {
		// Check if trustline exists to indicate subaccount
		if account.Balances[i].Asset_code == "SUBACCOUNT" {
			// Check the account is authorised as a subaccount
			if account.Balances[i].Is_authorized {
				return account.Balances[i].Asset_issuer, nil
			}
		}
	}
	return serviceKey, nil
}
