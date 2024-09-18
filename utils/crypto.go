package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(data []byte) (result string) {
	const secret = "pass word"
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
