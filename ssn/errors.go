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
