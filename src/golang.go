package src

import (
	"net/url"
)

const (
	compile_endpoint = "https://go.dev/_/compile?backend="
	method		   = "POST"
)


func Golang(code string) (*Response, error) {
	formBody := url.Values{
		"version": {"2"},
		"body": {code},
		"withVet": {"true"},
	}
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept": "application/json",
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	}
	req := NewRequest(code, headers, formBody.Encode(), compile_endpoint)
	return req.Execute()

}

