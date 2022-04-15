package common

import (
	"encoding/base64"
	json "github.com/json-iterator/go"
)

func Encode(data interface{}) string {
	content, _ := json.Marshal(data)
	return base64.StdEncoding.EncodeToString(content)
}

func Decode(encodeString string, data interface{}) error {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		return err
	}
	return json.Unmarshal(decodeBytes, &data)
}
