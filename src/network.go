package src

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Code string
	Headers map[string]string
	FormBody string
	Url string
}

type Response struct {
	StatusCode int
	Headers http.Header
	Body string
}

func NewRequest(code string, headers map[string]string, formBody string, url string) *Request {
	return &Request{code, headers, formBody, url}
}

func  (r *Request) Execute() (*Response, error){
	resp := &Response{}
	compilerClient, err := NewCompilerClient()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, compile_endpoint,strings.NewReader(r.FormBody))

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