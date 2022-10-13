// Package hmap map 操作
package hmap

import "sort"

// IterateByOrder 字典序遍历map
func IterateByOrder[TV any](data map[string]TV, f func(key string, value TV)) {
	keys := make([]string, 0)
	for key := range data {
		keys = append(keys, key)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, key := range keys {
		f(key, data[key])
	}
}

// Equal 比较两个 map 是否相同. k/v 均相同
func Equal[TK, TV comparable](m1, m2 map[TK]TV) bool {
	// 比较 map1 中的数据是否在 map2 中存在且值相等
	checkFun := func(map1, map2 map[TK]TV) bool {
		for k, v := range map1 {
			m2v, ok := map2[k]
			if !ok {
				return false
			}
			if m2v != v {
				return false
			}
		}
		return true
	}
	return checkFun(m1, m2) && checkFun(m2, m1)
}
