package harray

import (
	"fmt"
	"testing"
)

func TestTopList(t *testing.T) {
	checkFunc := func(arr1, arr2 []int) {
		if len(arr1) != len(arr2) {
			t.Error("数据不一致")
			return
		}
		size := len(arr1)
		for i := 0; i < size; i++ {
			if arr1[i] != arr2[i] {
				t.Error("数据不一致")
				return
			}
		}
	}
	comparator := func(a int, b int) int {
		return a - b
	}
	testList := []struct {
		Size       int
		Input      []int
		Output     []int
		Repeat     bool
		Comparator func(a, b int) int
	}{
		{ // 数量不足
			Size:       4,
			Input:      []int{1},
			Output:     []int{1},
			Comparator: comparator,
		},
		{ // 刚刚好
			Size:       4,
			Input:      []int{1, 3, 2, 9},
			Output:     []int{9, 3, 2, 1},
			Comparator: comparator,
		},
		{ // 超出
			Size:       5,
			Input:      []int{11, 2, 5, 8, 4, 9, 7, 20, 6},
			Output:     []int{20, 11, 9, 8, 7},
			Comparator: comparator,
		},
		{
			Size:       4,
			Input:      []int{1, 2, 3, 4, 3, 2},
			Output:     []int{4, 3, 2, 1},
			Repeat:     false,
			Comparator: comparator,
		},
		{
			Size:       4,
			Input:      []int{1, 2, 3, 4, 3, 2},
			Output:     []int{4, 3, 3, 2},
			Repeat:     true,
			Comparator: comparator,
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			topList := NewTopList(test.Size, test.Repeat, test.Comparator)
			topList.Add(test.Input...)
			checkFunc(topList.GetData(), test.Output)
		})
	}
}
