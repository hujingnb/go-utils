// Package hstring 字符串处理
package hstring

import (
	"bytes"
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
