package client

import (
	"fmt"
	"net/http"
)

// Here is an example of a mock client that can be used to do pretend http requests.
// Notice our client has no knowledge of the other decorators that have orthogonal concerns.
type MockClient struct{}

func (mc MockClient) Do(*http.Request) (*http.Response, error) {
	fmt.Println("[client] Doing my ting")
	return &http.Response{}, nil
}
