package util

import (
	"crypto/md5"
	"encoding/hex"
)

// DigestMD5 转换字符串到md5
func DigestMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
