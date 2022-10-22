package src

import (
	_"net/url"
	_"strings"
	"github.com/oyamo/telegram-compiler/config"
)


func CPlus(code string) (*Response, error) {

	payload, err := ConstructPayload(code, "cpp")
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	}
	req := NewRequest(config.METHOD, headers, payload, config.ENDPOINT)
	return req.Execute()
} 