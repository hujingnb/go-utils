package harray

import (
	hstring "github.com/hujingnb/go-utils/string"
	"testing"
)

// 排序使用的测试用例
func getTestList() []struct {
	input      []int
	comparator func(int, int) int
	output     []int
} {
	comparator := func(a int, b int) int {
		return a - b
	}
	return []struct {
		input      []int
		comparator func(int, int) int
		output     []int
	}{
		{
			input:      []int{1},
			comparator: comparator,
			output:     []int{1},
		},
		{
			input:      []int{3, 1},
			comparator: comparator,
			output:     []int{1, 3},
		},
		{
			input:      []int{3, 1, 2},
			comparator: comparator,
			output:     []int{1, 2, 3},
		},
		{
			input:      []int{3, 1, 9, 2},
			comparator: comparator,
			output:     []int{1, 2, 3, 9},
		},
		{
			input:      []int{1, 5, 7, 8, 6},
			comparator: comparator,
			output:     []int{1, 5, 6, 7, 8},
		},
		{
			input:      []int{1, 5, 7, 5, 2, 9},
			comparator: comparator,
			output:     []int{1, 2, 5, 5, 7, 9},
		},
		{
			input:      []int{1, 5, 7, 5, 10, 2, 9},
			comparator: comparator,
			output:     []int{1, 2, 5, 5, 7, 9, 10},
		},
		{
			input:      []int{3, 1, 2, 5, 6, 43, 4, 1, 2, 3, 4, 86, 3, 2, 23, 3, 6, 10, 4, 1, 7, 4, 0},
			comparator: comparator,
			output:     []int{0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 6, 6, 7, 10, 23, 43, 86},
		},
	}
}

func TestSortBubble(t *testing.T) {
	for _, test := range getTestList() {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortBubble(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("bubble sort error")
			}
		})
	}
}

func TestSortInsert(t *testing.T) {
	for _, test := range getTestList() {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortInsert(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("insert sort error")
			}
		})
	}
}

func TestSortSelect(t *testing.T) {
	for _, test := range getTestList() {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortSelect(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("select sort error")
			}
		})
	}
}

func TestSortQuick(t *testing.T) {
	for _, test := range getTestList() {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortQuick(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("quick sort error")
			}
		})
	}
}

func TestSortMerge(t *testing.T) {
	for _, test := range getTestList() {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortMerge(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("merge sort error")
			}
		})
	}
}
