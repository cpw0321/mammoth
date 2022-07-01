package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 ...
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 ...
func SHA256(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
