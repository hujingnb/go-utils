package hmap

import "fmt"

func ExampleIterateByOrder() {
	m := map[string]int{
		"b": 2,
		"c": 3,
	}
	IterateByOrder(m, func(k string, v int) {
		fmt.Printf("%s-%d\n", k, v)
	})
	/*
		Output:
			b-2
			c-3
	*/
}
