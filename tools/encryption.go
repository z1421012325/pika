package tools

import (
	"crypto/md5"
	"os"

	"pika/config"
)

var salf = os.Getenv(config.SALF_NAME)

// 密码加密
func EnCryptionPassword(password string) (crypStr string) {

	password = password + salf
	m := md5.New()
	crypStr = string(m.Sum([]byte(password)))
	return
}

// 密码检测
func CheckPassword(encryptionStr, VerifyPassword string) bool {
	return EnCryptionPassword(VerifyPassword) == encryptionStr
}
