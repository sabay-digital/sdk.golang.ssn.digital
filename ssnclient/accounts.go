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

// AuthorizeSubAccount() returns true if the Service/Merchant has trust line `SUBACCOUNT` else false
func AuthorizeSubAccount(serviceKey, api string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, api+"/accounts/"+serviceKey, nil)
	if ssn.Log(err, "AuthorizeSubAccount: Build HTTP request") {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if ssn.Log(err, "AuthorizeSubAccount: Send HTTP request") {
		return "", err
	}
	defer res.Body.Close()

	// Read the request response
	body, err := ioutil.ReadAll(res.Body)
	if ssn.Log(err, "AuthorizeSubAccount: Read HTTP response body") {
		return "", err
	}

	// Take the JSON apart
	account := ssn.Account{}
	err = json.Unmarshal(body, &account)
	if ssn.Log(err, "AuthorizeSubAccount: Unmarshal response body") {
		return "", err
	}

	for i := range account.Balances {
		// Check asset code matches
		if account.Balances[i].Asset_code == "SUBACCOUNT" {
			// Check the account is authorised to hold the asset
			if account.Balances[i].Is_authorized {
				return account.Balances[i].Asset_issuer, nil
			}
		}
	}
	return serviceKey, nil
}
