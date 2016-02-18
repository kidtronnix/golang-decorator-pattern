package client

import (
	"fmt"
	"net/http"
)

// This function will create and Audit decorator. We could use this for adding
// this action to an audit trail
func Audit() ClientDecorator {
	return func(c Client) Client {
		return DoFunc(func(r *http.Request) (*http.Response, error) {
			fmt.Println("[audit] added to audit trail")
			return c.Do(r)
		})
	}
}
