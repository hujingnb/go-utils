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
	// Output: [[1 2] [3 4]]
}

// 随机打乱数组顺序
func ExampleShuffle() {
	arr := []int{1, 2, 3}
	Shuffle(arr)
	for _, item := range arr {
		fmt.Println(item)
	}
	// Unordered Output:
	// 1
	// 2
	// 3
}

// 判断元素是否在数组中
func ExampleInArray() {
	arr := []int{1, 2, 3}
	ret := InArray(arr, 2)
	fmt.Println(ret)
	// Output: true
}

// 数组去重
func ExampleUnique() {
	arr := []int{1, 2, 2, 3}
	ret := Unique(arr)
	fmt.Println(ret)
	// Output: [1 2 3]
}

// 判断数组是否相等
func ExampleEqual() {
	arr1 := []int{1, 2}
	arr2 := []int{1, 2}
	ret := Equal(arr1, arr2)
	fmt.Println(ret)
	// Output: true
}
