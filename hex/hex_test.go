package hex

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	testList := []struct {
		numStr string
		radix  int
		result int64
		err    string
	}{
		{ // 空字符串
			numStr: "",
			radix:  10,
			err:    "string is empty",
		},
		{ // 基数范围向下超出
			numStr: "10",
			radix:  1,
			err:    fmt.Sprintf("radix min %d and max %d", minRadix, maxRadix),
		},
		{ // 基数范围向上超出
			numStr: "10",
			radix:  50,
			err:    fmt.Sprintf("radix min %d and max %d", minRadix, maxRadix),
		},
		{ // 小数
			numStr: "10.1",
			radix:  10,
			err:    "string is float",
		},
		{ // 包含非法字符
			numStr: "102",
			radix:  2,
			err:    "string has error char",
		},
		{
			numStr: "123",
			radix:  3,
			err:    "string has error char",
		},
		{
			numStr: "1005",
			radix:  5,
			err:    "string has error char",
		},
		{
			numStr: "102m",
			radix:  20,
			err:    "string has error char",
		},
		{
			numStr: "10&",
			radix:  36,
			err:    "string has error char",
		},
		{ // 传入大写字母
			numStr: "10B",
			radix:  12,
			result: 155,
		},
		{ // 各个进制的正常数据
			numStr: "1111100111",
			radix:  2,
			result: 999,
		},
		{
			numStr: "1101000",
			radix:  3,
			result: 999,
		},
		{
			numStr: "33213",
			radix:  4,
			result: 999,
		},
		{
			numStr: "12444",
			radix:  5,
			result: 999,
		},
		{
			numStr: "4343",
			radix:  6,
			result: 999,
		},
		{
			numStr: "2625",
			radix:  7,
			result: 999,
		},
		{
			numStr: "1747",
			radix:  8,
			result: 999,
		},
		{
			numStr: "1330",
			radix:  9,
			result: 999,
		},
		{
			numStr: "999",
			radix:  10,
			result: 999,
		},
		{
			numStr: "829",
			radix:  11,
			result: 999,
		},
		{
			numStr: "6b3",
			radix:  12,
			result: 999,
		},
		{
			numStr: "5bb",
			radix:  13,
			result: 999,
		},
		{
			numStr: "515",
			radix:  14,
			result: 999,
		},
		{
			numStr: "469",
			radix:  15,
			result: 999,
		},
		{
			numStr: "3e7",
			radix:  16,
			result: 999,
		},
		{
			numStr: "37d",
			radix:  17,
			result: 999,
		},
		{
			numStr: "319",
			radix:  18,
			result: 999,
		},
		{
			numStr: "2eb",
			radix:  19,
			result: 999,
		},
		{
			numStr: "29j",
			radix:  20,
			result: 999,
		},
		{
			numStr: "25c",
			radix:  21,
			result: 999,
		},
		{
			numStr: "219",
			radix:  22,
			result: 999,
		},
		{
			numStr: "1ka",
			radix:  23,
			result: 999,
		},
		{
			numStr: "1hf",
			radix:  24,
			result: 999,
		},
		{
			numStr: "1eo",
			radix:  25,
			result: 999,
		},
		{
			numStr: "1cb",
			radix:  26,
			result: 999,
		},
		{
			numStr: "1a0",
			radix:  27,
			result: 999,
		},
		{
			numStr: "17j",
			radix:  28,
			result: 999,
		},
		{
			numStr: "15d",
			radix:  29,
			result: 999,
		},
		{
			numStr: "139",
			radix:  30,
			result: 999,
		},
		{
			numStr: "117",
			radix:  31,
			result: 999,
		},
		{
			numStr: "v7",
			radix:  32,
			result: 999,
		},
		{
			numStr: "u9",
			radix:  33,
			result: 999,
		},
		{
			numStr: "td",
			radix:  34,
			result: 999,
		},
		{
			numStr: "sj",
			radix:  35,
			result: 999,
		},
		{
			numStr: "rr",
			radix:  36,
			result: 999,
		},
	}
	// 进行测试
	for _, test := range testList {
		runName := fmt.Sprintf("test_%s_%d", test.numStr, test.radix)
		t.Run(runName, func(t *testing.T) {
			value, err := Input(test.numStr, test.radix)
			if test.err != "" {
				if err == nil || err.Error() != test.err {
					t.Error("this input need has error")
				}
			} else {
				if err != nil {
					t.Error("this input don't need has error")
				}
			}
			if value != test.result {
				t.Error("this input result is error")
			}
		})
	}
}
