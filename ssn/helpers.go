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
	Blocks []section `json:"blocks"`
}

type section struct {
	Type   string  `json:"type"`
	Text   field   `json:"text"`
	Fields []field `json:"fields"`
}

type field struct {
	Type string `json:"type"`
	Text string `json:"test"`
}

// SendSlackMesg sends a message to a Slack webhook
func SendSlackMesg(mesg, ll, app, env, url string) {
	header := []field{
		{Type: "mrkdwn", Text: "*App Name*"},
		{Type: "mrkdwn", Text: "*Environment*"},
		{Type: "mrkdwn", Text: app},
		{Type: "mrkdwn", Text: env},
		{Type: "mrkdwn", Text: "*Log Level*"},
		{Type: "mrkdwn", Text: ll},
	}

	section := []section{
		{
			Type:   "section",
			Fields: header,
		},
		{
			Type: "section",
			Text: field{
				Type: "mrkdwn",
				Text: "*Message*",
			},
		},
		{
			Type: "section",
			Text: field{
				Type: "mrkdwn",
				Text: mesg,
			},
		},
	}

	block := slackMesg{
		Blocks: section,
	}

	slackBody, _ := json.Marshal(block)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(slackBody))
	req.Header.Add("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
