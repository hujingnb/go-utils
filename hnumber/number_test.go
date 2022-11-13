package hnumber

import (
	"fmt"
	"testing"
)

func TestHexInput(t *testing.T) {
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
			err:    fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix),
		},
		{ // 基数范围向上超出
			numStr: "10",
			radix:  50,
			err:    fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix),
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
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			value, err := HexInput(test.numStr, test.radix)
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

func TestHexOutput(t *testing.T) {
	testList := []struct {
		num    int64
		radix  int
		result string
		err    string
	}{
		{ // 0
			num:    0,
			radix:  10,
			result: "0",
		},
		{ // 基数范围向下超出
			num:   10,
			radix: 1,
			err:   fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix),
		},
		{ // 基数范围向上超出
			num:   10,
			radix: 50,
			err:   fmt.Sprintf("radix min %d and max %d", hexMinRadix, hexMaxRadix),
		},
		{ // 各进制正常数据
			num:    999,
			radix:  2,
			result: "1111100111",
		},
		{
			num:    999,
			radix:  3,
			result: "1101000",
		},
		{
			num:    999,
			radix:  4,
			result: "33213",
		},
		{
			num:    999,
			radix:  5,
			result: "12444",
		},
		{
			num:    999,
			radix:  6,
			result: "4343",
		},
		{
			num:    999,
			radix:  7,
			result: "2625",
		},
		{
			num:    999,
			radix:  8,
			result: "1747",
		},
		{
			num:    999,
			radix:  9,
			result: "1330",
		},
		{
			num:    999,
			radix:  10,
			result: "999",
		},
		{
			num:    999,
			radix:  11,
			result: "829",
		},
		{
			num:    999,
			radix:  12,
			result: "6b3",
		},
		{
			num:    999,
			radix:  13,
			result: "5bb",
		},
		{
			num:    999,
			radix:  14,
			result: "515",
		},
		{
			num:    999,
			radix:  15,
			result: "469",
		},
		{
			num:    999,
			radix:  16,
			result: "3e7",
		},
		{
			num:    999,
			radix:  17,
			result: "37d",
		},
		{
			num:    999,
			radix:  18,
			result: "319",
		},
		{
			num:    999,
			radix:  19,
			result: "2eb",
		},
		{
			num:    999,
			radix:  20,
			result: "29j",
		},
		{
			num:    999,
			radix:  21,
			result: "25c",
		},
		{
			num:    999,
			radix:  22,
			result: "219",
		},
		{
			num:    999,
			radix:  23,
			result: "1ka",
		},
		{
			num:    999,
			radix:  24,
			result: "1hf",
		},
		{
			num:    999,
			radix:  25,
			result: "1eo",
		},
		{
			num:    999,
			radix:  26,
			result: "1cb",
		},
		{
			num:    999,
			radix:  27,
			result: "1a0",
		},
		{
			num:    999,
			radix:  28,
			result: "17j",
		},
		{
			num:    999,
			radix:  29,
			result: "15d",
		},
		{
			num:    999,
			radix:  30,
			result: "139",
		},
		{
			num:    999,
			radix:  31,
			result: "117",
		},
		{
			num:    999,
			radix:  32,
			result: "v7",
		},
		{
			num:    999,
			radix:  33,
			result: "u9",
		},
		{
			num:    999,
			radix:  34,
			result: "td",
		},
		{
			num:    999,
			radix:  35,
			result: "sj",
		},
		{
			num:    999,
			radix:  36,
			result: "rr",
		},
	}
	// 进行测试
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			value, err := HexOutput(test.num, test.radix)
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

func FuzzRandRange(f *testing.F) {
	f.Add(0, 0)
	f.Fuzz(func(t *testing.T, min, max int) {
		ret, err := RandRange(min, max)
		if min > max {
			if err == nil {
				t.Error("can't check min less then max")
			}
		} else {
			if ret < min || ret > max {
				errStr := fmt.Sprintf("value out of range. min %d, max %d, ret %d", min, max, ret)
				t.Error(errStr)
			}
		}
	})
}
