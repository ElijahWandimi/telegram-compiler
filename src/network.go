package src

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/oyamo/telegram-compiler/config"
)

type Request struct {
	Method     string
	Headers  map[string]string
	FormBody string
	Url      string
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

func NewRequest(method string, headers map[string]string, formBody string, url string) *Request {
	return &Request{method, headers, formBody, url}
}

func (r *Request) Execute() (*Response, error) {
	resp := &Response{}
	compilerClient, err := NewCompilerClient()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(config.METHOD, r.Url, bytes.NewBuffer([]byte(r.FormBody)))


	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	if err != nil {
		return nil, err
	}

	res, err := compilerClient.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp.StatusCode = res.StatusCode
	resp.Headers = res.Header
	resp.Body = string(body)

	return resp, nil
}
