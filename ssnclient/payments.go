package ssnclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
*
* Create and Stream Payments
*
 */

type buildResponse struct {
	Envelope_xdr string `json:"envelope_xdr"`
}

// CreatePayment sends transaction information to the SSN API to build an XDR envelope
func CreatePayment(from, to, amount, assetCode, assetIssuer, memo, api string) string {
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
