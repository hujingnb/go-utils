package hmap

import (
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
