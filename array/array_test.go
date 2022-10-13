package harray

import (
	"fmt"
	hmap "github.com/hujingnb/go-utils/map"
	hstring "github.com/hujingnb/go-utils/string"
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

func TestIndexOf(t *testing.T) {
	testList := []struct {
		Input  []int
		Search int
		Result int
	}{
		{
			Input:  []int{1, 2, 3},
			Search: 2,
			Result: 1,
		},
		{
			Input:  []int{1, 2, 3},
			Search: 5,
			Result: -1,
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("test_%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := IndexOf(test.Input, test.Search)
			if ret != test.Result {
				t.Error("search array data index fail")
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

func TestIntersect(t *testing.T) {
	testList := []struct {
		Input  [][]int
		Output []int
	}{
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{3, 4, 5, 6},
			},
			Output: []int{3, 4},
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := Intersect(test.Input...)
			// 结果可能乱序, 比较内容是否一致
			tmpMap := make(map[int]bool)
			if len(ret) != len(test.Output) {
				t.Error("intersect result error")
			}
			for _, item := range ret {
				tmpMap[item] = true
			}
			for _, item := range test.Output {
				if _, ok := tmpMap[item]; !ok {
					t.Error("intersect result error")
					break
				}
			}
		})
	}
}

func TestDiff(t *testing.T) {
	testList := []struct {
		Input  [][]int
		Output []int
	}{
		{
			Input: [][]int{
				{1, 2, 3, 4, 5, 6},
				{4, 5, 6, 8, 9},
				{2, 10, 11},
			},
			Output: []int{1, 3},
		},
	}
	for index, test := range testList {
		runName := fmt.Sprintf("%d", index)
		t.Run(runName, func(t *testing.T) {
			ret := Diff(test.Input[0], test.Input[1:]...)
			// 结果可能乱序, 比较内容是否一致
			tmpMap := make(map[int]bool)
			if len(ret) != len(test.Output) {
				t.Error("diff result error")
			}
			for _, item := range ret {
				tmpMap[item] = true
			}
			for _, item := range test.Output {
				if _, ok := tmpMap[item]; !ok {
					t.Error("diff result error")
					break
				}
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	testList := []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			target: 3,
			output: 2,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			target: -1,
			output: -1,
		},
	}
	for _, test := range testList {
		runName := fmt.Sprintf("%s_%d", hstring.ToString(test.input), test.target)
		t.Run(runName, func(t *testing.T) {
			ret := BinarySearch(test.input, func(data int) int {
				return data - test.target
			})
			if ret != test.output {
				t.Error("search fail")
			}
		})
	}
}

func TestGetSureRandArr(t *testing.T) {
	testList := []struct {
		seed   int64
		n      int
		start  int
		end    int
		output []int
	}{
		{
			seed:   500,
			n:      20,
			start:  0,
			end:    21,
			output: []int{14, 7, 10, 17, 2, 5, 15, 16, 3, 12, 11, 4, 1, 8, 13, 19, 9, 18, 0, 6},
		},
		{
			seed:   500,
			n:      20,
			start:  0,
			end:    19,
			output: []int{14, 7, 10, 17, 2, 5, 15, 16, 3, 12, 11, 4, 1, 8, 13, 19, 9, 18, 0, 6},
		},
		{
			seed:   500,
			n:      20,
			start:  0,
			end:    3,
			output: []int{14, 7, 10, 17},
		},
		{
			seed:   500,
			n:      20,
			start:  17,
			end:    19,
			output: []int{18, 0, 6},
		},
		{
			seed:   500,
			n:      20,
			start:  5,
			end:    9,
			output: []int{5, 15, 16, 3, 12},
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			ret := GetSureRandArr(test.seed, test.n, test.start, test.end)
			if !Equal(ret, test.output) {
				t.Error("sure rand arr is error")
			}
		})
	}
}

func TestCount(t *testing.T) {
	testList := []struct {
		Input  []string
		Output map[string]int
	}{
		{
			Input: []string{"a", "b", "c", "a", "a", "b"},
			Output: map[string]int{
				"a": 3,
				"b": 2,
				"c": 1,
			},
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			ret := Count(test.Input)
			if !hmap.Equal(ret, test.Output) {
				t.Error("count is error")
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testList := []struct {
		Input  []int
		Output []int
	}{
		{ // 单个
			Input:  []int{1},
			Output: []int{1},
		},
		{ // 奇数
			Input:  []int{1, 2, 3},
			Output: []int{3, 2, 1},
		},
		{ // 偶数
			Input:  []int{1, 2, 3, 4},
			Output: []int{4, 3, 2, 1},
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			in := test.Input
			Reverse(in)
			if !Equal(in, test.Output) {
				t.Error("array reverse fail")
			}
		})
	}
}
