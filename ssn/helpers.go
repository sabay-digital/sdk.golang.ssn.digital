package ssn

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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

// SendSlackMesg sends a message to a Slack webhook
func SendSlackMesg(mesg, url string) {
	slackBody, _ := json.Marshal(slackMesg{Text: mesg})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(slackBody))
	req.Header.Add("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
