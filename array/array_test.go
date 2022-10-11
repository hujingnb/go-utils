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
