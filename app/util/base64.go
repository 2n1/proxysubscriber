package util

import "encoding/base64"

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Base64Decode(data string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
