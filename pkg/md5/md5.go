package md5

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "www.liuziying.com"

// EnCryptPassword 使用md5给密码加密
func EnCryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
