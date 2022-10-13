package hmap

import "fmt"

// 按key字典序遍历 map
func ExampleIterateByOrder() {
	m := map[string]int{
		"b": 2,
		"c": 3,
	}
	IterateByOrder(m, func(k string, v int) {
		fmt.Printf("%s-%d\n", k, v)
	})
	// Output:
	// b-2
	// c-3
}

// 比较两 map 是否相等
func ExampleEqual() {
	var m1 map[string]string
	var m2 map[string]string
	var ret bool
	m1 = map[string]string{
		"a": "aa",
	}
	m2 = map[string]string{
		"a": "aa",
	}
	ret = Equal(m1, m2)
	fmt.Println(ret) // true
	m1 = map[string]string{
		"a": "aab",
	}
	m2 = map[string]string{
		"a": "aa",
	}
	ret = Equal(m1, m2)
	fmt.Println(ret) // false
	// Output:
	// true
	// false
}
