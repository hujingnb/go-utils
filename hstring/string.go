// Package hstring 字符串处理
package hstring

import (
	"bytes"
	"fmt"
	"github.com/hujingnb/go-utils/harray"
	"strconv"
	"strings"
	"unicode"
)

// Reverse 字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	harray.Reverse(runes)
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

type PadType int32

const (
	PadLeft PadType = iota
	PadRight
)

// Pad 填充字符串到指定长度
// 参数如下:
//  input 待处理数据, 可以是int/string/float 等等
//  length 指定长度
//  padStr 填充的字符串
//  _type 指定填充类型, 填充在左或在右
func Pad(input interface{}, length int, padStr string, _type PadType) string {
	s := ToString(input)
	// 长度足够, 直接返回
	l := len(s)
	if l >= length {
		return s
	}
	// 达到指定长度需要补充的字符串
	output := ""
	for i := 1; i <= length-l; i = i + len(padStr) {
		output += padStr
	}
	if _type == PadLeft {
		return output + s
	} else if _type == PadRight {
		return s + output
	} else {
		panic("type is error")
	}
}

// Join 将数组转为字符串, 按照 seq 进行分割
// 因为系统库提供的 strings.Join 函数, 参数只支持 string
// 故这里重写一个方法, 可接收任意参数
func Join[T any](array []T, seq string) string {
	joinArray := make([]string, 0, len(array))
	for _, item := range array {
		joinArray = append(joinArray, ToString(item))
	}
	return strings.Join(joinArray, seq)
}
