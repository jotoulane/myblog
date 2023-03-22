package util

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "bilibili.com"

func EncryptPassword(Password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(Password)))
}
