package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(data []byte) (result string) {
	// 修改到配置文件里面 todo
	const secret = "pass code"
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
