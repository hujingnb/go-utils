// Package hstring 字符串处理
package hstring

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"unicode"
)

// Reverse 字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CamelToSnake 驼峰转下划线
func CamelToSnake(s string) string {
	buffer := bytes.NewBuffer([]byte{})
	for i, c := range s {
		if unicode.IsUpper(c) { // 碰到大写字母, 则在前面添加下划线
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteRune(unicode.ToLower(c))
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

// SnakeToCamel 下划线转驼峰
func SnakeToCamel(s string) string {
	buffer := bytes.NewBuffer([]byte{})
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '_' {
			continue
		}
		if i > 0 && s[i-1] == '_' {
			buffer.WriteRune(unicode.ToUpper(rune(s[i])))
		} else {
			buffer.WriteRune(rune(s[i]))
		}
	}
	return buffer.String()
}

// ToString 转字符串
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int:
		return strconv.FormatInt(int64(value.(int)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return fmt.Sprintf("%g", value.(float32))
	case float64:
		return fmt.Sprintf("%g", value.(float64))
	case bool:
		return strconv.FormatBool(value.(bool))
	case []byte:
		return string(value.([]byte))
	default:
		return fmt.Sprintf("%+v", value)
	}
}

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
