// Package hex 用于各个进制之间的转换
package hex

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

const (
	minRadix = 2  // 支持的最小进制
	maxRadix = 36 // 支持的最大进制
)

// 每次进位的值
const baseIndex = "0123456789abcdefghijklmnopqrstuvwxyz"

// Input 按照指定进制读取字符串
func Input(numStr string, radix int) (int64, error) {
	numStr = strings.ToLower(numStr)
	// 空字符串
	if len(numStr) <= 0 {
		return 0, errors.New("string is empty")
	}
	// 检查进制基数范围
	if radix < minRadix || radix > maxRadix {
		return 0, errors.New(fmt.Sprintf("radix min %d and max %d", minRadix, maxRadix))
	}
	// 检查字符串必须为整数
	if strings.Index(numStr, ".") >= 0 {
		return 0, errors.New("string is float")
	}
	// 检查字符串是否包含非法字符
	baseIndexList := strings.Split(baseIndex, "")
	reg := regexp.MustCompile("^[" + strings.Join(baseIndexList[:radix], "|") + "]+$")
	if !reg.MatchString(numStr) {
		return 0, errors.New("string has error char")
	}

	// 开始进行转换
	v := 0.0
	length := len(numStr)
	for i := 0; i < length; i++ {
		s := string(numStr[i])
		// 获取当前这个字符的基数
		index := strings.Index(baseIndex, s)
		v += float64(index) * math.Pow(float64(radix), float64(length-1-i)) // 倒序
	}
	return int64(v), nil
}
