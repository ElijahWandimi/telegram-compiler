package src

import (
	"net/http"
	"net/url"
	"time"
)



func NewCompilerClient() (*http.Client, error) {
	proxy := NextProxy()
	proxy_url, err := url.Parse(proxy.FullAddress())
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy_url),
	}

	_ = transport

	client := &http.Client{
		//Transport: transport,
		Timeout: 120 * time.Second,
	}

	return (client), nil
}