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

// 数组取交集
func ExampleIntersect() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5}
	ret := Intersect(arr1, arr2)
	for _, item := range ret {
		fmt.Println(item)
	}
	// Unordered Output:
	// 3
	// 4
}

// 数组取差集
func ExampleDiff() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{2, 3, 5}
	ret := Diff(arr1, arr2)
	for _, item := range ret {
		fmt.Println(item)
	}
	// Unordered Output:
	// 1
	// 4
}

// 二分查找
func ExampleBinarySearch() {
	// 升序列表二分查找
	arr1 := []int{1, 2, 3, 4}
	target1 := 2
	ret1 := BinarySearch(arr1, func(data int) int {
		return data - target1
	})
	fmt.Println(ret1)
	// 降序列表二分查找
	arr2 := []int{4, 3, 2, 1}
	target2 := 2
	ret2 := BinarySearch(arr2, func(data int) int {
		return target2 - data
	})
	fmt.Println(ret2)
	// Output:
	// 1
	// 2
}

// 冒泡排序
func ExampleSortBubble() {
	arr := []int{1, 7, 5, 8, 6}
	SortBubble(arr, func(a int, b int) int {
		return a - b
	})
	fmt.Println(arr)
	// Output: [1 5 6 7 8]
}

// 插入排序
func ExampleSortInsert() {
	arr := []int{1, 7, 5, 8, 6}
	SortInsert(arr, func(a int, b int) int {
		return a - b
	})
	fmt.Println(arr)
	// Output: [1 5 6 7 8]
}

// 选择排序
func ExampleSortSelect() {
	arr := []int{1, 7, 5, 8, 6}
	SortSelect(arr, func(a int, b int) int {
		return a - b
	})
	fmt.Println(arr)
	// Output: [1 5 6 7 8]
}

// 快排
func ExampleSortQuick() {
	arr := []int{1, 7, 5, 8, 6}
	SortQuick(arr, func(a int, b int) int {
		return a - b
	})
	fmt.Println(arr)
	// Output: [1 5 6 7 8]
}

// 归并排序
func ExampleSortMerge() {
	arr := []int{1, 7, 5, 8, 6}
	SortMerge(arr, func(a int, b int) int {
		return a - b
	})
	fmt.Println(arr)
	// Output: [1 5 6 7 8]
}
