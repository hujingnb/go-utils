package harray

import "fmt"

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
