package hstring

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5By32 32位md5
func Md5By32(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

// Md5By16 16位md5
func Md5By16(s string) string {
	return Md5By32(s)[8:24]
}

// Sha1 sha1
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
