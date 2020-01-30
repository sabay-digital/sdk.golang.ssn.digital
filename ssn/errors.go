package ssn

import (
	"fmt"
	"time"
)

// Log provides time stamped output for stack tracing errors
func Log(err error, cause string) bool {
	if err != nil {
		t := time.Now()
		fmt.Printf("%s: Error: %s Cause: %s\n", t.Format(time.RFC3339), error.Error(err), cause)
		return true
	}
	return false
}

// ParseErrorPayload returns a RedirectPayload object that can be used to display a template error
func ParseErrorPayload(cause, detail, host string) RedirectPayload {
	return RedirectPayload{
		RedirectURL: "https://" + host + "/v1/error",
		Payload: []PayloadItem{
			{Key: "cause", Value: cause},
			{Key: "detail", Value: detail},
		},
	}
}