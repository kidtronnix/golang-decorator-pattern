package client

import (
	"log"
	"net/http"
	"os"
)


// Our client is defined as something that does a http request and gets a response and error
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// We define what a decorator to our client will look like
type ClientDecorator func(Client) Client

// Singature of DoFunc
type DoFunc func(*http.Request) (*http.Response, error)

// Add method to DoFunc type to satisfy Client interface
func (d DoFunc) Do(r *http.Request) (*http.Response, error) {
	return d(r)
}

// Our client that will be composed of many layers of functionality
// In this example when we call Do(r), we want to do some logging, write to the audit trail 
// and do a mock http request.
func NewClient() Client {
	l := log.New(os.Stdout, "", 0)
	c := Client(MockClient{})
	c = Logging(l)(c)
	c = Audit()(c)
	return c
}
