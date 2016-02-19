package main

import (
	"net/http"

	"github.com/smaxwellstewart/golang-decorator-pattern/client"
)

func main() {
	c := client.NewClient()
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	c.Do(req)
}
