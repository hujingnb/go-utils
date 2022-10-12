package harray

import (
	hstring "github.com/hujingnb/go-utils/string"
	"testing"
)

func TestSortBubble(t *testing.T) {
	testList := []struct {
		input      []int
		comparator func(int, int) int
		output     []int
	}{
		{
			input: []int{1, 5, 7, 8, 6},
			comparator: func(a int, b int) int {
				return a - b
			},
			output: []int{1, 5, 6, 7, 8},
		},
	}
	for _, test := range testList {
		t.Run(hstring.ToString(test.input), func(t *testing.T) {
			in := test.input
			SortBubble(in, test.comparator)
			if !Equal(in, test.output) {
				t.Error("bubble sort error")
			}
		})
	}
}
