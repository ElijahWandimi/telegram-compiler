package src

import (
	"io/ioutil"
	"net/url"
)

const (
	compile_endpoint = "https://go.dev/_/compile?backend="
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

	req, err := compilerClient.PostForm(compile_endpoint, formBody)
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

