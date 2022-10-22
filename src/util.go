package src

import "encoding/json"

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
		"args":   nil,
		"stdin":  nil,
		"language": language,
		"versionIndex": 4,
		"libs": nil,
		"projectKey": 1001,
		"hasInputFiles": false,
	}

	// convert map to json
	jsonPayload, err := json.Marshal(codeMap)
	if err != nil {
		return "", err
	}

	return string(jsonPayload), nil

}