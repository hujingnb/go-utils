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
