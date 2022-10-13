package hmap

import (
	"fmt"
	"testing"
)

func TestIterateByOrder(t *testing.T) {
	m := map[string]int{
		"b": 2,
		"c": 3,
		"f": 6,
		"d": 4,
		"a": 1,
	}
	var varList []int
	IterateByOrder(m, func(k string, v int) {
		varList = append(varList, v)
	})
	// 检查遍历顺序
	for i := 1; i < len(varList); i++ {
		if varList[i] < varList[i-1] {
			t.Error("order is error")
		}
	}
	// 检查结果数量
	if len(varList) != len(m) {
		t.Error("value is miss")
	}
}

func TestEqual(t *testing.T) {
	testList := []struct {
		input1 map[string]string
		input2 map[string]string
		output bool
	}{
		{
			input1: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			output: true,
		},
		{
			input1: map[string]string{
				"a": "aa",
				"b": "bb",
				"c": "cc",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			output: false,
		},

		{
			input1: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bb",
				"c": "cc",
			},
			output: false,
		},

		{
			input1: map[string]string{
				"a": "aa",
				"b": "bbb",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			output: false,
		},

		{
			input1: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bbb",
			},
			output: false,
		},

		{
			input1: map[string]string{
				"a": "aa",
				"b": "bbb",
				"c": "cc",
			},
			input2: map[string]string{
				"a": "aa",
				"b": "bb",
			},
			output: false,
		},

		{
			input1: map[string]string{
				"a": "aa",
			},
			input2: map[string]string{},
			output: false,
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			ret := Equal(test.input1, test.input2)
			if ret != test.output {
				t.Error("map equal fail")
			}
		})
	}
}
