package cashier

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// StatusResponse is the root level JSON object for the cashier endpoint /v1/status
type StatusResponse struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
}

// InfoResponse is the root level JSON object for the cashier endpoint /v1/info
type InfoResponse struct {
	Assets_issued    []Asset `json:"assets_issued"`
	Payment_provider string  `json:"payment_provider"`
	Payment_type     string  `json:"payment_type"`
	Authorization    string  `json:"authorization"`
}

// Asset is the JSON object that describes assets issued by a cashier
type Asset struct {
	Asset_code   string `json:"asset_code"`
	Limit_amount int    `json:"limit_amount,omitempty"`
	Asset_Issuer string `json:"asset_issuer"`
}

// RedirectPayload describes the payload to be added to the cashier redirect form
type RedirectPayload struct {
	RedirectURL string
	Payload     []PayloadItem
}

// PayloadItem is a single payload item to be redirected
type PayloadItem struct {
	Key   string
	Value string
}

type slackMesg struct {
	Text string `json:"text"`
}

// ParseErrorPayload returns a RedirectPayload object that can be used to display an error in the cashier
func ParseErrorPayload(cause, detail, host string) RedirectPayload {
	return RedirectPayload{
		RedirectURL: "https://" + host + "/v1/error",
		Payload: []PayloadItem{
			{Key: "cause", Value: cause},
			{Key: "detail", Value: detail},
		},
	}
}

// SendSlackMesg sends a message to a Slack webhook
func SendSlackMesg(mesg, url string) {
	slackBody, _ := json.Marshal(slackMesg{Text: mesg})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(slackBody))
	req.Header.Add("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
