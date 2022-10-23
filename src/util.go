package src

import (
	"encoding/json"

	"github.com/oyamo/telegram-compiler/config"
)

// Escape : Escape special characters like <, >, &, ", ' and /.
func Escape(s string) string {
	stringBuilder := ""
	for _, c := range s {
		switch c {
		case '"':
			stringBuilder += `\"`
		case '\n':
			stringBuilder += `\n`
		default:
			stringBuilder += string(c)
		}
	}
	return stringBuilder
}

func ConstructPayload(code string, language string) (string, error) {
	codeMap := map[string]interface{}{
		"script": code,
		"stdin":  nil,
		"language": language,
		"versionIndex": 0,
		"clientId": config.CLIENT_ID,
   		 "clientSecret": config.CLIENT_SECRET,
	}

	// convert map to json
	jsonPayload, err := json.Marshal(codeMap)
	if err != nil {
		return "", err
	}

	return string(jsonPayload), nil

}