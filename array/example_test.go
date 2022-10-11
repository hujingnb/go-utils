package harray

import "fmt"

// 按照指定长度切割数组
func ExampleChunk() {
	arr := []int{1, 2, 3, 4}
	size := 2
	ret, err := Chunk(arr, size)
	if err != nil {
		return
	}
	fmt.Println(ret)
	/*
		Output:
			[[1, 2], [3, 4]]
	*/
}

// 随机打乱数组顺序
func ExampleShuffle() {
	arr := []int{1, 2, 3}
	Shuffle(arr)
	fmt.Println(arr)
}

// 判断元素是否在数组中
func ExampleInArray() {
	arr := []int{1, 2, 3}
	ret := InArray(arr, 2)
	fmt.Println(ret)
	// Output:
	//  true
}
