package hstring

import (
	"fmt"
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
		runName := fmt.Sprintf("test_%s", test.Input)
		t.Run(runName, func(t *testing.T) {
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
		runName := fmt.Sprintf("test_%s", test.Input)
		t.Run(runName, func(t *testing.T) {
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
		runName := fmt.Sprintf("test_%s", test.Input)
		t.Run(runName, func(t *testing.T) {
			ret := CamelToSnake(test.Input)
			if ret != test.Output {
				t.Error("output is error")
			}
		})
	}
}
