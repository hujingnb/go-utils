package harray

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	testList := []struct {
		input []interface{}
		size  int
		err   string
	}{
		{ // 错误size
			input: []interface{}{1, 2, 3},
			size:  0,
			err:   "size min 1",
		},
		{ // 正常数据
			input: []interface{}{1, 2, 3, 4},
			size:  1,
		},
		{
			input: []interface{}{1, 2, 3, 4},
			size:  2,
		},
		{
			input: []interface{}{1, 2, 3, 4},
			size:  3,
		},
		{
			input: []interface{}{1, 2, 3, 4},
			size:  4,
		},
		{
			input: []interface{}{1, 2, 3, 4},
			size:  5,
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("test_%d", index)
		t.Run(runName, func(t *testing.T) {
			ret, err := Chunk(test.input, test.size)
			if test.err != "" {
				if err == nil || err.Error() != test.err {
					t.Error("this input need nil error")
				}
			} else {
				if err != nil {
					t.Error("this input don't need error")
				}
			}
			if err == nil {
				retLen := 0
				// 检查结果中每个数组的数量
				for i, item := range ret {
					// 最后一组数量不定, 不进行检查
					if len(item) != test.size && i < len(ret)-1 {
						t.Error("chunk array length is error")
					}
					retLen += len(item)
				}
				if retLen != len(test.input) {
					t.Error("chunk array all length is error")
				}
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	testList := [][]interface{}{
		{1, 2, 3, 4},
		{2, 2, 2, 2},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("test_%d", index)
		t.Run(runName, func(t *testing.T) {
			Shuffle(test)
		})
	}
}

func TestInArray(t *testing.T) {
	testList := []struct {
		Input  []int
		Search int
		Result bool
	}{
		{
			Input:  []int{1, 2, 3},
			Search: 2,
			Result: true,
		},
		{
			Input:  []int{1, 2, 3},
			Search: 5,
			Result: false,
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("test_%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := InArray(test.Input, test.Search)
			if ret != test.Result {
				t.Error("search in array fail")
			}
		})
	}
}

func TestUnique(t *testing.T) {
	testList := []struct {
		Input  []int
		Output []int
	}{
		{
			Input:  []int{1, 2, 3, 2, 3},
			Output: []int{1, 2, 3},
		},
		{
			Input:  []int{1, 2, 2, 2, 3},
			Output: []int{1, 2, 3},
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := Unique(test.Input)
			if !Equal(test.Output, ret) {
				t.Error("unique is error")
			}
		})
	}
}

func TestEqual(t *testing.T) {
	testList := []struct {
		Input1 []int
		Input2 []int
		Result bool
	}{
		{
			Input1: []int{1, 2, 3},
			Input2: []int{1, 2, 3},
			Result: true,
		},
		{
			Input1: []int{1, 2, 3},
			Input2: []int{1, 2, 3, 4},
			Result: false,
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := Equal(test.Input1, test.Input2)
			if ret != test.Result {
				t.Error("equal is error")
			}
		})
	}
}
