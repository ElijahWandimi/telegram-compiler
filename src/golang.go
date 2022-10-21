package src

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	compile_endpoint = "http://go.dev/_/compile?backend="
	method		   = "POST"
)


func Golang(code string) (string, error) {
	compilerClient, err := NewCompilerClient()
	if err != nil {
		return "", err
	}

	formBody := url.Values{
		"body": {code},
		"version": {"2"},
		"withVet": {"true"},
	}

	req, err := http.NewRequest(method, compile_endpoint,strings.NewReader(formBody.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}


	res, err := compilerClient.Do(req)
	if err != nil {
		return "", err
	}


	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

