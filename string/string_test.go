package hstring

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{ // 单字符
			Input:  "a",
			Output: "a",
		},
		{ // 奇数
			Input:  "abc",
			Output: "cba",
		},
		{ // 偶数
			Input:  "abcd",
			Output: "dcba",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Reverse(test.Input)
			if ret != test.Output {
				t.Error("this input is error")
			}
		})
	}
}

func TestSnakeToCamel(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "a",
			Output: "a",
		},
		{
			Input:  "abc",
			Output: "abc",
		},
		{
			Input:  "abc_def",
			Output: "abcDef",
		},
		{
			Input:  "_def",
			Output: "Def",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := SnakeToCamel(test.Input)
			if ret != test.Output {
				t.Error("output is error")
			}
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "a",
			Output: "a",
		},
		{
			Input:  "abc",
			Output: "abc",
		},
		{
			Input:  "abcDef",
			Output: "abc_def",
		},
		{
			Input:  "Def",
			Output: "def",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := CamelToSnake(test.Input)
			if ret != test.Output {
				t.Error("output is error")
			}
		})
	}
}

func TestToString(t *testing.T) {
	check := func(data interface{}, s string) {
		ret := ToString(data)
		if ret != s {
			t.Error("to string error", data, s)
		}
	}
	check("hello", "hello")
	check(int8(1), "1")
	check(int16(1), "1")
	check(int32(1), "1")
	check(int64(1), "1")
	check(int(1), "1")
	check(uint8(1), "1")
	check(uint16(1), "1")
	check(uint32(1), "1")
	check(uint64(1), "1")
	check(float32(1.1), "1.1")
	check(float64(1.1), "1.1")
	check(true, "true")
	check(false, "false")
	check([]byte{'h', 'e', 'l', 'l', 'o'}, "hello")
	m := map[string]string{"a": "a"}
	check(m, "map[a:a]")
}

func TestGet32Md5(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Get32Md5(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}

func TestGet16Md5(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "bc4b2a76b9719d91",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Get16Md5(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}
