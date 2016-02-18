package client

import (
	"log"
	"net/http"
)

// This function will return a logging decorator, we give it a logger so we
// can have create a custom logger to our liking
func Logging(l *log.Logger) ClientDecorator {
	// Here we return a function that satifies out ClientDecorator type
	return func(c Client) Client {
		return DoFunc(func(r *http.Request) (*http.Response, error) {
			l.Println("[logging] Client do called: " + r.Host)
			return c.Do(r)
		})
	}
}
