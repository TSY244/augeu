package utils

import "encoding/base64"

func Base64Encode(originStr []byte) string {
	// base64
	return base64.StdEncoding.EncodeToString(originStr)
}

func Base64Decode(originStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(originStr)
}
