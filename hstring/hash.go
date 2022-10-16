package hstring

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Get32Md5 32位md5
func Get32Md5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

// Get16Md5 16位md5
func Get16Md5(s string) string {
	return Get32Md5(s)[8:24]
}

// GetSha1 sha1
func GetSha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
