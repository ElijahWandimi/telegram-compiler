package src

import (
	"fmt"
	_"net/url"
	_"strings"
)

const (
	kt_c_endpoint = "https://api.kotlinlang.org//api/1.7.20/compiler/run"
	kt_c_method  = "POST"
)

func Kotlin(code string) (*Response, error) {

	payload := fmt.Sprintf(`{
		"args": "344",
		"files": [
		  {
			"name": "File.kt",
			"text": "%s",
			"publicId": ""
		  }
		],
		"confType": "java"
	  }`, code)

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		"referer":      "https://play.kotlinlang.org/",
		"origin":       "https://play.kotlinlang.org",
	}
	req := NewRequest(kt_c_method, headers, payload, kt_c_endpoint)
	return req.Execute()
} 