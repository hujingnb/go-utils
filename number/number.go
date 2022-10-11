// Package hnumber 数字处理
package hnumber

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	hexMinRadix = 2  // 支持的最小进制
	hexMaxRadix = 36 // 支持的最大进制
)

// 每次进位的值
const hexBaseIndex = "0123456789abcdefghijklmnopqrstuvwxyz"

/*
HexInput 按照指定进制读取字符串

参数如下:

 numStr: 输入字符串
 radix: 进制基数
*/
func HexInput(numStr string, radix int) (int64, error) {
	numStr = strings.ToLower(numStr)
	// 空字符串
	if len(numStr) <= 0 {
		return 0, errors.New("string is empty")
	}
	// 检查进制基数范围
	if radix < hexMinRadix || radix > hexMaxRadix {
		return 0, errors.New(fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix))
	}
	// 检查字符串必须为整数
	if strings.Index(numStr, ".") >= 0 {
		return 0, errors.New("string is float")
	}
	// 检查字符串是否包含非法字符
	baseIndexList := strings.Split(hexBaseIndex, "")
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
		index := strings.Index(hexBaseIndex, s)
		v += float64(index) * math.Pow(float64(radix), float64(length-1-i)) // 倒序
	}
	return int64(v), nil
}

/*
HexOutput 将10进制数字按照特定进制进行输出

参数如下:

 num: 10进制数字
 radix: 要转的目标进制
*/
func HexOutput(num int64, radix int) (string, error) {
	if num == 0 {
		return "0", nil
	}
	// 检查进制基数范围
	if radix < hexMinRadix || radix > hexMaxRadix {
		return "", errors.New(fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix))
	}
	// 计算进制字符串
	numStr := ""
	for num != 0 {
		yu := num % int64(radix)
		numStr = string(hexBaseIndex[yu]) + numStr
		num = num / int64(radix)
	}
	return numStr, nil
}

// RandRange 获得随机数,范围: [min, max]
func RandRange(min int, max int) (int, error) {
	if min > max {
		return 0, errors.New(fmt.Sprintf("min %d need less max %d", min, max))
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(max + 1 - min)
	return randNum + min, nil
}
