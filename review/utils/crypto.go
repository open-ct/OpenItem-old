package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return ConvertToBase64(h.Sum(nil))
}

func ConvertToBase64(str []byte) string {
	enc := base64.StdEncoding.EncodeToString(str)
	return enc
}
