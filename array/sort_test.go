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
			input:      []int{1, 5, 7, 8, 6},
			comparator: comparator,
			output:     []int{1, 5, 6, 7, 8},
		},
		{
			input:      []int{1, 5, 7, 5, 2, 9, 2},
			comparator: comparator,
			output:     []int{1, 2, 2, 5, 5, 7, 9},
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
				t.Error("select sort error")
			}
		})
	}
}
