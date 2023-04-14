package src

import (
	_"fmt"
	_ "net/url"
	"github.com/oyamo/telegram-compiler/config"
)


func Golang(code string) (*Response, error) {
	
	payload, err := ConstructPayload(code, "go", 4)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	}
	req := NewRequest(config.METHOD, headers, payload, config.ENDPOINT)
	return req.Execute()

}
